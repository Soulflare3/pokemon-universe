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
	"net"
	pnet "network" // PU Network packet
	"fmt"
)

type Connection struct {
	Socket net.Conn
	Tranceiver *pnet.Tranceiver
	IsOpen bool
	Owner  *Player
}

func NewConnection(_socket net.Conn) *Connection {
	return &Connection{Socket: _socket, Tranceiver : pnet.NewTranceiver(_socket)}
}

func (c *Connection) HandleConnection() {
	c.IsOpen = true

	for {
		if message, err := c.Tranceiver.Receive(); err != "" {
			fmt.Printf("Error receiving message: %s\n", err)
			break
		} else {
			c.ProcessMessage(message)
		}
	}

	c.IsOpen = false
	c.Owner.removeConnection()
}

func (c *Connection) ProcessMessage(_message *pnet.Message) {
	switch _message.Header {
	case pnet.HEADER_LOGIN:
		c.Send_PlayerData()
		
	case pnet.HEADER_WALK:
		c.Receive_Walk(_message)
		
	case pnet.HEADER_TURN:
		c.Receive_Turn(_message)
		
	case pnet.HEADER_REFRESHWORLD:
		c.Receive_RefreshWorld()
	}
}

func (c *Connection) SendMessage(_message *pnet.Message) {
	c.Tranceiver.Send(_message)
}

func (c *Connection) SendPacket(_message pnet.INetMessageWriter) {
	packet, _ := _message.WritePacket()
	packet.SetHeader()

	c.Socket.Write(packet.Buffer[0:packet.MsgSize])
}
