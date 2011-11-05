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
package main

import (
	"os"
	"mysql"
)

type PokemonList 			map[int]*Pokemon
type PokemonSpeciesList 	map[int]*PokemonSpecies
type MoveList 				map[int]*Move
type AbilityList 			map[int]*Ability
type PokemonStatArray 		[]*PokemonStat
type PokemonTypeArray		[]int
type PokemonAbilityList 	map[int]*PokemonAbility
type PokemonMoveList 		map[int]*PokemonMove

var g_PokemonManager *PokemonManager = NewPokemonManager()

type PokemonManager struct {
	pokemon				PokemonList // All pokemon including different forms
	pokemonSpecies		PokemonSpeciesList // All pokemon species
	moves				MoveList
	abilities			AbilityList
}

func NewPokemonManager() *PokemonManager {
	return &PokemonManager { pokemon: make(PokemonList),
							 pokemonSpecies: make(PokemonSpeciesList),
							 moves: make(MoveList),
							 abilities: make(AbilityList) }
}

func (m *PokemonManager) Load() {
	// Load moves
	if !m.loadMoves() {
		return
	}
	
	// Load abilities
	if !m.loadAbilities() {
		return
	}
	
	// Load PokemonSpecies
	if m.loadPokemonSpecies() {
		return
	}
	
	// Load Pokemon
	if m.loadPokemon() {
		return
	}
}

func (m *PokemonManager) loadMoves() bool {
	var err os.Error
	var query string = "SELECT id, identifier, type_id, power, accuracy, priority, target_id, damage_class_id," +
		" effect_id, effect_chance, contest_type_id, contest_effect_id, super_contest_effect_id" +
		" FROM moves"
		
	if err = g_db.Query(query); err != nil {
		g_logger.Printf("[ERROR] SQL error while executing query pokemon moves: %s\n\r", err)
		return false
	}
	
	var result *mysql.Result
	result, err = g_db.UseResult()
	if err != nil {
		g_logger.Println("[ERROR] SQL error while fetching result pokemon moves: %s\n\r", err)
		return false
	}
	
	defer result.Free()
	g_logger.Println(" - Processing moves")
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		move := NewMove()
		move.MoveId = row[0].(int)
		move.Identifier = row[1].(string)
		move.TypeId = row[2].(int)
		move.Power = row[3].(int)
		move.Accuracy = row[4].(int)
		move.Priority = row[5].(int)
		move.TargetId = row[6].(int)
		move.DamageClassId = row[7].(int)
		move.EffectId = row[8].(int)
		move.EffectChance = row[9].(int)
		move.ContestType = row[10].(int)
		move.ContestEffect = row[11].(int)
		move.SuperContestEffect = row[12].(int)
		
		// Add to map
		m.moves[move.MoveId] = move
	}
	
	return true
}

func (m *PokemonManager) loadAbilities() bool {
	var err os.Error
	var query string = "SELECT id, identifier FROM abilities"
	
	if err = g_db.Query(query); err != nil {
		g_logger.Println("[ERROR] SQL error while executing query pokemon abilities: %s\n\r", err)
		return false
	}
	
	var result *mysql.Result
	result, err = g_db.UseResult()
	if err != nil {
		g_logger.Println("[ERROR] SQL error while fetching result pokemon abilities: %s\n\r", err)
		return false
	}
	
	defer result.Free()
	g_logger.Println(" - Processing abilities")
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		ability := NewAbility()
		ability.AbilityId = row[0].(int)
		ability.Identifier = row[1].(string)
		
		// Add to map
		m.abilities[ability.AbilityId] = ability
	}
	
	return true
}

func (m *PokemonManager) loadPokemonSpecies() bool {
	// Select all pokemon including their evolution parameters
	var err os.Error
	var query string = "SELECT `ps`.id, `ps`.identifier, `ps`.evolves_from_species_id, `ps`.color_id, `ps`.shape_id, `ps`.habitat_id, `ps`.gender_rate," +
		" `ps`.capture_rate, `ps`.base_happiness, `ps`.is_baby, `ps`.hatch_counter, `ps`.has_gender_differences, `ps`.growth_rate_id, `ps`.forms_switchable," +
		" `pe`.evolved_species_id, `pe`.evolution_trigger_id, `pe`.trigger_item_id, `pe`.minimum_level, `pe`.gender, `pe`.location_id, `pe`.held_item_id, `pe`.time_of_day," + 
		" `pe`.known_move_id, `pe`.minimum_happiness, `pe`.minimum_beauty, `pe`.relative_physical_stats, `pe`.party_species_id, `pe`.trade_species_id" +
		" FROM pokemon_species AS `ps` LEFT JOIN pokemon_evolution AS `pe`" +
		" ON `pe`.evolved_species_id = (SELECT `ps2`.id FROM pokemon_species AS `ps2`" +
											" WHERE `ps2`.evolves_from_species_id = `ps`.id LIMIT 1)"
	
	if err = g_db.Query(query); err != nil {
		g_logger.Println("[ERROR] SQL error while executing query pokemon species: %s\n\r", err)
		return false
	}
	
	var result *mysql.Result
	result, err = g_db.UseResult()
	if err != nil {
		g_logger.Println("[ERROR] SQL error while fetching result pokemon species: %s\n\r", err)
		return false
	}
	
	defer result.Free()
	g_logger.Println(" - Processing pokemon species")
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		// Create EvolutionChain object
		evoChain := NewEvolutionChain()
		evoChain.EvolvedSpeciesId = row[14].(int)
		evoChain.EvolutionTriggerId = row[15].(int)
		evoChain.TriggerItemId = row[16].(int) // TODO: Change this to the actual Item objet
		evoChain.MinimumLevel = row[17].(int)
		evoChain.Gender = row[18].(int)
		evoChain.LocationId = row[19].(int)
		evoChain.HeldItemId = row[20].(int)
		evoChain.TimeOfDay = row[21].(int)
		evoChain.KnownMoveId = row[22].(int) // TODO: Change to move object
		evoChain.MinimumHappiness = row[23].(int)
		evoChain.MinimumBeauty = row[24].(int)
		evoChain.RelativePhysicalStats = row[25].(int)
		evoChain.PartySpeciesId = row[26].(int)
		evoChain.TradeSpeciesId = row[27].(int)
		
		// Creat PokemonSpecies object
		pokemon := NewPokemonSpecies()
		pokemon.SpeciesId = row[0].(int)
		pokemon.Identifier = row[1].(string)
		pokemon.EvolvesFromSpeciesId = row[2].(int)
		pokemon.EvolutionChain = evoChain
		pokemon.ColorId = row[3].(int)
		pokemon.ShapeId = row[4].(int)
		pokemon.HabitatId = row[5].(int)
		pokemon.GenderRate = row[6].(int)
		pokemon.CaptureRate = row[7].(int)
		pokemon.BaseHappiness = row[8].(int)
		pokemon.IsBaby = row[9].(int)
		pokemon.HatchCounter = row[10].(int)
		pokemon.HasGenderDifferences = row[11].(int)
		pokemon.GrowthRateId = row[12].(int)
		pokemon.FormsSwitchable = row[12].(int)
		
		// Add to map
		m.pokemonSpecies[pokemon.SpeciesId] = pokemon
	}
	
	return true
}

func (m *PokemonManager) loadPokemon() bool {
	var query string = "SELECT id, species_id, height, weight, base_experience, order, is_default FROM pokemon"
	result, err := DBQuerySelect(query)
	if err != nil {
		return false
	}
	
	defer result.Free()
	g_logger.Println(" - Processing pokemon")
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		pokemon := NewPokemon()
		pokemon.PokemonId = row[0].(int)
		pokemon.Species = m.GetPokemonSpecies(row[1].(int))
		pokemon.Height = row[2].(int)
		pokemon.Weight = row[3].(int)
		pokemon.BaseExperience = row[4].(int)
		pokemon.Order = row[5].(int)
		pokemon.IsDefault = row[6].(int)
		
		// Load base stats for this pokemon
		pokemon.loadStats()
		
		// Fetch available abilities for this pokemon
		pokemon.loadAbilities()
		
		// Fetch available forms for this pokemon
		pokemon.loadForms()
		
		// Fetch learnable moves for this pokemon
		pokemon.loadMoves()
		
		// Fetch pokemon types
		pokemon.loadTypes()
		
		// Add to map
		m.pokemon[pokemon.PokemonId] = pokemon
	}
	
	return true
}

func (m *PokemonManager) GetPokemon(_pokemonId int) *Pokemon {
	return m.pokemon[_pokemonId]
}

func (m *PokemonManager) GetPokemonSpecies(_speciesId int) *PokemonSpecies {
	return m.pokemonSpecies[_speciesId]
}

// TODO: Do something with the _formId variable, now we just get the name by only _speciesId
func (m *PokemonManager) GetPokemonName(_speciesId, _formId int) string {
	var name string = ""
	species := m.GetPokemonSpecies(_speciesId)
	if species != nil {
		name = species.Identifier
	}
	return name
}

// TODO: Do something with the _FormId variable
func (m *PokemonManager) GetPokemonTypes(_pokemonId, _formId int) PokemonTypeArray {
	pokemon := m.GetPokemon(_pokemonId)
	return pokemon.Types
}

func (m *PokemonManager) GetAbilityById(_abilityId int) *Ability {
	return m.abilities[_abilityId]
}

func (m *PokemonManager) GetMoveById(_moveId int) *Move {
	return m.moves[_moveId]
}