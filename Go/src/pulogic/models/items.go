package models

const (
	Items_Id            string = "items.id"
	Items_Identifier    string = "items.identifier"
	Items_CategoryId    string = "items.category_id"
	Items_Cost          string = "items.cost"
	Items_FlingPower    string = "items.fling_power"
	Items_FlingEffectId string = "items.fling_effect_id"
)

type Items struct {
	Id            int `PK`
	Identifier    string
	CategoryId    int
	Cost          int
	FlingPower    int
	FlingEffectId int
}
