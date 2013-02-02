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
	log "nonamelib/log"
	
	"pulogic/models"
)

type PlayerPokemonList map[int]*PlayerPokemon

type PlayerPokemon struct {
	IdDb		int
	PlayerId	int
	Base		*Pokemon
	Nickname	string
	IsBound		bool // Can (not) trade if 1
	Experience	int64
	Stats		[]int
	Happiness	int
	Gender		int
	Ability		*Ability
	Moves		[]*PlayerPokemonMove
	IsShiny		bool
	InParty		bool
	Slot		int
	Nature		int
	DamagedHp	int
	
	IsDeleted	bool
	IsModified	bool
}

func NewPlayerPokemon(_playerId int) *PlayerPokemon {
	return &PlayerPokemon{ Stats: make([]int, 6),
							Moves: make([]*PlayerPokemonMove, 4),
							Nature: 0,
							PlayerId: _playerId,
							IsDeleted: false }
}

func (p *PlayerPokemon) LoadMoves() {
	pokemonEntities := []models.PlayerPokemonMove{}
	err := G_orm.FindAll(&pokemonEntities)
	if err != nil {
		log.Error("PlayerPokemon", "LoadMoves", "Error loading pokemon moves. IdplayerPokemon: %d, Error: %s", p.IdDb, err.Error())
	} else {
		index := 0
		for _, entity := range(pokemonEntities) {
			p.Moves[index] = NewPlayerPokemonMove(entity.IdplayerPokemonMove, entity.Idmove, entity.PpUsed)
			index++
		}
		
		if index == 0 {
			log.Warning("PlayerPokemon", "LoadMoves", "Pokemon has zero moves. IdPlayerPokemon: %d", p.IdDb)
		}
	}
}

func (p *PlayerPokemon) SaveMoves() {
//	for i := 0; i < 4; i++ {
//		
//		if move := p.Moves[i]; move != nil {
//			
//			query := "UPDATE player_pokemon_move SET idmove=%d, pp_used=%d WHERE idplayer_pokemon_move=%d"
//			puh.DBQuery(fmt.Sprintf(query, move.Move.MoveId, move.CurrentPP, move.DbId))
//		}
//	}
}

func (p *PlayerPokemon) GetNickname() string {
	if len(p.Nickname) == 0 {
		return p.Base.Species.Identifier
	}
	return p.Nickname
}

func (p *PlayerPokemon) GetLevel() int {
	return CalculateLevelFromExperience(p.Experience)
}

func (p *PlayerPokemon) GetTotalHp() int {
	return HpForLevel(p.Base.Stats[POKESTAT_HP].BaseStat, p.Stats[POKESTAT_HP], p.GetLevel())
}

func (p *PlayerPokemon) IsFainted() bool {
	return (p.DamagedHp >= p.GetTotalHp())
}