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
	"fmt" 
	
	"github.com/astaxie/beedb"
	
	"nonamelib/log"
	"pulogic/models"
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
var G_orm *beedb.Model

func NewPokemonManager() *PokemonManager {
	if G_orm == nil {
		panic("G_orm NOT DEFINED!")
	}
	
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
	
	// Load ability messages
	if !m.loadAbilityMessages() {
		return false
	}
	
	return true
}

func (m *PokemonManager) loadMoves() bool {
	var moves []models.MovesJoinMoveFlavorText
	err := G_orm.SetTable("moves").Join("LEFT", "move_flavor_text", fmt.Sprintf("%v = %v", models.MoveFlavorText_IdMove, models.Moves_Id)).FindAll(&moves)
	if err != nil {
		log.Error("PokeManager", "loadMoves", "Failed to load moves. Error: %v", err.Error())
		return false
	}
	
	log.Verbose("PokeManager", "loadMoves", "Processing moves")
	for _, row := range(moves) {
		move := NewMoveFromEntity(row)	
		m.moves[move.MoveId] = move 
	}
	
	return true
}

func (m *PokemonManager) loadAbilities() bool {
	var abilities []models.Abilities
	err := G_orm.FindAll(&abilities)
	if err != nil {
		log.Error("PokeManager", "loadAbilities", "Failed to load abilities. Error: %v", err.Error())
		return false
	}
	
	log.Verbose("PokeManager", "loadAbilities", "Processing abiliites")
	for _, row := range(abilities) {
		ability := NewAbilityFromEntity(row)
		m.abilities[ability.AbilityId] = ability
	}
	
	return true
}

func (m *PokemonManager) loadPokemonSpecies() bool {
	var pokes []models.PokemonSpeciesJoinPokemonEvolution
	err := G_orm.SetTable("pokemon_species").Join("LEFT", "pokemon_evolution", "pokemon_evolution.evolved_species_id = (SELECT `ps2`.id FROM pokemon_species AS `ps2` WHERE `ps2`.evolves_from_species_id = pokemon_species.id LIMIT 1)").FindAll(&pokes)
	if err != nil {
		log.Error("PokeManager", "loadPokemonSpecies", "Failed to load pokemon species. Error: %v", err.Error())
		return false
	}
	
	log.Verbose("PokeManager", "loadPokemonSpecies", "Processing pokemon species")
	for _, row := range(pokes) {
		pokemon := NewPokemonSpecesFromEntity(row)
		m.pokemonSpecies[pokemon.SpeciesId] = pokemon
	}
	
	return true
}

func (m *PokemonManager) loadPokemon() bool {
	var pokes []models.Pokemon
	err := G_orm.FindAll(&pokes)
	if err != nil {
		log.Error("PokeManager", "loadPokemon", "Failed to load pokemon. Error: %v", err.Error())
		return false
	}

	log.Verbose("PokeManager", "loadPokemon", "Processing pokemon")
	for _, row := range(pokes) {
		pokemon := NewPokemonFromEntity(row)
				
		// Add to map
		m.pokemon[pokemon.PokemonId] = pokemon
	}
	
	return true
}

func (m *PokemonManager) loadMoveMessages() bool {
	var move_messages []models.MoveMessages
	err := G_orm.FindAll(&move_messages)
	if err != nil {
		log.Error("PokeManager", "loadMoveMessages", "Failed to load move messages. Error: %v", err.Error())
		return false
	}
	
	log.Verbose("PokeManager", "loadMoveMessages", "Processing move messages")
	for _, row := range(move_messages) {
		messages := strings.Split(row.Message, "|")
		messageMap := make(map[int]string)
		for index, msg := range(messages) {
			messageMap[index] = msg
		}
		
		m.moveMessages[row.MoveEffectId] = messageMap
	}	
	
	return true
}

func (m *PokemonManager) loadAbilityMessages() bool {
	var ability_messages []models.AbilityMessages
	err := G_orm.FindAll(&ability_messages)
	if err != nil {
		log.Error("PokeManager", "loadAbilityMessages", "Failed to load ability messages. Error: %v", err.Error())
		return false
	}
	
	log.Verbose("PokeManager", "loadAbilityMessagess", "Processing ability messages")
	for _, row := range(ability_messages) {
		messages := strings.Split(row.Message, "|")
		messageMap := make(map[int]string)
		for index, msg := range(messages) {
			messageMap[index] = msg
		}
		
		m.abilityMessages[row.AbilityId] = messageMap
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
		log.Error("PokemonManager", "GetPokemonTypes", "Could not find pokemon with id: %d", _pokemonId)
		return nil
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
	move, found := m.moves[_moveId]
	
	if !found {
		log.Error("PokeManager", "GetMoveById", "Could not find move with id: %d", _moveId)
		return nil
	}
	
	return move
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