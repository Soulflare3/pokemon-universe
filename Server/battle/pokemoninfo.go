package main

var g_PokemonInfo *PokemonInfo = NewPokemonInfo()

type PokemonInfo struct {
	Names	map[*PokemonUniqueId]string
}

func NewPokemonInfo() *PokemonInfo {
	return &PokemonInfo{ Names: make(map[*PokemonUniqueId]string) }
}