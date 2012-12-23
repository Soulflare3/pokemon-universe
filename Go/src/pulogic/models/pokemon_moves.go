package models

const (
	PokemonMoves_PokemonId           string = "pokemon_moves.pokemon_id"
	PokemonMoves_VersionGroupId      string = "pokemon_moves.version_group_id"
	PokemonMoves_MoveId              string = "pokemon_moves.move_id"
	PokemonMoves_PokemonMoveMethodId string = "pokemon_moves.pokemon_move_method_id"
	PokemonMoves_Level               string = "pokemon_moves.level"
	PokemonMoves_Order               string = "pokemon_moves.order"
)

type PokemonMoves struct {
	PokemonId           int `PK`
	VersionGroupId      int `PK`
	MoveId              int `PK`
	PokemonMoveMethodId int `PK`
	Level               int `PK`
	//Order               int
}
