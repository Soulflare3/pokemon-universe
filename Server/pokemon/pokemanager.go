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
	"fmt"
	"strings"
)

type PokemonList 			map[int]*Pokemon
type PokemonSpeciesList 	map[int]*PokemonSpecies
type MoveList 				map[int]*Move
type AbilityList 			map[int]*Ability
type PokemonStatArray 		[]*PokemonStat
type PokemonTypeArray		[]int
type PokemonAbilityList 	map[int]*PokemonAbility
type PokemonMoveList 		map[int]*PokemonMove
type MessageList			map[int]map[int]string

type PokemonManager struct {
	pokemon				PokemonList // All pokemon including different forms
	pokemonSpecies		PokemonSpeciesList // All pokemon species
	moves				MoveList
	abilities			AbilityList
	moveMessages		MessageList
	abilityMessages		MessageList
}

func NewPokemonManager() *PokemonManager {
	return &PokemonManager { pokemon: make(PokemonList),
							 pokemonSpecies: make(PokemonSpeciesList),
							 moves: make(MoveList),
							 abilities: make(AbilityList),
							 moveMessages: make(MessageList),
							 abilityMessages: make(MessageList) }
}

func (m *PokemonManager) Load() bool {
	// Load moves
	if !m.loadMoves() {
		return false
	}
	
	// Load abilities
	if !m.loadAbilities() {
		return false
	}
	
	// Load PokemonSpecies
	if !m.loadPokemonSpecies() {
		return false
	}
	
	// Load Pokemon
	if !m.loadPokemon() {
		return false
	}
	
	// Load move messages
	if !m.loadMoveMessages() {
		return false
	}
	
	if !m.loadAbilityMessages() {
		return false
	}
	
	return true
}

func (m *PokemonManager) loadMoves() bool {
	var query string = "SELECT id, identifier, type_id, power, accuracy, priority, target_id, damage_class_id," +
		" effect_id, effect_chance, contest_type_id, contest_effect_id, super_contest_effect_id, pp, flavor_text" +
		" FROM moves" +
		" RIGHT JOIN move_flavor_text ON move_id  = id" + 
		" WHERE version_group_id = 11"
	result, err := DBQuerySelect(query)
	if err != nil {
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
		move.MoveId = int(row[0].(int64))
		move.Identifier = DBGetString(row[1])
		move.TypeId = DBGetInt(row[2])
		move.Power = DBGetInt(row[3])
		move.Accuracy = DBGetInt(row[4])
		move.Priority = DBGetInt(row[5])
		move.TargetId = DBGetInt(row[6])
		move.DamageClassId = DBGetInt(row[7])
		move.EffectId = DBGetInt(row[8])
		move.EffectChance = DBGetInt(row[9])
		move.ContestType = DBGetInt(row[10])
		move.ContestEffect = DBGetInt(row[11])
		move.SuperContestEffect = DBGetInt(row[12])
		move.PP = DBGetInt(row[13])
		move.FlavorText = DBGetString(row[14])
		
		// Add to map
		m.moves[move.MoveId] = move
	}
	
	return true
}

func (m *PokemonManager) loadAbilities() bool {
	var query string = "SELECT id, identifier FROM abilities"
	result, err := DBQuerySelect(query)
	if err != nil {
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
		ability.AbilityId = DBGetInt(row[0])
		ability.Identifier = DBGetString(row[1])
		
		// Add to map
		m.abilities[ability.AbilityId] = ability
	}
	
	return true
}

func (m *PokemonManager) loadPokemonSpecies() bool {
	// Select all pokemon including their evolution parameters
	var query string = "SELECT `ps`.id, `ps`.identifier, `ps`.evolves_from_species_id, `ps`.color_id, `ps`.shape_id, `ps`.habitat_id, `ps`.gender_rate," +
		" `ps`.capture_rate, `ps`.base_happiness, `ps`.is_baby, `ps`.hatch_counter, `ps`.has_gender_differences, `ps`.growth_rate_id, `ps`.forms_switchable," +
		" `pe`.evolved_species_id, `pe`.evolution_trigger_id, `pe`.trigger_item_id, `pe`.minimum_level, `pe`.gender, `pe`.location_id, `pe`.held_item_id, `pe`.time_of_day," + 
		" `pe`.known_move_id, `pe`.minimum_happiness, `pe`.minimum_beauty, `pe`.relative_physical_stats, `pe`.party_species_id, `pe`.trade_species_id" +
		" FROM pokemon_species AS `ps` LEFT JOIN pokemon_evolution AS `pe`" +
		" ON `pe`.evolved_species_id = (SELECT `ps2`.id FROM pokemon_species AS `ps2`" +
											" WHERE `ps2`.evolves_from_species_id = `ps`.id LIMIT 1)"
	result, err := DBQuerySelect(query)
	if err != nil {
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
		evoChain.EvolvedSpeciesId = DBGetInt(row[14])
		evoChain.EvolutionTriggerId = DBGetInt(row[15])
		evoChain.TriggerItemId = DBGetInt(row[16]) // TODO: Change this to the actual Item objet
		evoChain.MinimumLevel = DBGetInt(row[17])
		evoChain.Gender = DBGetString(row[18])
		evoChain.LocationId = DBGetInt(row[19])
		evoChain.HeldItemId = DBGetInt(row[20])
		evoChain.TimeOfDay = DBGetString(row[21])
		evoChain.KnownMoveId = DBGetInt(row[22]) // TODO: Change to move object
		evoChain.MinimumHappiness = DBGetInt(row[23])
		evoChain.MinimumBeauty = DBGetInt(row[24])
		evoChain.RelativePhysicalStats = DBGetInt(row[25])
		evoChain.PartySpeciesId = DBGetInt(row[26])
		evoChain.TradeSpeciesId = DBGetInt(row[27])
		
		// Creat PokemonSpecies object
		pokemon := NewPokemonSpecies()
		pokemon.SpeciesId = DBGetInt(row[0])
		pokemon.Identifier = DBGetString(row[1])
		pokemon.EvolvesFromSpeciesId = DBGetInt(row[2])
		pokemon.EvolutionChain = evoChain
		pokemon.ColorId = DBGetInt(row[3])
		pokemon.ShapeId = DBGetInt(row[4])
		pokemon.HabitatId = DBGetInt(row[5])
		pokemon.GenderRate = DBGetInt(row[6])
		pokemon.CaptureRate = DBGetInt(row[7])
		pokemon.BaseHappiness = DBGetInt(row[8])
		pokemon.IsBaby = DBGetInt(row[9])
		pokemon.HatchCounter = DBGetInt(row[10])
		pokemon.HasGenderDifferences = DBGetInt(row[11])
		pokemon.GrowthRateId = DBGetInt(row[12])
		pokemon.FormsSwitchable = DBGetInt(row[12])
		
		// Add to map
		m.pokemonSpecies[pokemon.SpeciesId] = pokemon
	}
	
	return true
}

func (m *PokemonManager) loadPokemon() bool {
	var query string = "SELECT `id`, `species_id`, `height`, `weight`, `base_experience`, `order`, `is_default` FROM pokemon"
	result, err := DBQuerySelect(query)
	if err != nil {
		return false
	}
	
	g_logger.Println(" - Processing pokemon")
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		pokemon := NewPokemon()
		pokemon.PokemonId = DBGetInt(row[0])
		pokemon.Species = m.GetPokemonSpecies(DBGetInt(row[1]))
		pokemon.Height = DBGetInt(row[2])
		pokemon.Weight = DBGetInt(row[3])
		pokemon.BaseExperience = DBGetInt(row[4])
		pokemon.Order = DBGetInt(row[5])
		pokemon.IsDefault = DBGetInt(row[6])
				
		// Add to map
		m.pokemon[pokemon.PokemonId] = pokemon
	}
	result.Free()
	
	// We load the actual pokemon data after we loaded the pokemon
	for _, pokemon := range(m.pokemon) {
		// Load base stats for this pokemon
		if !pokemon.loadStats() {
			return false
		}
		
		// Fetch available abilities for this pokemon
		if !pokemon.loadAbilities() {
			return false
		}
		
		// Fetch available forms for this pokemon
		if !pokemon.loadForms() {
			return false
		}
		
		// Fetch learnable moves for this pokemon
		if !pokemon.loadMoves() {
			return false
		}
		
		// Fetch pokemon types
		if !pokemon.loadTypes() {
			return false
		}
	}
	
	return true
}

func (m *PokemonManager) loadMoveMessages() bool {
	var query string = "SELECT `move_effect_id`, `message` FROM move_messages"
	result, err := DBQuerySelect(query)
	if err != nil {
		return false
	}
	
	g_logger.Println(" - Processing Move Messages")
	defer result.Free()
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		effect_id := DBGetInt(row[0])
		message := DBGetString(row[1])
		
		messages := strings.Split(message, "|")
		messageMap := make(map[int]string)
		for index, msg := range(messages) {
			messageMap[index] = msg
		}
		
		m.moveMessages[effect_id] = messageMap
	}
	
	return true
}

func (m *PokemonManager) loadAbilityMessages() bool {
	var query string = "SELECT `ability_id`, `message` FROM ability_messages"
	result, err := DBQuerySelect(query)
	if err != nil {
		return false
	}
	
	g_logger.Println(" - Processing Ability Messages")
	defer result.Free()
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		ability_id := DBGetInt(row[0])
		message := DBGetString(row[1])
		
		messages := strings.Split(message, "|")
		messageMap := make(map[int]string)
		for index, msg := range(messages) {
			messageMap[index] = msg
		}
		
		m.moveMessages[ability_id] = messageMap
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
	if pokemon == nil {
		fmt.Printf("PokemonManager::GetPokemonTypes - Could not find pokemon with id: %d\n\r", _pokemonId)
	}
	return pokemon.Types
}

func (m *PokemonManager) GetAbilityById(_abilityId int) *Ability {
	return m.abilities[_abilityId]
}

func (m *PokemonManager) GetAbilityNameById(_abilityId int) (toReturn string) {
	ability := m.GetAbilityById(_abilityId)
	if ability != nil {
		toReturn = ability.Identifier
	} else {
		toReturn = "Unknown"
	}
	return
}

func (m *PokemonManager) GetMoveById(_moveId int) *Move {
	return m.moves[_moveId]
}

func (m *PokemonManager) GetMoveNameById(_moveId int) (toReturn string) {
	move := m.GetMoveById(_moveId)
	if move != nil {
		toReturn = move.Identifier
	} else {
		toReturn = "Unknown"
	}
	return
}

func (m *PokemonManager) GetMoveMessage(_id, _part int) string {
	var toReturn string = ""
	if value, found := m.moveMessages[_id]; found {
		if _part < len(value) {
			if partValue, ok := value[_part]; ok {
				toReturn = partValue
			}
		}
	}
	
	return toReturn
}

func (m *PokemonManager) GetAbilityMessage(_id, _part int) string {
	var toReturn string = ""
	if value, found := m.abilityMessages[_id]; found {
		if _part < len(value) {
			if partValue, ok := value[_part]; ok {
				toReturn = partValue
			}
		}
	}
	
	return toReturn
}

func (m *PokemonManager) GetItemNameById(_id int) string {
	// TODO: We don't have items yet, so return unknown
	return "Unknown"
}