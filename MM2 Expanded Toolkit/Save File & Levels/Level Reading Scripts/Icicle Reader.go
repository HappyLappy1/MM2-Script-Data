package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"log"

	"github.com/aead/cmac"
	"honnef.co/go/spew"
	"github.com/fatih/color"
    "github.com/rodaine/table"
)

var bcdTable = []uint32{
	0x7ab1c9d2, 0xca750936, 0x3003e59c, 0xf261014b,
	0x2e25160a, 0xed614811, 0xf1ac6240, 0xd59272cd,
	0xf38549bf, 0x6cf5b327, 0xda4db82a, 0x820c435a,
	0xc95609ba, 0x19be08b0, 0x738e2b81, 0xed3c349a,
	0x045275d1, 0xe0a73635, 0x1debf4da, 0x9924b0de,
	0x6a1fc367, 0x71970467, 0xfc55abeb, 0x368d7489,
	0x0cc97d1d, 0x17cc441e, 0x3528d152, 0xd0129b53,
	0xe12a69e9, 0x13d1bdb7, 0x32eaa9ed, 0x42f41d1b,
	0xaea5f51f, 0x42c5d23c, 0x7cc742ed, 0x723ba5f9,
	0xde5b99e3, 0x2c0055a4, 0xc38807b4, 0x4c099b61,
	0xc4e4568e, 0x8c29c901, 0xe13b34ac, 0xe7c3f212,
	0xb67ef941, 0x08038965, 0x8afd1e6a, 0x8e5341a3,
	0xa4c61107, 0xfbaf1418, 0x9b05ef64, 0x3c91734e,
	0x82ec6646, 0xfb19f33e, 0x3bde6fe2, 0x17a84cca,
	0xccdf0ce9, 0x50e4135c, 0xff2658b2, 0x3780f156,
	0x7d8f5d68, 0x517cbed1, 0x1fcddf0d, 0x77a58c94,
}

type Random struct {
	s0 uint32
	s1 uint32
	s2 uint32
	s3 uint32
}

func (r *Random) u32() uint32 {
	// https://github.com/kinnay/NintendoClients/blob/d41d394065900c7214906b5a87da35f561b16763/nintendo/sead/random.py
	temp := r.s0
	temp = (temp ^ (temp << 11)) & 0xFFFFFFFF
	temp ^= temp >> 8
	temp ^= r.s3
	temp ^= r.s3 >> 19
	r.s0 = r.s1
	r.s1 = r.s2
	r.s2 = r.s3
	r.s3 = temp
	return temp
}

func (r *Random) uint(max uint32) uint64 {
	return (uint64(r.u32()) * uint64(max)) >> 32
}

func createKey(r *Random, table []uint32, size uint32, writer *bytes.Buffer) {
	for i := uint32(0); i < size/4; i++ {
		value := uint32(0)
		for e := 0; e < 4; e++ {
			index := r.uint(uint32(len(table)))
			shift := r.uint(4) * 8
			b := (table[index] >> shift) & 0xFF
			value = (value << 8) | b
		}
		binary.Write(writer, binary.LittleEndian, uint32(value))
	}
}

func DecryptLevel(buf []byte) ([]byte, error) {
	if len(buf) != 0x5c000 {
		return []byte{}, fmt.Errorf("invalid buf size %d != %d", len(buf), 0x5c000)
	}

	end := 0x5BFD0
	writer := new(bytes.Buffer)

	// Create random instance
	r := &Random{
		binary.LittleEndian.Uint32(buf[end+0x10 : end+0x14]),
		binary.LittleEndian.Uint32(buf[end+0x14 : end+0x18]),
		binary.LittleEndian.Uint32(buf[end+0x18 : end+0x1C]),
		binary.LittleEndian.Uint32(buf[end+0x1C : end+0x20]),
	}

	cmacWant := buf[end+0x20 : end+0x30]
	crcWant := buf[8:12]

	// Construct AES instance
	aesKey := new(bytes.Buffer)
	createKey(r, bcdTable, 0x10, aesKey)

	aesBlock, err := aes.NewCipher(aesKey.Bytes())
	if err != nil {
		return nil, err
	}

	aesMode := cipher.NewCBCDecrypter(aesBlock, buf[end:end+0x10])
	decrypted := make([]byte, 0x5BFC0)
	aesMode.CryptBlocks(decrypted, buf[0x10:0x5BFD0])

	// crc check
	if crc32.ChecksumIEEE(decrypted) != binary.LittleEndian.Uint32(crcWant) {
		return nil, fmt.Errorf("crc invalid")
	}

	// cmac check
	cmacKey := new(bytes.Buffer)
	createKey(r, bcdTable, 0x10, cmacKey)
	cmacBlock, err := aes.NewCipher(cmacKey.Bytes())
	if err != nil {
		return nil, err
	}
	cmacDigest, err := cmac.Sum(decrypted, cmacBlock, 0x10)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(cmacDigest, cmacWant) {
		return nil, fmt.Errorf("cmac invalid")
	}

	if false {
		// write bcd header
		_, err = writer.Write(buf[:0x10])
		if err != nil {
			return nil, err
		}
	}

	// Decrypted course data
	_, err = writer.Write(decrypted)
	if err != nil {
		return nil, err
	}

	return writer.Bytes(), nil
}

func EncryptLevel(buf []byte) ([]byte, error) {
	var withoutBcdHeader bool

	if len(buf) == 0x5BFD0-0x10 {
		withoutBcdHeader = true
	} else if len(buf) == 0x5BFD0 {
		// with bcd header
	} else {
		return []byte{}, fmt.Errorf("invalid buf size %d != %d (%d)", len(buf), 0x5BFD0, 0x5BFC0-0x10)
	}

	var reader *bytes.Reader
	if withoutBcdHeader {
		reader = bytes.NewReader(buf)
	} else {
		reader = bytes.NewReader(buf[0x10:])
	}

	decrypted := make([]byte, 0x5BFC0)
	_, err := io.ReadFull(reader, decrypted)
	if err != nil {
		return nil, err
	}

	writer := new(bytes.Buffer)

	if withoutBcdHeader {
		var data = []any{
			uint32(0x1),
			uint16(0x10),
			uint16(0x0),
			uint32(crc32.ChecksumIEEE(decrypted)),
			uint8(0x53),
			uint8(0x43),
			uint8(0x44),
			uint8(0x4C),
		}
		for _, v := range data {
			err := binary.Write(writer, binary.LittleEndian, v)
			if err != nil {
				return nil, err
			}
		}
	} else {
		binary.LittleEndian.PutUint32(buf[0x8:(0x8+4)], crc32.ChecksumIEEE(decrypted)) // shouldn't be needed, only when modifing course data
		writer.Write(buf[:0x10])
	}

	// Technically random bytes, can make deterministic here
	randomSeed := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	r := &Random{
		binary.LittleEndian.Uint32(randomSeed[0:4]),
		binary.LittleEndian.Uint32(randomSeed[4:8]),
		binary.LittleEndian.Uint32(randomSeed[8:12]),
		binary.LittleEndian.Uint32(randomSeed[12:16]),
	}
	aesIv := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	aesKey := new(bytes.Buffer)
	createKey(r, bcdTable, 0x10, aesKey)
	aesBlock, err := aes.NewCipher(aesKey.Bytes())
	if err != nil {
		return nil, err
	}

	aesMode := cipher.NewCBCEncrypter(aesBlock, aesIv)
	encrypted := make([]byte, 0x5BFC0)
	aesMode.CryptBlocks(encrypted, decrypted)

	cmacKey := new(bytes.Buffer)
	createKey(r, bcdTable, 0x10, cmacKey)
	cmacBlock, err := aes.NewCipher(cmacKey.Bytes())
	if err != nil {
		return nil, err
	}

	cmacDigest, err := cmac.Sum(decrypted, cmacBlock, 0x10)
	if err != nil {
		return nil, err
	}

	_, err = writer.Write(encrypted)
	if err != nil {
		return nil, err
	}

	_, err = writer.Write(aesIv)
	if err != nil {
		return nil, err
	}

	_, err = writer.Write(randomSeed)
	if err != nil {
		return nil, err
	}

	_, err = writer.Write(cmacDigest)
	if err != nil {
		return nil, err
	}

	return writer.Bytes(), nil
}

func DecryptReplay(buf []byte) ([]byte, error) {
	return nil, nil
}

type Header struct {
	YStart                  uint8  // Starting Y position of level
	YGoal                   uint8  // Y position of goal (in tiles)
	XGoal                   uint16 //	X position of goal (in tiles * 10)
	TimeLimit               uint16 // Time limit; defaults to 300
	ClearConditionMagnitude uint16 // Target amount required for clear condition
	CreationYear            uint16 //	Creation year
	CreationMonth           uint8  //	Creation month
	CreationDay             uint8  //	Creation day
	CreationHour            uint8  //	Creation hour
	CreationMinute          uint8  //	Creation minute
	AutoscrollSpeed         uint8
	ClearConditionCategory  uint8  //	Clear condition type (more below)
	ClearConditionObject    uint32 // Clear condition object (CRC32 of more below)
	UnkGameVer              uint32
	ManagementFlags         uint32 // Management flags: &1 always seems to be set, &2 shows that a level has passed its Clear Check, &0x10 shows that the level may not be uploaded
	ClearAttemmpts          uint32
	ClearCheckTime          uint32 // Time taken in Clear Check (units unknown), or 0xFFFFFFFF if the level has not been cleared
	CreationId              uint32 // Initialised to a random value when a level is created
	UploadId                uint64
	GameVersion             uint32
	Unk1                    [0xBD]byte
	GameStyle               [0x2]byte // Game style: M1, M3, MW, WU, 3W
	Unk2                    uint8
	Name                    [0x42]byte // Course name, null-terminated, UCS-2
	Description             [0xCA]byte //	Course description, null-terminated, UCS-2 (game only lets you enter up to 75, but there's space for 100)
}

type Object struct {
	X      uint32
	Y      uint32
	Unk1   uint16
	Width  uint8
	Height uint8
	Flag   uint32
	CFlag  uint32
	Ex     uint32
	Id     uint16
	CId    uint16
	LId    uint16
	SId    uint16
}

type Sound struct {
	Id   uint8
	X    uint8
	Y    uint8
	Unk1 uint8
}

type SnakeNode struct {
	Index     uint16
	Direction uint16
	Id        uint32
}

type Snake struct {
	Index     uint8
	NodeCount uint8
	Unk1      uint16
	Nodes     [120]SnakeNode
}

type ClearPipeNode struct {
	Type      uint8
	Index     uint8
	X         uint8
	Y         uint8
	Width     uint8
	Height    uint8
	Unk1      uint8
	Direction uint8
}

type ClearPipe struct {
	Index     uint8
	NodeCount uint8
	Unk       uint16
	Nodes     [36]ClearPipeNode
}

type PiranhaCreeperNode struct {
	Unk1      uint8
	Direction uint8
	Unk2      uint16
}

type PiranhaCreeper struct {
	Unk1      uint8
	Index     uint8
	NodeCount uint8
	Unk2      uint8
	Nodes     [20]PiranhaCreeperNode
}

type ExclamationBlockNode struct {
	Unk1      uint8
	Direction uint8
	Unk2      uint16
}

type ExclamationBlock struct {
	Unk1      uint8
	Index     uint8
	NodeCount uint8
	Unk2      uint8
	Nodes     [10]ExclamationBlockNode
}

type TrackBlockNode struct {
	Unk1      uint8
	Direction uint8
	Unk2      uint16
}

type TrackBlock struct {
	Unk1      uint8
	Index     uint8
	NodeCount uint8
	Unk2      uint8
	Nodes     [10]TrackBlockNode
}

type Ground struct {
	X            uint8
	Y            uint8
	Id           uint8
	BackgroundId uint8
}

type Track struct {
	Unk1  uint16
	Flags uint8
	X     uint8
	Y     uint8
	Type  uint8
	LId   uint16
	Unk2  uint16
	Unk3  uint16
}

type Icicle struct {
	X    uint8
	Y    uint8
	Type uint8
	Unk1 uint8
}

type LevelArea struct {
	Theme                     uint8
	AutoscrollType            uint8
	BoundaryType              uint8
	Orientation               uint8
	LiquidEndHeight           uint8
	LiquidMode                uint8
	LiquidSpeed               uint8
	LiquidStartHeight         uint8
	BoundaryRight             uint32
	BoundaryTop               uint32
	BoundaryLeft              uint32
	BoundaryBottom            uint32
	UnkFlag                   uint32
	ObjectCount               uint32
	SoundEffectCount          uint32
	SnakeBlockCount           uint32
	ClearPipeCount            uint32
	PiranhaCreeperCount       uint32
	ExclamationMarkBlockCount uint32
	TrackBlockCount           uint32
	Unk1                      uint32
	GroundCount               uint32
	TrackCount                uint32
	IceCount                  uint32
	Objects                   [2600]Object
	Sounds                    [300]Sound
	Snakes                    [5]Snake
	ClearPipes                [200]ClearPipe
	PiranhaCreepers           [10]PiranhaCreeper
	ExclamationBlocks         [10]ExclamationBlock
	TrackBlocks               [10]TrackBlock
	Ground                    [4000]Ground
	Tracks                    [1500]Track
	Icicles                   [300]Icicle
	Unk2                      [0xDBC]byte
}

type BCD struct {
	//VersionNumber1 uint32  // Assumed version number -- must be 1 in v1.0.0
	//VersionNumber2 uint16  // Assumed version number -- must be 0x10 in v1.0.0
	//Padding1       [2]byte // 2 empty bytes
	//CRC32          uint32  // CRC32 of the decrypted level file from offset 0x10 to 0x5BFD0 (non-inclusive)
	//Magic          [4]byte // Magic SCDL (53 53 44 4C)
	Header    Header // header 0x200
	OverWorld LevelArea
	SubWorld  LevelArea
}

func LoadBCD(buf []byte) (*BCD, error) {
	s := &BCD{}
	err := s.Load(buf)
	return s, err
}

func (s *BCD) Load(buf []byte) error {
	var err error
	buf, err = DecryptLevel(buf)
	if err != nil {
		return err
	}
	return binary.Read(bytes.NewReader(buf), binary.LittleEndian, s)
}

func (s *BCD) Save() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.LittleEndian, s)
	if err != nil {
		return []byte{}, err
	}
	out, err := EncryptLevel(buf.Bytes())
	return out, err
}

func RemoveUploadedFlag(rawBCD []byte) ([]byte, error) {
	level, err := LoadBCD(rawBCD)
	if err != nil {
		return []byte{}, err
	}

	level.Header.UploadId = 0
	level.Header.ManagementFlags = (level.Header.ManagementFlags &^ (1 << 16))
	level.Header.ManagementFlags = (level.Header.ManagementFlags &^ (1 << 4))
	level.Header.ManagementFlags = (level.Header.ManagementFlags &^ (1 << 6))

	newBCD, err := level.Save()
	if err != nil {
		return []byte{}, err
	}
	return newBCD, nil
}

func debug_upload_ban() error {
	buf, err := ioutil.ReadFile("data/level_tests/clean_short.bcd")
	if err != nil {
		return err
	}

	level, err := LoadBCD(buf)
	if err != nil {
		return err
	}

	//spew.Dump(level.Header)
	fmt.Println(fmt.Sprintf("%032b", level.Header.ManagementFlags), level.Header.ManagementFlags)

	buf, err = ioutil.ReadFile("data/level_tests/upload_banned.bcd")
	if err != nil {
		return err
	}

	level, err = LoadBCD(buf)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("%032b", level.Header.ManagementFlags), level.Header.ManagementFlags)
	level.Header.ManagementFlags = (level.Header.ManagementFlags &^ (1 << 4))
	level.Header.ManagementFlags = (level.Header.ManagementFlags &^ (1 << 6))
	fmt.Println(fmt.Sprintf("%032b", level.Header.ManagementFlags), level.Header.ManagementFlags)

	return nil
}

func test_roundtrip() error {
	//buf, err := ioutil.ReadFile("course_data_000.bcd")
	buf, err := ioutil.ReadFile("course_data_000.bcd")
	if err != nil {
		return err
	}

	level, err := LoadBCD(buf)
	if err != nil {
		return err
	}
	if false {
		level.Header.ClearConditionCategory = 2
		level.Header.ClearConditionObject = 46219146 // stone in goal
		level.OverWorld.Objects[0].Id = 75           // stone, sadly fails to pass corrupt check on upload
		//level.OverWorld.Objects[0].Id = 5 // ? block
	}

	spew.Dump(level.Header)
	//spew.Dump(level.Header, level.OverWorld.Objects[:int(level.OverWorld.ObjectCount)])
	out, err := level.Save()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("course_data_001.bcd", out, 0600)
	if err != nil {
		return err
	}

	if true {
		// round-trip check
		level2, err := LoadBCD(out)
		if err != nil {
			return err
		}
		//spew.Dump(level2.Header, level2.OverWorld.Objects[:int(level2.OverWorld.ObjectCount)])
		out2, err := level2.Save()
		if err != nil {
			return err
		}

		spew.Dump(bytes.Equal(out, out2))
	}

	return err
}

func main() {
  // save bcd
  // Choose which level file to read (such as "course_data_000.bcd")
  rawBCD, err := ioutil.ReadFile("course_data_000.bcd")
  if err != nil {
    log.Fatal(err)
  }

  level, err := LoadBCD(rawBCD)
  if err != nil {
    log.Fatal(err)
  }
   
   	
	fmt.Println("OverWorld Icicles: ", level.OverWorld.IceCount)
    headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
    columnFmt := color.New(color.FgYellow).SprintfFunc()
    //X-Pos (X): X position of the object in full blocks (I think?)
	//Y-Pos (Y): Y position of the object in full blocks (I think?)
	//Icicle Type (Type): Is the icicle light blue, or dark blue? 
	tbl := table.New("Index", 
	"X-Pos", 
	"Y-Pos",
	"Type")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
    for i := 0; i < int(level.OverWorld.IceCount); i++ {
		tbl.AddRow(i, 
		level.OverWorld.Icicles[i].X, 
		level.OverWorld.Icicles[i].Y,
		level.OverWorld.Icicles[i].Type)
		}
	tbl.Print()
	fmt.Println("SubWorld Icicles: ", level.SubWorld.IceCount)
	tbl = table.New("Index", 
	"X-Pos", 
	"Y-Pos",
	"Type")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
    for i := 0; i < int(level.SubWorld.IceCount); i++ {
		tbl.AddRow(i, 
		level.SubWorld.Icicles[i].X, 
		level.SubWorld.Icicles[i].Y,
		level.SubWorld.Icicles[i].Type)
		}
    tbl.Print()	
  // save bcd
  newBCD, err := level.Save()
  if err != nil {
    log.Fatal(err)
  }
  // Output should be unchanged from initial file, will distinguish anyway in case it matters
  err = ioutil.WriteFile("course_data_00N.bcd", newBCD, 0600)
  if err != nil {
    log.Fatal(err)
  }
}