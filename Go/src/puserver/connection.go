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
	pul "pulogic"
	pnetmsg "pulogic/netmsg"
	pnet "network" // PU Network packet
	pos "putools/pos"
	"net/websocket"
)

type Connection struct {
	Socket     *websocket.Conn

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
			// println(err.Error())
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
		
	case pnet.HEADER_FRIENDUPDATE:
		c.ReceiveFriendUpdate(_packet)
		
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
	playerData := &pnetmsg.SendPlayerData{}
	playerData.UID			= c.Owner.GetUID()
	playerData.Name			= c.Owner.GetName()
	playerData.Position		= c.Owner.GetPosition()
	playerData.Direction 	= c.Owner.Direction
	playerData.Money		= c.Owner.Money
	playerData.Outfit		= c.Owner.Outfit
	c.SendMessage(playerData)
	
	// Send PkMn
	pokemonData := &pnetmsg.SendPokemonData{}
	pokemonData.Pokemon = c.Owner.PokemonParty
	c.SendMessage(pokemonData)
	
	// TODO: Send items
	// ----
	
	// Send list to client
	c.SendFriendList(c.Owner.Friends)
	
	// Send map
	c.SendMapData(DIR_NULL, c.Owner.GetPosition())
	
	// ready
	readyMessage := &pnetmsg.LoginMessage{}
	readyMessage.Status = pnetmsg.LOGINSTATUS_READY
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
	
	tilesMessage := pnetmsg.NewSendTilesMessage()
	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			index := pos.Hash(positionX + x, positionY + y, z)
			if tile, found := g_map.GetTile(index); found {
				tilesMessage.AddTile(tile)
			}
		}
		
		if _direction == DIR_NULL {
			c.SendMessage(tilesMessage)
		}
	}
	
	if _direction != DIR_NULL {
		c.SendMessage(tilesMessage)
	}
}

func (c *Connection) SendCreatureMove(_creature pul.ICreature, _from *Tile, _to *Tile) {
	msg := pnetmsg.NewWalkMessage(_creature)
	msg.AddPositions(_from.Position, _to.Position)
	c.SendMessage(msg)
}

func (c *Connection) SendCreatureTurn(_creature pul.ICreature, direction int) {
	msg := pnetmsg.NewCreatureTurnMessage(_creature)
	msg.AddDirection(direction)
	c.SendMessage(msg)
}

func (c *Connection) SendCreatureAdd(_creature pul.ICreature) {
	if _creature.GetUID() != c.Owner.GetUID() {
		msg := pnetmsg.NewCreatureAddMessage(_creature)
		c.SendMessage(msg)
	}
}

func (c *Connection) SendCreatureRemove(_creature pul.ICreature) {
	if _creature.GetUID() != c.Owner.GetUID() {
		msg := pnetmsg.NewCreatureRemoveMessage(_creature)
		c.SendMessage(msg)
	}
}

func (c *Connection) SendPlayerWarp(_position pos.Position) {
	msg := pnetmsg.NewPlayerWarpMessage()
	msg.AddDestination(_position)
	c.SendMessage(msg)
}

func (c *Connection) SendCreatureSay(_creature pul.ICreature, _speakType int, _text string, _channelId int, _time int) {
	msg := pnetmsg.NewChatMessageExt(_creature, _speakType, _text, _channelId, _time)
	c.SendMessage(msg)
}

func (c *Connection) SendFriendList(_friends FriendList) {
	msg := pnetmsg.NewFriendListMessage()
	
	for _, friend := range(_friends) {
		if !friend.IsRemoved {
			msg.AddFriend(friend.Name, friend.Online)
		}
	}
	
}

func (c *Connection) SendFriendUpdate(_name string, _online bool) {
	msg := pnetmsg.NewFriendUpdateMessageExt(_name, _online, false)
	c.SendMessage(msg)
}

func (c *Connection) SendFriendRemove(_name string) {
	msg := pnetmsg.NewFriendUpdateMessageExt(_name, false, true)
	c.SendMessage(msg)
}

// ------------------------------------------------------ //
//                     RECEIVE
// ------------------------------------------------------ //
func (c *Connection) ReceivePlayerWalk(_packet *pnet.Packet) {
	msg := pnetmsg.NewWalkMessage(c.Owner)
	msg.ReadPacket(_packet)
	
	g_game.OnPlayerMove(msg.Creature, msg.Direction, msg.SendMap)
}

func (c *Connection) ReceivePlayerTurn(_packet *pnet.Packet) {
	msg := pnetmsg.NewCreatureTurnMessage(c.Owner)
	msg.ReadPacket(_packet)
}

func (c *Connection) ReceiveRefreshWorld() {
	// Send whole screen
	c.SendMapData(DIR_NULL, c.Owner.GetPosition())
	
	msg := pnetmsg.NewRefreshWorldMessage()
	c.SendMessage(msg)
}

func (c *Connection) ReceiveChat(_packet *pnet.Packet) {
	msg := pnetmsg.NewChatMessage(c.Owner)
	msg.ReadPacket(_packet)
	
	g_game.OnPlayerSay(msg.From.(*Player), msg.ChannelId, msg.SpeakType, msg.Receiver, msg.Text)
}

func (c *Connection) ReceiveFriendUpdate(_packet *pnet.Packet) {
	msg := pnetmsg.NewFriendUpdateMessage()
	msg.ReadPacket(_packet)
	
	if msg.Removed == 1 {
		c.Owner.RemoveFriend(msg.Name)
	} else {
		c.Owner.AddFriend(msg.Name)
	}
}