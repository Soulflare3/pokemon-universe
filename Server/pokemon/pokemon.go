package main

type Pokemon struct {
	PokemonId				int
	Species					*PokemonSpecies
	Height					int
	Weight					int
	BaseExperince			int
	Order					int
	IsDefault				bool
	
	Stats					[]*PokemonStat // 6
	
	Abilities				map[int]*PokemonAbility
	Forms					map[int]*PokemonForm
}

func NewPokemon() *Pokemon {
	return &Pokemon{ Stats: make([]*PokemonStat, 6),
					 Abilities: make(map[int]*PokemonAbility),
					 Forms: make(map[int]*PokemonForm) }
}