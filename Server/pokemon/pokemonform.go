package main

type PokemonForm struct {
	Id				int
	Identifier		string
	IsDefault		int
	IsBattleOnly	int
	Order			int
}

func NewPokemonForm() *PokemonForm {
	return &PokemonForm{}
}