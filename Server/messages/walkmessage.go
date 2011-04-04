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
	"os"
	pnet "network"
	pos "position"
)

type WalkMessage struct {
	// Receive
	creature 	ICreature
	direction	uint16
	sendMap		bool
	
	// Send
	from		pos.Position
	to			pos.Position
}

func NewWalkMessage(_creature ICreature) *WalkMessage {
	message := &WalkMessage{}
	message.creature = _creature
	
	return message
}

// GetHeader returns the header value of this message
func (m *WalkMessage) GetHeader() uint8 {
	return pnet.HEADER_WALK
}

// ReadPacket reads all data from a packet and puts it in the object
func (m *WalkMessage) ReadPacket(_packet *pnet.Packet) os.Error {
	m.direction = _packet.ReadUint16()
	if _packet.ReadUint16() == 1 {
		m.sendMap = true
	}
	
	g_game.OnPlayerMove(m.creature, m.direction, m.sendMap)
	
	return nil
}

func (m *WalkMessage) AddPositions(_from pos.Position, _to pos.Position) {
	m.from	= _from
	m.to	= _to
}

// WritePacket write the needed object data to a Packet and returns it
func (m *WalkMessage) WritePacket() (*pnet.Packet, os.Error) {
	packet := pnet.NewPacketExt(m.GetHeader())
	packet.AddUint64(m.creature.GetUID())
	packet.AddUint16(uint16(m.from.X))
	packet.AddUint16(uint16(m.from.Y))
	packet.AddUint16(uint16(m.to.X))
	packet.AddUint16(uint16(m.to.Y))
	
	return packet, nil
}
