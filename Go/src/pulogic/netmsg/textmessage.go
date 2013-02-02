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
package netmsg

import (
	pnet "nonamelib/network"
)

type TextMessage struct {	
	// Send
	Class	int
	Message string	
}

func NewTextMessage(_messageClass int, _message string) *TextMessage {
	message := &TextMessage{ Class: _messageClass,
							 Message: _message }
	
	return message
}

// GetHeader returns the header value of this message
func (m *TextMessage) GetHeader() uint8 {
	return pnet.HEADER_TEXT_MSG
}

// WritePacket write the needed object data to a Packet and returns it
func (m *TextMessage) WritePacket() pnet.IPacket {
	packet := pnet.NewPacketExt(m.GetHeader())
	packet.AddUint16(uint16(m.Class))
	packet.AddString(m.Message)
	
	return packet
}