package models

const (
	Abilities_Id           string = "abilities.id"
	Abilities_Identifier   string = "abilities.identifier"
	Abilities_GenerationId string = "abilities.generation_id"
)

type Abilities struct {
	Id           int `PK`
	Identifier   string
	GenerationId int
}
