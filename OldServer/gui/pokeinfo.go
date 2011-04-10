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

//import "fmt"

var POKEINFO_RECT *PU_Rect = NewRect(331, 354, 631, 257)

type PU_PokeInfo struct {
	PU_GuiElement
	
	pokemon int
}

func NewPokeInfo() *PU_PokeInfo {
	info := &PU_PokeInfo{}
	info.visible = true
	g_gui.AddElement(info)
	
	return info
}

func (g *PU_PokeInfo) Draw() {
	if !g.visible {
		return
	}
	
	if pokemon := g_game.self.pokemon[g.pokemon]; pokemon != nil {
		img := g_game.GetGuiImage(IMG_GUI_POKEINFOBG)
		if img != nil {
			img.Draw(POKEINFO_RECT.x, POKEINFO_RECT.y)
		}
		/*pokemonhp := pokemon.hp
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
		}*/
	} 
}

func (g *PU_PokeInfo) MouseDown(_x int, _y int) {

}

func (g *PU_PokeInfo) MouseUp(_x int, _y int) {

}

func (g *PU_PokeInfo) MouseMove(_x int, _y int) {

}

func (g *PU_PokeInfo) MouseScroll(_dir int) {

}

func (g *PU_PokeInfo) Focusable() bool {
	return false
}

func (g *PU_PokeInfo) KeyDown(_keysym int, _scancode int) {

}

