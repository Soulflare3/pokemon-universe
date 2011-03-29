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

const (
	FONT_PURITANBOLD_10 = iota
	FONT_PURITANBOLD_12
	FONT_PURITANBOLD_14
	FONT_ARIALBLACK_8
	FONT_ARIALBLACK_9
	FONT_ARIALBLACK_10
	FONT_ARIALBLACK_14
	FONT_ARIALBLACK_18
	FONT_ARIALBLACK_48
)

type PU_Font struct {
	font *sdl.Font
	fontmap map[uint32]map[uint16]*PU_Image
	color *sdl.Color
	size int
	style uint32
	alpha uint8
}

func NewFont(_file string, _size int) *PU_Font {
	sdlfont := sdl.LoadFont(_file, _size)
	if sdlfont == nil {
		fmt.Printf("Error loading Font: %v", sdl.GetError())
	}
	f := &PU_Font{font : sdlfont,
				  fontmap : make(map[uint32]map[uint16]*PU_Image),
				  color : &sdl.Color{255,255,255,255},
				  alpha : 255,
				  size : _size}
	f.Build()
	return f	
}

func (f *PU_Font) Release() {
	f.font.Release()
}

func (f *PU_Font) SetColor(_red uint8, _green uint8, _blue uint8) {
	f.color.R = _red
	f.color.G = _green
	f.color.B = _blue
}

func (f *PU_Font) SetAlpha(_alpha uint8) {
	f.alpha = _alpha
}

func (f *PU_Font) SetStyle(_bold bool, _italic bool, _underline bool) {
	var b, i, u uint32 = 0, 0, 0
	if _bold {
		b = 1
	}
	if _italic {
		i = 1
	}
	if _underline {
		u = 1
	}
	
	f.style = (b << 16) | (i << 8) | (u)
	
	_, present := f.fontmap[f.style]
	if !present {
		f.font.SetStyle(_bold, _italic, _underline)
		f.Build()
	}
}

func (f *PU_Font) Build() {
	f.fontmap[f.style] = make(map[uint16]*PU_Image)
	for c := 32; c <= 127; c++ {
		surface := f.font.RenderText_Blended(fmt.Sprintf("%c",c), ColorKeyToSDL(ColorCodeToKey(COLOR_WHITE)))
		img := NewImageFromSurface(surface)
		
		f.fontmap[f.style][uint16(c)] = img
	}
}

func (f *PU_Font) DrawText(_text string, _x int, _y int) {
	prev_char := -1
	for c := 0; c < len(_text); c++ {
		_, _, _, _, advance := f.font.GetMetrics(uint16(_text[c]))
		
		if prev_char != -1 {
			kerning := f.font.GetKerning(prev_char, int(_text[c]))
			_x += kerning
			
			prev_char = int(_text[c])
		}
		
		img := f.fontmap[f.style][uint16(_text[c])]
		if img != nil {
			img.SetColorMod(f.color.R, f.color.G, f.color.B)
			img.SetAlphaMod(f.alpha)
			img.Draw(_x, _y)
			_x += advance
		}
	}
}

func (f *PU_Font) DrawBorderedText(_text string, _x int, _y int) {
	r, g, b := f.color.R, f.color.R, f.color.B
	f.SetColor(0, 0 , 0)
	f.DrawText(_text, _x-1, _y-1)
	f.DrawText(_text, _x+1, _y-1)
	f.DrawText(_text, _x-1, _y+1)
	f.DrawText(_text, _x+1, _y+1)
	f.SetColor(r, g, b)
	f.DrawText(_text, _x, _y)
}

func (f *PU_Font) DrawTextInRect(_text string, _x int, _y int, _rect *PU_Rect) {
	prev_char := -1
	for c := 0; c < len(_text); c++ {
		_, _, _, _, advance := f.font.GetMetrics(uint16(_text[c]))
		
		if prev_char != -1 {
			kerning := f.font.GetKerning(prev_char, int(_text[c]))
			_x += kerning
			
			prev_char = int(_text[c])
		}
		
		img := f.fontmap[f.style][uint16(_text[c])]
		if img != nil {
			img.SetColorMod(f.color.R, f.color.G, f.color.B)
			img.SetAlphaMod(f.alpha)
			img.DrawInRect(_x, _y, _rect)
			_x += advance
		}
	}
}

func (f *PU_Font) DrawTextCentered(_text string, _x int, _width int, _y int) {
	 x := (_x+(_width/2))-(f.GetStringWidth(_text)/2);
	 f.DrawText(_text, x, _y)
}

func (f *PU_Font) GetStringWidth(_text string) int {
	w := 0
	prev_char := -1
	for c := 0; c < len(_text); c++ {
		_, _, _, _, advance := f.font.GetMetrics(uint16(_text[c]))
		
		if prev_char != -1 {
			kerning := f.font.GetKerning(prev_char, int(_text[c]))
			w += kerning
			
			prev_char = int(_text[c])
		}
		
		w += advance
	}
	return w
}

func (f *PU_Font) GetStringHeight() int {
	return f.font.GetHeight()-4
}

