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
	"container/list"
)

type Pokemon struct {
	PokemonId				int
	Species					*PokemonSpecies
	Height					int
	Weight					int
	BaseExperience			int
	Order					int
	IsDefault				int
	
	Stats					PokemonStatArray // Size = 6
	
	Abilities				PokemonAbilityList
	Forms					*list.List
	Moves					PokemonMoveList
	Types					PokemonTypeArray
}

func NewPokemon() *Pokemon {
	pokemon := Pokemon{ Stats: make(PokemonStatArray, 6),
					 Abilities: make(PokemonAbilityList),
					 Moves: make(PokemonMoveList),
					 Types: make(PokemonTypeArray, 2),
					 Forms: new(list.List) }
	// pokemon.Forms.Init()
	
	pokemon.Types[0] = 0
	pokemon.Types[1] = 0
	
	return &pokemon
}

func (p *Pokemon) loadStats() bool {
	var query string = "SELECT stat_id, base_stat, effort FROM pokemon_stats WHERE pokemon_id='%d'"
	result, err := DBQuerySelect(fmt.Sprintf(query, p.PokemonId))
	if err != nil {
		return false
	}
	
	defer result.Free()
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		stat := NewPokemonStat()
		stat.StatType = DBGetInt(row[0])
		stat.BaseStat = DBGetInt(row[1])
		stat.Effort = DBGetInt(row[2])
		p.Stats[stat.StatType-1] = stat
	}
	
	return true
}

func (p *Pokemon) loadAbilities() bool {
	var query string = "SELECT ability_id, is_dream, slot FROM pokemon_abilities WHERE pokemon_id='%d'"
	result, err := DBQuerySelect(fmt.Sprintf(query, p.PokemonId))
	if err != nil {
		return false
	}
	
	defer result.Free()
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		ability := NewPokemonAbility()
		id := DBGetInt(row[0])
		ability.Ability = g_PokemonManager.GetAbilityById(id)
		ability.IsDream = DBGetInt(row[1])
		ability.Slot = DBGetInt(row[2])

		if ability.Ability != nil {
			p.Abilities[id] = ability
		}
	}		
	
	return true
}

func (p *Pokemon) loadForms() bool {
	var query string = "SELECT `id`, `form_identifier`, `is_default`, `is_battle_only`, `order` FROM pokemon_forms WHERE pokemon_id='%d'"
	result, err := DBQuerySelect(fmt.Sprintf(query, p.PokemonId))
	if err != nil {
		return false
	}
	
	defer result.Free()
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		form := NewPokemonForm()
		form.Id = DBGetInt(row[0])
		form.Identifier = DBGetString(row[1])
		form.IsDefault = DBGetInt(row[2])
		form.IsBattleOnly = DBGetInt(row[3])
		form.Order = DBGetInt(row[4])
		
		p.Forms.PushBack(form)
	}
	
	return true
}

func (p *Pokemon) loadMoves() bool {
	var query string = "SELECT `version_group_id`, `move_id`, `pokemon_move_method_id`, `level`, `order` FROM pokemon_moves" + 
						" WHERE pokemon_id='%d' AND version_group_id=11"
	result, err := DBQuerySelect(fmt.Sprintf(query, p.PokemonId))
	if err != nil {
		return false
	}
	
	defer result.Free()
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		pmove := NewPokemonMove()
		pmove.Pokemon = p
		pmove.VersionGroup = DBGetInt(row[0])
		moveId := DBGetInt(row[1])
		pmove.Move = g_PokemonManager.GetMoveById(moveId)
		pmove.PokemonMoveMethod = DBGetInt(row[2])
		pmove.Level = DBGetInt(row[3])
		pmove.Order = DBGetInt(row[4])
		
		if pmove.Move != nil {
			p.Moves[moveId] = pmove
		}
	}
	
	return true
}

func (p *Pokemon) loadTypes() bool {
	var query string = "SELECT type_id, slot FROM pokemon_types WHERE pokemon_id='%d' ORDER BY slot"
	result, err := DBQuerySelect(fmt.Sprintf(query, p.PokemonId))
	if err != nil {
		return false
	}
	
	defer result.Free()
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		slot := DBGetInt(row[1])
		p.Types[slot - 1] = DBGetInt(row[0])
	}
	
	return true
}