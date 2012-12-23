package models

const (
	Types_Id            string = "types.id"
	Types_Identifier    string = "types.identifier"
	Types_GenerationId  string = "types.generation_id"
	Types_DamageClassId string = "types.damage_class_id"
)

type Types struct {
	Id            int `PK`
	Identifier    string
	GenerationId  int
	DamageClassId int
}
