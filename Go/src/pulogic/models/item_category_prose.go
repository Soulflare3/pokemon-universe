package models

const (
	ItemCategoryProse_ItemCategoryId  string = "item_category_prose.item_category_id"
	ItemCategoryProse_LocalLanguageId string = "item_category_prose.local_language_id"
	ItemCategoryProse_Name            string = "item_category_prose.name"
)

type ItemCategoryProse struct {
	ItemCategoryId  int `PK`
	LocalLanguageId int `PK`
	Name            string
}
