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
	pnet "network" // PU Network packet
	pos "position"
	"websocket"
)

type Connection struct {
	Socket     *websocket.Conn
	Tranceiver *pnet.Tranceiver
	IsOpen     bool
	Owner      *Player
}

func NewConnection(_socket *websocket.Conn) *Connection {
	return &Connection{Socket: _socket}
}

func (c *Connection) HandleConnection() {
	c.IsOpen = true
	
	for {
		packet := pnet.NewPacket()
		var buffer []uint8
		err := websocket.Message.Receive(c.Socket, &buffer)
		if err == nil {
			copy(packet.Buffer[0:len(buffer)], buffer[0:len(buffer)])
			packet.GetHeader()
			c.ProcessPacket(packet)
		} else {
			println(err.Error())
			break;
		}
	}

	c.IsOpen = false
	c.Owner.removeConnection()
}

func (c *Connection) ProcessPacket(_packet *pnet.Packet) {
	header := _packet.ReadUint8()
	switch header {
	case pnet.HEADER_LOGIN:
		c.SendPlayerData()

	case pnet.HEADER_WALK:
		c.ReceivePlayerWalk(_packet)

	case pnet.HEADER_TURN:
		c.ReceivePlayerTurn(_packet)

	case pnet.HEADER_REFRESHWORLD:
		c.ReceiveRefreshWorld()
		
	case pnet.HEADER_CHAT:
		c.ReceiveChat(_packet)
	}	
}

func (c *Connection) SendMessage(_message pnet.INetMessageWriter) {
	packet := _message.WritePacket()
	packet.SetHeader()

	buffer := packet.GetBuffer()
	data := buffer[0:packet.GetMsgSize()]

	websocket.Message.Send(c.Socket, data)
}

// ------------------------------------------------------ //
//                     SEND
// ------------------------------------------------------ //
func (c *Connection) SendPlayerData() {
	playerData := &SendPlayerData{}
	playerData.UID			= c.Owner.GetUID()
	playerData.Name			= c.Owner.GetName()
	playerData.Position		= c.Owner.GetPosition()
	playerData.Direction 	= c.Owner.Direction
	playerData.Money		= c.Owner.Money
	playerData.Outfit		= c.Owner.Outfit
	c.SendMessage(playerData)
	
	//ToDo: Send PkMn
	
	//ToDo: Send items
	
	// Send map
	c.SendMapData(DIR_NULL, c.Owner.GetPosition())
	
	// ready
	readyMessage := &LoginMessage{}
	readyMessage.Status = LOGINSTATUS_READY
	c.SendMessage(readyMessage)
}

func (c *Connection) SendCancel(_message string) {

}

func (c *Connection) SendMapData(_direction int, _centerPosition pos.Position) {
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
	
	// Top-left coordinates
	positionX 	:= (_centerPosition.X - CLIENT_VIEWPORT_CENTER.X)
	positionY 	:= (_centerPosition.Y - CLIENT_VIEWPORT_CENTER.Y)
	z			:= _centerPosition.Z
	
	tilesMessage := NewSendTilesMessage()
	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			index := pos.Hash(positionX + x, positionY + y, z)
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

func (c *Connection) SendCreatureMove(_creature ICreature, _from *Tile, _to *Tile) {
	msg := NewWalkMessage(_creature)
	msg.AddPositions(_from.Position, _to.Position)
	c.SendMessage(msg)
}

func (c *Connection) SendCreatureTurn(_creature ICreature, direction int) {
	msg := NewCreatureTurnMessage(_creature)
	msg.AddDirection(direction)
	c.SendMessage(msg)
}

func (c *Connection) SendCreatureAdd(_creature ICreature) {
	if _creature.GetUID() != c.Owner.GetUID() {
		msg := NewCreatureAddMessage(_creature)
		c.SendMessage(msg)
	}
}

func (c *Connection) SendCreatureRemove(_creature ICreature) {
	if _creature.GetUID() != c.Owner.GetUID() {
		msg := NewCreatureRemoveMessage(_creature)
		c.SendMessage(msg)
	}
}

func (c *Connection) SendPlayerWarp(_position pos.Position) {
	msg := NewPlayerWarpMessage()
	msg.AddDestination(_position)
	c.SendMessage(msg)
}

func (c *Connection) SendCreatureSay(_creature ICreature, _speakType int, _text string, _channelId int, _time int) {
	msg := NewChatMessageExt(_creature, _speakType, _text, _channelId, _time)
	c.SendMessage(msg)
}

// ------------------------------------------------------ //
//                     RECEIVE
// ------------------------------------------------------ //
func (c *Connection) ReceivePlayerWalk(_packet *pnet.Packet) {
	msg := NewWalkMessage(c.Owner)
	msg.ReadPacket(_packet)
}

func (c *Connection) ReceivePlayerTurn(_packet *pnet.Packet) {
	msg := NewCreatureTurnMessage(c.Owner)
	msg.ReadPacket(_packet)
}

func (c *Connection) ReceiveRefreshWorld() {
	// Send whole screen
	c.SendMapData(DIR_NULL, c.Owner.GetPosition())
	
	msg := NewRefreshWorldMessage()
	c.SendMessage(msg)
}

func (c *Connection) ReceiveChat(_packet *pnet.Packet) {
	msg := NewChatMessage(c.Owner)
	msg.ReadPacket(_packet)
}