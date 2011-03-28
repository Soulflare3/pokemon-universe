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

import "fmt"

const (
	ICON_FRIENDS = 0
	ICON_PARTY = 1
	ICON_POKEDEX = 2
	ICON_BACKPACK = 3
	ICON_POKEMON = 4
	ICON_CHARACTER = 5
	ICON_MENU = 6
	ICON_OPTIONS = 7
	ICON_POKEGEAR = 8
)

type PU_GameUI struct {
	PU_GuiElement
}

func NewGameUI() *PU_GameUI {
	ui := &PU_GameUI{}
	ui.visible = false
	g_gui.AddElement(ui)
	
	return ui
}

func (g *PU_GameUI) Draw() {
	if !g.visible {
		return
	}
	
	dockx, docky := 546, 608
	img := g_game.GetGuiImage(IMG_GUI_ICONDOCK)
	if img != nil {
		img.Draw(dockx, docky)
	}
	
	img = g_game.GetGuiImage(IMG_GUI_ICON_POKEBALL)
	if img != nil {
		iconFont := g_engine.GetFont(FONT_PURITANBOLD_10)
		iconFont.SetColor(255,255,255)
		
		for i := 0; i < 9; i++ {
			img.Draw(dockx+10+(i*46), docky+8)
			switch i {
			case ICON_FRIENDS:
				iconFont.DrawTextCentered("Friends", dockx+(i*46), 50, docky+37)
				
			case ICON_PARTY:
				iconFont.DrawTextCentered("Party", dockx+(i*46), 50, docky+37)
				
			case ICON_POKEDEX:
				iconFont.DrawTextCentered("Pokedex", dockx+(i*46), 50, docky+37)
				
			case ICON_BACKPACK:
				iconFont.DrawTextCentered("Backpack", dockx+(i*46), 50, docky+37)
				
			case ICON_POKEMON:
				iconFont.DrawTextCentered("Pokemon", dockx+(i*46), 50, docky+37)
				
			case ICON_CHARACTER:
				iconFont.DrawTextCentered("Character", dockx+(i*46), 50, docky+37)
				
			case ICON_MENU:
				iconFont.DrawTextCentered("Menu", dockx+(i*46), 50, docky+37)

			case ICON_OPTIONS:
				iconFont.DrawTextCentered("Options", dockx+(i*46), 50, docky+37)
				
			case ICON_POKEGEAR:
				iconFont.DrawTextCentered("Pokegear", dockx+(i*46), 50, docky+37)
			}
		}
	}
	
	pokecount := 0
	for p := 0; p < NUM_POKEMON; p++ {
		if g_game.self.pokemon[p] != nil {
			pokecount++
		} else {
			break
		}
	}
	for i := 0; i < NUM_POKEMON; i++ {
		x, y := (963-(pokecount*74))+(i*74), 664
		if pokemon := g_game.self.pokemon[i]; pokemon != nil {
			pokemonhp := pokemon.hp
			pokemonmaxhp := pokemon.hpmax
			
			if pokemonhp > 0 {
				var hpbar *PU_Image
				hpperc := int((float32(pokemonhp)/float32(pokemonmaxhp))*100.0)
				switch {
				case hpperc <= 20:
					hpbar = g_game.GetGuiImage(IMG_GUI_POKEMON_REDHPBAR)
					
				case hpperc > 20 && hpperc <= 40:
					hpbar = g_game.GetGuiImage(IMG_GUI_POKEMON_YELLOWHPBAR)
					
				default:
					hpbar = g_game.GetGuiImage(IMG_GUI_POKEMON_GREENHPBAR)
				}
				
				var temp *PU_Image
				if i == 0 {
					temp = g_game.GetGuiImage(IMG_GUI_POKEMON_SELECTED)
				} else {
					temp = g_game.GetGuiImage(IMG_GUI_POKEMON)
				}
				if temp != nil {
					temp.Draw(x, y)
				}
				
				temp = g_game.GetGuiImage(IMG_GUI_POKEMON_HPBAR)
				if temp != nil {
					temp.Draw(x+6, y+27)
				}
				
				if hpbar != nil {
					hpbarwidth := int((float32(hpperc)/100.0)*float32(hpbar.w))
					hpbar.DrawRectInRect(NewRect(x+6, y+27, int(hpbarwidth)+2, int(hpbar.h)), NewRect(0, 0, hpbarwidth, int(hpbar.h)))
				}
			} else {
				var temp *PU_Image
				temp = g_game.GetGuiImage(IMG_GUI_POKEMON)
				if temp != nil {
					temp.Draw(x, y)
				}
			
				temp = g_game.GetGuiImage(IMG_GUI_POKEMON_HPBAR)
				if temp != nil {
					temp.Draw(x+6, y+27)
				}			
			}
			
			//icon here
			
			font := g_engine.GetFont(FONT_ARIALBLACK_9)
			font.SetColor(255,255,255)
			font.DrawText(pokemon.name, x+6, y+2)
			
			font = g_engine.GetFont(FONT_ARIALBLACK_8)
			font.SetColor(255,255,255)
			font.DrawText(fmt.Sprintf("%d", pokemon.level), x+6, y+16)
		} else {
			temp := g_game.GetGuiImage(IMG_GUI_POKEMON)
			if temp != nil {
				temp.SetAlphaMod(100)
				temp.Draw(x, y)
				temp.SetAlphaMod(255)
			}
		}
	}
}

func (g *PU_GameUI) MouseDown(_x int, _y int) {

}

func (g *PU_GameUI) MouseUp(_x int, _y int) {

}

func (g *PU_GameUI) MouseMove(_x int, _y int) {

}

func (g *PU_GameUI) MouseScroll(_dir int) {

}

func (g *PU_GameUI) Focusable() bool {
	return false
}

func (g *PU_GameUI) KeyDown(_keysym int, _scancode int) {

}

