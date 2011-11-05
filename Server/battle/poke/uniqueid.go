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

type UniqueId struct {
	PokeNum int
	SubNum int
}

func NewUniqueId() *UniqueId {
	return &UniqueId { PokeNum: 173, SubNum: 0 }
}

func NewUniqueIdExt(_pokeNum, _subNum int) *UniqueId {
	uniqueId := UniqueId { PokeNum: _pokeNum,
							SubNum: _subNum }
	return &uniqueId
}

func NewUniqueIdFromPacket(_packet *pnet.QTPacket) *UniqueId {
	uniqueId := UniqueId { PokeNum: int(_packet.ReadUint16()),
							SubNum: int(_packet.ReadUint8()) }
	return &uniqueId
}

func (u *UniqueId) WritePacket() pnet.IPacket {
	packet := pnet.NewQTPacket()
	packet.AddUint16(uint16(u.PokeNum))
	packet.AddUint8(uint8(u.SubNum))
	return packet
}