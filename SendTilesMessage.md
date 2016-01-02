# Introduction #
SendTilesMessage is used by the server to send all viewable tiles to the client.
# Client #

## Read ##
-

# Server #

## Send ##
Packet setup
| **Value** | **Type** | **Comment** |
|:----------|:---------|:------------|
| Header    | uint8    | HEADER\_TILES |
| Total tiles | uint16   |             |
| **- Foreach tile -** |
| Position X | uint16   |             |
| Position Y | uint16   |             |
| Blocking  | uint32   |             |
| **-- Foreach tile layer --** |
| Layer     | uint16   |             |
| Sprite    | uint32   |             |
| **-- End --** |
| Location id | uint16   |             |
| Location name | string   |             |
| **- End -**|