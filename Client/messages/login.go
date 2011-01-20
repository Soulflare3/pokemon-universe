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

type PU_Message_Login struct {
	username string
	password string
	version uint16
}

func NewLoginMessage() *PU_Message_Login {
	return &PU_Message_Login{}
}

func NewLoginMessageExt(_username string, _password string, _version uint16) *PU_Message_Login {
	return &PU_Message_Login{username : _username,
							password : _password,
							version : _version}
}

func (m *PU_Message_Login) WritePacket() (*punet.Packet, os.Error) {
	packet := punet.NewPacketExt(punet.HEADER_LOGIN)
	packet.AddString(m.username)
	packet.AddString(m.password)
	packet.AddUint16(m.version)
	return packet, nil
}

//This message requests tiles and the player's identify from the server after logging in
type PU_Message_LoginRequest struct {
}

func NewLoginRequestMessage() *PU_Message_LoginRequest {
	return &PU_Message_LoginRequest{}
}

func (m *PU_Message_LoginRequest) WritePacket() (*punet.Packet, os.Error) {
	packet := punet.NewPacket()
	packet.AddUint8(punet.HEADER_LOGIN)
	return packet, nil
}
