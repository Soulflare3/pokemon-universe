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

//#include "SDL.h"
import "C" 
import "unsafe"

type cast unsafe.Pointer

func Delay(_ticks uint32) {
	C.SDL_Delay(C.Uint32(_ticks))
}

func GetTicks() uint32 { 
    return uint32(C.SDL_GetTicks())
}   

func Quit() {
	C.SDL_Quit()
}

func GetError() (ret string) {
	ret = C.GoString(C.SDL_GetError())
	C.SDL_ClearError()	
	return
}

func KeyDown(_key int) bool {
	zero := C.int(0)
	var state = uintptr(unsafe.Pointer(C.SDL_GetKeyboardState(&zero)))+uintptr(_key)
	down := (*uint8)(cast(state))
	if *down == 1 {
		return true
	}
	return false
}

func Init() (error string) {
	flags := int64(C.SDL_INIT_VIDEO)
    if C.SDL_Init(C.Uint32(flags)) != 0 {
        error = C.GoString(C.SDL_GetError())
        return
    }
    return ""
}

type Window struct { 
	window *C.SDL_Window
}

func CreateWindow(_title string, _width int, _height int) (*Window, string) {
	ctitle := C.CString(_title)
    var window *C.SDL_Window = C.SDL_CreateWindow(ctitle, 
					 							  C.SDL_WINDOWPOS_CENTERED, 
												  C.SDL_WINDOWPOS_CENTERED,
												  C.int(_width), 
												  C.int(_height), 
												  C.SDL_WINDOW_SHOWN | C.SDL_WINDOW_OPENGL)
	C.free(unsafe.Pointer(ctitle))
    if window == nil {
        return nil, GetError()
    }
    return &Window{window}, ""
}

func DestroyWindow(_window *Window) {
	C.SDL_DestroyWindow(_window.window)
}

func PollEvent() (*SDLEvent, bool) {
	var ev *SDLEvent = &SDLEvent{}
    if C.SDL_PollEvent((*C.SDL_Event)(cast(ev))) != 0 {
		return ev, true
    }
	return nil, false
}


