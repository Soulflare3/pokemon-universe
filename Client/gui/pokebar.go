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
	"math"
)

const (
	POKEBAR_NONE = -1
)

type PU_Pokebar struct {
	PU_GuiElement
	
	mouseDown bool
	offset int
	dragStart int
	dragging int
	dragX int
}

func NewPokebar() *PU_Pokebar {
	bar := &PU_Pokebar{}
	bar.visible = true
	g_gui.AddElement(bar)
	
	bar.dragging = POKEBAR_NONE
	
	return bar
}

func (g *PU_Pokebar) Draw() {
	if !g.visible {
		return
	}
	
	pokecount := g_game.self.GetPokemonCount()
	for i := 0; i < NUM_POKEMON; i++ {
		x, y := (963-(pokecount*74))+(i*74), 664
		if g.dragging == i {
			x = g.dragX
		} 
		if pokemon := g_game.self.pokemon[i]; pokemon != nil {
			pokemonhp := pokemon.hp
			pokemonmaxhp := pokemon.hpmax
			
			if pokemonhp > 0 {
				var hpbar *PU_Image
				hpperc := int(math.Floor((float64(pokemonhp)/float64(pokemonmaxhp))*100.0))
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
					hpbar.DrawRectClip(NewRect(x+6, y+27, int(hpbarwidth)+2, int(hpbar.h)), NewRect(0, 0, hpbarwidth, int(hpbar.h)))
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
			
			icon := g_game.GetPokeImage(int(pokemon.id), POKEIMAGE_ICON)
            if icon != nil {
                    icon.Draw(x+42, y+4)
            }
			
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

func (g *PU_Pokebar) MouseDown(_x int, _y int) {
	g.mouseDown = true
	pokecount := 0
	if g_game.self != nil {
		pokecount = g_game.self.GetPokemonCount()
	}

	for i := 0; i < NUM_POKEMON; i++ {
		x, y, w, h := (963-(pokecount*74))+(i*74), 664, 74, 39
		if _x >= x && _x <= x+w {
			if _y >= y && _y <= y+h {
				if self, pkmn := g_game.self, g_game.self.pokemon[i]; self != nil && pkmn != nil {
					g.dragX = x
					g.offset = _x - x
					g.dragging = i
					g.dragStart = i
				}
				return;
			}
		}
	}
	g.dragging = POKEBAR_NONE
}

func (g *PU_Pokebar) MouseUp(_x int, _y int) {
	if g.dragStart != g.dragging && g.dragging != POKEBAR_NONE {
		//g_conn.Game().SendSlotChange(g.dragStart, g.dragging)
	}
	g.mouseDown = false
	g.dragging = POKEBAR_NONE
}

func (g *PU_Pokebar) MouseMove(_x int, _y int) {
	if g.mouseDown {
		if g.dragging != POKEBAR_NONE {
			pokecount := 0
			if g_game.self != nil {
				pokecount = g_game.self.GetPokemonCount()
			}
			
			x, w := (963-(pokecount*74))+(g.dragging*74), 74
			g.dragX = _x-g.offset
			midx := g.dragX+37

			for i := 0; i < NUM_POKEMON; i++ {
				x = (963-(pokecount*74))+(i*74)
				if midx >= x && midx <= x+w {
					if i != g.dragging {
						if self, pkmn := g_game.self, g_game.self.pokemon[i]; self != nil && pkmn != nil {
							dragPokemon := g_game.self.pokemon[g.dragging]
							g_game.self.pokemon[g.dragging] = g_game.self.pokemon[i]
							g_game.self.pokemon[i] = dragPokemon

							g.dragging = i;
						}
						return;
					}
				}
			}
		}
	} 
}

func (g *PU_Pokebar) MouseScroll(_dir int) {

}

func (g *PU_Pokebar) Focusable() bool {
	return false
}

func (g *PU_Pokebar) KeyDown(_keysym int, _scancode int) {

}

