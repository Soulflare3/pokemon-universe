package models

const (
	ItemFlagProse_ItemFlagId      string = "item_flag_prose.item_flag_id"
	ItemFlagProse_LocalLanguageId string = "item_flag_prose.local_language_id"
	ItemFlagProse_Name            string = "item_flag_prose.name"
	ItemFlagProse_Description     string = "item_flag_prose.description"
)

type ItemFlagProse struct {
	ItemFlagId      int `PK`
	LocalLanguageId int `PK`
	Name            string
	Description     string
}
