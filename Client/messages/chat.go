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

func (p *PU_GameProtocol) Send_Chat(_speaktype int, _channelId int, _text string) {
	message := punet.NewData_Chat()
	message.Chat.SpeakType = _speaktype
	message.Chat.ChannelId = _channelId
	message.Chat.Message = _text
	g_conn.SendMessage(message)
}

func (p *PU_GameProtocol) Receive_CreatureChat(_message *punet.Message) {
	data := _message.Chat
	text := NewText(FONT_PURITANBOLD_14)
	color := uint32(16773632)
	if data.Receiver == g_game.self.name {
		color = CreateColorKey(39,175,197)
	}
	text.Add(data.Receiver + ": ", color)
	text.Add(data.Message, 16777215)
	g_game.chat.AddMessage(data.ChannelId, text)
	g_game.onscreenchat.Add(data.Receiver, data.Message)
}
