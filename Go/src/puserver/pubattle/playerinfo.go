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
package pubattle

import (
	pnet "network"
)

type PlayerInfo struct {
	Id int
	Nick string
	Info string
	Auth int
	Rating int
	Pokes []*UniqueId
	Tier string
	
	flags int
	rating int
	avatar int
	color *QColor
	gen int
}

func NewPlayerInfo() *PlayerInfo {
	return &PlayerInfo{ Pokes: make([]*UniqueId, 6) }
}

func NewPlayerInfoFromFullPlayerInfo(_info *FullPlayerInfo) *PlayerInfo {
	return &PlayerInfo{ Nick: _info.Nick() }
}

func NewPlayerInfoFromPacket(_packet *pnet.QTPacket) *PlayerInfo {
	playerInfo := NewPlayerInfo();
	playerInfo.Id = int(_packet.ReadUint32())
	playerInfo.Nick = _packet.ReadString()
	playerInfo.Info = _packet.ReadString()
	playerInfo.Auth = int(_packet.ReadUint8())
	playerInfo.flags = int(_packet.ReadUint8())
	playerInfo.rating = int(_packet.ReadUint16())
	
	for i := 0; i < 6; i++ {
		playerInfo.Pokes[i] = NewUniqueIdFromPacket(_packet)
	}
	
	playerInfo.avatar = int(_packet.ReadUint8())
	playerInfo.Tier = _packet.ReadString()
	playerInfo.color = NewQColorFromPacket(_packet)
	playerInfo.gen = int(_packet.ReadUint8())
	
	return playerInfo
}