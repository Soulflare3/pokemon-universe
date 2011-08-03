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

type PU_BattleEvent_Dialogue struct {
	npc      int
	question string
	options  map[int]string
}

func NewBattleEvent_Dialogue(_question string, _options map[int]string) *PU_BattleEvent_Dialogue {
	event := &PU_BattleEvent_Dialogue{}
	event.npc = -1
	event.question = _question
	for id, text := range _options {
		event.options[id] = text
	}
	return event
}

func NewBattleEvent_DialogueNPC(_npc int, _question string, _options map[int]string) *PU_BattleEvent_Dialogue {
	event := NewBattleEvent_Dialogue(_question, _options)
	event.npc = _npc
	return event
}

func (e *PU_BattleEvent_Dialogue) Execute() {
	if e.npc >= 0 {
		g_game.dialogue.SetDialogueNPC(e.npc, e.question, e.options)
	} else {
		if e.question != "" {
			g_game.dialogue.SetDialogueQuestion(e.question, e.options)
		} else {
			g_game.dialogue.SetDialogueOptions(e.options)
		}
	}
}
