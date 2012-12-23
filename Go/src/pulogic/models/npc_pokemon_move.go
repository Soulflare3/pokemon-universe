package models

const (
	NpcPokemonMove_IdnpcPokemonMove string = "npc_pokemon_move.idnpc_pokemon_move"
	NpcPokemonMove_IdnpcPokemon     string = "npc_pokemon_move.idnpc_pokemon"
	NpcPokemonMove_Idmove           string = "npc_pokemon_move.idmove"
)

type NpcPokemonMove struct {
	IdnpcPokemonMove int `PK`
	IdnpcPokemon     int
	Idmove           int
}
