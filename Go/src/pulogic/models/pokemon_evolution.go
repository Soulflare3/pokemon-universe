package models

const (
	PokemonEvolution_Id                    string = "pokemon_evolution.id"
	PokemonEvolution_EvolvedSpeciesId      string = "pokemon_evolution.evolved_species_id"
	PokemonEvolution_EvolutionTriggerId    string = "pokemon_evolution.evolution_trigger_id"
	PokemonEvolution_TriggerItemId         string = "pokemon_evolution.trigger_item_id"
	PokemonEvolution_MinimumLevel          string = "pokemon_evolution.minimum_level"
	PokemonEvolution_Gender                string = "pokemon_evolution.gender"
	PokemonEvolution_LocationId            string = "pokemon_evolution.location_id"
	PokemonEvolution_HeldItemId            string = "pokemon_evolution.held_item_id"
	PokemonEvolution_TimeOfDay             string = "pokemon_evolution.time_of_day"
	PokemonEvolution_KnownMoveId           string = "pokemon_evolution.known_move_id"
	PokemonEvolution_MinimumHappiness      string = "pokemon_evolution.minimum_happiness"
	PokemonEvolution_MinimumBeauty         string = "pokemon_evolution.minimum_beauty"
	PokemonEvolution_RelativePhysicalStats string = "pokemon_evolution.relative_physical_stats"
	PokemonEvolution_PartySpeciesId        string = "pokemon_evolution.party_species_id"
	PokemonEvolution_TradeSpeciesId        string = "pokemon_evolution.trade_species_id"
)

type PokemonEvolution struct {
	IdpokemonEvolution    int `PK`
	EvolvedSpeciesId      int
	EvolutionTriggerId    int
	TriggerItemId         int
	MinimumLevel          int
	Gender                string
	LocationId            int
	HeldItemId            int
	TimeOfDay             string
	KnownMoveId           int
	MinimumHappiness      int
	MinimumBeauty         int
	RelativePhysicalStats int
	PartySpeciesId        int
	TradeSpeciesId        int
}
