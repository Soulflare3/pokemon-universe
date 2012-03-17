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
	pnet "network"
)

type FriendUpdateMessage struct {
	Name	string
	Removed	int // uint8
	Online	int // uint8
}

func NewFriendUpdateMessage() *FriendUpdateMessage {
	return &FriendUpdateMessage{}
}

func NewFriendUpdateMessageExt(_name string, _online bool, _removed bool) *FriendUpdateMessage {
	onlineInt := 0
	removedInt := 0
	
	if _online {
		onlineInt = 1
	}
	if _removed { 
		removedInt = 1
	}
	
	return &FriendUpdateMessage { Name: _name,
								  Online: onlineInt,
								  Removed: removedInt }
}

// GetHeader returns the header value of this message
func (m *FriendUpdateMessage) GetHeader() uint8 {
	return pnet.HEADER_FRIENDUPDATE
}

func (m *FriendUpdateMessage) ReadPacket(_packet pnet.IPacket) error {
	m.Name = _packet.ReadString()
	m.Removed = int(_packet.ReadUint8())
	
	return nil
}

// WritePacket write the needed object data to a Packet and returns it
func (m *FriendUpdateMessage) WritePacket() pnet.IPacket {
	packet := pnet.NewPacketExt(m.GetHeader())
	packet.AddString(m.Name)
	packet.AddUint8(uint8(m.Removed))
	packet.AddUint8(uint8(m.Online))
	
	return packet
}