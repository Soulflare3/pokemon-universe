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
	"fmt"
	pnet "network"
)

type POClient struct {
	player *Player
	socket *POClientSocket
	
	meLoginPlayer *FullPlayerInfo
	mePlayer *PlayerInfo
}

func NewPOClient(_player *Player) (*POClient, os.Error) {
	poClient := POClient{ player: _player }
	
	///
	// TODO: Convert Player object to FullPlayerInfo
	//
	
	return &poClient, nil
}

func (c *POClient) Connect() {
	c.socket = NewPOClientSocket(c)
	c.socket.Connect("127.0.0.1", 5080) // TODO: Put this in server config
}

func (c *POClient) ProcessPacket(_packet *pnet.QTPacket) {
	header := int(_packet.ReadUint8())
	switch header {
		case COMMAND_ChallengeStuff:
			// TODO
		default:
			fmt.Printf("UNIMPLEMENTED PACKET: %v\n", header)
	}
}