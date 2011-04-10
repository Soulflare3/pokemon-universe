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

type PU_Message_ReceiveDialogue struct {

}

func NewReceiveDialogueMessage(_packet *punet.Packet) *PU_Message_ReceiveDialogue {
	msg := &PU_Message_ReceiveDialogue{}
	msg.ReadPacket(_packet)
	return msg
}

func (m *PU_Message_ReceiveDialogue) ReadPacket(_packet *punet.Packet) os.Error {
	dialoguetype := int(_packet.ReadUint8())
	npc := -1
	if dialoguetype == DIALOGUE_NPC || dialoguetype == DIALOGUE_NPCTEXT {
		npc = int(_packet.ReadUint16())
	}
	if dialoguetype == DIALOGUE_NPC || dialoguetype == DIALOGUE_QUESTION || dialoguetype == DIALOGUE_OPTIONS {
		question := ""
		if dialoguetype != DIALOGUE_OPTIONS {
			question = _packet.ReadString()
		}
		numOptions := int(_packet.ReadUint8())
		options := make(map[int]string)
		for i := 0; i < numOptions; i++ {
			optionid := int(_packet.ReadUint16())
			optionstr := _packet.ReadString()
			
			options[optionid] = optionstr
		}
		switch dialoguetype {
		case DIALOGUE_NPC:
			g_game.dialogue.SetDialogueNPC(npc, question, options)
		
		case DIALOGUE_QUESTION:
			g_game.dialogue.SetDialogueQuestion(question, options)
			
		case DIALOGUE_OPTIONS:
			g_game.dialogue.SetDialogueOptions(options)
		}
		return nil
	}
	if dialoguetype == DIALOGUE_NPCTEXT {
		text := _packet.ReadString()
		
		pNpc := g_map.GetCreatureByID(uint32(npc)).(*PU_Player)
		if pNpc != nil {
			message := NewText(FONT_PURITANBOLD_14)
			message.Add(pNpc.name+": ", CreateColorKey(0, 255, 255))
			message.Add(text, CreateColorKey(255, 255, 255))
			g_game.chat.AddMessage(CHANNEL_LOG, message)
		}
		return nil
	}
	if dialoguetype == DIALOGUE_CLOSE {
		g_game.dialogue.Close()
		return nil
	}
	return nil
}

type PU_Message_DialogueAnswer struct {
	answer int
}

func NewDialogueAnswerMessage() *PU_Message_DialogueAnswer {
	return &PU_Message_DialogueAnswer{}
}

func (m *PU_Message_DialogueAnswer) WritePacket() (*punet.Packet, os.Error) {
	packet := punet.NewPacketExt(0x20)
	packet.AddUint16(uint16(m.answer))
	return packet, nil
}

