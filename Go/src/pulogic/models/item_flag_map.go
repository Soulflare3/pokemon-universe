package models

const (
	ItemFlagMap_ItemId     string = "item_flag_map.item_id"
	ItemFlagMap_ItemFlagId string = "item_flag_map.item_flag_id"
)

type ItemFlagMap struct {
	ItemId     int `PK`
	ItemFlagId int `PK`
}
