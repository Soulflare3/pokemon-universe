package models

const (
	ItemFlavorText_ItemId         string = "item_flavor_text.item_id"
	ItemFlavorText_VersionGroupId string = "item_flavor_text.version_group_id"
	ItemFlavorText_LanguageId     string = "item_flavor_text.language_id"
	ItemFlavorText_FlavorText     string = "item_flavor_text.flavor_text"
)

type ItemFlavorText struct {
	ItemId         int `PK`
	VersionGroupId int `PK`
	LanguageId     int `PK`
	FlavorText     string
}
