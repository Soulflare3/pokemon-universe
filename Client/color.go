package main

import "sdl"

//Colorcodes used by the server
const (
	COLOR_WHITE = iota
	COLOR_YELLOW
	COLOR_GRAY
	COLOR_RED
	COLOR_GREEN
	COLOR_ORANGE
	COLOR_PURPLE
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

func ColorCodeToKey(_code int) uint32 {
	switch _code {
		case COLOR_WHITE:
			return CreateColorKey(255,255,255)

		case COLOR_YELLOW:
			return CreateColorKey(255,242,0)

		case COLOR_GRAY:
			return CreateColorKey(137,137,137)

		case COLOR_RED:
			return CreateColorKey(255,0,0)

		case COLOR_GREEN:
			return CreateColorKey(0,255,0)

		case COLOR_ORANGE:
			return CreateColorKey(242,108,79)

		case COLOR_PURPLE:
			return CreateColorKey(210,53,255)
			
	}
	return CreateColorKey(255,255,255)
}

