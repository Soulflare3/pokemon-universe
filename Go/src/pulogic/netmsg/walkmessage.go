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
package netmsg

import (
	pnet "nonamelib/network"
	pos "nonamelib/pos"
	pul "pulogic"
)

type WalkMessage struct {
	// Receive
	Creature 	pul.ICreature
	Direction	int // uint16
	SendMap		bool
	
	// Send
	From		pos.Position
	To			pos.Position
}

func NewWalkMessage(_creature pul.ICreature) *WalkMessage {
	message := &WalkMessage{}
	message.Creature = _creature
	
	return message
}

// GetHeader returns the header value of this message
func (m *WalkMessage) GetHeader() uint8 {
	return pnet.HEADER_WALK
}

// ReadPacket reads all data from a packet and puts it in the object
func (m *WalkMessage) ReadPacket(_packet pnet.IPacket) (err error) {
	direction, err := _packet.ReadUint16()
	if err != nil {
		return
	}
	m.Direction = int(direction)
	
	sendMap, err := _packet.ReadUint16()
	if err != nil {
		return err
	}
	
	if sendMap == 1 {
		m.SendMap = true
	}
	
	return nil
}

func (m *WalkMessage) AddPositions(_from pos.Position, _to pos.Position) {
	m.From	= _from
	m.To	= _to
}

// WritePacket write the needed object data to a Packet and returns it
func (m *WalkMessage) WritePacket() pnet.IPacket {
	packet := pnet.NewPacketExt(m.GetHeader())
	packet.AddUint64(m.Creature.GetUID())
	packet.AddUint16(uint16(m.From.X))
	packet.AddUint16(uint16(m.From.Y))
	packet.AddUint16(uint16(m.To.X))
	packet.AddUint16(uint16(m.To.Y))
	
	return packet
}