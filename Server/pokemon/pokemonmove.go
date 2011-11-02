package main

const (
	MOVEMETHOD_LEVELUP		int = 1
	MOVEMETHOD_EGG				= 2
	MOVEMETHOD_TUTOR		    = 3
	MOVEMETHOD_MACHINE			= 4
	MOVEMETHOD_LIGHTBALLEGG		= 6
	MOVEMETHOD_FORMCHANGE		= 10
)

type PokemonMove struct {
	Pokemon				*Pokemon
	VersionGroup		int
	Move				*Move
	PokemonMoveMethod	int
	Level				int
	Order				int
}