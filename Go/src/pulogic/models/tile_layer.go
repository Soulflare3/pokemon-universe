package models

const (
	TileLayer_IdtileLayer string = "tile_layer.idtilelayer"
	TileLayer_Idtile      string = "tile_layer.idtile"
	TileLayer_Sprite      string = "tile_layer.sprite"
	TileLayer_Layer       string = "tile_layer.layer"
)

type TileLayer struct {
	IdtileLayer int `PK`
	Idtile      int `PK`
	Sprite      int
	Layer       int
}
