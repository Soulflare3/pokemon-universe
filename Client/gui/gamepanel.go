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

//not really a typical GUI element as there will be only one instance of it and it has very specific implementations.. but it uses the same events as GUI elements

type PU_GamePanel struct {
	PU_GuiElement

	chatInput *PU_Textfield
	
	gameUI *PU_GameUI
}

func NewGamePanel() *PU_GamePanel {
	panel := &PU_GamePanel{}
	panel.visible = true
	g_gui.AddElement(panel)
	
	panel.chatInput = NewTextfield(NewRect(2,690,375,17), FONT_PURITANBOLD_12)
	panel.chatInput.SetColor(255,255,255)
	panel.chatInput.KeyDownCallback = ChatKeydown
	g_gui.SetFocus(panel.chatInput)
	
	panel.gameUI = NewGameUI()
	panel.gameUI.visible = true
	
	return panel
}

func (g *PU_GamePanel) Destroy() {
	if g.chatInput != nil {
		g_gui.RemoveElement(g.chatInput)
	}
}

func ChatKeydown(_keysym int, _scancode int) {
	if _scancode == 13 { // enter/return
		if g_game.self != nil && g_game.panel != nil {
			text := g_game.panel.chatInput.text
			if text != "" {
				g_game.chat.SendMessage(text)
				g_game.panel.chatInput.text = ""
			}
		}
	}
}

func (g *PU_GamePanel) Draw() {
	if !g.visible {
		return
	}

	//draw the background
	img := g_game.GetGuiImage(IMG_GUI_BOTTOM)
	if img != nil {
		img.Draw(0, 277)
	}
	
	
}

func (g *PU_GamePanel) MouseDown(_x int, _y int) {

}

func (g *PU_GamePanel) MouseUp(_x int, _y int) {

}

func (g *PU_GamePanel) MouseMove(_x int, _y int) {

}

func (g *PU_GamePanel) MouseScroll(_dir int) {

}

func (g *PU_GamePanel) Focusable() bool {
	return false
}

func (g *PU_GamePanel) KeyDown(_keysym int, _scancode int) {

}

