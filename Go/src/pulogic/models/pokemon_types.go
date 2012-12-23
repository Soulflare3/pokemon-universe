package models

const (
	PokemonTypes_PokemonId string = "pokemon_types.pokemon_id"
	PokemonTypes_TypeId    string = "pokemon_types.type_id"
	PokemonTypes_Slot      string = "pokemon_types.slot"
)

type PokemonTypes struct {
	PokemonId int `PK`
	TypeId    int
	Slot      int `PK`
}
