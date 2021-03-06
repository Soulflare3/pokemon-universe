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
	pul "pulogic"
)

type CreatureTurnMessage struct {
	Creature pul.ICreature
	Direction int // uint16
}

func NewCreatureTurnMessage(_creature pul.ICreature) *CreatureTurnMessage {
	return &CreatureTurnMessage { Creature: _creature }
}

// GetHeader returns the header value of this message
func (m *CreatureTurnMessage) GetHeader() uint8 {
	return pnet.HEADER_TURN
}

func (m *CreatureTurnMessage) AddDirection(_dir int) {
	m.Direction = _dir
}

func (m *CreatureTurnMessage) ReadPacket(_packet pnet.IPacket) error {
	direction, err := _packet.ReadUint16()
	if err != nil {
		return err
	}
	m.Direction = int(direction)
	
	return nil
}

// WritePacket write the needed object data to a Packet and returns it
func (m *CreatureTurnMessage) WritePacket() pnet.IPacket {
	packet := pnet.NewPacketExt(m.GetHeader())
	packet.AddUint64(m.Creature.GetUID())
	packet.AddUint16(uint16(m.Direction))
	
	return packet
}
