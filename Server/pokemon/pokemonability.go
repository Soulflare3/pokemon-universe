package main

type PokemonAbility struct {
	Ability			*Ability
	IsDream			int
	Slot			int
}

func NewPokemonAbility() *PokemonAbility {
	return &PokemonAbility{}
}