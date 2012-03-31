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

type QuestUpdateMessage struct {
	Id			int64 // player quest id
	Name		string
	Description	string
	Status		int	// uint8
	Removed		bool
}

func NewQuestUpdateMessage() *QuestUpdateMessage {
	return &QuestUpdateMessage{}
}

func NewQuestUpdateMessageExt(_questId int64, _name string, _description string, _status int) *QuestUpdateMessage {
	return &QuestUpdateMessage { Id: _questId,
								 Name: _name,
								 Description: _description,
								 Status: _status }
}

// GetHeader returns the header value of this message
func (m *QuestUpdateMessage) GetHeader() uint8 {
	return pnet.HEADER_QUESTUPDATE
}

func (m *QuestUpdateMessage) ReadPacket(_packet pnet.IPacket) error {
	m.Id = int64(_packet.ReadUint64())
	m.Removed = (_packet.ReadUint8() == 1)
	
	return nil
}

// WritePacket write the needed object data to a Packet and returns it
func (m *QuestUpdateMessage) WritePacket() pnet.IPacket {
	packet := pnet.NewPacketExt(m.GetHeader())
	packet.AddUint64(uint64(m.Id))
	packet.AddString(m.Name)
	packet.AddString(m.Description)
	packet.AddUint8(uint8(m.Status))
	
	if m.Removed {
		packet.AddUint8(1)
	} else {
		packet.AddUint8(0)
	}
	
	return packet
}