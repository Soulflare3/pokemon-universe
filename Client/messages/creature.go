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

type PU_Message_RemoveCreature struct {

}

func NewRemoveCreatureMessage(_packet *punet.Packet) *PU_Message_RemoveCreature {
	msg := &PU_Message_RemoveCreature{}
	msg.ReadPacket(_packet)
	return msg
}

func (m *PU_Message_RemoveCreature) ReadPacket(_packet *punet.Packet) os.Error {
	id := _packet.ReadUint32()
	
	creature := g_map.GetCreatureByID(id)
	if creature != nil {
		g_map.RemoveCreature(creature)
	}

	return nil
}

