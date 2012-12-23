package models

const (
	PokemonStats_PokemonId string = "pokemon_stats.pokemon_id"
	PokemonStats_StatId    string = "pokemon_stats.stat_id"
	PokemonStats_BaseStat  string = "pokemon_stats.base_stat"
	PokemonStats_Effort    string = "pokemon_stats.effort"
)

type PokemonStats struct {
	PokemonId int `PK`
	StatId    int `PK`
	BaseStat  int
	Effort    int
}
