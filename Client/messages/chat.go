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
	"os"
)

type PU_Message_ReceiveChat struct {
	name string
	speaktype int 
	channel int
	message string
}

func NewReceiveChatMessage(_packet *punet.Packet) *PU_Message_ReceiveChat {
	msg := &PU_Message_ReceiveChat{}
	msg.ReadPacket(_packet)
	return msg
}

func (m *PU_Message_ReceiveChat) ReadPacket(_packet *punet.Packet) os.Error {
	m.name = _packet.ReadString()
	m.speaktype = int(_packet.ReadUint8())
	m.channel = int(_packet.ReadUint16())
	m.message = _packet.ReadString()
	
	if m.message != "" {
		if m.speaktype == SPEAK_PRIVATE {
			//add PM	
		} else {
			message := NewText(FONT_PURITANBOLD_14)
			
			if m.name == g_game.self.name {
				message.Add(m.name+": ", CreateColorKey(39,175,197))
			} else {
				message.Add(m.name+": ", 16773632) //TODO: change to CreateColorKey .. 
			}
			
			message.Add(m.message, 16777215) //TODO: change to CreateColorKey as well .. 
			
			g_game.chat.AddMessage(m.channel, message)
		}
	}
	return nil
}


type PU_Message_SendChat struct {
	speaktype int
	channel int 
	message string 
}

func NewSendChatMessage() *PU_Message_SendChat {
	return &PU_Message_SendChat{}
}

func (m *PU_Message_SendChat) WritePacket() (*punet.Packet, os.Error) {
	packet := punet.NewPacketExt(0x10)
	packet.AddUint8(uint8(m.speaktype))
	packet.AddUint16(uint16(m.channel))
	packet.AddString(m.message)
	return packet, nil
}

