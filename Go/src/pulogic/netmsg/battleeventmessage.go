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
	pnet "network"
)

type BattleEventMessage struct {
	EventType	uint32
	
	PokemonId	uint32
	MoveSlotId	uint32
	NewPP		uint8
	NewHP		uint16
	Text		string
}

func NewBattleEventMessage(_eventType uint32) *BattleEventMessage {
	return &BattleEventMessage { EventType: _eventType }
}

// GetHeader returns the header value of this message
func (m *BattleEventMessage) GetHeader() uint8 {
	return pnet.HEADER_BATTLEMESSAGE
}

// WritePacket write the needed object data to a Packet and returns it
func (m *BattleEventMessage) WritePacket() pnet.IPacket {
	packet := pnet.NewPacketExt(m.GetHeader())
	packet.AddUint32(m.EventType)
	
	switch m.EventType {
		case pnet.BATTLEEVENT_TEXT:
			packet = m.writeMessage(packet)
		case pnet.BATTLEEVENT_CHANGEPP:
			packet = m.writeChangePP(packet)
		case pnet.BATTLEEVENT_CHANGEHP:
			packet = m.writeChangeHP(packet)
	}
	
	return packet
}

func (m *BattleEventMessage) writeChangePP(_packet pnet.IPacket) pnet.IPacket {
	_packet.AddUint32(m.PokemonId)
	_packet.AddUint32(m.MoveSlotId)
	_packet.AddUint8(m.NewPP)
	
	return _packet
}

func (m *BattleEventMessage) writeChangeHP(_packet pnet.IPacket) pnet.IPacket {
	_packet.AddUint32(m.PokemonId)
	_packet.AddUint16(m.NewHP)
	
	return _packet
}

func (m *BattleEventMessage) writeMessage(_packet pnet.IPacket) pnet.IPacket {
	_packet.AddString(m.Text)
	
	return _packet
}