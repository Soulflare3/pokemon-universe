package models

const (
	ItemCategories_Id         string = "item_categories.id"
	ItemCategories_PocketId   string = "item_categories.pocket_id"
	ItemCategories_Identifier string = "item_categories.identifier"
)

type ItemCategories struct {
	Id         int `PK`
	PocketId   int
	Identifier string
}
