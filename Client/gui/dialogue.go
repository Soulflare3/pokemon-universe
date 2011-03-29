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

const (
	DIALOGUE_STATE_IDLE = iota
	DIALOGUE_STATE_NORMAL
	DIALOGUE_STATE_ANSWERED
)

const (
	DIALOGUE_QUESTION = 0
	DIALOGUE_NPC = 1
	DIALOGUE_CLOSE = 2
	DIALOGUE_OPTIONS = 3
	DIALOGUE_NPCTEXT = 4
)

var DIALOGUE_RECT *PU_Rect = NewRect(452, 295, 514, 302)

type PU_Dialogue struct {
	PU_GuiElement
	
	dialoguetype int
	options map[int]string
	question []string
	
	selected int
	state int 
	npcname string
}

func NewDialogue() *PU_Dialogue {
	dialogue := &PU_Dialogue{}
	dialogue.visible = true
	g_gui.AddElement(dialogue)
	
	dialogue.state = DIALOGUE_STATE_IDLE
	
	return dialogue
}

func (g *PU_Dialogue) Reset() {
	g.options = nil
	g.selected = 0
	g.npcname = ""
	g.question = nil
	g.dialoguetype = 0
}

func (g *PU_Dialogue) Close() {
	g.Reset()
	g.state = DIALOGUE_STATE_IDLE
}

func (g *PU_Dialogue) Answer(_id int) {
	g.state = DIALOGUE_STATE_ANSWERED
	g_conn.Game().SendDialogueAnswer(_id)
}

func (g *PU_Dialogue) SetOptions(_options map[int]string) {
	g.options = make(map[int]string)
	g.selected = 0
	for id, text := range _options {
		g.options[id] = text
	}
}

func (g *PU_Dialogue) SetDialogueNPC(_npc int, _question string, _options map[int]string) {
	g.Reset()
	
	npc := g_map.GetCreatureByID(uint32(_npc)).(*PU_Player)
	if npc != nil {
		g.npcname = npc.name
		g.question = ClipText(_question, FONT_ARIALBLACK_18, 470)
		
		g.SetOptions(_options)
		
		message := NewText(FONT_PURITANBOLD_14)
		message.Add(npc.name+": ", CreateColorKey(0, 255, 255))
		message.Add(_question, CreateColorKey(255, 255, 255))
		g_game.chat.AddMessage(CHANNEL_LOG, message)
		
		g.state = DIALOGUE_STATE_NORMAL
		g.dialoguetype = DIALOGUE_NPC
	}
}

func (g *PU_Dialogue) SetDialogueQuestion(_question string, _options map[int]string) {
	g.Reset()
	
	g.npcname = ""
	g.question = ClipText(_question, FONT_ARIALBLACK_18, 470)
	
	g.SetOptions(_options)
	
	message := NewText(FONT_PURITANBOLD_14)
	message.Add(_question, CreateColorKey(255, 255, 255))
	g_game.chat.AddMessage(CHANNEL_LOG, message)

	g.state = DIALOGUE_STATE_NORMAL
	g.dialoguetype = DIALOGUE_QUESTION
}

func (g *PU_Dialogue) SetDialogueQuestionText(_question *PU_Text, _options map[int]string) {
	g.Reset()
	
	g.npcname = ""
	g.question = ClipText(_question.GetAll(), FONT_ARIALBLACK_18, 470)
	
	g.SetOptions(_options)
	
	g_game.chat.AddMessage(CHANNEL_LOG, _question)

	g.state = DIALOGUE_STATE_NORMAL
	g.dialoguetype = DIALOGUE_QUESTION
}

func (g *PU_Dialogue) SetDialogueOptions(_options map[int]string) {
	g.Reset()
	
	g.npcname = ""
	g.question = nil
	
	g.SetOptions(_options)

	g.state = DIALOGUE_STATE_NORMAL
	g.dialoguetype = DIALOGUE_OPTIONS
}

func (g *PU_Dialogue) Draw() {
	if !g.visible {
		return
	}
	if g.state != DIALOGUE_STATE_IDLE {
		g_game.GetGuiImage(IMG_GUI_DIALOGUEWINDOW).Draw(DIALOGUE_RECT.x, DIALOGUE_RECT.y)
		
		font := g_engine.GetFont(FONT_ARIALBLACK_48)
		font.SetColor(255, 255, 255)
		font.DrawText(g.npcname, DIALOGUE_RECT.x+13, DIALOGUE_RECT.y+7)

		font = g_engine.GetFont(FONT_ARIALBLACK_14)
		font.SetColor(127, 127, 127)
		lineskip := 0
		for _, line := range g.question {
			font.DrawText(line, DIALOGUE_RECT.x+15, DIALOGUE_RECT.y+70+lineskip)
			lineskip += 17
		}
		
		font.SetColor(255, 255, 255)
		lineHeight := 25
		i := 0
		for id, text := range g.options {
			if id == g.selected {
				font.DrawText("> "+text, DIALOGUE_RECT.x+15, DIALOGUE_RECT.y+149+(i*lineHeight))
			} else {
				font.DrawText("- "+text, DIALOGUE_RECT.x+15, DIALOGUE_RECT.y+149+(i*lineHeight))
			}
			i++
		}
	}
}

func (g *PU_Dialogue) MouseDown(_x int, _y int) {

}

func (g *PU_Dialogue) MouseUp(_x int, _y int) {
	if !g.visible {
		return
	}
	if g.state != DIALOGUE_STATE_IDLE {
		if !DIALOGUE_RECT.Contains(_x, _y) {
			return
		}

		if g.state == DIALOGUE_STATE_NORMAL	{
			lineHeight := 33
			i := 0
			for id, text := range g.options {
				curY := DIALOGUE_RECT.y+129+(i*lineHeight)
				if _y >= curY && _y < curY+lineHeight {
					lineWidth := g_engine.GetFont(FONT_PURITANBOLD_14).GetStringWidth("> "+text)
					if _x >= DIALOGUE_RECT.x+15 && _x <= DIALOGUE_RECT.x+15+lineWidth {
						if g.dialoguetype == DIALOGUE_NPC || g.dialoguetype == DIALOGUE_QUESTION {
							answerText := NewText(FONT_PURITANBOLD_14)
							answerText.Add(g_game.self.name+": ", CreateColorKey(39,175,197))
							answerText.Add(text, 16777215)
							g_game.chat.AddMessage(CHANNEL_LOCAL, answerText)	
						}
						g.Answer(id)
						return						
					}
				}
				i++
			}
		}
	}
}

func (g *PU_Dialogue) MouseMove(_x int, _y int) {
	if !g.visible {
		return
	}
	if g.state != DIALOGUE_STATE_IDLE {
		if !DIALOGUE_RECT.Contains(_x, _y) {
			return
		}

		if g.state == DIALOGUE_STATE_NORMAL	{
			lineHeight := 33
			i := 0
			for id, text := range g.options {
				curY := DIALOGUE_RECT.y+129+(i*lineHeight)
				if _y >= curY && _y < curY+lineHeight {
					lineWidth := g_engine.GetFont(FONT_PURITANBOLD_14).GetStringWidth("> "+text)
					if _x >= DIALOGUE_RECT.x+15 && _x <= DIALOGUE_RECT.x+15+lineWidth {
						g.selected = id
						return					
					}
				}
				i++
			}
		}
	}
}

func (g *PU_Dialogue) MouseScroll(_dir int) {

}

func (g *PU_Dialogue) Focusable() bool {
	return false
}

func (g *PU_Dialogue) KeyDown(_keysym int, _scancode int) {

}

