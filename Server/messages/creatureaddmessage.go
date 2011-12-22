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

type CreatureAddMessage struct {
	creature ICreature
}

func NewCreatureAddMessage(_creature ICreature) *CreatureAddMessage {
	return &CreatureAddMessage { creature: _creature }
}

// GetHeader returns the header value of this message
func (m *CreatureAddMessage) GetHeader() uint8 {
	return pnet.HEADER_ADDCREATURE
}

// WritePacket write the needed object data to a Packet and returns it
func (m *CreatureAddMessage) WritePacket() pnet.IPacket {
	packet := pnet.NewPacketExt(m.GetHeader())
	packet.AddUint64(m.creature.GetUID())
	packet.AddString(m.creature.GetName())
	packet.AddUint16(uint16(m.creature.GetPosition().X))
	packet.AddUint16(uint16(m.creature.GetPosition().Y))
	packet.AddUint16(uint16(m.creature.GetDirection()))
	
	// Outfit
	packet.AddUint8(uint8(m.creature.GetOutfit().GetOutfitStyle(OUTFIT_UPPER)))
	packet.AddUint32(uint32(m.creature.GetOutfit().GetOutfitColour(OUTFIT_UPPER)))
	packet.AddUint8(uint8(m.creature.GetOutfit().GetOutfitStyle(OUTFIT_NEK)))
	packet.AddUint32(uint32(m.creature.GetOutfit().GetOutfitColour(OUTFIT_NEK)))
	packet.AddUint8(uint8(m.creature.GetOutfit().GetOutfitStyle(OUTFIT_HEAD)))
	packet.AddUint32(uint32(m.creature.GetOutfit().GetOutfitColour(OUTFIT_HEAD)))
	packet.AddUint8(uint8(m.creature.GetOutfit().GetOutfitStyle(OUTFIT_UPPER)))
	packet.AddUint32(uint32(m.creature.GetOutfit().GetOutfitColour(OUTFIT_UPPER)))
	packet.AddUint8(uint8(m.creature.GetOutfit().GetOutfitStyle(OUTFIT_FEET)))
	packet.AddUint32(uint32(m.creature.GetOutfit().GetOutfitColour(OUTFIT_FEET)))
	packet.AddUint8(uint8(m.creature.GetOutfit().GetOutfitStyle(OUTFIT_LOWER)))
	packet.AddUint32(uint32(m.creature.GetOutfit().GetOutfitColour(OUTFIT_LOWER)))
	
	return packet
}
