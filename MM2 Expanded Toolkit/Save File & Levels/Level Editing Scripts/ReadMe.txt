Credit to Kramer for setting me up with the script template and much more!
Credit to The Great Rambler for a lot of things
Credit to the Open Course World Dev team for many more things (AAAAAAAA)



Set Up The Scripts
1. Install golang, the language all of this is coded in: https://go.dev/doc/install
2. Install some sort of text/code viewer/editor to modify the scripts as needed
	Note: I use Notepad++: https://notepad-plus-plus.org/downloads/
3. Unless you have some other way to run golang code (like an IDE or something), open up Terminal
4. Navigate to the "MM2 Expanded Toolkit\Save File & Levels\Level Editing Scripts" Folder
	Note: "ls" can be used to see what files and folders are in the current directory, and "cd ___" will move you to the specified folder
5. Create a go mod folder using the commands "go mod init learn" then "go mod tidy"
	Note: The scripts will look for level files wherever you create the go.mod, NOT where the script files are!
6. Repeat steps 4 and 5 for the "MM2 Expanded Toolkit\Save File & Levels\Level Reading Scripts" Folder
	Note: You may be able to get away with only making a single go.mod if you put it in "Save File & Levels".
	Note2: Might even be wise to put this entire folder inside your Ryujinx save file, and place the go.mod so it edits your save data directly... (you will need to do this for BOTH 0 and 1)
7. You're all set up! 
	To view a script, simply open it with Notepad++ (I've listed where the important lines are in each script at the top)
	To run a script, simply type "go run [Drag script here]" in terminal
	Once satisfied with your level file, paste it into your Ryujinx save data where you obtained it from. Paste in BOTH 0 AND 1!!!
	
Other Notes: I plan to add to the editor scripts as I go, but if you need something badly, and don't feel capable of making it yourself, give me a ping, and I'll see what I can do.


Documentation: See https://github.com/TheGreatRambler/toost/blob/main/level.ksy for an overview of useful stuff

Other stuff I've learned: Flags are u32, and appear to be a bitfield in some capacity. 
0x06040040 is right-facing (at least for a one-way), while 0x06440040, 0x06840040, and 0x06C40040 are for the other 3 directions. 
Toads use the same system: 0x06040040, 0x06440040, 0x06840040, 0x06C40040, 0x07040040, 0x07440040 are for each of the 6 forms
Object size is also affected by this field. 0x06000040 is a default state for objects, but 0x06004040 makes them large.
