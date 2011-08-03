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
)

const (
	LOGINSTATUS_IDLE            = 0
	LOGINSTATUS_WRONGACCOUNT    = 1
	LOGINSTATUS_SERVERERROR     = 2
	LOGINSTATUS_DATABASEERROR   = 3
	LOGINSTATUS_ALREADYLOGGEDIN = 4
	LOGINSTATUS_READY           = 5
	LOGINSTATUS_CHARBANNED      = 6
	LOGINSTATUS_SERVERCLOSED    = 7
	LOGINSTATUS_WRONGVERSION    = 8
	LOGINSTATUS_FAILPROFILELOAD = 9
)

type LoginMessage struct {
	// Receive
	Username      string
	Password      string
	ClientVersion uint16

	// Send
	Status  uint32
	Servers map[string]ServerInfo
}

func NewLoginMessage() *LoginMessage {
	return &LoginMessage{Servers: make(map[string]ServerInfo)}
}

// GetHeader returns the header value of this message
func (m *LoginMessage) GetHeader() uint8 {
	return pnet.HEADER_LOGIN
}

// ReadPacket reads all data from a packet and puts it in the object
func (m *LoginMessage) ReadPacket(_packet *pnet.Packet) os.Error {
	m.Username = _packet.ReadString()
	m.Password = _packet.ReadString()
	m.ClientVersion = _packet.ReadUint16()

	return nil
}

// WritePacket write the needed object data to a Packet and returns it
func (m *LoginMessage) WritePacket() (*pnet.Packet, os.Error) {
	packet := pnet.NewPacketExt(m.GetHeader())
	packet.AddUint32(m.Status)

	if m.Status == LOGINSTATUS_READY {
		packet.AddUint32(uint32(len(m.Servers)))

		for name, server := range m.Servers {
			packet.AddString(name)
			packet.AddString(server.Name)
			packet.AddString(server.Ip)
			packet.AddUint8(server.Online)
		}
	}

	return packet, nil
}
