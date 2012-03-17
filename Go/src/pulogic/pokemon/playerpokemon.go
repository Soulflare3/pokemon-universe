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
	
	puh "puhelper"
	log "putools/log"
)

type PlayerPokemonList map[int]*PlayerPokemon

type PlayerPokemon struct {
	IdDb		int
	PlayerId	int
	Base		*Pokemon
	Nickname	string
	IsBound		int // Can (not) trade if 1
	Experience	float64
	Stats		[]int
	Happiness	int
	Gender		int
	Ability		*Ability
	Moves		[]*PlayerPokemonMove
	IsShiny		int
	InParty		int
	Slot		int
	Nature		int
	TotalHp		int
	DamagedHp	int
}

func NewPlayerPokemon(_playerId int) *PlayerPokemon {
	return &PlayerPokemon{ Stats: make([]int, 6),
							Moves: make([]*PlayerPokemonMove, 4),
							Nature: 0,
							PlayerId: _playerId }
}

func (p *PlayerPokemon) LoadMoves() {
	var query string = "SELECT idplayer_pokemon_move, idmove, pp_used FROM player_pokemon_move WHERE idplayer_pokemon='%d'"
	result, err := puh.DBQuerySelect(fmt.Sprintf(query, p.IdDb))
	if err != nil {
		return
	}
	
	defer puh.DBFree()
	if result.RowCount() == 0 {
		log.Printf("[WARNING] Pokemon (db id: %d) has zero moves\n", p.IdDb)
	}
	
	var index int = 0
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		uniqueId := puh.DBGetInt64(row[0])
		moveId := puh.DBGetInt(row[1])
		ppUsed := puh.DBGetInt(row[2])
		p.Moves[index] = NewPlayerPokemonMove(uniqueId, moveId, ppUsed)
		index++
	}
}

func (p *PlayerPokemon) SaveMoves() {
	
	for i := 0; i < 4; i++ {
		
		if move := p.Moves[i]; move != nil {
			
			query := "UPDATE player_pokemon_move SET idmove=%d, pp_used=%d WHERE idplayer_pokemon_move=%d"
			puh.DBQuery(fmt.Sprintf(query, move.Move.MoveId, move.CurrentPP, move.DbId))
		}
	}
}

func (p *PlayerPokemon) GetNickname() string {
	if len(p.Nickname) == 0 {
		return p.Base.Species.Identifier
	}
	return p.Nickname
}

func (p *PlayerPokemon) GetLevel() int {
	return puh.CalculateLevelFromExperience(p.Experience)
}