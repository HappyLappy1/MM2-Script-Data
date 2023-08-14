


## Flags for objects

### General Structure 
| Flag      | Description | Extended Description  |
|-----------|-------------|--------------------------------------|
| 0x1       | In Pipe     | Object thinks it's in a pipe |
| 0x2       | Wings       | Object thinks it has wings | 
| 0x4       | Alt Form 2/2| One of the places an object can store an alternate form. Typically used for Binary alt-forms (like goombas and goombrats) |
| 0x8       | Direction A1| One of the places an object can store an alternate form/direction. May be a bitfield with 0x4 and/or 0x10. (like Thwomp Direction)|
| 0x10      | Direction A2| One of the places an object can store an alternate form/direction. May be a bitfield with 0x4 and/or 0x8. (like Thwomp Direction)|
| 0x20      | Face Left / Pipe Left  | Bitfield with 0x40. One of the places to store object direction, or to store direction of pipe for editor sprite. A goomba in a sideways pipe appears sideways because this flag is set. |
| 0x40      | Face Left / Pipe Up / Always | Bitfield with 0x40. One of the places to store object direction, or to store direction of pipe for editor sprite. All objects, even those not in pipes or unable to enter them, will have 0x40 by default. |
| 0x80      | Is Container/Has Contents| Used by Pipes/Clowncars/Clouds/Claws to detect if something has been placed inside. "Is Container" is used on objects that don't have this flag normally, while "Has Contents" is used on objects that can. If "Unused", this flag does not affect the object|
| 0x100     | Stretch | Only boos make use of this flag, and it detects whether the boo is a stretch. Weird, huh? |
| 0x200     | In Clowncar | Object thinks it's in a clowncar |
| 0x400     | On a Track  | Object thinks it's on tracks |
| 0x800     | Unused      | Unused globally |
| 0x1000    | In Stack    | Object thinks it's in a stack |
| 0x2000    | Medium      | Medium Sized enemies are a brand new discovery. Generally 1.5x in size, and have behavior consistent with their small counterparts |
| 0x4000    | Big / 2x2   | Object believes it has ingested a mushroom, and grown to 2x in size. Many blocks react to this by becoming a 2x2 grid of themselves|
| 0x8000    | Parachute   | Object believes it has a parachute|
| 0x10000   | In Cloud    | Object believes it is in a cloud. This is the one the game always uses|
| 0x20000   | In Cloud(?) | Object ALSO believes it is in a cloud. Could be a bitfield, or a duplicate flag|
| 0x40000   | Alt Form 2/3| One of the places an object can store an alternate form. Typically used for Trinary alt-forms (like 10/30/50 coin type). | 
| 0x80000   | Alt Form 2/3| One of the places an object can store an alternate form. Typically used for Trinary alt-forms (like 10/30/50 coin type). |
| 0x100000  | Left Tracks / Begin Entry Index | If the object is on tracks, it would travel to the left on those tracks; is a bitfield with 0x200000. If a door, warp box, or pipe, this is a bitfield of variable size that pairs doors/warp-boxes/pipes| 
| 0x200000  | Vert Tracks / Cont. Entry Index | If the object is on tracks, it is travelling vertically on those tracks; Is a bitfield with 0x200000. If a door, warp box, or pipe, this is a bitfield of variable size that pairs doors/warp-boxes/pipes| 
| 0x400000  | Alt Form 1/N / Cont. Entry Index | Beginning of a bitfield of variable size for storing alt-forms of objects with many of them (such as arrows with 8 unique directions). If a door, warp box, or pipe, this is a bitfield of variable size that pairs doors/warp-boxes/pipes. |
| 0x800000  | Alt Form 3/N / Cont. Entry Index | A bitfield of variable size for storing alt-forms of objects with many of them (such as arrows with 8 unique directions). If a door, warp box, or pipe, this is a bitfield of variable size that pairs doors/warp-boxes/pipes. |
| 0x1000000 | Alt Form 5/N / Cont. Entry Index | of a bitfield of variable size for storing alt-forms of objects with many of them (such as arrows with 8 unique directions). If a warp pipe, this is a bitfield of variable size that pairs warp pipes. |
| 0x2000000 | L-Wall Hang | Object is hanging from the left wall. Bitfield with 0x4000000 | 
| 0x4000000 | Roof Hang   | Object is hanging from the ceiling. Bitfield with 0x2000000 |  
| 0x6000000 | Ground Hang/Always| All objects have this flag, unless they are capable of clinging to a wall or hanging from a ceiling, and are currently doing so. |  
| 0x8000000 | In Claw     | Object believes its' in a claw. For some objects, this means moving up half a tile and no other behavioral change. |
| 0x10000000| Unused      | Unused Globally |
| 0x20000000| Unused      | Unused Globally |
| 0x40000000| Unused      | Unused Globally |
| 0x80000000| Unused      | Unused Globally |

### [0] Goomba/Goombrat/Galoomba/Galoombud
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Goombrat    |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x20, 0x40, 0x60: If unset, faces right. Primarily used for objects in pipes. Set by default to 0x40 if not in a pipe.
0x100000: Moves left on horizontal or diagonal tracks, right if unset. Moves Down on vertical tracks (Up if Unset). 
0x6000000: Always set on this enemy


### [1] Koopa Troopa
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Red Shell   |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x20, 0x40, 0x60: If unset, faces right. Primarily used for objects in pipes. Set by default to 0x40 if not in a pipe.
0x100000: Moves left on horizontal or diagonal tracks, right if unset. Moves Down on vertical tracks (Up if Unset). 
0x6000000: Always set on this enemy

### [2] Piranha Plant
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Fire Plant  |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | L-Wall Hang |
| 0x4000000 | Roof Hang   |
| 0x6000000 | Ground Hang |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x20, 0x40, 0x60: If unset, faces right. Primarily used for objects in pipes. Set by default to 0x40 if not in a pipe.
0x2000000, 0x4000000, 0x6000000: Enemy can cling to walls/ceiling. Hanging on Right Wall (R-Wall) if unset (0x0000000)
0x100000: Moves left on horizontal or diagonal tracks, right if unset. Moves Down on vertical tracks (Up if Unset). 

### [3] Hammer Bro
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Fire Bro    |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Big Bro     |
| 0x4000    | Sledge Bro  |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x20, 0x40, 0x60: If unset, faces right. Primarily used for objects in pipes. Set by default to 0x40 if not in a pipe.
Note: Fire Bros Crash outside 3DW


### [4] Brick Block 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | 2x2 Bricks  |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Winged Blocks are VERY different to normal blocks, and their behavior has been tested independently. 
0x1: No wings = empty pipe
0x20, 0x40, 0x60: Wings rotate in editor, but not while playing.
0x80: Causes block to despawn if winged. This flag is typically used on pipes, claws, clowncars, and clouds to signify "some other object is riding/inside me".
0x4000: 1x1 if winged. Objects inside (CId) can be released from all four bricks. The individual blocks act like their own brick blocks. 
0x8000,0x10000, 0x20000: Despawns if NOT winged. No unique behavior otherwise.
0x8000000: Despawns if not winged. seems to work in claw. 
 
### [5] Question Block 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | 2x2 Blocks  |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Winged Blocks are VERY different to normal blocks, and their behavior has been tested independently. 
0x1: No wings = empty pipe
0x20, 0x40, 0x60: Wings rotate in editor, but not while playing.
0x80: Causes block to despawn if winged. This flag is typically used on pipes, claws, clowncars, and clouds to signify "some other object is riding/inside me".
0x4000: 1x1 if winged. Objects inside (CId) can be released from all four blocks. The individual blocks act like their own blocks. 
0x8000,0x10000, 0x20000: Despawns if NOT winged. No unique behavior otherwise.
0x8000000: Despawns if not winged. seems to work in claw. 

### [6] Hard Block
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | 2x2 Blocks  |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Winged Blocks are VERY different to normal blocks, and their behavior has been tested independently. 
0x1: No wings = empty pipe
0x20, 0x40, 0x60: Wings rotate in editor, but not while playing.
0x80: Causes block to despawn if winged. This flag is typically used on pipes, claws, clowncars, and clouds to signify "some other object is riding/inside me".
0x4000: 1x1 if winged. The individual blocks act like their own blocks. 
0x8000,0x10000, 0x20000: Despawns if NOT winged. No unique behavior otherwise.
0x8000000: Despawns if not winged. seems to work in claw. 

### [7] Ground Tile
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | 2x2 Blocks  |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Ground Tile objects are mostly place-holders/used in the editor. Ground tile objects are removed upon saving a level, and entered into their own data type.
I imagine none of this jank will survive a save/reload of the level.
0x1: Doesn't actually work, but causes some funny sprite shenanigans
0x2, 0x400, 0x800, 0x8000, 0x10000, 0x20000, 0x8000000: Spawns Ground tile sprite, but NOT the collision, similar to ghost blocks.
0x4000: Produces a 1x1 visible ground tile, and a 2x2 invisible ground tile. 

### [8] Coin/Frozen Coin
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Frozen      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Coins, Winged Coins, Frozen Coins, and Winged Frozen coins have each been tested, as they act quite differently. 
All Coins that aren't "loose" or winged are affected by a P-switch.
0x1: Coins work as expected, Frozen coins dispense identically to normal coins (as do winged frozen coins to winged coins). 
0x20, 0x40, 0x60: These look funny. If unset, faces right. Primarily used for objects in pipes. Set by default to 0x40 if not in a pipe.
0x200, 0x400, 0x8000, 0x10000, 0x20000, : Frozen immune to P-switch, loose coin when melted. Same for frozen winged coin.
0x2000: Coin is normal sized, winged coin is actually medium, Frozen & Frozen Winged unaffected.
0x4000: Coin is a 2x2 grid of coins (like a block), Winged coin is actually large, Frozen coin is a 2x2 block, Winged frozen coin is normal.
0x80000000: Works on all coin types, even frozen coins. Frozen coins bumped up half a tile. Immune to p-switch, loose coin released when melted.

### [9] Pipe
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | BIIIIG      |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Red Pipe    |
| 0x80000   | Blue Pipe   |
| 0xC0000   | Yellow Pipe |
| 0x100000  | Entry #1    |
| 0xF00000  | Entry #15   |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | Unused      |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: I only conducted this test on 1x1 Pipes, which may have glitchy behavior. 
0x1, 0x200, 0x10000, 0x20000, 0x400000, 0x8000000: Could not place objects inside. Suggestive of remnant editor behavior of Wings/Clowncar/Cloud/Etc.
0x20, 0x40, 0x60: Turns the pipe to face that direction. For 1x1 pipes, rotates center of sprite, but hitbox is rotated at a corner (or where the center of a 2x2 pipe would be). 
0x80: Has no weird behavior if artificially applied, but in normal game circumstances is used to signify "a goomba/mushroom/coin/etc. is in this pipe!". 
0x4000: Pipe is 3x3, though the sprite looks like several pipes got overlapped. Functions fine.
0x40000, 0x80000, 0xC0000: Green if unset. 
0x100000 to 0xF00000: Sends the player to the oldest pipe with the same "entry flag" in the subworld, if one exists. Softlocks otherwise. 
The Pipe limit of ten could be arbitrary, with #11-15 being valid. 

### [10] Spring  
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Horizontal  |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |


0x20, 0x40, 0x60: These look funny. If unset, faces right. Primarily used for objects in pipes. Set by default to 0x40 if not in a pipe.
0x100000: Moves left on horizontal or diagonal tracks, right if unset. Moves Down on vertical tracks (Up if Unset).

### [11] Lift  
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Blue Lift   |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Move Left   |
| 0x40      | Move Up     |
| 0x60      | Move Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | "Big"       |
| 0x8000    | Parachute   |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: one-tile lifts are funky, and are hard to differentiate. 
0x20, 0x40, 0x60: If unset, faces right. Lifts use these to know which direction to move in. 
0x4000: Not really bigger, just max length. Unused in-game.
0x8000000: Floats half a tile higher.

### [12] Thwomp  
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Crush Down  |
| 0x10      | Crush Right |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x1: The thwomp is evidently in a pipe, but does not exit it under any circumstances: Unfortunate.
0x20, 0x40, 0x60: These work, and are a joy to see.

### [13] Bullet Bill Blaster
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Wings       |
| 0x4       | Red Blaster |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |
0x1: The blaster is evidently in a pipe, but does not exit it under any circumstances: Unfortunate.
0x2: The blaster sprouts wings, and chases mario down like a winged goomba. It still has bill blaster collision, and is only capable of crushing Mario, but it will never fire.
0x20, 0x40, 0x60: These work, and are a joy to see.
0x2000: Only sprite is bigger. Looks funky if extended up or down. Bullet bills are not bigger, and are shot from below barrel.
0x4000: Is bigger, collision and all, but resets height after returning to editor. Bullet bills are not bigger, and are shot from below barrel.
0x8000: Invisible Parachute, but it DOES work. does not fire while parachuting. 
0x8000000: Half a tile higher.

### [14] Mushroom Platform
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Unused      |
| 0x40      | Always      |
| 0x60      | Unused      |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Green       |
| 0x80000   | Yellow      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | Unused      |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

All testing conducted on 1x1 Mushroom platforms. 
0x2000, 0x4000: Medium is offset half a tile, and Big became 5x5. 
0x40000, 0x80000: Wowza Bowza, the semisolid changed color!

### [15] Bob Omb 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Pre-Ignited |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |


0x20, 0x40, 0x60: These look funny. If unset, faces right. Primarily used for objects in pipes. Set by default to 0x40 if not in a pipe.
0x100000: Moves left on horizontal or diagonal tracks, right if unset. Moves Down on vertical tracks (Up if Unset).

### [16] Semisolid Platform (2DW) 
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Unused      |
| 0x40      | Always      |
| 0x60      | Unused      |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Type #2     |
| 0x80000   | Type #3     |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | Unused      |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

All testing conducted on 1x1 Semisolid platforms. 
0x2000, 0x4000: Medium is offset half a tile, and Big became 5x5. 
0x40000, 0x80000: Wowza Bowza, the semisolid transformed!

### [17] Bridge  
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Unused      |
| 0x40      | Always      |
| 0x60      | Unused      |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

All testing conducted on 1x1 Bridges. 
0x2000, 0x4000: Medium is offset half a tile, and Big became a 3x5 sprite, with a 1x5 collision. Annoying self-reset habit.
 
### [18] P-Switch 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | L-Wall Hang |
| 0x4000000 | Roof Hang   |
| 0x6000000 | Ground Hang |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x20, 0x40, 0x60: These look funny. If unset, faces right. Primarily used for objects in pipes. Set by default to 0x40 if not in a pipe.
0x100000: Moves left on horizontal or diagonal tracks, right if unset. Moves Down on vertical tracks (Up if Unset).
0x2000000, 0x4000000, 0x6000000: P-switches cannot cling to walls, but the same flag that works on Piranha Plants will turn the p-switch sideways. It will not stick to anything (at least in 3.0.2)

### [19] POW Block
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Red Pow     |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x4: Only for 3DW
0x20, 0x40, 0x60: These look funny. If unset, faces right. Primarily used for objects in pipes. Set by default to 0x40 if not in a pipe.
0x100000: Moves left on horizontal or diagonal tracks, right if unset. Moves Down on vertical tracks (Up if Unset).

### [20] Super Mushroom/Master Sword
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Master Sword|
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x20, 0x40, 0x60: These look funny. If unset, faces right. Primarily used for objects in pipes. Set by default to 0x40 if not in a pipe.
0x100000: Moves left on horizontal or diagonal tracks, right if unset. Moves Down on vertical tracks (Up if Unset).

### [21] Donut Block 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | 2x2 Bricks  |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Winged Blocks are VERY different to normal blocks, and their behavior has been tested independently. 
0x1: No wings = empty pipe
0x20, 0x40, 0x60: Wings rotate in editor, but not while playing.
0x80: Causes block to despawn if winged. This flag is typically used on pipes, claws, clowncars, and clouds to signify "some other object is riding/inside me".
0x4000: 1x1 if winged. Objects inside (CId) can be released from all four bricks. The individual blocks act like their own blocks. 
0x8000,0x10000, 0x20000: Despawns if NOT winged. No unique behavior otherwise.
0x8000000: Despawns if not winged. seems to work in claw. 

### [22] Cloud  
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | 2x2 Blocks  |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Winged Blocks are VERY different to normal blocks, and their behavior has been tested independently. 
0x1: No wings = empty pipe
0x20, 0x40, 0x60: Wings rotate in editor, but not while playing.
0x80: Causes block to despawn if winged. This flag is typically used on pipes, claws, clowncars, and clouds to signify "some other object is riding/inside me".
0x4000: 1x1 if winged.
0x8000,0x10000, 0x20000: Despawns if NOT winged. No unique behavior otherwise.
0x8000000: Despawns if not winged. seems to work in claw. 

### [23] Note Block 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Pink        |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | 2x2 Blocks  |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Winged Blocks are VERY different to normal blocks, and their behavior has been tested independently. 
0x1: No wings = empty pipe
0x20, 0x40, 0x60: Wings rotate in editor, but not while playing.
0x80: Causes block to despawn if winged. This flag is typically used on pipes, claws, clowncars, and clouds to signify "some other object is riding/inside me".
0x4000: 1x1 if winged.
0x8000,0x10000, 0x20000: Despawns if NOT winged. No unique behavior otherwise.
0x8000000: Despawns if not winged. seems to work in claw. 

### [24] Fire Bar 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | C-Clockwise |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x400000  | +2 Fireballs|
| 0x800000  | +4 Fireballs|
| 0x1000000 | +8 Fireballs|
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x1: Yes, Firebars spawn from pipes! 
0x20, 0x40, 0x60: Hard to see unless you zoom out (arrows in the way), but they are rotated. easier to see in SMB1.
0x80: Despawns.
0x2000, 0x4000: Editor sprite is bigger, no changes after hitting Play.
0x8000000: up half a tile, unsure if it'd stay in a real claw.

### [25] Spiny / Spiny Shellmet
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Shellmet    |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | L-Wall Hang |
| 0x4000000 | Roof Hang   |
| 0x6000000 | Ground Hang |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x2: Spiny Shellmet with wings can shoot spikes while on a track or in a claw. 
0x20, 0x40, 0x60: These look funny. If unset, faces right. Primarily used for objects in pipes. Set by default to 0x40 if not in a pipe.
0x2000, 0x4000: Medium Spiny shellmet can be held but not worn. Large Shellmet cannot be held nor worn. 
0x100000: Moves left on horizontal or diagonal tracks, right if unset. Moves Down on vertical tracks (Up if Unset).
0x2000000, 0x4000000, 0x6000000: Spinies cannot cling to walls, but the same flag that works on Piranha Plants will turn the Spiny sideways. It will not stick to anything (at least in 3.0.2)

### [26] Goal Ground 
| Flag    | Description |
|---------|-------------|
| All     | Unused      |

Note: This is unsurprising, as Goal Ground is (probably) part of the Ground tile data structure, and is furthermore never placable in the editor. 
Probably Always has 0x6000000 and 0x40, or it might be corrupt. Also, Flagpoles can *sometimes* be corrupt if Goal Ground is not present near/beneath it. Not thoroughly tested.
Generates Phantom Ground tiles beneath it all the way down. 
### [27] Goal/Flagpole 
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Unused      |
| 0x40      | Always      |
| 0x60      | Unused      |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | Unused      |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x80: Despawns if set.
0x8000000: Flagpole is moved up half a tile, and therefore *might* stay in a claw. 
### [28] Buzzy Beetle / Buzzy Sbell
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Shellmet    |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | L-Wall Hang |
| 0x4000000 | Roof Hang   |
| 0x6000000 | Ground Hang |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x2: Winged buzzy beetle shells are strange. On tracks, they act like parabeetles. they will just fall infinitely if left as is. Funky stuff. 
0x20, 0x40, 0x60: These look funny. If unset, faces right. Primarily used for objects in pipes. Set by default to 0x40 if not in a pipe.
0x2000, 0x4000: Medium Buzzy beetle shells can be carried, but not worn. Large buzzy beetle shells cannot be carried nor worn. 
0x100000: Moves left on horizontal or diagonal tracks, right if unset. Moves Down on vertical tracks (Up if Unset).
0x2000000, 0x4000000, 0x6000000: Spinies cannot cling to walls, but the same flag that works on Piranha Plants will turn the Buzzy sideways. It will not stick to anything (at least in 3.0.2)


### [29] Hidden Block 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | 2x2 Bricks  |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Winged Blocks are VERY different to normal blocks, and their behavior has been tested independently. 
0x1: No wings = empty pipe
0x20, 0x40, 0x60: Wings rotate in editor, but not while playing.
0x80: Causes block to despawn if winged. This flag is typically used on pipes, claws, clowncars, and clouds to signify "some other object is riding/inside me".
0x4000: 1x1 if winged. Objects inside (CId) can be released from all four bricks. The individual blocks act like their own blocks. 
0x8000,0x10000, 0x20000: Despawns if NOT winged. No unique behavior otherwise.
0x8000000: Despawns if not winged. seems to work in claw. 

### [30] Lakitu
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x1000: Will unstack itself if any action is taken in the editor. 
0x10000, 0x20000: Lakitu spits out spinies in both cloud types. I expected that to be the difference, but evidently that isn't it. 0x10000 is the one used in-game.

### [31] Lakitu Cloud
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |


### [32] Banzai Bill 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x2: Winged Banzai bills can be bounced on once. This may change their trajectory. Also, they have default winged goomba behavior if coming out of a block. 
0x8000: Parachuting Banzai Bills will fly to the right, but keep their orientation upon landing.

### [33] One Up 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x20, 0x40, 0x60: These look funny. If unset, faces right. Primarily used for objects in pipes. Set by default to 0x40 if not in a pipe.
0x100000: Moves left on horizontal or diagonal tracks, right if unset. Moves Down on vertical tracks (Up if Unset).

### [34] Fire Flower 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Super-Ball  |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Progressive |
| 0x80000   | Progressive?|
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x20, 0x40, 0x60: These look funny. If unset, faces right. Primarily used for objects in pipes. Set by default to 0x40 if not in a pipe.
0x100000: Moves left on horizontal or diagonal tracks, right if unset. Moves Down on vertical tracks (Up if Unset).
0x40000, 0x80000: Both for conditional progressive Power-Ups. Both are mushrooms, as far as I can tell. 0x40000 is used in-game. 

### [35] Super Star 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x20, 0x40, 0x60: These look funny. If unset, faces right. Primarily used for objects in pipes. Set by default to 0x40 if not in a pipe.
0x100000: Moves left on horizontal or diagonal tracks, right if unset. Moves Down on vertical tracks (Up if Unset). 

### [36] Lava Lift 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Unused      |
| 0x4       | Blue Lift   |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Unused      |
| 0x40      | Always      |
| 0x60      | Unused      |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x400: Will stay perfectly still until the player lands on it, even if a track is nowhere nearby. Pretty useful.
0x8000000: Is half a tile higher.
### [37] Starting Brick 
| Flag    | Description |
|---------|-------------|
| All     | Unused      |

Note: This is unsurprising, as Starting Bricks are (probably) part of the Ground tile data structure, and are furthermore never placable in the editor. 
Probably needs 0x6000000 and 0x40, or it might be corrupt.
Generates Phantom Ground tiles beneath it all the way down. 
### [38] Starting Arrow 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Unused      |
| 0x40      | Always      |
| 0x60      | Unused      |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | Unused      |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x1: Spawns out of Pipes! 
0x80: Despawns if set
0x8000000: Is half a tile higher.

### [39] Magikoopa  
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x20, 0x40, 0x60: These look funny. If unset, faces right. Primarily used for objects in pipes. Set by default to 0x40 if not in a pipe.
0x2000: Normal sized magic bullets, with normal pool of transformations. 
0x100000: Moves left on horizontal or diagonal tracks, right if unset. Moves Down on vertical tracks (Up if Unset). 
### [40] Spike Top 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Blue Shell  |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | L-Wall Hang |
| 0x4000000 | Roof Hang   |
| 0x6000000 | Ground Hang |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x20, 0x40, 0x60: If unset, faces right. Primarily used for objects in pipes. Set by default to 0x40 if not in a pipe.
0x2000: Medium Spike Tops appear to have the AI of normal Spike Tops, which means their size will prevent them from climbing walls correctly. 
0x2000000, 0x4000000, 0x6000000: Enemy can cling to walls/ceiling. Hanging on Right Wall (R-Wall) if unset (0x0000000)
0x100000: Moves left on horizontal or diagonal tracks, right if unset. Moves Down on vertical tracks (Up if Unset).
### [41] Boo  
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Boo Ring    |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Stretch     |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | L-Wall Hang |
| 0x4000000 | Roof Hang   |
| 0x6000000 | Ground Hang |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x20, 0x40, 0x60: If unset, faces right. Primarily used for objects in pipes. Set by default to 0x40 if not in a pipe.
0x100: Boo becomes a stretch, and will despawn completely if not attached to a floor or ceiling. 
0x2000: Medium        Stretches have the AI of normal Stretches, which means they will attempt to enter/exit in small spaces, which could kill them.  
0x100000: Moves left on horizontal or diagonal tracks, right if unset. Moves Down on vertical tracks (Up if Unset).
0x2000000, 0x4000000, 0x6000000: Boos cannot cling to walls, but the same flag that works on Piranha Plants will turn the Boo sideways. It will not stick to anything (at least in 3.0.2). (Combine with Stretch Flag)

### [42] Clown Car 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Fire Car    |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Has Contents|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x2: Winged Cars act like Winged Goombas, hurt Mario, and are immune to basically everything. They can jump into other Clowncars.
0x80: No visible effects, but used in-game to signify "something inside this clowncar"
0x200: Clowncars inside other Clowncars hurt Mario, despite not being winged. 
0x2000, 0x4000: Medium Clowncars are used in-game, and are 3x3. Large Clowncars are 4x4, and are used for large Bowser. 
0x10000, 0x20000: Clowncars in Clouds hurt mario, despite not being winged.
0x8000000: Half a tile higher. 

### [43] Spikes  
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Unused      |
| 0x40      | Always      |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | 2x2 Blocks  |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | Unused      |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x1, 0x2, 0x200, 0x400, 0x8000, 0x10000, 0x20000, 0x8000000: Becomes a Phantom spike block. No collision or hurtbox.
0x2000: Offset by half a tile in the editor only. 

### [44] Big Mushroom / Tanooki Leaf / Cape Feather / Propellor Mushroom / Super Bell 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Progressive |
| 0x80000   | Progressive?|
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x20, 0x40, 0x60: These look funny. If unset, faces right. Primarily used for objects in pipes. Set by default to 0x40 if not in a pipe.
0x2000: Same size as normal for Big Mushroom. Bigger for all other styles. 
0x10000, 0x20000: Likely true for all objects in clouds, but despawns if set in 3DW. Only tested on Super Bell so far though. Clowncar Works fine though. 
0x40000, 0x80000: Both for conditional progressive Power-Ups. Both are mushrooms, but the small mushroom editor sprite is missing on 0x80000. 0x40000 is used in-game. 
0x100000: Moves left on horizontal or diagonal tracks, right if unset. Moves Down on vertical tracks (Up if Unset).

### [45] Shoe Goomba / Yoshi
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Stiletto    |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x4, 0x1000: Unused for Yoshi
0x2000: Boot is medium sized while goomba is inside, but becomes small after killed. Yoshi egg is medium sized, but hatches into a normal Yoshi. 
0x4000: Red Yoshi Egg.

### [46] Dry Bones 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Dry Shell   |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x20, 0x40, 0x60: These look funny. If unset, faces right. Primarily used for objects in pipes. Set by default to 0x40 if not in a pipe.
0x2000, 0x4000: Medium Bones has normal sized properties, No SFX when bounced on. Dry Shell is medium/large sized until player enters it.  
0x100000: Moves left on horizontal or diagonal tracks, right if unset. Moves Down on vertical tracks (Up if Unset). 
### [47] Cannon 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Red Cannon  |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x400000  | Barrel 180° |
| 0x800000  | Barrel 225° |
| 0x1000000 | Barrel 90°  |
| 0x2000000 | L-Wall Hang |
| 0x4000000 | Roof Hang   |
| 0x6000000 | Ground Hang |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x1: Spawns from pipes! Spawns up to 3 Cannons from a single pipe. Large cannons do not exit pipes, but medium ones likely would.
0x20, 0x40, 0x60: If unset, faces right. Primarily used for objects in pipes. Set by default to 0x40 if not in a pipe.
0x80: Despawns
0x2000: Medium        cannons shoot normal sized bullets. 
0x400000, 0x800000, 0x1000000: Enough bitspace to map 8 unique directions the cannon can fire in. Only 5 are used: 0x1400000, 0x1800000, and 0x1C00000 are unused.
In counterclockwise degrees from the base: 0x0 = 135°, 0x400000 = 180°, 0x800000 = 225°, 0xC00000 = 270°, and 0x1000000 = 90°. 
In the editor, the unused 3 directions display incorrect directions, but do as follows after hitting play: 0x1400000 = 135°, 0x1800000 = 180°, 0x1C00000 = 225°.
0x2000000, 0x4000000, 0x6000000: Enemy can cling to walls/ceiling. Hanging on Right Wall (R-Wall) if unset (0x0000000)
0x100000: Moves left on horizontal or diagonal tracks, right if unset. Moves Down on vertical tracks (Up if Unset).

### [48] Blooper  
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Nanny       |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x20, 0x40, 0x60: These look funny. If unset, faces right. Primarily used for objects in pipes. Set by default to 0x40 if not in a pipe.
0x2000, 0x4000: Blooper Nannies always have 4 baby bloopers of the same size, regardless of the Nanny's size. 
0x100000: Moves left on horizontal or diagonal tracks, right if unset. Moves Down on vertical tracks (Up if Unset). 

### [49] Castle Bridge 
| Flag    | Description |
|---------|-------------|
| All     | Unused      |

Invisible unless in Castle theme.
Note: This is unsurprising, as Castle Bridges are (probably) part of the Ground tile data structure, and is furthermore never placable in the editor. 
Probably Always has 0x6000000 and 0x40, or it might be corrupt. Castle Flagpoles (Axes) may require this object to not be corrupt. Never tested. 

### [50] Hop-Chop
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Referred to as "Jumping Machine" in Toost. Technically true lmao. 
Tested in 3DW to prevent crashes. Clouds despawn in 3DW
0x200: Works perfectly fine, kinda. Mario will become locked in place if he attempts to bounce on any hop-chop inside a clowncar. 
0x1000: Works as expected, kinda. De-stacks if anything is touched, bouncing on it locks mario in place if not the bottom enemy in the stack.

### [51] Skipsqueak
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Spiky       |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Tested in 3DW to prevent crashes. Clouds despawn in 3DW
0x1000: Works as expected, but de-stacks if anything is touched. 

### [52] Wiggler
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Red         |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x80: Despawns wiggler
0x2000: Only the head is medium sized. The segments grow to become medium sized if bounced on. 

### [53] Horizontal Conveyor Belt
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Clockwise   |
| 0x10      | Clockwise(?)|
| 0x20      | Face Left   |
| 0x40      | Face Up     |
| 0x60      | Face Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Fast Speed  |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | On/Off      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x1, 0x200, 0x400, 0x10000, 0x20000, 0x8000000: In editor, direction arrow is missing. 
0x2, : Objects are unable to cling to the ceiling of a conveyor belt with this flag set in the editor. 
0x4, 0x8, 0x10: There may be more going on here than I understand. I suspect this is some kind of bitfield for conveyor belt orientation. 
0x20, 0x40, 0x60: These work after you hit play. Vertical Conveyor belts are possible with these flags. You would also need to edit the width/height to make collision match sprite, but neat af.
0x2000, 0x4000: Medium is only offset, Big is 6x2 for some reason. 
0x40000: This makes the conveyor belt spin at a faster speed. 
0x400000: Makes conveyor belt switch directions after an on/off switch is triggered.

### [54] Burner
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Unused      |
| 0x4       | Face Right  |
| 0x8       | Face Down   |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Parachute   |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x1: It exits pipes, but always faces down out of upward facing pipes. perhaps the pipe orientation fixes that? Needs further testing.
0x4, 0x8: Suggestive of a bitfield: 0x0 = face
0x2000: 0x4000: Editor sprite only. 
0x8000000: Up half a tile. 

### [55] Warp Door
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | P-Door      |
| 0x80000   | Key Door    |
| 0xC0000   | Unused      |
| 0x100000  | Entry #1    |
| 0x700000  | Entry #8    |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: I am not going to test these doors for enterability, because I do not want to crash my game 32 times. 
If there are many valid exit doors, you will crash upon exiting. I think there's a general door cap that causes crashes too.
0x1: Yes, doors can come out of pipes! Probably only door ID 0, but it's still something! I'd bet it functions like my crate doors, in that it's a temporary point of entry. 
0x80: Despawns the door.
0x2000, 0x40000: Editor Sprite only.
0x100000-0x700000: Doors are linked together by their flags. One door object links itself to a different door object via these flags. 
Door 0x000000 is connected to Door 0x100000 (Blue Spade)
Door 0x200000 is connected to Door 0x300000 (Purple Club)
Door 0x400000 is connected to Door 0x500000 (Yellow Diamond)
Door 0x600000 is connected to Door 0x700000 (Red Heart)
Temporary doors (such as those from pipes, crates, blocks, etc.) will spit the player out on the other end (IE: a temporary door with flag 0x0 will send the player to door 0x100000 if it exists)
0x8000000: Up half a block

### [56] Cheep Cheep 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Red         |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x80: Despawns

### [57] Muncher  
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | "Oinker"    |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x4: The game is oinking at me when I return to the editor, and this muncher is the source of that... Somehow. 
I can only assume this is an unfinished alt-form of the muncher.
0x2000: has collision of small muncher, but has the sprite of a medium muncher. 
0x8000: invisible but functional parachute. Works like piranha plants inside pipes. 

### [58] Rocky Wrench 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Monty Mole  |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x4: Appears as a Rocky Wrench in the editor, but very clearly is a Monty Mole once play is hit. Must be a leftover from MM1. 
0x80: Despawns
0x2000: a 1x1 Mole in all ways except sprite. will stick out of the wall slightly if attempting to dig, and does not fully submerge into the ground. Amusing to see. 

### [59] Track  
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Unused      |
| 0x40      | Always      |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Vertical    |
| 0x200000  | Diagonal    |
| 0x400000  | Curved      |
| 0x800000  | On/Off      |
| 0x1000000 | Diagonal????|
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | Unused      |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Tracks have their own data type, so any and all interaction with them as objects is wonky and off-putting.
Upon first load-in, they were invisible. All of them were invisible (even after hitting play), except for 0x4000, which had incorrectly drawn 6x2 track. 
As I suspected, a save + reload caused the tracks to put themselves in as their expected data-type, and suddenly none of them were invisible. 

0x4000: Visible as an object as a 2x6 track thing (Wowza!), and indistinguishable after conversion. 
0x100000: Track is vertical. 
0x200000: Track is diagonal, top left to bottom right. 
0x400000: Track is curved diagonal, top left to bottom to bottom right. 
0x800000: Track is on/off, horizontal, with the on/off end on the right side. 
0x1000000: Track ends are diagonal, but very clearly disconnected. Hard to describe in words. 

### [60] Lava Bubble 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x80: Despawns

### [61] Chain Chomp 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unchained   |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x1: Spawns as an unchained chomp. 0x5 also spawns an unchained chomp from the pipe. 
0x80: Despawns

### [62] Bowser
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Very hard to test bowser, he has a hard cap of ~16, which makes it hard to tell which bowsers actually spawned. I *think* I got it right though. 
29 total are showing up, so I'd guess 0x80, 0x8000000, and 0x1 are unaccounted for. 
There are 40 total bowsers according to the clear-condition, but there should only be 32. 
Pipe is likely adding an extra 9, and one isn't a "bowser" to the clear condition anymore.

### [63] Ice Block 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | 2x2 Blocks  |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Winged Blocks are VERY different to normal blocks, and their behavior has been tested independently. 
0x1: No wings = empty pipe
0x20, 0x40, 0x60: Wings rotate in editor, but not while playing.
0x80: Causes block to despawn if winged. This flag is typically used on pipes, claws, clowncars, and clouds to signify "some other object is riding/inside me".
0x4000: 1x1 if winged. The individual blocks act like their own blocks. 
0x8000,0x10000, 0x20000: Despawns if NOT winged. No unique behavior otherwise.
0x8000000: Despawns if not winged. seems to work in claw. 

### [64] Vine  
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Face Left   |
| 0x40      | Face Up     |
| 0x60      | Face Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | Unused      |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Tests conducted with 1x1 vines. 1x1 vines are invisible unless tiles near them are updated. 
0x1: Spawns a growing vine that stops after 2 tiles. 
0x20, 0x40, 0x60: The vine sprite remains rotated even outside the editor. With correct height/width edits, this could make horizontal vines. 
0x2000, 0x4000: Medium is a half-block shift in the editor only, Big is a 2x6 block of vines. 
In SMB1 only, the right stalk of the "Big" vine is unclimbable, and instead will warp the player to the left vine if attempted.

### [65] Stingby
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Tested in 3DW to prevent crashes
0x2: Stingbys have wings on their bodies, so they can't normally be given extra wings. Has winged goomba AI while winged, and retuns to normal if bounced on.
0x80: Despawns
0x2000: No SFX when bouncing on, no SFX while chasing. 
0x8000000: Up half a tile. 

### [66] Arrow
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | Point Dn+Rt |
| 0x800000  | Point Down  |
| 0x1000000 | Point Left  |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x1: Arrows spawn from pipes! They also retain their direction!
0x80: Despawns
0x2000, 0x4000: Sprite bigger in editor only. 
0x400000, 0x800000, 0x1000000: Bitfield for direction arrow points (0x1C00000).
0x000000 = Right, 0x400000 = Down & Right, 0x800000 = Down, 0xC00000 = Down & Left, 0x1000000 = Left, 0x1400000 = Up & Left, 0x1800000 = Up, 0x1C00000 = Up & Right
0x8000000: Up Half a tile. 

### [67] One Way
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | Point Down  |
| 0x800000  | Point Left  |
| 0x1000000 | ??????????  |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x1: One-ways can come out of pipes! Their direction is preserved too!
0x80: Despawns
0x2000, 0x4000: Bigger only in editor sprite. 
0x400000, 0x800000: Bitfield for one-way orientation. 0x000000: Point Right, 0x400000: Point Down, 0x800000: Point Left, 0xC00000: Point Up. 
0x1000000: Has no arrow on it, but can be rotated, presumably extending the above bitfield. May be ACE potential? 
Causes the editor to assign it some funky labels, like "enterable pipe", "diagonal directions", "item inside", etc.
0x8000000: Up half a tile. 

### [68] Saw  
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x1: Saws can come out of pipes!
0x80: Despawns
0x2000, 0x4000: Bigger only in editor sprite. 
0x8000000: Up half a tile. 

### [69] Player
| Flag      | Description |
|-----------|-------------|
| All       | Unused      |
Note: This is Mario/Luigi/Toad/Toadette as an object. Primarily used in the editor for stuff such as "assign power-up" and "use XY coords for spawn location".
Seen no apparent effect after hitting play in past tests.
None of them spawned in the editor. I had minor success with putting them in blocks then taking them back out, so I tried that. 
Only the right-most player will actually work, and if more than 5 players are present, the game will either crash, or teleport to (0, 0). 
Perhaps the flags actually do something? I've also tested this in multiplayer, and still nothing. 
I would imagine the player is normally stored in the ground structure, alongside the start-bricks.

### [70] Big Coin
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On Track    |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | 30 Coin     |
| 0x80000   | 50 Coin     |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x1: Spits out exactly 1 coin. 
0x80: Despawns.
0x2000, 0x4000: coin is bigger in sprite and collision, medium is 3x3, large is 4x4. 
0x8000000: half a tile higher, and a loose coin.

### [71] Semisolid Platform (3DW) 
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Unused      |
| 0x40      | Always      |
| 0x60      | Unused      |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | Unused      |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x2000, 0x4000: medium offset half a tile, big is 6x6

### [72] Koopa Car 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x1: Koopa cars will exit pipes!
0x2: Has AI of a winged goomba, immune to bouncing.
0x80: Despawns
0x2000, 0x4000: If mario enters a koopa car of magnified size, Mario's sprite increases to the size of the Koopa Car while driving. 

### [73] Toad (NSMBU Only)
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Unused      |
| 0x40      | Always      |
| 0x60      | Unused      |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | Unused      |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | Red Hat     |
| 0x800000  | Blue Hat    |
| 0x1000000 | Yellow Hat  |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Referred to as "Cinobio" in Toost.  
Toad has a global cap of 16, and only the first 16 toads from the bottom left will spawn. 
Testing these was a nightmare, and I likely missed stuff. They're also invisible in the editor.


0x2000, 0x4000: If they had editor sprites, maybe something would happen in editor.

0x400000, 0x800000, 0x1000000: Determines toad hat color: 0x1800000 and 0x1C00000 are "unused", relative to the functional toad hats. 
0x000000 = No Hat, 0x400000 = Red Hat, 0x800000 = Blue Hat, 0xC00000 = Green Hat, 0x1000000 = Yellow Hat, 0x1400000 = Purple Hat. 


### [74] Spike/Spike Ball
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Spike Ball  |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Spike is sufficiently different to Spike Ball for me to test this twice. 

0x2: Spike Balls with wings have goomba AI, and can safely be spin-jumped on. 
0x80: Despawns.
0x2000: Spike is medium sized. He will throw a spike ball that is 2.25x normal size (1.5 ^ 2), and once it leaves his hands it will be a 1.5x spike ball. 
Medium Spike Balls roll as if they were normal sized. 

### [75] Stone/Story Stone 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x1: Works as expected. 
0x80: Despawns.
0x2000, 0x4000: medium is carryable, shrinks in Mario's hands. Large is only carryable with edgeclips and such. 

### [76] Twister
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Unused      |
| 0x40      | Always      |
| 0x60      | Unused      |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x1: Functions normally
0x2000, 0x4000: Editor Sprite only. 
0x8000000: Up half a tile. 

### [77] Boom Boom / Pom Pom
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Wings       |
| 0x4       | Pom Pom     |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Hard to test because of low boom-boom cap. May be incomplete or blatantly wrong.

0x1: Could not get boom-boom to spawn, even after deleting the other 31 boom-booms. 
0x4: 3DW only, does nothing in 2DW. 
0x2000: Medium is 3x3. seems normal otherwise. Hard to test. 

### [78] Pokey  
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Most testing done with 1 high pokeys. 

0x2000, 0x4000: head is medium/large, segments are normal sized in editor but medium/large after hitting play (squished together), hitbox is normal sized. Big Pokey max height loading in. 
Snow Pokeys drop snowballs of expected sizes. While I forgot to test those for spike/spikeballs, they evidently work fine. 

### [79] P-Block
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Unused      |
| 0x4       | Filled-In   |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | 2x2 Blocks  |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x2, 0x80, 0x8000, 0x10000, 0x20000: Despawns

### [80] Dash Panel
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |


0x2000, 0x4000: Editor Sprite Only. 
0x8000000: Up half a tile. 

### [81] Smb2 Mushroom / Frog Suit / P-Balloon / Acorn Suit / Boomerang Flower
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Progressive |
| 0x80000   | Progressive?|
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x20, 0x40, 0x60: These look funny. If unset, faces right. Primarily used for objects in pipes. Set by default to 0x40 if not in a pipe.
0x2000: Same size as normal for Big Mushroom. Bigger for all other styles. 
0x10000, 0x20000: Likely true for all objects in clouds, but despawns if set in 3DW. Clowncar Works fine though. 
0x40000, 0x80000: Both for conditional progressive Power-Ups. Both are mushrooms, but the small mushroom editor sprite is missing on 0x80000. 0x40000 is used in-game. 
0x100000: Moves left on horizontal or diagonal tracks, right if unset. Moves Down on vertical tracks (Up if Unset).
### [82] Donut  
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | On Track    |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x2000, 0x4000: Editor Sprite only. 
0x8000000: Up half a tile.

### [83] Skewer  
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | On Track    |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | Face Left   |
| 0x800000  | Face Up     |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x400000, 0x800000: Bitfield for Skewer direction.
0x8000000: Up half a tile. 

### [84] Snake Block 
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Wings       |
| 0x4       | Blue Snake  |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Snake Blocks REALLY do not like to exist without snake block nodes. Lots of crashes attempting to test this. 
Attempted to add a snake block node from scratch, and that appeared to do nothing... Will be using observation first, then doing more drastic tests.
Fine, will try with 5 snake blocks at a time, with pre-existing nodes. Slow af. 

0x2: The snake block appears to have wings in the editor, no apparent difference after pressing play. The wings expand with the snake block length though. 
0x20, 0x40, 0x60: offsets snake blocks on grid after pressing play. this can screw with snake block nodes. 
0x80: Despawns

0x2000, 0x4000: Medium sprite bigger in editor only, but the large is glitchy af (editor only), with a 2x5 sprite with some green mesa-like spikes hanging off of it. 
With nodes, medium is unaffected, but tail end of large is confused for a little bit. Hard to explain in text. 
0x8000000: Half a tile higher. 

### [85] Track Block 
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Unused      |
| 0x4       | Blue Block  |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Track Blocks REALLY do not like to exist without track block nodes. Lots of crashes attempting to test this. Going to just do 5 at a time, with nodes. 

0x2000, 0x4000: Both sprites bigger in editor only.
0x8000000: Half a tile up. 

### [86] Charvaargh  
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | Unused      |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x2000, 0x4000: Editor Sprite only. 

### [87] Slight Slope 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Face Left   |
| 0x40      | Face Up     |
| 0x60      | Face Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | Unused      |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: 1x1 slopes all look identical. Hard to test, and hard to reverse test due to slopes saving as the ground data type.
0x20, 0x40, 0x60: messes with slope collision & sprites in glitchy ways. 
0x2000, 0x4000: Medium offset half a tile, Big is 4x6.

### [88] Steep Slope 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Face Left   |
| 0x40      | Face Up     |
| 0x60      | Face Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | Unused      |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: 1x1 slopes all look identical. Hard to test, and hard to reverse test due to slopes saving as the ground data type.
0x20, 0x40, 0x60: messes with slope collision & sprites in glitchy ways. 
0x2000, 0x4000: Medium offset half a tile, Big is 4x5.

### [89] Custom Autoscroll Bird
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Unused      |
| 0x40      | Always      |
| 0x60      | Unused      |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | Unused      |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Medium Auto |
| 0x80000   | Fast Auto   |
| 0x100000  | Node #1     |
| 0x200000  | Node #2     |
| 0x400000  | Node #4     |
| 0x800000  | Node #8     |
| 0x1000000 | Node #10(?) |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Unused      |
| 0x8000000 | Unused      |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Called "Reel Camera" in Toost. Does not spawn unless custom auto-scroll is turned on.
0x40000, 0x80000: Speed of the autoscroll node. 
0x100000, 0x200000, 0x400000, 0x800000: Bitfield for autoscroll node counts. I suspect it holds up to 15 nodes, and it's lying to me about 0x1000000 being node #10. 

### [90] Checkpoint Flag
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | Point Down  |
| 0x800000  | Point Left  |
| 0x1000000 | ??????????  |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x1: CPs spawn from pipes, yes. 
0x80: Despawns
0x2000, 0x4000: Bigger only in editor sprite.
0x400000, 0x800000: Bitfield for one-way orientation. 0x000000: Point Right, 0x400000: Point Down, 0x800000: Point Left, 0xC00000: Point Up. 
0x1000000: Has no arrow on it, but can be rotated, presumably extending the above bitfield. May be ACE potential? 
Causes the editor to assign it some funky labels, like "enterable pipe", "diagonal directions", "item inside", etc.
0x8000000: Up half a tile. 
### [91] See-Saw
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Unused      |
| 0x40      | Always      |
| 0x60      | Unused      |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | On Tracks   |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x2000, 0x4000: Medium shifted half a tile, big is just a 1x10. 


### [92] Red Coin 
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Unused      |
| 0x40      | Always      |
| 0x60      | Unused      |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | 2x2 Blocks  |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | Unused      |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |


### [93] Clear Pipe
| Flag      | Description |
|-----------|-------------|
| All       | Unused      |

Note: This one was a shocker for me. Clear pipes do not work without nodes, and thus did not spawn without them.
Even with nodes though, none of my 32 clear pipes were anything short of functional. No large sprites, no disappearing pipes, no half-tile offsets, nothing. 
That said, the interesting stuff that can be done with clear pipes is inside the nodes themselves, but big clear pipes do not exist. 

### [94] Diagonal Conveyor Belt 
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Clockwise   |
| 0x10      | Clockwise(?)|
| 0x20      | Face Left   |
| 0x40      | Face Up     |
| 0x60      | Face Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Fast Speed  |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | On/Off      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x1, 0x200, 0x400, 0x10000, 0x20000, 0x8000000: In editor, direction arrow is missing. 
0x2, : Objects are unable to cling to the ceiling of a conveyor belt with this flag set in the editor. 
0x4, 0x8, 0x10: There may be more going on here than I understand. I suspect this is some kind of bitfield for conveyor belt orientation. 
0x20, 0x40, 0x60: These work after you hit play. Vertical Conveyor belts are possible with these flags. Funkier with diagonal. 
You would also need to edit the width/height to make collision match sprite, but neat af.
0x2000, 0x4000: Medium is only offset, Big is 6x2 for some reason. 
0x40000: This makes the conveyor belt spin at a faster speed. 
0x400000: Makes conveyor belt switch directions after an on/off switch is triggered.

### [95] Key  
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Unused      |
| 0x4       | Phanto Key  |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x4: SMB1 only. 
0x80: Despawns.
0x2000, 0x4000: Editor sprite only. 

### [96] Ant Trooper 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Spiky       |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachuting |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | L-Wall Hang |
| 0x4000000 | Roof Hang   |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |
0x1000: Works as expected. 
Note: Crashes outside 3DW

### [97] Warp Box
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Unused      |
| 0x4       | Key Box     |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Entry #1    |
| 0x700000  | Entry #8    |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: I am not going to test these boxes for enterability, because I do not want to crash my game 32 times. 
If there are many valid exit boxes, you will crash upon exiting. I think there's a general box cap that causes crashes too.
0x1: Yes, boxes can come out of pipes! Probably only box ID 0, but it's still something! I'd bet it functions like my crate boxes, in that it's a temporary point of entry. 
0x80: Despawns the box.
0x2000, 0x40000: Editor Sprite only.
0x100000-0x700000: Boxes are linked together by their flags. One door object links itself to a different door object via these flags. 
Box 0x000000 is connected to Box 0x100000 (Blue Spade)
Box 0x200000 is connected to Box 0x300000 (Purple Club)
Box 0x400000 is connected to Box 0x500000 (Yellow Diamond)
Box 0x600000 is connected to Box 0x700000 (Red Heart)
Temporary Boxes (such as those from pipes, crates, blocks, etc.) will spit the player out on the other end (IE: a temporary Box with flag 0x0 will send the player to Box 0x100000 if it exists)
0x8000000: Up half a block

### [98] Bowser Jr 
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachuting |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Hard to test for same reason as bowser. Gah
0x80: Despawns

### [99] On Off Block
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | 2x2 Blocks  |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | Unused      |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x4000: Looks like 1x1 in editor, expands after hitting play. 

### [100] Dotted Line Block
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Unused      |
| 0x4       | Red         |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | 2x2 Blocks  |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | Unused      |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |


### [101] Water Marker 
| Flag      | Description |
|-----------|-------------|
| All       | Unused      |

Note: Never seen this object play nice with me before. No crashes or anything, but I suspect it's either completely inert, or converted to ground data instantly. 
All of the info it handles takes place in the map section of the level file anyway, in theory. 

### [102] Monty Mole 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachuting |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x4: Interesting that it doesn't become a rocky wrench. Monty Mole must have been added late in development as a unique enemy. 

### [103] Fish Bone 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachuting |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x80: Despawns
0x2000: No SFX when breaking. 

### [104] Angry Sun 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Unused      |
| 0x4       | Moon        |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x80: Despawns
0x2000, 0x4000: Editor Sprite only. 
0x8000000: Half a tile higher. 

### [105] Swinging Claw 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | "Pipe Left" |
| 0x40      | "Pipe Up"   |
| 0x60      | "Pipe Down" |
| 0x80      | Has Contents|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x1: Spawns from pipe, seems normal. 
0x20, 0x40, 0x60: Edit claw height, but does not spin claw in editor. Unfortunate. 
0x80: No effect on claw, used to tell the game "this claw is holding something". 
0x4000: Editor is very confused, sprite is just a normal claw, but hitbox suggests an 8x4 claw. is normal when pressing play though. 
0x8000000: Half a tile higher. Amusing that it works. 
### [106] Tree  
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Face Left   |
| 0x40      | Face Up     |
| 0x60      | Face Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Trees were 1x1 at the beginning of the test, expanding each back to a reasonable size after examining everything. 
The trunk of the tree seems to be 
0x1: Only capable of spawning a 1 block tall tree.
0x20, 0x40, 0x60: Trunk remains facing this direction after pushing play. The leaves return to where they should be, and the tree has collision where a normal tree would.
0x80: Leaves despawn, Trunk sprite still present, unable to climb. Seems useful for decoration. 
0x2000, 0x4000: In editor leaves grow bigger and trunk grows taller + leaves offset from trunk. 
Once play is hit, medium is normal sized, large is normal but max height. 
0x8000000: Up half a tile. 
### [107] Piranha Creeper 
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Unused      |
| 0x4       | Blue Creeper|
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Face Left   |
| 0x40      | Face Up     |
| 0x60      | Face Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | L-Wall Hang |
| 0x4000000 | Roof Hang   |
| 0x6000000 | Ground Hang |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: As expected, despawn without nodes. Touching them crashes the game. Will retry 10 at a time, with proper nodes. 
0x20, 0x40, 0x60: Spin creeper in editor, offset them after hitting play, which causes minor discontinuities in the node path. probably happens with offset creepers too. 
0x80: Despawns.
0x2000, 0x4000: Editor sprite only. 
0x10000, 0x20000: Spawn, despite being "in clouds". Rare for 3DW, so they must have no cloud functionality. 
0x2000000, 0x4000000: Determines Piranha creeper orientation, and affects pathing. 
If in opposite direction to first node, despawns creeper. If perpendicular to first node, causes major node discontinuity.
0x8000000: Up half a tile, causes minor node discontinuity. 

### [108] Blinking Block 
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Unused      |
| 0x4       | Blue Block  |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Unused      |
| 0x40      | Always      |
| 0x60      | Unused      |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | 2x2 Block   |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | Unused      |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x80: Despawns. 

### [109] Sound Effect
| Flag      | Description |
|-----------|-------------|
| All       | Unused      |

Note: Never seen this object play nice with me before. No crashes or anything, but I suspect it's either completely inert, or converted to ground data instantly. 
Sound Effects have their own data structure, and I cannot get these to convert. If you want to add 0x100000000 unique custom SFX though, this would be how. 
Could even be 0x10000000000000000 if you used object flag data. Would likely necessitate LID structure to make it happen though. 

### [110] Spike Block 
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Unused      |
| 0x4       | Blue Block  |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Unused      |
| 0x40      | Always      |
| 0x60      | Unused      |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | 'Big'       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Blue        |
| 0x80000   | Yellow      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | Unused      |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: 1x1 Spikes used for this test. Spikes that would despawn (such as 0x1, 0x2, 0x80, 0x10000, and 0x20000) will become phantom spikes, with only collision. 

0x4000: Sprite is the same, but the hitbox is 4x4. 

### [111] Mechakoopa  
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachuting |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Blasta      |
| 0x80000   | Zappa       |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x80: Despawns.

### [112] Crate  
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x1: Works fine, contents not stored. 
0x2: Wings invisible, works like any other flying carryable. CAN crush mario. 
0x80: Despawns.
0x2000, 0x4000: Medium carryable, becomes small until released, has small collision. Large carryable via clipping, same deal. 
0x8000000: Up half a tile. 

### [113] Mushroom Trampoline 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Blue        |
| 0x8       | Move Up     |
| 0x10      | Move Right  |
| 0x20      | Unused      |
| 0x40      | Unused      |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | "Big"       |
| 0x8000    | Parachute   |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x2: Acts like a lift with wings. 
0x8, 0x10: Bitfield: 0x0 = Move Left, 0x8 = Move Up, 0x10 = Move Right, 0x18 = Move Down. 
Choosing to use these bits instead of 0x20-0x60 tells me that mushroom trampolines may have been intended to go in pipes earlier in development. 
0x80: Despawns. 
0x4000: 1x8 when pressing play. 
0x8000000: half a tile higher. 

### [114] Porkupuffer  
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x8000000: Half a tile higher. 

### [115] Toadette Cage  
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Unused      |
| 0x40      | Always      |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Referred to as "Cinobic" in Toost. Despawns mario if 4 or more are outside containers. Sometimes(?) crashes at exactly 2, and Cap seems to be at 5.
If mario dies while despawned, no death animation or SFX plays, and it's a silent fade to black. 
Annoying, since I will need to test in batches of 3 to see if they do something noteworthy after obtaining the flag (which can't happen if mario despawns).
Wasted effort: None of them do anything new. I swear I've seen a larger toadette before, but evidently I'm mistaken. 
Even tested it in ? blocks to see if spawning from a block made a difference, but nope! 
0x1: Spawns from pipes, regardless of limit
0x80: Despawns.
0x4000: Does not seem to do anything, but the yellow outline from multi-select says it's 2x2. 
0x8000000: Higher by half a tile. 

### [116] Super Hammer
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Unused      |
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Progressive |
| 0x80000   | Progressive?|
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Despawns outside of 3DW. would theoretically work inside a CP in new Soup, but cannot manually put it inside in the editor. 
0x20, 0x40, 0x60: These look funny. If unset, faces right. Primarily used for objects in pipes. Set by default to 0x40 if not in a pipe. 
0x10000, 0x20000: Likely true for all objects in clouds, but despawns if set in 3DW. Only tested on Super Bell so far though. Clowncar Works fine though. 
0x40000, 0x80000: Both for conditional progressive Power-Ups. Both are mushrooms, but the small mushroom editor sprite is missing on 0x80000. 0x40000 is used in-game. 
0x100000: Moves left on horizontal or diagonal tracks, right if unset. Moves Down on vertical tracks (Up if Unset).

### [117] Bully  
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x80: Despawns.
0x2000: No SFX when landed on. 
0x8000000: Half a tile higher. 

### [118] Icicle  
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Unused      |
| 0x4       | Dark Blue   |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Unused      |
| 0x40      | Always      |
| 0x60      | Unused      |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | "Big"       |
| 0x8000    | Parachute   |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Icicles have their own unique data type. 
Any attempt to save/reload icicles of any kind convert them to this data type, removing most interesting properties. Would require coursebot upload to function. 

0x1, 0x200, 0x400, 0x10000, 0x20000, 0x8000000: Icicle is blue colored until tampered with.
0x1: Spawns from pipe, can be either light or dark blue. 
0x80: Despawns
0x4000: 2x2 in editor, 2x1 when pressing play. 
0x8000000: up half a tile. 

### [119] Exclamation Block 
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Unused      |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Face Left   |
| 0x40      | Face Up     |
| 0x60      | Face Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | "Medium"    |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x400000  | Unused      |
| 0x800000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: As expected, despawn without nodes. Touching them crashes the game. Will retry 10 at a time, with proper nodes. 
0x20, 0x40, 0x60: Spin blocks in editor, offset them after hitting play, which causes minor discontinuities in the node path. probably happens with offset blocks too. 
0x80: Despawns.
0x2000, 0x4000: Editor sprite only. 
0x10000, 0x20000: Spawn, despite being "in clouds". Rare for 3DW, so they must have no cloud functionality. 
0x8000000: Up half a tile, causes minor node discontinuity. 

### [120] Lemmy  
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x80: Despawns
0x2000: Medium Lemmy rides on a normal ball, and dispenses normal balls. 
0x8000000: Half a tile higher. 

### [121] Morton  
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x80: Despawns
0x2000: Morton's fire and magic are normal sized. 
0x8000000: Half a tile higher. 

### [122] Larry  
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x80: Despawns
0x2000: Larry's magic is normal sized. 
0x8000000: Half a tile higher. 

### [123] Wendy  
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x80: Despawns
0x2000: Wendy's rings are normal sized. 
0x8000000: Half a tile higher. 

### [124] Iggy  
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x80: Despawns
0x2000: Iggy's magic is normal sized. 
0x8000000: Half a tile higher. 

### [125] Roy  
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x80: Despawns
0x2000: Roy's magic is normal sized. Can spawn anywhere a normal sized Roy can after digging. 
0x8000000: Half a tile higher. 

### [126] Ludwig  
| Flag      | Description |
|-----------|-------------|
| 0x1       | Unused      |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | In Stack    |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | In Cloud    |
| 0x20000   | In Cloud(?) |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x80: Despawns
0x2000: Ludwig's magic is normal sized. 
0x8000000: Half a tile higher. 

### [127] Cannon Box 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x80: Despawns
0x8000000: Half a tile higher

### [128] Propeller Box 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x80: Despawns
0x8000000: Half a tile higher

### [129] Goomba Mask 
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x80: Despawns
0x8000000: Half a tile higher

### [130] Bullet Bill Mask
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x80: Despawns
0x8000000: Half a tile higher

### [131] Red Pow Box
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Wings       |
| 0x4       | Unused      |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Pipe Left   |
| 0x40      | Pipe Up     |
| 0x60      | Pipe Down   |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | In Clowncar |
| 0x400     | On a Track  |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Medium      |
| 0x4000    | Big         |
| 0x8000    | Parachute   |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Left Tracks |
| 0x200000  | Vert Tracks |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

0x80: Despawns
0x8000000: Half a tile higher

### [132] On/Off Trampoline
| Flag      | Description |
|-----------|-------------|
| 0x1       | In Pipe     |
| 0x2       | Unused      |
| 0x4       | Off         |
| 0x8       | Unused      |
| 0x10      | Unused      |
| 0x20      | Unused      |
| 0x40      | Always      |
| 0x80      | Is Container|
| 0x100     | Unused      |
| 0x200     | Unused      |
| 0x400     | Unused      |
| 0x800     | Unused      |
| 0x1000    | Unused      |
| 0x2000    | Unused      |
| 0x4000    | "Big"       |
| 0x8000    | Unused      |
| 0x10000   | Unused      |
| 0x20000   | Unused      |
| 0x40000   | Unused      |
| 0x80000   | Unused      |
| 0x100000  | Unused      |
| 0x200000  | Unused      |
| 0x800000  | Unused      |
| 0x400000  | Unused      |
| 0x1000000 | Unused      |
| 0x2000000 | Unused      |
| 0x4000000 | Unused      |
| 0x6000000 | Always      |
| 0x8000000 | In Claw     |
| 0x10000000| Unused      |
| 0x20000000| Unused      |
| 0x40000000| Unused      |
| 0x80000000| Unused      |

Note: Tested in 3DW, on 1x1 on/off trampolines. 
0x80: Despawns.
0x4000: 1x8, nothing else of note.
0x8000000: half a tile higher. 
