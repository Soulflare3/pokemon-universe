package main

const (
	POKESTAT_HP		int = iota
	POKESTAT_ATTACK
	POKESTAT_DEFENSE
	POKESTAT_SPECIALATTACK
	POKESTAT_SPECIALDEFENCE
	POKESTAT_SPEED
	POKESTAT_ACCURACY
	POKESTAT_EVASION
)

type PokemonStat struct {
	StatType	int
	BaseStat	int
	Effort		int
}

func NewPokemonStat() *PokemonStat {
	return &PokemonStat{}
}