package models

const (
	ItemFlags_Id         string = "item_flags.id"
	ItemFlags_Identifier string = "item_flags.identifier"
)

type ItemFlags struct {
	Id         int `PK`
	Identifier string
}
