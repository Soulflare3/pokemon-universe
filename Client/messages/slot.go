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

type PU_Message_SlotChange struct {
	oldSlot int
	newSlot int
}

func NewSlotChangeMessage() *PU_Message_SlotChange {
	return &PU_Message_SlotChange{}
}

func (m *PU_Message_SlotChange) WritePacket() (*punet.Packet, os.Error) {
	packet := punet.NewPacketExt(0xD2) //temporarily not using a header const from punet because this might change
	packet.AddUint16(uint16(m.oldSlot))
	packet.AddUint16(uint16(m.newSlot))		
	return packet, nil
}

