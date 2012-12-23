package models

const (
	TileEvents_IdtileEvents string = "tile_events.idtile_events"
	TileEvents_Eventype     string = "tile_events.eventtype"
	TileEvents_Param1       string = "tile_events.param1"
	TileEvents_Param2       string = "tile_events.param2"
	TileEvents_Param3       string = "tile_events.param3"
	TileEvents_Param4       string = "tile_events.param4"
	TileEvents_Param5       string = "tile_events.param5"
)

type TileEvents struct {
	IdtileEvents int `PK`
	Eventtype     int
	Param1       string
	Param2       string
	Param3       string
	Param4       string
	Param5       string
}
