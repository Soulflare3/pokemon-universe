package models

const (
	Tile_Idtile      string = "tile.idtile"
	Tile_X           string = "tile.x"
	Tile_Y           string = "tile.y"
	Tile_Z           string = "tile.z"
	Tile_Idlocation  string = "tile.idlocation"
	Tile_Movement    string = "tile.movement"
	Tile_Script      string = "tile.script"
	Tile_IdtileEvent string = "tile.idtile_event"
)

type Tile struct {
	Idtile      int `PK`
	X           int
	Y           int
	Z           int
	Idlocation  int
	Movement    int
	Script      string
	IdtileEvent int
}
