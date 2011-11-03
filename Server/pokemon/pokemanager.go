package main

import (
	"os"
	"container/list"
	"mysql"
)

type PokemonList 			map[int]*Pokemon
type PokemonSpeciesList 	map[int]*PokemonSpecies
type MoveList 				map[int]*Move
type AbilityList 			map[int]*Ability
type PokemonStatArray 		[]*PokemonStat
type PokemonAbilityList 	map[int]*PokemonAbility
type PokemonFormList 		*list.List
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
		g_logger.Println("[ERROR] SQL error while executing query pokemon moves: " + err)
		return false
	}
	
	var result *mysql.Result
	result, err = g_db.UseResult()
	if err != nil {
		g_logger.Println("[ERROR] SQL error while fetching result pokemon moves: " + err)
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
		g_logger.Println("[ERROR] SQL error while executing query pokemon abilities: " + err)
		return false
	}
	
	var result *mysql.Result
	result, err = g_db.UseResult()
	if err != nil {
		g_logger.Println("[ERROR] SQL error while fetching result pokemon abilities: " + err)
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
		g_logger.Println("[ERROR] SQL error while executing query pokemon species: " + err)
		return false
	}
	
	var result *mysql.Result
	result, err = g_db.UseResult()
	if err != nil {
		g_logger.Println("[ERROR] SQL error while fetching result pokemon species: " + err)
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
		pokemon.EvolvedFromSpeciesId = row[2].(int)
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
}

func (m *PokemonManager) loadPokemon() {
	var query string = "SELECT id, species_id, height, weight, base_experience, order, is_default FROM pokemon"
	result, err := DBQuerySelect(query)
	if err != nil {
		return
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
		pokemon.Species = GetPokemonSpecies(row[1].(int))
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
		
		// Add to map
		m.pokemon[pokemon.PokemonId] = pokemon
	}
}

func (m *PokemonManager) GetPokemonSpecies(_speciesId int) *PokemonSpecies {
	return m.pokemonSpecies[_speciesId]
}

func (m *PokemonManager) GetAbilityById(_abilityId int) *Ability {
	return m.abilities[_abilityId]
}

func (m *PokemonManager) GetMoveById(_moveId int) *Move {
	return m.moves[_moveId]
}