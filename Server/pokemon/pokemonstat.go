package main

const (
	_ 					= iota // Truncate first value
	POKESTAT_HP		int = 1
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