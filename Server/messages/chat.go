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
	pnet "network"
)

func (c *Connection) Send_CreatureChat(_creature ICreature, _channelId int, _speakType int, _message string) {
	msg := pnet.NewData_Chat()
	msg.ChannelId = _channelId
	msg.SpeakType = _speakType
	msg.Receiver = _creature.GetName()
	msg.Message = _message
	c.SendMessage(msg)
}

func (c *Connection) Receive_Chat(_message *pnet.Message) {
	data := _message.Chat
	g_game.OnPlayerSay(c.Owner, data.ChannelId, data.SpeakType, data.Receiver, data.Message)
}
