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
)

type PlayerPokemonList map[int]*PlayerPokemon

type PlayerPokemon struct {
	IdDb		int
	PlayerId	int
	Base		*Pokemon
	Nickname	string
	IsBound		int // Can (not) trade if 1
	Experience	int
	Stats		[]int
	Happiness	int
	Gender		int
	Ability		*Ability
	Moves		[]*Move
	IsShiny		int
	InParty		int
	Slot		int
	Nature		int
}

func NewPlayerPokemon(_playerId int) *PlayerPokemon {
	return &PlayerPokemon{ Stats: make([]int, 6),
							Moves: make([]*Move, 4),
							Nature: 0,
							PlayerId: _playerId }
}

func (p *PlayerPokemon) LoadMoves() {
	var query string = "SELECT idmove FROM player_pokemon_move WHERE idplayer_pokemon='%d'"
	derp := fmt.Sprintf(query, p.IdDb)
	result, err := DBQuerySelect(derp)
	if err != nil {
		return
	}
	
	defer result.Free()
	var index int = 0
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		moveId := DBGetInt(row[0])
		p.Moves[index] = g_PokemonManager.GetMoveById(moveId)
		index++
	}
}

func (p *PlayerPokemon) GetNickname() string {
	if len(p.Nickname) == 0 {
		return p.Base.Species.Identifier
	}
	return p.Nickname
}

func (p *PlayerPokemon) GetLevel() int {
	// TODO: Calculate pokemon level from exp
	
	return 25
}