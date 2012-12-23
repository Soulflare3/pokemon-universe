package models

const (
	MapchangeLayer_IdmapchangeLayer string = "mapchange_layer.idmapchange_layer"
	MapchangeLayer_IdmapchangeTile  string = "mapchange_layer.idmapchange_tile"
	MapchangeLayer_Index            string = "mapchange_layer.index"
	MapchangeLayer_Sprite           string = "mapchange_layer.sprite"
)

type MapchangeLayer struct {
	IdmapchangeLayer int `PK`
	IdmapchangeTile  int
	Index            int
	Sprite           int
}
