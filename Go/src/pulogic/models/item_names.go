package models

const (
	ItemNames_ItemId          string = "item_names.item_id"
	ItemNames_LocalLanguageId string = "item_names.local_language_id"
	ItemNames_Name            string = "item_names.name"
)

type ItemNames struct {
	ItemId          int `PK`
	LocalLanguageId int `PK`
	Name            string
}
