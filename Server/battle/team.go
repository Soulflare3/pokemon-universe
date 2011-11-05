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

type Team struct {
	Gen int
	Pokes []*TeamPoke
}

func NewTeam() *Team {
	team := Team { Pokes: make([]*TeamPoke, 6) }
	for i := 0; i < 6; i++ {
		team.Pokes[i] = NewTeamPoke()
	}
	
	return &team
}

func NewTeamFromParty(_party *PokemonParty) *Team {
	team := Team { Pokes: make([]*TeamPoke, 6) }
	for i := 0; i < 6; i++ {
		if _party.Party[i] != nil {
			team.Pokes[i] = NewTeamPokeFromPokemon(_party.Party[i])
		}
	}
	return &team
}

func NewTeamFromPacket(_packet *pnet.QTPacket) *Team {
	team := Team { Pokes: make([]*TeamPoke, 6) }
	team.Gen = int(_packet.ReadUint8())
	for i := 0; i < 6; i++ {
		team.Pokes[i] = NewTeamPokeFromPacket(_packet)
	}
	return &team
}

func (t *Team) WritePacket() pnet.IPacket {
	packet := pnet.NewQTPacket()
	packet.AddUint8(uint8(t.Gen))
	for i := 0; i < 6; i++ {
		packet.AddBuffer(t.Pokes[i].WritePacket().GetBufferSlice())
	}
	
	return packet
}