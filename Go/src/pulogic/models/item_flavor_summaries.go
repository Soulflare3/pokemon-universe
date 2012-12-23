package models

const (
	ItemFlavorSummaries_ItemId          string = "item_flavor_summaries.item_id"
	ItemFlavorSummaries_LocalLanguageId string = "item_flavor_summaries.local_language_id"
	ItemFlavorSummaries_FlavorSummary   string = "item_flavor_summaries.flavor_summary"
)

type ItemFlavorSummaries struct {
	ItemId          int `PK`
	LocalLanguageId int `PK`
	FlavorSummary   string
}
