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

type PU_Message_Identity struct {
	player *PU_Player
}

func NewIdentityMessage(_packet *punet.Packet) *PU_Message_Identity {
	msg := &PU_Message_Identity{}
	msg.ReadPacket(_packet)
	return msg
}

func (m *PU_Message_Identity) ReadPacket(_packet *punet.Packet) os.Error {
	m.player = NewPlayer(0)
	m.player.id = _packet.ReadUint32()
	m.player.name = _packet.ReadString()
	m.player.x = int16(_packet.ReadUint16())
	m.player.y = int16(_packet.ReadUint16())
	m.player.direction = int(_packet.ReadUint16())
	m.player.money = _packet.ReadUint32()
	
	for part := BODY_UPPER; part <= BODY_LOWER; part++ {
		m.player.bodyParts[part].id = int(_packet.ReadUint8())
		color := _packet.ReadUint32()
		red := uint8(color >> 16)
		green := uint8(color >> 8)
		blue := uint8(color)
		m.player.bodyParts[part].SetColor(int(red), int(green), int(blue))
	}
	return nil
}

