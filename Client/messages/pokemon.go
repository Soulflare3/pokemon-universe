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
	punet "network"
	"os"
)

type PU_Message_ReceivePokemon struct {
	player *PU_Player
}

func NewReceivePokemonMessage(_packet *punet.Packet) *PU_Message_ReceivePokemon {
	msg := &PU_Message_ReceivePokemon{}
	msg.ReadPacket(_packet)
	return msg
}

func (m *PU_Message_ReceivePokemon) ReadPacket(_packet *punet.Packet) os.Error {
	if g_game.self == nil {
		return nil
	}
	
	count := int(_packet.ReadUint16())
	for i := 0; i < count; i++ {
		slot := _packet.ReadUint16()
		if slot < 0 || slot > 5 {
			break
		}
		
		pokemon := g_game.self.pokemon[slot]
		if pokemon == nil {
			g_game.self.pokemon[slot] = NewPokemon()
			pokemon = g_game.self.pokemon[slot]
		}
		
		pokemon.uid = _packet.ReadUint32()
		pokemon.id = int16(_packet.ReadUint16())
		pokemon.name = _packet.ReadString()
		pokemon.level = int16(_packet.ReadUint16())
		pokemon.hpmax = int16(_packet.ReadUint16())
		pokemon.hp = int16(_packet.ReadUint16())
		pokemon.expPerc = int16(_packet.ReadUint16())
		pokemon.expCurrent = int32(_packet.ReadUint32())
		pokemon.expTnl = int32(_packet.ReadUint32())-pokemon.expCurrent
		pokemon.type1 = _packet.ReadString()
		pokemon.type2 = _packet.ReadString()
		pokemon.flavor = _packet.ReadString()
		pokemon.sex = int16(_packet.ReadUint16())
		
		pokemon.stats[POKESTAT_ATTACK] = int16(_packet.ReadUint16())
		pokemon.stats[POKESTAT_DEFENSE] = int16(_packet.ReadUint16())
		pokemon.stats[POKESTAT_SPECIALATTACK] = int16(_packet.ReadUint16())
		pokemon.stats[POKESTAT_SPECIALDEFENSE] = int16(_packet.ReadUint16())
		pokemon.stats[POKESTAT_SPEED] = int16(_packet.ReadUint16())
		
		attackCount := int(_packet.ReadUint16())
		for a := 0; a < attackCount; a++ {
			num := int(_packet.ReadUint16())
			name := _packet.ReadString()
			description := _packet.ReadString()
			poketype := _packet.ReadString()
			pp := uint16(_packet.ReadUint16())
			ppmax := uint16(_packet.ReadUint16())
			power := uint16(_packet.ReadUint16())
			accuracy := uint16(_packet.ReadUint16())
			category := _packet.ReadString()
			target := _packet.ReadString()
			contact := _packet.ReadString()
			
			pokemon.SetAttack(num, name, description, poketype, pp, ppmax, power, accuracy, category, target, contact)
		}
	}
	return nil
}

type PU_Message_RefreshPokemon struct {
}

func NewRefreshPokemonMessage() *PU_Message_RefreshPokemon {
	return &PU_Message_RefreshPokemon{}
}

func (m *PU_Message_RefreshPokemon) WritePacket() (*punet.Packet, os.Error) {
	packet := punet.NewPacket()
	packet.AddUint8(0xD1)
	return packet, nil
}
