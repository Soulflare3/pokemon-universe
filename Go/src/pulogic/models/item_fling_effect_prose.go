package models

const (
	ItemFlingEffectProse_ItemFlingEffectId string = "item_fling_effect_prose.item_fling_effect_id"
	ItemFlingEffectProse_LocalLanguageId   string = "item_fling_effect_prose.local_language_id"
	ItemFlingEffectProse_Effect            string = "item_fling_effect_prose.effect"
)

type ItemFlingEffectProse struct {
	ItemFlingEffectId int `PK`
	LocalLanguageId   int `PK`
	Effect            string
}
