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

type Quest struct {
	Id			int64
	Name		string
	Description	string
	Status		int	// uint8
}

type QuestListMessage struct {
	Quests map[int64]*Quest
}

func NewQuestListMessage() *QuestListMessage {
	return &QuestListMessage { }
}

// GetHeader returns the header value of this message
func (m *QuestListMessage) GetHeader() uint8 {
	return pnet.HEADER_QUESTLIST
}

func (m *QuestListMessage) AddQuest(_questId int64, _name string, _description string, _status int) {
	quest := &Quest { Id: _questId,
					  Name: _name, 
					  Description: _description,
					  Status: _status }
	m.Quests[_questId] = quest
}

func (m *QuestListMessage) ReadPacket(_packet pnet.IPacket) error {
	return nil
}

// WritePacket write the needed object data to a Packet and returns it
func (m *QuestListMessage) WritePacket() pnet.IPacket {
	packet := pnet.NewPacketExt(m.GetHeader())
	packet.AddUint32(uint32(len(m.Quests)))
	
	for _, quest := range(m.Quests) {
		packet.AddUint64(uint64(quest.Id))
		packet.AddString(quest.Name)
		packet.AddString(quest.Description)
		packet.AddUint8(uint8(quest.Status))
	}
	
	return packet
}
