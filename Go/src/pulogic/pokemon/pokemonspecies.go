/*Pokemon Universe MMORPG
Copyright (C) 2010 the Pokemon Universe Authors

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.*/
package pokemon

import "pulogic/models"

type PokemonSpecies struct {
	SpeciesId				int
	Identifier				string
	EvolvesFromSpeciesId	int
	EvolutionChain			*EvolutionChain
	ColorId					int
	ShapeId					int
	HabitatId				int
	GenderRate				int
	CaptureRate				int
	BaseHappiness			int
	IsBaby					bool
	HatchCounter			int
	HasGenderDifferences	bool
	GrowthRateId			int
	FormsSwitchable			bool
}

func NewPokemonSpecies() *PokemonSpecies {
	return &PokemonSpecies{}
}

func NewPokemonSpecesFromEntity(_entity models.PokemonSpeciesJoinPokemonEvolution) *PokemonSpecies {
	// Create EvolutionChain object
	evoChain := NewEvolutionChain()
	evoChain.EvolvedSpeciesId = _entity.IdpokemonEvolution
	evoChain.EvolutionTriggerId = _entity.EvolutionTriggerId
	evoChain.TriggerItemId = _entity.TriggerItemId // TODO: Change this to the actual Item objet
	evoChain.MinimumLevel = _entity.MinimumLevel
	evoChain.Gender = _entity.Gender
	evoChain.LocationId = _entity.LocationId
	evoChain.HeldItemId = _entity.HeldItemId
	evoChain.TimeOfDay = _entity.TimeOfDay
	evoChain.KnownMoveId = _entity.KnownMoveId // TODO: Change to move object
	evoChain.MinimumHappiness = _entity.MinimumHappiness
	evoChain.MinimumBeauty = _entity.MinimumBeauty
	evoChain.RelativePhysicalStats = _entity.RelativePhysicalStats
	evoChain.PartySpeciesId = _entity.PartySpeciesId
	evoChain.TradeSpeciesId = _entity.TradeSpeciesId
	
	// Creat PokemonSpecies object
	pokemon := NewPokemonSpecies()
	pokemon.SpeciesId = _entity.Id
	pokemon.Identifier = _entity.Identifier
	pokemon.EvolvesFromSpeciesId = _entity.EvolvesFromSpeciesId
	pokemon.EvolutionChain = evoChain
	pokemon.ColorId = _entity.ColorId
	pokemon.ShapeId = _entity.ShapeId
	pokemon.HabitatId = _entity.HabitatId
	pokemon.GenderRate = _entity.GenderRate
	pokemon.CaptureRate = _entity.CaptureRate
	pokemon.BaseHappiness = _entity.BaseHappiness
	pokemon.IsBaby = _entity.IsBaby
	pokemon.HatchCounter = _entity.HatchCounter
	pokemon.HasGenderDifferences = _entity.HasGenderDifferences
	pokemon.GrowthRateId = _entity.GrowthRateId
	pokemon.FormsSwitchable = _entity.FormsSwitchable	
	
	return pokemon
}