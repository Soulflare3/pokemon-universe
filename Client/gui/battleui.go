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
	ICON_BATTLE_POKEMON = 0
	ICON_BATTLE_ATTACK = 1
	ICON_BATTLE_BAG = 2
	ICON_BATTLE_RUN = 3
)

type PU_BattleUI struct {
	PU_GuiElement
}

func NewBattleUI() *PU_BattleUI {
	ui := &PU_BattleUI{}
	ui.visible = false
	g_gui.AddElement(ui)
	
	return ui
}

func (g *PU_BattleUI) Draw() {
	if !g.visible {
		return
	}
	
	img := g_game.GetGuiImage(IMG_GUI_BATTLEBG)
	if img != nil {
		img.Draw(0, 0)
	}
	
	img = g_game.GetGuiImage(IMG_GUI_BATTLETOP)
	if img != nil {
		img.Draw(0, 0)
	}
	
	img = g_game.GetGuiImage(IMG_GUI_BATTLENAME)
	if img != nil {
		img.Draw(1, 2)
		img.Draw(775, 2)
	}
	
	img = g_game.GetGuiImage(IMG_GUI_BATTLELVLLEFT)
	if img != nil {
		img.Draw(190, 2)
	}
	
	img = g_game.GetGuiImage(IMG_GUI_BATTLELVLRIGHT)
	if img != nil {
		img.Draw(648, 2)
	}
	
	img = g_game.GetGuiImage(IMG_GUI_BATTLEHPBAR)
	if img != nil {
		font := g_engine.GetFont(FONT_PURITANBOLD_14)
		font.SetColor(255,255,255)
		font.DrawText("HP", 10, 40)
		font.DrawText("HP", 940, 40)
		
		img.Draw(30, 43)
		img.Draw(770, 43)
	}
	
	dockx, docky := 546, 608
	img = g_game.GetGuiImage(IMG_GUI_ICONDOCK)
	if img != nil {
		img.Draw(dockx, docky)
	}
	
	img = g_game.GetGuiImage(IMG_GUI_ICON_POKEBALL)
	if img != nil {
		iconFont := g_engine.GetFont(FONT_PURITANBOLD_14)
		iconFont.SetColor(255,255,255)
		
		for i := 0; i < 4; i++ {
			img.Draw(dockx+30+(i*100), docky+12)
			switch i {
			case ICON_BATTLE_POKEMON:
				iconFont.DrawText("Pokemon", dockx+(i*100)+65, docky+19)
				
			case ICON_BATTLE_ATTACK:
				iconFont.DrawText("Attack", dockx+(i*100)+65, docky+19)
				
			case ICON_BATTLE_BAG:
				iconFont.DrawText("Bag", dockx+(i*100)+65, docky+19)
				
			case ICON_BATTLE_RUN:
				iconFont.DrawText("Run", dockx+(i*100)+65, docky+19)
			}
		}
	}
}

func (g *PU_BattleUI) MouseDown(_x int, _y int) {

}

func (g *PU_BattleUI) MouseUp(_x int, _y int) {

}

func (g *PU_BattleUI) MouseMove(_x int, _y int) {

}

func (g *PU_BattleUI) MouseScroll(_dir int) {

}

func (g *PU_BattleUI) Focusable() bool {
	return false
}

func (g *PU_BattleUI) KeyDown(_keysym int, _scancode int) {

}

