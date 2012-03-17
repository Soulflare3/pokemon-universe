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

import (
	"strings"
	
	puh "puhelper"
	log "putools/log"
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

var manager *PokemonManager

func NewPokemonManager() *PokemonManager {
	if manager == nil {
		manager =  &PokemonManager { pokemon: make(PokemonList),
									pokemonSpecies: make(PokemonSpeciesList),
									moves: make(MoveList),
									abilities: make(AbilityList),
									moveMessages: make(MessageList),
									abilityMessages: make(MessageList) }
	}
	
	return manager
}

func GetInstance() *PokemonManager {
	if manager == nil {
		manager = NewPokemonManager()
	}
	
	return manager
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
		" RIGHT JOIN move_flavor_text ON id_move  = id" + 
		" WHERE version_group_id = 11"
	result, err := puh.DBQuerySelect(query)
	if err != nil {
		return false
	}
	
	defer puh.DBFree()
	log.Println(" - Processing moves")
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		move := NewMove()
		move.MoveId = int(row[0].(int64))
		move.Identifier = puh.DBGetString(row[1])
		move.TypeId = puh.DBGetInt(row[2])
		move.Power = puh.DBGetInt(row[3])
		move.Accuracy = puh.DBGetInt(row[4])
		move.Priority = puh.DBGetInt(row[5])
		move.TargetId = puh.DBGetInt(row[6])
		move.DamageClassId = puh.DBGetInt(row[7])
		move.EffectId = puh.DBGetInt(row[8])
		move.EffectChance = puh.DBGetInt(row[9])
		move.ContestType = puh.DBGetInt(row[10])
		move.ContestEffect = puh.DBGetInt(row[11])
		move.SuperContestEffect = puh.DBGetInt(row[12])
		move.PP = puh.DBGetInt(row[13])
		move.FlavorText = puh.DBGetString(row[14])
		
		// Add to map
		m.moves[move.MoveId] = move
	}
	
	return true
}

func (m *PokemonManager) loadAbilities() bool {
	var query string = "SELECT id, identifier FROM abilities"
	result, err := puh.DBQuerySelect(query)
	if err != nil {
		return false
	}
	
	defer puh.DBFree()
	log.Println(" - Processing abilities")
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		ability := NewAbility()
		ability.AbilityId = puh.DBGetInt(row[0])
		ability.Identifier = puh.DBGetString(row[1])
		
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
	result, err := puh.DBQuerySelect(query)
	if err != nil {
		return false
	}
	
	defer puh.DBFree()
	log.Println(" - Processing pokemon species")
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		// Create EvolutionChain object
		evoChain := NewEvolutionChain()
		evoChain.EvolvedSpeciesId = puh.DBGetInt(row[14])
		evoChain.EvolutionTriggerId = puh.DBGetInt(row[15])
		evoChain.TriggerItemId = puh.DBGetInt(row[16]) // TODO: Change this to the actual Item objet
		evoChain.MinimumLevel = puh.DBGetInt(row[17])
		evoChain.Gender = puh.DBGetString(row[18])
		evoChain.LocationId = puh.DBGetInt(row[19])
		evoChain.HeldItemId = puh.DBGetInt(row[20])
		evoChain.TimeOfDay = puh.DBGetString(row[21])
		evoChain.KnownMoveId = puh.DBGetInt(row[22]) // TODO: Change to move object
		evoChain.MinimumHappiness = puh.DBGetInt(row[23])
		evoChain.MinimumBeauty = puh.DBGetInt(row[24])
		evoChain.RelativePhysicalStats = puh.DBGetInt(row[25])
		evoChain.PartySpeciesId = puh.DBGetInt(row[26])
		evoChain.TradeSpeciesId = puh.DBGetInt(row[27])
		
		// Creat PokemonSpecies object
		pokemon := NewPokemonSpecies()
		pokemon.SpeciesId = puh.DBGetInt(row[0])
		pokemon.Identifier = puh.DBGetString(row[1])
		pokemon.EvolvesFromSpeciesId = puh.DBGetInt(row[2])
		pokemon.EvolutionChain = evoChain
		pokemon.ColorId = puh.DBGetInt(row[3])
		pokemon.ShapeId = puh.DBGetInt(row[4])
		pokemon.HabitatId = puh.DBGetInt(row[5])
		pokemon.GenderRate = puh.DBGetInt(row[6])
		pokemon.CaptureRate = puh.DBGetInt(row[7])
		pokemon.BaseHappiness = puh.DBGetInt(row[8])
		pokemon.IsBaby = puh.DBGetInt(row[9])
		pokemon.HatchCounter = puh.DBGetInt(row[10])
		pokemon.HasGenderDifferences = puh.DBGetInt(row[11])
		pokemon.GrowthRateId = puh.DBGetInt(row[12])
		pokemon.FormsSwitchable = puh.DBGetInt(row[12])
		
		// Add to map
		m.pokemonSpecies[pokemon.SpeciesId] = pokemon
	}
	
	return true
}

func (m *PokemonManager) loadPokemon() bool {
	var query string = "SELECT `id`, `species_id`, `height`, `weight`, `base_experience`, `order`, `is_default` FROM pokemon"
	result, err := puh.DBQuerySelect(query)
	if err != nil {
		return false
	}
	
	defer puh.DBFree()
	log.Println(" - Processing pokemon")
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		pokemon := NewPokemon()
		pokemon.PokemonId = puh.DBGetInt(row[0])
		pokemon.Species = m.GetPokemonSpecies(puh.DBGetInt(row[1]))
		pokemon.Height = puh.DBGetInt(row[2])
		pokemon.Weight = puh.DBGetInt(row[3])
		pokemon.BaseExperience = puh.DBGetInt(row[4])
		pokemon.Order = puh.DBGetInt(row[5])
		pokemon.IsDefault = puh.DBGetInt(row[6])
				
		// Add to map
		m.pokemon[pokemon.PokemonId] = pokemon
	}
	
	return true
}

func (m *PokemonManager) loadMoveMessages() bool {
	var query string = "SELECT `move_effect_id`, `message` FROM move_messages"
	result, err := puh.DBQuerySelect(query)
	if err != nil {
		return false
	}
	
	log.Println(" - Processing Move Messages")
	defer puh.DBFree()
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		effect_id := puh.DBGetInt(row[0])
		message := puh.DBGetString(row[1])
		
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
	result, err := puh.DBQuerySelect(query)
	if err != nil {
		return false
	}
	
	log.Println(" - Processing Ability Messages")
	defer puh.DBFree()
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		ability_id := puh.DBGetInt(row[0])
		message := puh.DBGetString(row[1])
		
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
		log.Printf("PokemonManager::GetPokemonTypes - Could not find pokemon with id: %d\n\r", _pokemonId)
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