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
	"container/list" 
	pnet "network"
)

type ChatUsersMap map[uint64]*Player
type NormalChannelMap map[int]*ChatChannel
type GuildChannelMap map[int]*ChatChannel
type PrivateChannelMap map[int]*ChatPrivateChannel

type Chat struct {
	NormalChannels		NormalChannelMap
	GuildChannels		GuildChannelMap
	PrivateChannels		PrivateChannelMap
	
	dummyPrivate		*ChatChannel
}

func NewChat() *Chat {
	chat := Chat{ NormalChannels: make(NormalChannelMap),
				  GuildChannels: make(GuildChannelMap),
				  PrivateChannels: make(PrivateChannelMap) }

	// Create default channels	
	worldChannel := NewChatChannel(pnet.CHANNEL_WORLD, "World")
	chat.NormalChannels[pnet.CHANNEL_WORLD] = worldChannel
	
	tradeChannel := NewChatChannel(pnet.CHANNEL_TRADE, "Trade")
	chat.NormalChannels[pnet.CHANNEL_TRADE] = tradeChannel
	
	privateChannel := NewChatChannel(pnet.CHANNEL_PRIVATE, "Private Chat Channel")
	chat.dummyPrivate = privateChannel
	
	return &chat
}

func (c *Chat) GetFreePrivateChannelId() int {
	for i := 100; i < 10000; i++ {
		if _, found := c.PrivateChannels[i]; !found {
			return i
		}
	}
	
	return 0
}

func (c *Chat) GetChannel(_player *Player, _channelId int) IChatChannel {
	if _channelId == pnet.CHANNEL_GUILD {
		return nil
	} else {
		if value, found := c.NormalChannels[_channelId]; found {
			return value
		}
		
		if value, found := c.PrivateChannels[_channelId]; found {
			return value
		}
	}
	
	return nil
}

func (c *Chat) GetChannelById(_channelId int) *ChatChannel {
	if channel, found := c.NormalChannels[_channelId]; found {
		return channel
	}
	return nil
}

func (c *Chat) GetChannelName(_player *Player, _channelId int) (name string) {
	channel := c.GetChannel(_player, _channelId)
	if channel != nil {
		name = channel.GetName()
	} else {
		name = ""
	}
	
	return
}

func (c *Chat) GetPrivateChannel(_player *Player) *ChatPrivateChannel {
	for _, channel := range(c.PrivateChannels) {
		if channel.GetOwner() == _player.GetUID() {
			return channel
		}
	}
	
	// TODO: Check if player has a guild
	
	
	return nil
}

func (c *Chat) GetChannelList(_player *Player) *list.List {
	clist := list.New()
	gotPrivate := false
	
	for _, channel := range(c.NormalChannels) {
		clist.PushBack(channel)
	}
	
	for _, channel := range(c.PrivateChannels) {
		if channel != nil {
			if channel.IsInvited(_player) {
				clist.PushBack(channel)
			}
			
			if channel.GetOwner() == _player.GetUID() {
				gotPrivate = true
			}
		}
	}
	
	if !gotPrivate {
		clist.PushBack(c.dummyPrivate)
	}
	
	return clist
}

func (c *Chat) CreateChannel(_player *Player, _channelId int) IChatChannel {
	if c.GetChannel(_player, _channelId) != nil {
		return nil
	}
	
	if _channelId == pnet.CHANNEL_GUILD {
		return nil
	} else if _channelId == pnet.CHANNEL_PRIVATE {
		// only 1 private channel for each player
		if c.GetPrivateChannel(_player) != nil {
			return nil
		}
		
		// find a free private channel slot
		slot := c.GetFreePrivateChannelId()
		if slot != 0 {
			newChannel := NewChatPrivateChannel(slot, _player.GetName() + "'s Channel")
			newChannel.Owner = _player.GetUID()
			
			c.PrivateChannels[slot] = newChannel
			return newChannel
		}
	}
	
	return nil
}

func (c *Chat) DeleteChannel(_player *Player, _channelId int) bool {
	if _channelId == pnet.CHANNEL_GUILD {
		return false
	} else {
		if channel, found := c.PrivateChannels[_channelId]; found {
			channel.CloseChannel()
			
			delete(c.PrivateChannels, channel.Id)
			return true
		} 
	}
	
	return false
}

func (c *Chat) AddUserToChannel(_player *Player, _channelId int) bool {
	channel := c.GetChannel(_player, _channelId)
	if channel == nil {
		return false
	}
	
	return channel.AddUser(_player)
}

func (c *Chat) RemoveUserFromChannel(_player *Player, _channelId int) bool {
	channel := c.GetChannel(_player, _channelId)
	if channel == nil {
		return false
	}
	
	if channel.RemoveUser(_player, false) {
		if channel.GetOwner() == _player.GetUID() {
			c.DeleteChannel(_player, channel.GetId())
		}
		
		return true
	}
	
	return false
}

func (c *Chat) RemoveUserFromAllChannels(_player *Player) bool {
	clist := c.GetChannelList(_player)
	for e := clist.Front(); e != nil; e = e.Next() {
		channel := e.Value.(IChatChannel)
		channel.RemoveUser(_player, false)
		
		if channel.GetOwner() == _player.GetUID() {
			c.DeleteChannel(_player, channel.GetId())
		}
	}
	
	return true
}

func (c *Chat) TalkToChannel(_player *Player, _type int, _text string, _channelId int) bool {
	channel := c.GetChannel(_player, _channelId)
	if channel == nil {
		return false
	}
	
	return channel.Talk(_player, _type, _text, 0)
}