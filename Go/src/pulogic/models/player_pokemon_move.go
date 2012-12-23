package models

const (
	PlayerPokemonMove_IdplayerPokemonMove string = "player_pokemon_move.idplayer_pokemon_move"
	PlayerPokemonMove_IdplayerPokemon     string = "player_pokemon_move.idplayer_pokemon"
	PlayerPokemonMove_Idmove              string = "player_pokemon_move.idmove"
	PlayerPokemonMove_PpUsed              string = "player_pokemon_move.pp_used"
)

type PlayerPokemonMove struct {
	IdplayerPokemonMove int `PK`
	IdplayerPokemon     int
	Idmove              int
	PpUsed              int
}
