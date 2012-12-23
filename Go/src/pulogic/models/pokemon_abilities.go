package models

const (
	PokemonAbilities_PokemonId string = "pokemon_abilities.pokemon_id"
	PokemonAbilities_AbilityId string = "pokemon_abilities.ability_id"
	PokemonAbilities_IsDream   string = "pokemon_abilities.is_dream"
	PokemonAbilities_Slot      string = "pokemon_abilities.slot"
)

type PokemonAbilities struct {
	PokemonId int `PK`
	AbilityId int
	IsDream   bool
	Slot      int `PK`
}
