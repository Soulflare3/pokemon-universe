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

type FullPlayerInfo struct {
	Team *PlayerTeam
	IsDefault bool
	
	ladderEnabled bool
	showTeam bool
	nameColor *QColor
}

func NewFullPlayerInfo(_team *PlayerTeam, _ladderEnabled, _showTeam bool) *FullPlayerInfo {
	fullPlayerInfo := &FullPlayerInfo { Team: _team,
										IsDefault: true,
										ladderEnabled: _ladderEnabled,
										showTeam: _showTeam }
	return fullPlayerInfo
}

func NewFullPlayerInfoFromPacket(_packet *pnet.QTPacket) *FullPlayerInfo {
	fullPlayerInfo := &FullPlayerInfo{}
	fullPlayerInfo.Team = NewPlayerTeamFromPacket(_packet)
	fullPlayerInfo.IsDefault = true
	fullPlayerInfo.ladderEnabled = _packet.ReadBool()
	fullPlayerInfo.nameColor = NewQColorFromPacket(_packet)
	
	return fullPlayerInfo
}

func (p *FullPlayerInfo) Nick() string {
	return p.Team.Nick
}

func (p *FullPlayerInfo) WritePacket() pnet.IPacket {
	packet := pnet.NewQTPacket()
	packet.AddBuffer(p.Team.WritePacket().GetBufferSlice())
	packet.AddBool(p.ladderEnabled)
	packet.AddBool(p.showTeam)
	packet.AddBuffer(p.nameColor.WritePacket().GetBufferSlice())

	return packet
}