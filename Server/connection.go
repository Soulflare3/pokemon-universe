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
	pos "position"
)

type Connection struct {
	Socket net.Conn
	IsOpen bool
	Owner  *Player
}

func NewConnection(_socket net.Conn) *Connection {
	return &Connection{Socket: _socket}
}

func (c *Connection) HandleConnection() {
	c.IsOpen = true

	for {
		var headerbuffer [2]uint8
		recv, err := c.Socket.Read(headerbuffer[0:])
		if err != nil || recv == 0 {
			g_logger.Printf("Error while reading socket: %v", err)
			break
		}

		packet := pnet.NewPacket()
		copy(packet.Buffer[0:2], headerbuffer[0:2])
		packet.GetHeader()

		databuffer := make([]uint8, packet.MsgSize)
		recv, err = c.Socket.Read(databuffer[0:])
		if recv == 0 {
			continue
		} else if err != nil {
			g_logger.Printf("Connection read error: %v", err)
			continue
		}

		copy(packet.Buffer[2:], databuffer[:])
		c.ProcessPacket(packet)
	}

	c.IsOpen = false
	g_logger.Println("Connection closed")
}

func (c *Connection) ProcessPacket(_packet *pnet.Packet) {
	header := _packet.ReadUint8()
	switch header {
	case pnet.HEADER_LOGIN:
		c.SendPlayerData()
	}
}

func (c *Connection) SendMessage(_message pnet.INetMessageWriter) {
	packet, _ := _message.WritePacket()
	packet.SetHeader()

	c.Socket.Write(packet.Buffer[0:packet.MsgSize])
}

/*****************
BATCH SEND FUNCTIONS
******************/
func (c *Connection) SendPlayerData() {
	playerData := &SendPlayerData{}
	playerData.UID		= c.Owner.GetUID()
	playerData.Name		= c.Owner.GetName()
	playerData.Position	= c.Owner.GetPosition()
	playerData.Direction = c.Owner.Direction
	playerData.Money	= c.Owner.Money
	playerData.Outfit	= c.Owner.Outfit
	c.SendMessage(playerData)
	
	//ToDo: Send PkMn
	
	//ToDo: Send items
	
	// Send map
	c.GenerateMapData(DIR_NULL)
	
	// ready
	readyMessage := &LoginMessage{}
	readyMessage.Status = LOGINSTATUS_READY
	c.SendMessage(readyMessage)
}

func (c *Connection) GenerateMapData(_direction int) {
	xMin := 1
	xMax := CLIENT_VIEWPORT.X
	yMin := 1
	yMax := CLIENT_VIEWPORT.Y
	
	if _direction != DIR_NULL {
		switch _direction {
		case DIR_NORTH:
			yMax = 1
		case DIR_EAST:
			xMin = CLIENT_VIEWPORT.X
		case DIR_SOUTH:
			yMin = CLIENT_VIEWPORT.Y
		case DIR_WEST:
			xMax = 1
		}
	}
	
	position := c.Owner.GetPosition()
	positionX := (position.X - CLIENT_VIEWPORT_CENTER.X)
	positionY := (position.Y - CLIENT_VIEWPORT_CENTER.Y)
	
	tilesMessage := NewSendTilesMessage()
	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			index := pos.Hash(positionX + x, positionY + y, position.Z)
			tilesMessage.AddTile(index)
		}
		
		if _direction == DIR_NULL {
			c.SendMessage(tilesMessage)
		}
	}
	
	if _direction != DIR_NULL {
		c.SendMessage(tilesMessage)
	}
}

