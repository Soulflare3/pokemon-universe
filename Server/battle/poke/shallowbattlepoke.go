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

type ShallowBattlePoke struct {
	RNick string
	Nick string
	PokeName string
	UID *UniqueId
	Types []int
	Shiny bool
	Gender int
	LifePercent int
	Level int
	LastKnownPercent int
	Sub bool
	
	fullStatus uint32
}

func NewShallowBattlePoke() *ShallowBattlePoke {
	return &ShallowBattlePoke{ Types: make([]int, 2) }
}

func NewShallowBattlePokeFromPacket(_packet *pnet.QTPacket, _isMe bool) *ShallowBattlePoke {
	shallowPoke := ShallowBattlePoke{ Types: make([]int, 2) }
	shallowPoke.UID = NewUniqueIdFromPacket(_packet)
	shallowPoke.RNick = _packet.ReadString()
	shallowPoke.Nick = shallowPoke.RNick
	if !_isMe {
		shallowPoke.Nick = "the foe's " + shallowPoke.Nick
		
		shallowPoke.getName()
		shallowPoke.getTypes()
	}
	
	shallowPoke.LifePercent = int(_packet.ReadUint8())
	shallowPoke.fullStatus = _packet.ReadUint32()
	shallowPoke.Gender = int(_packet.ReadUint8())
	shallowPoke.Shiny = _packet.ReadBool()
	shallowPoke.Level = int(_packet.ReadUint32())
	
	return &shallowPoke
}

func (s *ShallowBattlePoke) getName() {
	s.PokeName = g_PokemonManager.GetPokemonName(s.UID.PokeNum, s.UID.SubNum)
}

func (s *ShallowBattlePoke) getTypes() {
	s.Types = g_PokemonManager.GetPokemonTypes(s.UID.PokeNum, s.UID.SubNum)
}

func (s *ShallowBattlePoke) ChangeStatus(_status uint) {
	// Clear past status
	s.fullStatus = s.fullStatus & ^(uint32(1 << STATUS_KEOD) | 0x3F)
	
	// Add new status
	s.fullStatus = s.fullStatus | (1 << _status)
	
}