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

type PlayerTeam struct {
	Nick string
	Info string
	LoseMessage string
	WinMessage string
	DefaultTier string
	Tier string
	Team *Team
	
	avatar int
}

func NewPlayerTeam() *PlayerTeam {
	return &PlayerTeam{}
}

func NewPlayerTeamFromPlayer(_player *Player) *PlayerTeam {
	playerTeam := PlayerTeam{}
	playerTeam.Nick = _player.GetName()
	playerTeam.Info = "PU Client"
	playerTeam.LoseMessage = "bgnt"
	playerTeam.WinMessage = "gg"
	playerTeam.avatar = 0
	playerTeam.DefaultTier = ""
	playerTeam.Team = NewTeamFromParty(_player.PokemonParty)
	
	return &playerTeam
}

func NewPlayerTeamFromPacket(_packet *pnet.QTPacket) *PlayerTeam {
	playerTeam := PlayerTeam{}
	playerTeam.Nick = _packet.ReadString()
	playerTeam.Info = _packet.ReadString()
	playerTeam.LoseMessage = _packet.ReadString()
	playerTeam.WinMessage = _packet.ReadString()
	playerTeam.avatar = int(_packet.ReadUint16())
	playerTeam.DefaultTier = _packet.ReadString()
	playerTeam.Team = NewTeamFromPacket(_packet)
	
	return &playerTeam
}

func (p *PlayerTeam) WritePacket() pnet.IPacket {
	packet := pnet.NewQTPacket()
	packet.AddString(p.Nick)
	packet.AddString(p.Info)
	packet.AddString(p.LoseMessage)
	packet.AddString(p.WinMessage)
	packet.AddUint16(uint16(p.avatar))
	packet.AddString(p.DefaultTier)
	packet.AddBuffer(p.Team.WritePacket().GetBufferSlice())
	
	return packet
}