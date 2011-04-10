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

//receive notification of a turning creature
type PU_Message_CreatureTurn struct {
	creature ICreature
	direction int
}

func NewCreatureTurnMessage(_packet *punet.Packet) *PU_Message_CreatureTurn {
	msg := &PU_Message_CreatureTurn{}
	msg.ReadPacket(_packet)
	return msg
}

func (m *PU_Message_CreatureTurn) ReadPacket(_packet *punet.Packet) os.Error {
	m.creature =  g_map.GetCreatureByID(_packet.ReadUint32())
	m.direction = int(_packet.ReadUint16())
	if m.creature != nil && m.creature != g_game.self {
		m.creature.SetDirection(m.direction)
	}
	return nil
}

//tell the server we're turning
type PU_Message_Turn struct {
	direction int
}

func NewTurnMessage() *PU_Message_Turn {
	return &PU_Message_Turn{}
}

func (m *PU_Message_Turn) WritePacket() (*punet.Packet, os.Error) {
	packet := punet.NewPacketExt(0xB2)
	packet.AddUint16(uint16(m.direction))		
	return packet, nil
}

