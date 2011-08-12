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
)

type PU_GameProtocol struct {

}

func NewGameProtocol() *PU_GameProtocol {
	return &PU_GameProtocol{}
}

func (p *PU_GameProtocol) ProcessMessage(_message *punet.Message) {
	switch _message.Header {
	case punet.HEADER_PING:
		p.Receive_Ping()

	case punet.HEADER_LOGIN:
		p.Receive_LoginStatus(_message)

	case punet.HEADER_IDENTITY:
		p.Receive_PlayerData(_message)

	case punet.HEADER_TILES:
		p.Receive_Tiles(_message)

	case punet.HEADER_ADDCREATURE:
		p.Receive_AddCreature(_message)

	case punet.HEADER_REMOVECREATURE:
		p.Receive_RemoveCreature(_message)

	case punet.HEADER_WALK:
		p.Receive_CreatureWalk(_message)

	case punet.HEADER_TURN:
		p.Receive_CreatureTurn(_message)

	case punet.HEADER_WARP:
		p.Receive_Warp(_message)

	case punet.HEADER_REFRESHCOMPLETE:
		p.Receive_TilesRefreshed()

	case punet.HEADER_CHAT:
		p.Receive_CreatureChat(_message)
	}
}

func (p *PU_GameProtocol) ProcessPacket(_packet *punet.Packet) {
}

func (p *PU_GameProtocol) SendLogin(_username string, _password string) {
	p.Send_Login(_username, _password)
}
