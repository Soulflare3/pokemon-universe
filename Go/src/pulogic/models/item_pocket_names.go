package models

const (
	ItemPocketNames_ItemPocketId    string = "item_pocket_names.item_pocket_id"
	ItemPocketNames_LocalLanguageId string = "item_pocket_names.local_language_id"
	ItemPocketNames_Name            string = "item_pocket_names.name"
)

type ItemPocketNames struct {
	ItemPocketId    int `PK`
	LocalLanguageId int `PK`
	Name            string
}
