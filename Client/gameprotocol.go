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
	"fmt"
	punet "network"
)

type PU_GameProtocol struct {
}

func NewGameProtocol() *PU_GameProtocol {
	return &PU_GameProtocol{}
}

func (p *PU_GameProtocol) ProcessPacket(_packet *punet.Packet) {
	header := _packet.ReadUint8()
	switch header {
		case punet.HEADER_LOGIN:
			p.ReceiveLoginStatus(_packet)
	}
}

func (p *PU_GameProtocol) ReceiveLoginStatus(_packet *punet.Packet) {
	g_conn.loginStatus = int(_packet.ReadUint8())
	fmt.Printf("Loginstatus: %v\n", g_conn.loginStatus)
}

func (p *PU_GameProtocol) SendLogin(_username string, _password string) {
	message := NewLoginMessage()
	message.username = _username
	message.password = _password
	message.version = CLIENT_VERSION
	g_conn.SendMessage(message)
} 
