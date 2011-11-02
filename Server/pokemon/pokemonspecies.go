package main

type PokemonSpecies struct {
	SpeciesId				int
	Identfier				int
	EvolvesFromSpeciesId	int
	EvolutionChain			*EvolutionChain
	ColorId					int
	ShapeId					int
	HabitatId				int
	GenderRate				int
	CaptureRate				int
	BaseHappiness			int
	IsBaby					int
	HatchCounter			int
	HasGenderDifferences	int
	GrowthRateId			int
	FormsSwitchable			int
}

func NewPokemonSpecies() *PokemonSpecies {
	return &PokemonSpecies{}
}