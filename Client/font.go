package main

import (
	"fmt"
	"sdl"
)

func CreateColorKey(_red uint8, _green uint8, _blue uint8) uint32 {
	r, g, b := uint32(_red), uint32(_green), uint32(_blue)
	return (r << 16) | (g << 8) | (b)
}

func ColorKeyToSDL(_color uint32) sdl.Color {
	var sdlcolor sdl.Color
	sdlcolor.R = (uint8)(_color >> 16);
	sdlcolor.G = (uint8)(_color >> 8);
	sdlcolor.B = (uint8)(_color);
	return sdlcolor;
}

type PU_Font struct {
	font *sdl.Font
	fontmap map[uint32]map[uint16]*PU_Image
	color uint32
}

func NewFont(_file string, _size int) *PU_Font {
	sdlfont := sdl.LoadFont(_file, _size)
	if sdlfont == nil {
		fmt.Printf("Error loading Font: %v", sdl.GetError())
	}
	f := &PU_Font{font : sdlfont,
				  fontmap : make(map[uint32]map[uint16]*PU_Image)}
	f.Build()
	return f	
}

func (f *PU_Font) SetColor(_red uint8, _green uint8, _blue uint8) {
	f.color = CreateColorKey(_red, _green, _blue)
}

func (f *PU_Font) Build() {
	f.fontmap[f.color] = make(map[uint16]*PU_Image)
	for c := 32; c <= 127; c++ {
		surface := f.font.RenderText_Blended(fmt.Sprintf("%c",c), ColorKeyToSDL(f.color))
		img := NewImageFromSurface(surface)
		
		f.fontmap[f.color][uint16(c)] = img
	}
}

func (f *PU_Font) DrawText(_text string, _x int, _y int) {
	_, present := f.fontmap[f.color]
	if !present {
		f.Build()
	}
	
	prev_char := -1
	for c := 0; c < len(_text); c++ {
		_, _, _, _, advance := f.font.GetMetrics(uint16(_text[c]))
		
		if prev_char != -1 {
			kerning := f.font.GetKerning(prev_char, int(_text[c]))
			_x += kerning
			
			prev_char = int(_text[c])
		}
		
		img := f.fontmap[f.color][uint16(_text[c])]
		if img != nil {
			img.Draw(_x, _y)
			_x += advance
		}
	}
}

