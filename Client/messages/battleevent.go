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

type PU_Message_BattleEvent struct {

}

func NewBattleEventMessage(_packet *punet.Packet) *PU_Message_BattleEvent {
	msg := &PU_Message_BattleEvent{}
	msg.ReadPacket(_packet)
	return msg
}

func (m *PU_Message_BattleEvent) ReadPacket(_packet *punet.Packet) os.Error {
	eventtype := int(_packet.ReadUint16())
	switch eventtype {
	case BATTLEEVENT_SLEEP:
		ticks := uint32(_packet.ReadUint16())
		
		event := NewBattleEvent_Sleep(ticks)
		g_game.battle.AddEvent(event)
		
	case BATTLEEVENT_TEXT:
		newText := NewText(FONT_PURITANBOLD_14)
		pieces := int(_packet.ReadUint16())
		for i := 0; i < pieces; i++ {
			color := int(_packet.ReadUint8())
			text := _packet.ReadString()
			
			newText.Add(text, ColorCodeToKey(color))
		}
		
		event := NewBattleEvent_Text(newText)
		g_game.battle.AddEvent(event)
	
	case BATTLEEVENT_CHANGEHP:
		fighter := int(_packet.ReadUint16())
		hp := int(_packet.ReadUint16())
		
		event := NewBattleEvent_ChangeHP(fighter, hp)
		g_game.battle.AddEvent(event)
		
	/*case BATTLEEVENT_ANIMATION:
		pos := int(_packet.ReadUint16())
		id := int(_packet.ReadUint16())
		
		event := NewBattleEvent_Animation(pos, id)*/
		
	case BATTLEEVENT_CHANGEPOKEMON_SELF:
		pokeid := int(_packet.ReadUint16())
		
		event := NewBattleEvent_ChangePokemon_Self(pokeid)
		g_game.battle.AddEvent(event)
		
	case BATTLEEVENT_CHANGEPOKEMON:
		fighter := int(_packet.ReadUint16())
		pokemon := int(_packet.ReadUint16())
		name := _packet.ReadString()
		hp := int(_packet.ReadUint16())
		level := int(_packet.ReadUint16())
		
		event := NewBattleEvent_ChangePokemon(fighter, pokemon, name, hp, level)
		g_game.battle.AddEvent(event)
		
	case BATTLEEVENT_CHANGEPP:
		pokemon := int(_packet.ReadUint16())
		attack := int(_packet.ReadUint16())
		value := int(_packet.ReadUint16())
		
		event := NewBattleEvent_ChangePP(pokemon, attack, value)
		g_game.battle.AddEvent(event)
		
	case BATTLEEVENT_STOPBATTLE:
		event := NewBattleEvent_StopBattle()
		g_game.battle.AddEvent(event)
		
	case BATTLEEVENT_CHANGELEVELSELF:
		pokemon := int(_packet.ReadUint16())
		level := int(_packet.ReadUint16())
		
		event := NewBattleEvent_ChangeLevelSelf(pokemon, level)
		g_game.battle.AddEvent(event)
		
	case BATTLEEVENT_CHANGELEVEL:
		fighter := int(_packet.ReadUint16())
		level := int(_packet.ReadUint16())
		
		event := NewBattleEvent_ChangeLevel(fighter, level)
		g_game.battle.AddEvent(event)
		
	case BATTLEEVENT_CHANGEATTACK:
		pokemon := int(_packet.ReadUint16())
		slot := int(_packet.ReadUint16())
		name := _packet.ReadString()
		description := _packet.ReadString()
		poketype := _packet.ReadString()
		pp := int(_packet.ReadUint16())
		ppmax := int(_packet.ReadUint16())
		power := int(_packet.ReadUint16())
		accuracy := int(_packet.ReadUint16())
		category := _packet.ReadString()
		target := _packet.ReadString()
		contact := _packet.ReadString()
		
		event := NewBattleEvent_ChangeAttack(pokemon, slot, name, description, poketype, pp, ppmax, power, accuracy, category, target, contact)
		g_game.battle.AddEvent(event)
		
	case BATTLEEVENT_ALLOWCONTROL:
		state := int(_packet.ReadUint16())
		
		event := NewBattleEvent_AllowControl(state)
		g_game.battle.AddEvent(event)
		
	case BATTLEEVENT_DIALOGUE:
		dialoguetype := int(_packet.ReadUint8())
		npc := -1
		if dialoguetype == DIALOGUE_NPC {
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
			
			event := NewBattleEvent_DialogueNPC(npc, question, options)
			g_game.battle.AddEvent(event)
		}
		
	case BATTLEEVENT_CHANGEEXP:
		pokemon := int(_packet.ReadUint16())
		exp := int(_packet.ReadUint16())
		
		event := NewBattleEvent_ChangeExp(pokemon, exp)
		g_game.battle.AddEvent(event)
	}

	return nil
}

