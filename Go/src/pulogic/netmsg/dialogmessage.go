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
	"strings"
	"strconv"
	
	pnet "nonamelib/network"
)

type DialogMessage struct {
	DialogType	int
	Question	string
	Answers		map[int]string
	NpcId		uint64
	
	// Receive
	AnswerId	int
}

func NewDialogMessage(_dialogType int) *DialogMessage {
	return &DialogMessage { DialogType: _dialogType,
							Answers: make(map[int]string) }
}

// GetHeader returns the header value of this message
func (m *DialogMessage) GetHeader() uint8 {
	return pnet.HEADER_DIALOG
}

func (m *DialogMessage) SetQuestion(_question string, _answers []string) {
	m.Question = _question
	
	if _answers != nil && len(_answers) > 0 {
		for _, answer := range(_answers) {
			answerSplit := strings.Split(answer, "-")
			
			optionId, _ := strconv.ParseInt(answerSplit[0], 10, 32)
			m.Answers[int(optionId)] = answerSplit[1]
		}
	}
}

func (m *DialogMessage) SetNpcId(_npcId uint64) {
	m.NpcId = _npcId
}

func (m *DialogMessage) ReadPacket(_packet pnet.IPacket) error {
	answer, err := _packet.ReadUint32()
	if err != nil {
		return err
	}
	m.AnswerId = int(answer)
	
	return nil
}

// WritePacket write the needed object data to a Packet and returns it
func (m *DialogMessage) WritePacket() pnet.IPacket {
	packet := pnet.NewPacketExt(m.GetHeader())
	packet.AddUint8(uint8(m.DialogType))
	
	if m.DialogType != pnet.DIALOG_CLOSE {
		if m.DialogType == pnet.DIALOG_NPC || m.DialogType == pnet.DIALOG_NPCTEXT {
			packet.AddUint64(m.NpcId)
		}
		
		if m.DialogType != pnet.DIALOG_OPTIONS {
			packet.AddString(m.Question)
		}
		
		if m.DialogType != pnet.DIALOG_NPCTEXT {
		
			for index, answer := range(m.Answers) {
				packet.AddUint32(uint32(index))
				packet.AddString(answer)
			}
		
		}
	}
	
	return packet
}
