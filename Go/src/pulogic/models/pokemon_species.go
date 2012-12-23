package models

const (
	PokemonSpecies_Id                   string = "pokemon_species.id"
	PokemonSpecies_Identifier           string = "pokemon_species.identifier"
	PokemonSpecies_GenerationId         string = "pokemon_species.generation_id"
	PokemonSpecies_EvolvesFromSpeciesId string = "pokemon_species.evolves_from_species_id"
	PokemonSpecies_EvolutionChainId     string = "pokemon_species.evolution_chain_id"
	PokemonSpecies_GenderRate           string = "pokemon_species.gender_rate"
	PokemonSpecies_CaptureRate          string = "pokemon_species.capture_rate"
	PokemonSpecies_BaseHappiness        string = "pokemon_species.base_happiness"
	PokemonSpecies_IsBaby               string = "pokemon_species.is_baby"
	PokemonSpecies_HatchCounter         string = "pokemon_species.hatch_counter"
	PokemonSpecies_HasGenderDifferences string = "pokemon_species.has_gender_differences"
	PokemonSpecies_GrowthRateId         string = "pokemon_species.growth_rate_id"
	PokemonSpecies_FormsSwitchable      string = "pokemon_species.forms_switchable"
	PokemonSpecies_ColorId              string = "pokemon_species.color_id"
	PokemonSpecies_ShapeId              string = "pokemon_species.shape_id"
	PokemonSpecies_HabitatId            string = "pokemon_species.habitat_id"
)

type PokemonSpecies struct {
	Id                   int `PK`
	Identifier           string
	GenerationId         int
	EvolvesFromSpeciesId int
	EvolutionChainId     int
	GenderRate           int
	CaptureRate          int
	BaseHappiness        int
	IsBaby               bool
	HatchCounter         int
	HasGenderDifferences bool
	GrowthRateId         int
	FormsSwitchable      bool
	ColorId              int
	ShapeId              int
	HabitatId            int
}

type PokemonSpeciesJoinPokemonEvolution struct {
	Id                   int `PK`
	Identifier           string
	GenerationId         int
	EvolvesFromSpeciesId int
	EvolutionChainId     int
	GenderRate           int
	CaptureRate          int
	BaseHappiness        int
	IsBaby               bool
	HatchCounter         int
	HasGenderDifferences bool
	GrowthRateId         int
	FormsSwitchable      bool
	ColorId              int
	ShapeId              int
	HabitatId            int
	
	IdpokemonEvolution	  int 
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