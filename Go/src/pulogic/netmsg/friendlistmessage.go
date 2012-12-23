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

type FriendListMessage struct {
	Friends map[string]int
}

func NewFriendListMessage() *FriendListMessage {
	return &FriendListMessage { }
}

// GetHeader returns the header value of this message
func (m *FriendListMessage) GetHeader() uint8 {
	return pnet.HEADER_FRIENDLIST
}

func (m *FriendListMessage) AddFriend(_name string, _online bool) {
	online := 0
	if _online {
		online = 1
	}
	m.Friends[_name] = online
}

func (m *FriendListMessage) ReadPacket(_packet pnet.IPacket) error {
	return nil
}

// WritePacket write the needed object data to a Packet and returns it
func (m *FriendListMessage) WritePacket() pnet.IPacket {
	packet := pnet.NewPacketExt(m.GetHeader())
	packet.AddUint32(uint32(len(m.Friends)))
	
	for name, online := range(m.Friends) {
		packet.AddString(name)
		packet.AddUint8(uint8(online))
	}
	
	return packet
}
