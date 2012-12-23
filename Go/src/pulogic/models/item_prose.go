package models

const (
	ItemProse_ItemId          string = "item_prose.item_id"
	ItemProse_LocalLanguageId string = "item_prose.local_language_id"
	ItemProse_ShortEffect     string = "item_prose.short_effect"
	ItemProse_Effect          string = "item_prose.effect"
)

type ItemProse struct {
	ItemId          int `PK`
	LocalLanguageId int `PK`
	ShortEffect     string
	Effect          string
}
