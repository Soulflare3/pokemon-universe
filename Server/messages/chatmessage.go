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

type ChatMessage struct {
	From		ICreature
	SpeakType	int
	Text		string
	ChannelId	int
	Receiver	string
	Time		int
}

func NewChatMessage(_from ICreature) *ChatMessage {
	return &ChatMessage { From: _from }
}

func NewChatMessageExt(_from ICreature, _type int, _text string, _channelId int, _time int) *ChatMessage {
	return &ChatMessage { From: _from,
						  SpeakType: _type,
						  Text: _text,
						  ChannelId: _channelId,
						  Time: _time }
}

// GetHeader returns the header value of this message
func (m *ChatMessage) GetHeader() uint8 {
	return pnet.HEADER_CHAT
}

func (m *ChatMessage) ReadPacket(_packet pnet.IPacket) error {
	m.SpeakType = int(_packet.ReadUint8())
	m.ChannelId = int(_packet.ReadUint16())
	m.Receiver = _packet.ReadString()
	m.Text = _packet.ReadString()
		
	g_game.OnPlayerSay(m.From.(*Player), m.ChannelId, m.SpeakType, m.Receiver, m.Text)
	
	return nil
}

// WritePacket write the needed object data to a Packet and returns it
func (m *ChatMessage) WritePacket() pnet.IPacket {
	packet := pnet.NewPacketExt(m.GetHeader())
	packet.AddUint64(m.From.GetUID())
	packet.AddString(m.From.GetName())
	packet.AddUint8(uint8(m.SpeakType))
	packet.AddUint16(uint16(m.ChannelId))
	packet.AddString(m.Text)

	return packet
}