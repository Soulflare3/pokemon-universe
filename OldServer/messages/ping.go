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

//send a ping back to the server to let it know we're alive
type PU_Message_Ping struct {
}

func NewPingMessage() *PU_Message_Ping {
	return &PU_Message_Ping{}
}

func (m *PU_Message_Ping) WritePacket() (*punet.Packet, os.Error) {
	packet := punet.NewPacket()
	packet.AddUint8(punet.HEADER_PING)
	return packet, nil
}

