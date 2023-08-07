## Contains data on ground tiles, Icicles, Tracks, and Sound Effects
"Snake-likes" (Snake blocks, Track Blocks, Clear Pipes, Piranha Creepers, ! Blocks, and Track Blocks) will get their own write-up. 
### Ground Tile Structure
| Field     | Description | Size |
|-----------|----------------------|------|
| Ground Id | Applied sequentially to each ground tile placed | 4000 | 
| X-Position| The X coordinate of the ground tile, in blocks | 0xFF |
| Y-Position| The Y coordinate of the ground tile, in blocks | 0xFF |
|Id| Used to indicate how the ground tile connects to other tiles, and if an internal doodad is present | 0xFF | 
|Background_Id| Used for external doodads, and seemingly nothing else. | 0xFF | 

Note: For all background IDs, it seems like the game only uses 0x0, 0x4, 0x8, 0xC, and 0x10 naturally. 
Manually assigning a background Id between those values functions fine
|Theme| 0x00-0x03 | 0x04-0x07 | 0x08-0x0B | 0x0C-0x0F | 0x10-0x13 | 0x14-0xFF |
|-|-|-|-|-|-|-|
| SMB1 Ground Day| No Doodad | 1x1 Light Green Bush| 2x1 Light Green Tree| 3x1 Light Green Tree| 1x3 Light Green Bush| Unused |
| SMB1 Ground Night| No Doodad | 1x1 Bush | 2x1 Dark Green Tree| 3x1 Dark Green Tree| 1x3 Dark Green Bush| Unused |
| SMB1 Sky Day| No Doodad | 1x1 White Flower | 2x1 Light Green Tree| 3x1 Light Green Tree| 1x3 Light Brown Fence| Unused |
| SMB1 Sky Night| No Doodad | 1x1 Blue Flower | 2x1 Dark Green Tree| 3x1 Dark Green Tree| 1x3 Dark Brown Fence| Unused |
| SMB1 Underground Day/Night| No Doodad | 1x1 Blue Mushroom| 2x1 Blue Vine| 3x1 Blue Vine| 1x3 Blue Skeleton | Unused |
| SMB1 Forest Day| No Doodad | 1x1 Brown Mushroom| 2x1 Dark Green Tree| 3x1 Dark Green Tree| 1x3 Dark Green Bush| Unused |
| SMB1 Forest Night| No Doodad | 1x1 Dark Brown Mushroom| 2x1 Spotted Beanstalk| 3x1 Spotted Beanstalk| 1x3 Dark Brown Skeleton| Unused |
| SMB1 Forest Day| No Doodad | 1x1 Brown Mushroom| 2x1 Dark Green Tree| 3x1 Dark Green Tree| 1x3 Dark Green Bush| Unused |
| SMB1 Underwater Day| No Doodad | 1x1 Blue Rock Pile| 2x1 Blue Rock Pile| 3x1 Blue Rock Pile| 1x3 Blue Rock Pile| Unused |
| SMB1 Underwater Night| No Doodad | 1x1 Blue Rock Pile| 2x1 Blue Rock Pile| 3x1 Blue Rock Pile| 1x3 Blue Rock Pile| Unused |
| SMB1 Ghost House Day| No Doodad | 1x1 Blue Mushroom| 2x1 Blue Lit Lantern| 3x1 Blue Grandfather Clock| 1x3 Blue Fence| Unused |
| SMB1 Ghost House Night| No Doodad | 1x1 Blue Mushroom| 2x1 Blue Unlit Lantern| 3x1 Blue Grandfather Clock| 1x3 Blue Fence| Unused |
| SMB1 Desert Day| No Doodad | 1x1 Yellow Plant| 2x1 Yellow Plant| 3x1 Yellow Plant| 1x3 Yellow Skeleton| Unused |
| SMB1 Desert Night| No Doodad | 1x1 Brown Plant| 2x1 Brown Plant| 3x1 Brown Plant| 1x3 Brown Skeleton| Unused |
| SMB1 Airship Day| No Doodad | 1x1 Light Gray Bolts x2| 2x1 Light Gray Pipes x2 | 3x1 Light Gray Flagpole w/h Black Flag & Rope| 1x3 Light Gray Fence w/h Black Rope| Unused |
| SMB1 Airship Night| No Doodad | 1x1 Dark Gray Bolts x2| 2x1 Dark Gray Pipes x2 | 3x1 Dark Gray Flagpole w/h Red Flag & Rope | 1x3 Dark Gray Fence w/h Red Rope| Unused |
| SMB1 Snow Day| No Doodad | 1x1 White Bush | 2x1 White Tree | 3x1 White Tree | 1x3 White Bush | Unused |
| SMB1 Snow Night| No Doodad | 1x1 Blue Bush | 2x1 Blue Tree | 3x1 Blue Tree | 1x3 Blue Bush | Unused |
| SMB1 Castle Day/Night| No Doodad | 1x1 Gray Pillar | 2x1 Gray Bowser Jr. Statue| 3x1 Gray Bowser Statue| 1x3 Gray Pillar Fence| Unused |
| SMB3 Ground Day/Night| No Doodad | 1x1 White Flower| 2x1 Spotted Bush| 3x1 Spotted Bush| 1x3 Spotted Bush| Unused |
| SMB3 Sky Day| No Doodad | 1x1 White Cloud Bulb| 2x1 White Cloud Bulbs| 3x1 White Cloud Bulbs| 1x3 White Cloud Bulbs x5| Unused |
| SMB3 Sky Night| No Doodad | 1x1 Black Cloud Bulb| 2x1 Black Cloud Bulbs| 3x1 Black Cloud Bulbs| 1x3 Black Cloud Bulbs x5| Unused |
To be continued
