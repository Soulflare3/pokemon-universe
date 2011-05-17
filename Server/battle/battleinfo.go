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
	"container/vector"
)

type BattleInfo struct {
	BaseBattleInfo

	myteam		*TeamBattle
	possible	bool
	sent		bool
	
	choices		vector.Vector
	choice		vector.Vector
	available	vector.Vector
	done		vector.Vector
	
	currentSlot	int8
	
	mystats		vector.Vector
	tempPoke	vector.Vector
	
	lastMove	[]int8
}

func NewBattleInfo(_team *TeamBattle, _me, _opp *PlayerInfo, _mode uint8, _my, _op int32) *BattleInfo {
	base := &BattleInfo{lastMove: make([]int8, 6) }
	base.Init(_me, _opp, _mode, _my, _op)
	
	base.possible = false
	base.myteam = _team
	base.sent = true
	
	base.currentSlot = base.Player(base.myself)
	//base.currentSlot = base.Slot(base.myself, 0)
	
	for i := 0; i < base.numberOfSlots/2; i++ {
		base.choices.Push(NewBattleChoices())
		base.choice.Push(NewBattleChoice())
		base.available.Push(false)
		base.done.Push(false)
		
		base.mystats.Push(NewBattleStats())
		base.tempPoke.Push(NewPokeBattle())
	}
	
	for i := 0; i < 6; i++ {
		base.pokemons[base.myself][i] = _team.Poke(int8(i))
	}
	
	return base
}

func (b *BattleInfo) GetTempPoke(_spot int8) *PokeBattle {
	return b.tempPoke.At(int(b.Number(_spot))).(*PokeBattle)
}

func (b *BattleInfo) SetTempPoke(_spot int8, _poke *PokeBattle) {
	b.tempPoke.Set(int(b.Number(_spot)), _poke)
}

func (b *BattleInfo) SwitchPokeExt(_spot int8, _poke int8, _own bool) {
	b.SwitchPoke(_spot, _poke)
	if _own {
		b.myteam.SwitchPokemon(b.Number(_spot), _poke)
		b.SetCurrentShallow(_spot, b.myteam.Poke(b.Number(_spot)))
		b.SetTempPoke(_spot, b.myteam.Poke(b.Number(_spot)))
	}
}

func (b *BattleInfo) CurrentPoke(_spot int8) *PokeBattle {
	return b.myteam.Poke(b.Number(_spot))
}