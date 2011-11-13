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
	pnet "network"
)

type BattlePoke struct {
	ShallowBattlePoke // Extends
	
	CurrentHP int
	TotalHP int
	ItemString string
	AbilityString string
	TeamNum int
	Stats []int
	Moves []*BattleMove
	
	item int
	ability int
	
	statusCount int
	originalStatusCount int
	nature int
	happiness int
	
	DVs []int
	EVs []int
}

func NewBattlePokeFromPacket(_packet *pnet.QTPacket) *BattlePoke {
	battlePoke := BattlePoke{}
	battlePoke.UID = NewUniqueIdFromPacket(_packet)	
	battlePoke.Nick = _packet.ReadString()
	battlePoke.TotalHP = int(_packet.ReadUint16())
	battlePoke.CurrentHP = int(_packet.ReadUint16())
	battlePoke.Gender = int(_packet.ReadUint8())
	battlePoke.Shiny = _packet.ReadBool()
	battlePoke.Level = int(_packet.ReadUint8())
	battlePoke.item = int(_packet.ReadUint16())
	battlePoke.ItemString = ""
	battlePoke.ability = int(_packet.ReadUint16())
	battlePoke.happiness = int(_packet.ReadUint8())
	
	battlePoke.Stats = make([]int, 5)
	for i := 0; i < 5; i++ {
		battlePoke.Stats[i] = int(_packet.ReadUint16())
	}
	battlePoke.Moves = make([]*BattleMove, 4)
	for i := 0; i < 4; i++ {
		battlePoke.Moves[i] = NewBattleMoveFromPacket(_packet)
	}
	battlePoke.EVs = make([]int, 6)
	for i := 0; i < 6; i++ {
		battlePoke.EVs[i] = int(_packet.ReadUint32())
	}
	battlePoke.DVs = make([]int, 6)
	for i := 0; i < 6; i++ {
		battlePoke.DVs[i] = int(_packet.ReadUint32())
	}
	
	if g_PokemonManager.GetPokemon(battlePoke.UID.PokeNum) == nil {
		return nil
	}
	
	battlePoke.getName()	
	battlePoke.getTypes()
		
	return &battlePoke
}