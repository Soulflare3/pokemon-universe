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
	"fmt"
	"container/list"
	
	"nonamelib/log"
	
	"pulogic/models"
)

type Pokemon struct {
	PokemonId				int
	Species					*PokemonSpecies
	Height					int
	Weight					int
	BaseExperience			int
	//Order					int
	IsDefault				bool
	
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
	
	pokemon.Types[0] = 0
	pokemon.Types[1] = 0
	
	return &pokemon
}

func NewPokemonFromEntity(_entity models.Pokemon) *Pokemon {
	pokemon := NewPokemon()
	pokemon.PokemonId = _entity.Id
	pokemon.Species = manager.GetPokemonSpecies(_entity.SpeciesId)
	pokemon.Height = _entity.Height
	pokemon.Weight = _entity.Weight
	pokemon.BaseExperience = _entity.BaseExperience
	//pokemon.Order = _entity.Order
	pokemon.IsDefault = _entity.IsDefault
	
	return pokemon
}

func (p *Pokemon) loadStats() bool {
	var base_stats []models.PokemonStats
	err := G_orm.Where(fmt.Sprintf("%v = %d", models.PokemonStats_PokemonId, p.PokemonId)).FindAll(&base_stats)
	if err != nil {
		log.Error("Pokemon", "loadStats", "Failed to load stats. Error: %v", err.Error())
		return false
	}

	for _, row := range(base_stats) {
		stat := NewPokemonStatFromEntity(row)
		p.Stats[stat.StatType-1] = stat
	}
	
	return true
}

func (p *Pokemon) loadAbilities() bool {
	var abilities []models.PokemonAbilities
	err := G_orm.Where(fmt.Sprintf("%v = %d", models.PokemonAbilities_PokemonId, p.PokemonId)).FindAll(&abilities)
	if err != nil {
		log.Error("Pokemon", "loadAbilities", "Failed to load stats. Error: %v", err.Error())
		return false
	}

	for _, row := range(abilities) {
		ability := NewPokemonAbilityFromEntity(row)
		
		if ability.Ability != nil {
			p.Abilities[ability.Ability.AbilityId] = ability
		}
	}		
	
	return true
}

func (p *Pokemon) loadForms() bool {
	var forms []models.PokemonForms
	err := G_orm.Where(fmt.Sprintf("%v = %d", models.PokemonForms_PokemonId, p.PokemonId)).FindAll(&forms)
	if err != nil {
		log.Error("Pokemon", "loadForms", "Failed to load forms. Error: %v", err.Error())
		return false
	}

	for _, row := range(forms) {
		form := NewPokemonFormFromEntity(row)		
		p.Forms.PushBack(form)
	}
	
	return true
}

func (p *Pokemon) loadMoves() bool {
	var moves []models.PokemonMoves
	err := G_orm.Where(fmt.Sprintf("%v = %d AND version_group_id=11", models.PokemonMoves_PokemonId, p.PokemonId)).FindAll(&moves)
	if err != nil {
		log.Error("Pokemon", "loadMoves", "Failed to load pokemon moves. Error: %v", err.Error())
		return false
	}

	for _, row := range(moves) {
		pmove := NewPokemonMoveFromEntity(p, row)
		
		if pmove.Move != nil {
			p.Moves[pmove.Move.MoveId] = pmove
		}
	}
	
	return true
}

func (p *Pokemon) loadTypes() bool {
	var types []models.PokemonTypes
	err := G_orm.Where(fmt.Sprintf("%v = %d", models.PokemonTypes_PokemonId, p.PokemonId)).OrderBy(models.PokemonTypes_Slot).FindAll(&types)
	if err != nil {
		log.Error("Pokemon", "loadTypes", "Failed to load pokemon types. Error: %v", err.Error())
		return false
	}
	
	for _, row := range(types) {		
		slot := row.Slot
		p.Types[slot - 1] = row.TypeId
	}
	
	return true
}