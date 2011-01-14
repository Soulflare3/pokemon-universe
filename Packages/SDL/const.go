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

/*
Simple SDL 1.3 wrapper
The goal is not to make a complete SDL wrapper, but to wrap only the SDL functions that the PU client needs.
*/
package sdl

const (
    SDL_INIT_VIDEO = 0x00000020
    SDL_INIT_TIMER = 0x00000001

    SDL_WINDOWPOS_CENTERED = 0x7FFFFFFE
    SDL_WINDOW_SHOWN = 4
    SDL_RENDERER_PRESENTVSYNC = 32
    SDL_RENDERER_ACCELERATED = 64

    SDL_BLENDMODE_NONE = 0x00000000
    SDL_BLENDMODE_MASK = 0x00000001
    SDL_BLENDMODE_BLEND = 0x00000002
    SDL_BLENDMODE_ADD = 0x00000004
    SDL_BLENDMODE_MOD = 0x00000008

	/* Window events */
    SDL_WINDOWEVENT_NONE = 0 
    SDL_WINDOWEVENT_SHOWN = 1
    SDL_WINDOWEVENT_HIDDEN = 2 
    SDL_WINDOWEVENT_EXPOSED = 3
    SDL_WINDOWEVENT_MOVED = 4
    SDL_WINDOWEVENT_RESIZED = 5
    SDL_WINDOWEVENT_MINIMIZED = 6 
    SDL_WINDOWEVENT_MAXIMIZED = 7
    SDL_WINDOWEVENT_RESTORED = 8
    SDL_WINDOWEVENT_ENTER = 9
    SDL_WINDOWEVENT_LEAVE = 10
    SDL_WINDOWEVENT_FOCUS_GAINED = 11
    SDL_WINDOWEVENT_FOCUS_LOST = 12
    SDL_WINDOWEVENT_CLOSE = 13
    SDL_QUIT = 0x100
    SDL_WINDOWEVENT = 0x200
    SDL_SYSWMEVENT = 0x201

    /* Keyboard events */
    SDL_KEYDOWN = 0x300
    SDL_KEYUP = 0x301 
    SDL_TEXTEDITING = 0x302
    SDL_TEXTINPUT = 0x303  

    /* Mouse events */
    SDL_MOUSEMOTION = 0x400
    SDL_MOUSEBUTTONDOWN = 0x401
    SDL_MOUSEBUTTONUP = 0x402 
    SDL_MOUSEWHEEL = 0x403
)
