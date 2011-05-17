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

type PokeBattle struct {
	nick			string
	num				PokemonUniqueId
	
	dvs				[]uint8
	evs				[]uint8
	
	normal_stats	[]uint16
	moves			[]*BattleMove
	
	lifePoints		uint16
	totalLifePoints uint16
	lifePercent		uint8
	
	item			uint16
	ability			uint16
	fullStatus		uint32
	statusCount		int8
	oriStatusCount	int8
	gender			uint8
	level			uint8
	nature			uint8
	happiness		uint8
	shiny			bool
}

func NewPokeBattle() *PokeBattle {
	return &PokeBattle{ dvs: make([]uint8, 6),
					evs: make([]uint8, 6),
					normal_stats: make([]uint16, 5),
					moves: make([]*BattleMove, 4) }
}

func NewPokeBattleFromPacket(_packet *pnet.QTPacket) *PokeBattle {
	poke := NewPokeBattle()

	pokeNum := _packet.ReadUint16()
	poke.num = NewPokemonUniqueIdFromRef(uint32(pokeNum))
	if poke.num.pokenum > 0 && poke.num.subnum >= 0 {
		poke.nick = _packet.ReadString()
		poke.lifePercent = _packet.ReadUint8()
		poke.fullStatus = _packet.ReadUint32()
		poke.gender = _packet.ReadUint8()
		poke.shiny = (_packet.ReadUint8() == 1) 
		poke.level = _packet.ReadUint8()
	}
	
	return poke	
}

func (p *PokeBattle) Init() {
	if p.totalLifePoints > 0 {
		p.lifePercent = uint8((p.lifePoints * 100) / p.totalLifePoints)
		if p.lifePercent == 0 && p.lifePoints > 0 {
			p.lifePercent = 1
		}
	}
}

func (p *PokeBattle) ChangeStatus(_status uint8) {
	// Clear past status
	p.fullStatus = p.fullStatus & ^(uint32(1 << PokemonStatus_Koed) | 0x3F)	
	// Add new status
	p.fullStatus = p.fullStatus | (1 << _status)
}