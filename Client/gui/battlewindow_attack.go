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
	"fmt"
	"sdl"
)

type PU_BattleWindow_Attack struct {
	PU_GuiElement
	PU_BattleWindow

	description []string
	updateDescription bool
}

func NewBattleWindow_Attack() *PU_BattleWindow_Attack {
	window := &PU_BattleWindow_Attack{}
	window.visible = true
	g_gui.AddElement(window)
	
	window.updateDescription = true
	window.windowtype = BATTLEWINDOW_ATTACK
	
	return window
}

func (g *PU_BattleWindow_Attack) Close() {
	g_gui.RemoveElement(g)
}

func (g *PU_BattleWindow_Attack) Draw() {
	if !g.visible {
		return
	}
	
	g_engine.DrawFillRect(NewRect(0, 0, 964, 575), &sdl.Color{0,0,0,0}, 150)	

	switch g.value {
	case MOVE_ATTACK1:
		img := g_game.GetGuiImage(IMG_GUI_BATTLEATTACK1)
		if img != nil {
			img.Draw(163, 29)
		}

	case MOVE_ATTACK2:
		img := g_game.GetGuiImage(IMG_GUI_BATTLEATTACK2)
		if img != nil {
			img.Draw(163, 29)
		}
		
	case MOVE_ATTACK3:
		img := g_game.GetGuiImage(IMG_GUI_BATTLEATTACK3)
		if img != nil {
			img.Draw(163, 29)
		}
		
	case MOVE_ATTACK4:
		img := g_game.GetGuiImage(IMG_GUI_BATTLEATTACK4)
		if img != nil {
			img.Draw(163, 29)
		}
	}
	
	img := g_game.GetGuiImage(IMG_GUI_BATTLEOPTIONSELECTED)
	if img != nil {
		img.Draw(171, 45)
	}
	
	img = g_game.GetGuiImage(IMG_GUI_BATTLEICONATTACK)
	if img != nil {
		img.Draw(178, 48)
	}
	
	img = g_game.GetGuiImage(IMG_GUI_BATTLECLOSEBUTTON)
	if img != nil {
		img.Draw(764, 48)
	}
	
	g.DrawAttackInfo(g.value)
	for i := 0; i < 4; i++ {
		g.DrawAttack(i)
	}
}

func (g *PU_BattleWindow_Attack) DrawAttackInfo(_attack int) {
	if g_game.battle == nil || g_game.battle.self.pokemon == nil {
		return
	}
	
	attack := g_game.battle.self.pokemon.attacks[_attack]
	if attack != nil {
		font := g_engine.GetFont(FONT_PURITANBOLD_34)
		font.SetColor(141, 198, 63)
		font.DrawText("Pokedex:", 189, 151)
		
		if g.updateDescription {
			g.description = ClipText(attack.description, FONT_PURITANBOLD_18, 236)
			g.updateDescription = false
		}
		
		font = g_engine.GetFont(FONT_PURITANBOLD_18)
		font.SetColor(255, 255, 255)
		lineHeight := font.GetStringHeight()+4
		for i, line := range g.description {
			font.DrawText(line, 189, 192+(i*lineHeight))
		}
		
		font.SetColor(141, 198, 63)
		font.DrawText("Statistics:", 189, 294)
		
		font = g_engine.GetFont(FONT_PURITANBOLD_14)
		font.SetColor(255, 255, 255)
		font.DrawText(fmt.Sprintf("Category: %s", attack.category), 189, 314)
		font.DrawText(fmt.Sprintf("Accuracy: %d", attack.accuracy), 189, 332)
		font.DrawText(fmt.Sprintf("PP: %d (Max %d)", attack.pp, attack.ppmax), 189, 350)
		font.DrawText(fmt.Sprintf("Target: %s", attack.target), 189, 368)
		font.DrawText(fmt.Sprintf("Power x Accuracy: %d", attack.power*attack.accuracy), 189, 386)
		
		font = g_engine.GetFont(FONT_PURITANBOLD_12)
		font.SetColor(255, 255, 255)
		if attack.contact == "no" {
			font.DrawText("* Does not make contact", 189, 424)
		} else {
			font.DrawText("* Makes contact", 189, 424)
		}
	}
}

func (g *PU_BattleWindow_Attack) DrawAttack(_attack int) {
	if g_game.battle == nil || g_game.battle.self.pokemon == nil {
		return
	}
	
	attack := g_game.battle.self.pokemon.attacks[_attack]
	if attack != nil {
		DrawType(attack.poketype, 437, 120+(_attack*102))
		
		font := g_engine.GetFont(FONT_PURITANBOLD_48)
		font.SetColor(255, 255, 255)
		font.DrawText(attack.name, 501, 135+(_attack*102))
		
		font = g_engine.GetFont(FONT_PURITANBOLD_14)
		font.SetColor(255, 255, 255)
		font.DrawText(fmt.Sprintf("Attack Type: %s", attack.poketype), 444, 195+(_attack*102))
		font.DrawText(fmt.Sprintf("PP: %d/%d", attack.pp, attack.ppmax), 583, 195+(_attack*102))
		font.DrawText(fmt.Sprintf("Power: %d", attack.power), 655, 195+(_attack*102))
		
	}
}

func (g *PU_BattleWindow_Attack) MouseDown(_x int, _y int) {

}

func (g *PU_BattleWindow_Attack) MouseUp(_x int, _y int) {
	closeRect := NewRect(764,48,39,39)
	if closeRect.Contains(_x, _y) {
		g_game.panel.battleUI.CloseWindow()
		return
	}
	
	if g_game.panel.battleUI.moveState == BATTLEUI_CHOOSEMOVE {
		if _x >= 436 && _x <= 436+375 {
			for i := 0; i < 4; i++ {
				curY := 114+(i*105)
				if _y >= curY && _y <= curY+79 {
					if g_game.battle.self.pokemon.attacks[i] != nil {
						g.value = i
						//g_conn.Game().SendBattleMove(MOVE_ATTACK, i, 0)
						g_game.panel.battleUI.ChooseMove(g.windowtype, i)
						g_game.panel.battleUI.CloseWindow()
					}
					break
				}
			}
		}
	}
}

func (g *PU_BattleWindow_Attack) MouseMove(_x int, _y int) {
	if g_game.panel.battleUI.moveState == BATTLEUI_CHOOSEMOVE {
		if _x >= 436 && _x <= 436+375 {
			for i := 0; i < 4; i++ {
				curY := 114+(i*105)
				if _y >= curY && _y <= curY+79 {
					g.value = i
					g.updateDescription = true
					break
				}
			}
		}
	}
}

func (g *PU_BattleWindow_Attack) MouseScroll(_dir int) {

}

func (g *PU_BattleWindow_Attack) Focusable() bool {
	return false
}

func (g *PU_BattleWindow_Attack) KeyDown(_keysym int, _scancode int) {

}

