package main

type EvolutionChain struct {
	EvolvedSpeciesId		int
	EvolutionTriggerId		int
	TriggerItemId			int
	MinimumLevel			int
	Gender					int
	LocationId				int
	HeldItemId				int
	TimeOfDay				int
	KnownMoveId				int
	MinimumHappiness		int
	MinimumBeauty			int
	RelativePhysicalStats	int
	PartySpeciesId			int
	TradeSpeciesId			int
}

func NewEvolutionChain() *EvolutionChain {
	return &EvolutionChain{}
}