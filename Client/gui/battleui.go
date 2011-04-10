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
	"sdl"
	"fmt"
	"math"
)

const (
	ICON_BATTLE_POKEMON = 0
	ICON_BATTLE_ATTACK = 1
	ICON_BATTLE_BAG = 2
	ICON_BATTLE_RUN = 3
)

const (
	BATTLEUI_INACTIVE = iota
	BATTLEUI_INITIALIZING
	BATTLEUI_ACTIVE
)

const (
	BATTLEUI_CHOOSEMOVE = 0
	BATTLEUI_MOVESENT = 1
	BATTLEUI_CHOOSEPOKEMON_ITEM = 2
	BATTLEUI_CHOOSEATTACK_ITEM = 3
	BATTLEUI_CHOOSEPOKEMON = 4
)



type PU_BattleUI struct {
	PU_GuiElement

	state int
	moveState int
	initProgress float32
	
	window IBattleWindow
	
	lastWindow int 
	lastValue int
}

func NewBattleUI() *PU_BattleUI {
	ui := &PU_BattleUI{}
	ui.visible = false
	g_gui.AddElement(ui)
	
	return ui
}

func (g *PU_BattleUI) Init() {
	g.initProgress = 0.0
	g.lastWindow = BATTLEWINDOW_NONE
	g.lastValue = 0
	
	g_game.ShowGameUI(false)
	g_game.ShowBattleUI(true)
	
	g_game.chat.SetActive(CHANNEL_BATTLE)
	g_game.state = GAMESTATE_BATTLE_INIT
	
	g.moveState = BATTLEUI_CHOOSEMOVE
	g.state = BATTLEUI_INITIALIZING
}

func (g *PU_BattleUI) Reset() {
	g.moveState = BATTLEUI_CHOOSEMOVE
	g.lastWindow = BATTLEWINDOW_NONE
	g.lastValue = 0
}

func (g *PU_BattleUI) ChooseMove(_window int, _value int) {
	g.moveState = BATTLEUI_MOVESENT
	g.lastWindow = _window
	g.lastValue = _value
}

func (g *PU_BattleUI) OpenWindow(_windowtype int) {
	g.CloseWindow()
	
	switch _windowtype {
	case BATTLEWINDOW_ATTACK:
		g.window = NewBattleWindow_Attack()
		
	case BATTLEWINDOW_ITEMS:
		
		
	case BATTLEWINDOW_POKEMON:
		
	}
	
	if g.window != nil {
		if _windowtype == g.lastWindow {
			g.window.SetValue(g.lastValue)
		}
	}
}

func (g *PU_BattleUI) CloseWindow() {
	if g.window != nil {
		if g.window.GetType() == BATTLEWINDOW_POKEMON && (g.moveState == BATTLEUI_CHOOSEPOKEMON_ITEM || g.moveState == BATTLEUI_CHOOSEATTACK_ITEM) {
			//g_conn.Game().SendBattleMove(MOVE_ANSWER, -1, 0)
		} else if g.window.GetType() == BATTLEWINDOW_POKEMON && g.moveState == BATTLEUI_CHOOSEPOKEMON {
			return //new pokemon MUST be chosen
		}
		
		g.window.Close()
		g.window = nil
	}
}


func (g *PU_BattleUI) Draw() {
	if !g.visible {
		return
	}
	
	switch g.state {
	case BATTLEUI_INITIALIZING:
		g.DrawInitializing()
		
	case BATTLEUI_ACTIVE:
		g.DrawActive()
	}
}

func (g *PU_BattleUI) DrawInitializing() {
	if g.initProgress < 1.0 { //fade out gamescreen
		alpha := int(math.Ceil(float64(g.initProgress)*255.0))
		g_engine.DrawFillRect(NewRect(0, 0, WINDOW_WIDTH, WINDOW_HEIGHT), &sdl.Color{0,0,0,0}, uint8(alpha))
	} else { //fade in background
		g_game.state = GAMESTATE_BATTLE
		
		img := g_game.GetGuiImage(IMG_GUI_BATTLEBG)
		if img != nil {
			img.Draw(0, 0)
		}
		
		alphaproc := float32(2.0-g.initProgress)
		alpha := int(math.Ceil(float64(alphaproc)*255.0))
		
		g_engine.DrawFillRect(NewRect(0, 0, WINDOW_WIDTH, 575), &sdl.Color{0,0,0,0}, uint8(alpha))
	}	

	g.initProgress += 1000./800. * (float32(g_frameTime)/1000.);
	if g.initProgress >= 2.0 {
		g.state = BATTLEUI_ACTIVE
	}
}

func (g *PU_BattleUI) DrawActive() {
	moveproc := float32(0.0)
	if g.initProgress < 3.0 {
		g.initProgress += 1000./800. * (float32(g_frameTime)/1000.);
		moveproc = 3.0-g.initProgress
	}

	img := g_game.GetGuiImage(IMG_GUI_BATTLEBG)
	if img != nil {
		img.Draw(0, 0)
	}
	
	visibleheight := int(math.Ceil(float64(moveproc)*float64(60 /*size of top control*/)))
	y := 0-visibleheight
	
	img = g_game.GetGuiImage(IMG_GUI_BATTLETOP)
	if img != nil {
		img.Draw(0, y)
	}
	
	img = g_game.GetGuiImage(IMG_GUI_BATTLENAME)
	if img != nil {
		img.Draw(1, y+2)
		img.Draw(775, y+2)
	}
	
	img = g_game.GetGuiImage(IMG_GUI_BATTLELVLLEFT)
	if img != nil {
		img.Draw(190, y+2)
	}
	
	img = g_game.GetGuiImage(IMG_GUI_BATTLELVLRIGHT)
	if img != nil {
		img.Draw(648, y+2)
	}
	
	img = g_game.GetGuiImage(IMG_GUI_BATTLEHPBAR)
	if img != nil {
		font := g_engine.GetFont(FONT_PURITANBOLD_14)
		font.SetColor(255,255,255)
		font.DrawText("HP", 10, y+40)
		font.DrawText("HP", 940, y+40)
		
		img.Draw(30, y+43)
		img.Draw(770, y+43)
	}
	
	//self
	self := g_game.battle.self
	if self != nil {
		font := g_engine.GetFont(FONT_ARIALBLACK_18)
		font.SetColor(255, 255, 255)
		font.DrawTextCentered(self.GetPokeName(), 1, 188, y+2)
		
		font = g_engine.GetFont(FONT_ARIALBLACK_16)
		font.SetColor(255, 255, 255)
		font.DrawText(fmt.Sprintf("LVL %d", self.GetLevel()), 200, y+6)
		
		hpperc := self.GetHPPerc()
		switch {
		case hpperc <= 20:
			img = g_game.GetGuiImage(IMG_GUI_BATTLEHPBAR_RED)
		
		case hpperc > 20 && hpperc <= 40:
			img = g_game.GetGuiImage(IMG_GUI_BATTLEHPBAR_YELLOW)
		
		default:
			img = g_game.GetGuiImage(IMG_GUI_BATTLEHPBAR_GREEN)
		}
		barwidth := int(math.Floor((float64(hpperc)/100.0)*float64(img.w))) 
		if img != nil {
			img.DrawRectClip(NewRect(30, y+43, barwidth, int(img.h)), NewRect(0, 0, barwidth, int(img.h)))
		} 
		
		hp := self.GetHP()
		hpmax := self.GetHPMax()
		font = g_engine.GetFont(FONT_ARIALBLACK_10)
		font.SetColor(0, 0, 0)
		font.DrawText(fmt.Sprintf("%d/%d", hp, hpmax), 93, y+41)
		
		img = g_game.GetGuiImage(IMG_GUI_BATTLEEXPBAR)
		if img != nil {
			img.Draw(0, y+62)
		}
		
		exp := self.GetExp()
		img = g_game.GetGuiImage(IMG_GUI_BATTLEEXPBARFILL)
		barwidth = int(math.Floor((float64(exp)/100.0)*float64(img.w))) 
		if img != nil {
			img.DrawRectClip(NewRect(0, y+62, barwidth, int(img.h)), NewRect(0, 0, barwidth, int(img.h)))
		} 
		
		img = g_game.GetPokeImage(self.GetPokeID(), POKEIMAGE_BACK)
		if img != nil {
			img.DrawRect(NewRect(20, 250, int(img.w*4), int(img.h*4)))
		}
	}
	
	//enemy
	enemy := g_game.battle.GetEnemy()
	if enemy != nil {
		font := g_engine.GetFont(FONT_ARIALBLACK_18)
		font.SetColor(255, 255, 255)
		font.DrawTextCentered(enemy.GetPokeName(), 775, 188, y+2)
		
		font = g_engine.GetFont(FONT_ARIALBLACK_16)
		font.SetColor(255, 255, 255)
		font.DrawText(fmt.Sprintf("LVL %d", enemy.GetLevel()), 698, y+6)
		
		hpperc := enemy.GetHPPerc()
		switch {
		case hpperc <= 20:
			img = g_game.GetGuiImage(IMG_GUI_BATTLEHPBAR_RED)
		
		case hpperc > 20 && hpperc <= 40:
			img = g_game.GetGuiImage(IMG_GUI_BATTLEHPBAR_YELLOW)
		
		default:
			img = g_game.GetGuiImage(IMG_GUI_BATTLEHPBAR_GREEN)
		}
		barwidth := int(math.Floor((float64(hpperc)/100.0)*float64(img.w)))  
		img.DrawRectClip(NewRect(770, y+43, barwidth, int(img.h)), NewRect(0, 0, barwidth, int(img.h)))
		
		img = g_game.GetPokeImage(enemy.GetPokeID(), POKEIMAGE_FRONT)
		if img != nil {
			img.DrawRect(NewRect(600, 100, int(img.w*4), int(img.h*4)))
		}
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
				iconFont.DrawText("Bag", dockx+(i*100)+70, docky+19)
				
			case ICON_BATTLE_RUN:
				iconFont.DrawText("Run", dockx+(i*100)+65, docky+19)
			}
		}
	}
}


func (g *PU_BattleUI) MouseDown(_x int, _y int) {

}

func (g *PU_BattleUI) MouseUp(_x int, _y int) {
	if g.visible {
		if g.window == nil {
			dockRect := NewRect(546, 608, 417, 52)
			if dockRect.Contains(_x, _y) {
				for i := 0; i < 4; i++ {
					iconRect := NewRect(dockRect.x+30+(i*100), dockRect.y+12, 80, 29)
					if iconRect.Contains(_x, _y) {
						switch i {
						case ICON_BATTLE_POKEMON:
							g.OpenWindow(BATTLEWINDOW_POKEMON)
			
						case ICON_BATTLE_ATTACK:
							g.OpenWindow(BATTLEWINDOW_ATTACK)
			
						case ICON_BATTLE_BAG:
							g.OpenWindow(BATTLEWINDOW_ITEMS)
						}
					}
				}
			}
		}
	}
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

