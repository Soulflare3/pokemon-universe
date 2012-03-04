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

const (
	LOGINSTATUS_IDLE = 0
	LOGINSTATUS_WRONGACCOUNT = 1
	LOGINSTATUS_SERVERERROR = 2
	LOGINSTATUS_DATABASEERROR = 3
	LOGINSTATUS_ALREADYLOGGEDIN = 4
	LOGINSTATUS_READY = 5
	LOGINSTATUS_CHARBANNED = 6
	LOGINSTATUS_SERVERCLOSED = 7
	LOGINSTATUS_WRONGVERSION = 8
	LOGINSTATUS_FAILPROFILELOAD = 9
)

type LoginMessage struct {
	// Receive
	Username 		string
	Password 		string
	ClientVersion 	int // uint16
	
	// Send
	Status			int // uint32
}

// GetHeader returns the header value of this message
func (m *LoginMessage) GetHeader() uint8 {
	return pnet.HEADER_LOGIN
}

// ReadPacket reads all data from a packet and puts it in the object
func (m *LoginMessage) ReadPacket(_packet pnet.IPacket) error {
	m.Username = _packet.ReadString()
	m.Password = _packet.ReadString()
	m.ClientVersion = int(_packet.ReadUint16())
	
	return nil
}

// WritePacket write the needed object data to a Packet and returns it
func (m *LoginMessage) WritePacket() pnet.IPacket {
	packet := pnet.NewPacketExt(m.GetHeader())
	packet.AddUint32(uint32(m.Status))
	
	return packet
}
