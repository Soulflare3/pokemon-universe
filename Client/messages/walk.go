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

//receive notification of a moving creature
type PU_Message_CreatureMove struct {
	creature ICreature
	fromTile *PU_Tile
	toTile *PU_Tile
}

func NewCreatureMoveMessage(_packet *punet.Packet) *PU_Message_CreatureMove {
	msg := &PU_Message_CreatureMove{}
	msg.ReadPacket(_packet)
	return msg
}

func (m *PU_Message_CreatureMove) ReadPacket(_packet *punet.Packet) os.Error {
	m.creature =  g_map.GetCreatureByID(_packet.ReadUint32())
	fromX := int16(_packet.ReadUint16())
	fromY := int16(_packet.ReadUint16())
	toX := int16(_packet.ReadUint16())
	toY := int16(_packet.ReadUint16())
	m.fromTile = g_map.GetTile(int(fromX), int(fromY))
	m.toTile = g_map.GetTile(int(toX), int(toY))
	
	if m.creature != nil {
		m.creature.ReceiveWalk(m.fromTile, m.toTile)
	} else {
	}
	return nil
}

//tell the server we're moving
type PU_Message_Move struct {
	direction int
	requestTiles bool
}

func NewMoveMessage() *PU_Message_Move {
	return &PU_Message_Move{}
}

func (m *PU_Message_Move) WritePacket() (*punet.Packet, os.Error) {
	packet := punet.NewPacketExt(0xB1) //temporarily not using a header const from punet because this might change
	packet.AddUint16(uint16(m.direction))
	tiles := 0
	if m.requestTiles {
		tiles = 1
	}
	packet.AddUint16(uint16(tiles))		
	return packet, nil
}

