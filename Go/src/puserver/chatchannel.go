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

type IChatChannel interface {
	ChatChannelType() int
	
	GetName() string
	GetId() int
	GetOwner() uint64
	
	AddUser(_player *Player) bool
	RemoveUser(_player *Player, _sendCloseChannel bool) bool
	Talk(_fromPlayer *Player, _type int, _text string, _time int) bool
}

type ChatChannel struct {
	Id		int
	Users	ChatUsersMap
	Name	string
}

func NewChatChannel(_id int, _name string) *ChatChannel {
	return &ChatChannel{ Id: _id,
						 Name: _name,
						 Users: make(ChatUsersMap) }
}

func (c *ChatChannel) ChatChannelType() int {
	return 1;
}

func (c *ChatChannel) GetName() string {
	return c.Name
}

func (c *ChatChannel) GetOwner() uint64 {
	return 0
}

func (c *ChatChannel) GetId() int {
	return c.Id
}

func (c *ChatChannel) AddUser(_player *Player) bool {
	_, found := c.Users[_player.GetUID()]
	if !found {
		c.Users[_player.GetUID()] = _player
	}
	
	return !found
}

func (c *ChatChannel) RemoveUser(_player *Player, _sendCloseChannel bool) bool {
	_, found := c.Users[_player.GetUID()]
	if found {
		delete(c.Users, _player.GetUID())
		
		if _sendCloseChannel {
			_player.sendClosePrivateChat(c.Id)
		}
	}
	
	return !found
}

func (c *ChatChannel) Talk(_fromPlayer *Player, _type int, _text string, _time int) (success bool) {
	success = false
	
	// Can't speak to a channel you're not connected to
	if _, found := c.Users[_fromPlayer.GetUID()]; found {
		for _, value := range(c.Users) {
			value.sendToChannel(_fromPlayer, _type, _text, c.Id, _time)
		}
		
		success = true
	}
	
	return
}

func (c *ChatChannel) SendInfo(_type int, _text string, _time int) {
	for _, value := range(c.Users) {
		value.sendToChannel(nil, _type, _text, c.Id, _time)
	}
}