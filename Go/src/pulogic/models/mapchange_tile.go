package models

const (
	MapchangeTile_IdmapchangeTile string = "mapchange_tile.idmapchange_tile"
	MapchangeTile_Idmapchange     string = "mapchange_tile.idmapchange"
	MapchangeTile_X               string = "mapchange_tile.x"
	MapchangeTile_Y               string = "mapchange_tile.y"
	MapchangeTile_Z               string = "mapchange_tile.z"
	MapchangeTile_Movement        string = "mapchange_tile.movement"
)

type MapchangeTile struct {
	IdmapchangeTile int `PK`
	Idmapchange     int
	X               int
	Y               int
	Z               int
	Movement        int
}
