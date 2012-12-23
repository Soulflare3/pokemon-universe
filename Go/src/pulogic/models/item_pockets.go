package models

const (
	ItemPockets_Id         string = "item_pockets.id"
	ItemPockets_Identifier string = "item_pockets.identifier"
)

type ItemPockets struct {
	Id         int `PK`
	Identifier string
}
