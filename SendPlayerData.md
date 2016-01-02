# Introduction #
SendPlayerData sends player info like name, position, money and outfit

# Server #

## Send ##
Packet build:
| **Value** | **Type** | **Comment** |
|:----------|:---------|:------------|
| Header    | uint8    | HEADER\_IDENTITY |
| UID       | uint64   | Game Unique ID |
| Name      | string   |             |
| Position X | uint16   |             |
| Position Y | uint16   |             |
| Direction | uint16   |             |
| Money     | uint32   | For the very rich |
| Outfit Upper | uint8    | Style       |
| Outfit Upper | uint32   | Colour      |
| Outfit Nek | uint8    | Style       |
| Outfit Nek | uint32   | Colour      |
| Outfit Head | uint8    | Style       |
| Outfit Head | uint32   | Colour      |
| Outfit Feet | uint8    | Style       |
| Outfit Feet | uint32   | Colour      |
| Outfit Lower | uint8    | Style       |
| Outfit Lower | uint32   | Colour      |