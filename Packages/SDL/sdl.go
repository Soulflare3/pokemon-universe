Pokemon Universe MMORPG
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
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.

/*
Simple SDL 1.3 wrapper
The goal is not to make a complete SDL wrapper, but to wrap only the SDL functions that the PU client needs.
*/

package sdl

//#include "SDL.h"
import "C"
import "unsafe"

type cast unsafe.Pointer

func GetTicks() uint32 { 
    return uint32(C.SDL_GetTicks())
}   

func Init() (error string) {
    if C.SDL_Init(SDL_INIT_VIDEO|SDL_INIT_TIMER) != 0 {
        error = C.GoString(C.SDL_GetError())
        return
    }
    return ""
}

func CreateWindow(_title string, _width int, _height int) (error string) {
    var window *C.struct_SDL_Window = C.SDL_CreateWindow(C.CString(_title), 
							 							 SDL_WINDOWPOS_CENTERED, 
															 SDL_WINDOWPOS_CENTERED,
														 C.int(_width), 
														 C.int(_height), 
														 SDL_WINDOW_SHOWN)
    if window == nil {
        error = C.GoString(C.SDL_GetError())
        return
    }
	
	//TODO: Make this piece actually work so that OpenGL is explicitly used (or possibly D3D on Windows)
	numRenderers := int(C.SDL_GetNumRenderDrivers())
	for i := 0; i < numRenderers; i++ {
		var rendererInfo *RendererInfo = &RendererInfo{}
		C.SDL_GetRenderDriverInfo(C.int(i), (*C.SDL_RendererInfo)(cast(rendererInfo)));
		
	}
	
    if C.SDL_CreateRenderer(window, 0, SDL_RENDERER_PRESENTVSYNC | SDL_RENDERER_ACCELERATED) != 0 {
        error = C.GoString(C.SDL_GetError())
        return
    }
    return ""
}

type RendererInfo struct {
    name *byte
    flags uint32
    mod_modes uint32
    blend_modes uint32
    scale_modes uint32
    num_texture_formats uint32
    texture_formats[50] [50]uint32
    max_texture_width int32
    max_texture_height int32
}

func StartEventLoop(_events chan<- *SDLEvent) {
    var ev *SDLEvent = &SDLEvent{}
    for{
        if C.SDL_WaitEvent((*C.SDL_Event)(cast(ev))) != 0 {
	    _events <- ev
        }
    }
}

type SDLEvent struct {
    Evtype uint32
    rest [48]byte
}

func (e *SDLEvent) Window() *WindowEvent {
    return (*WindowEvent)(cast(e))
}

func (e *SDLEvent) Keyboard() *KeyboardEvent {
    return (*KeyboardEvent)(cast(e))
}

func (e *SDLEvent) TextEdit() *TextEditingEvent {
    return (*TextEditingEvent)(cast(e))
}

func (e *SDLEvent) TextInput() *TextInputEvent {
    return (*TextInputEvent)(cast(e))
}

func (e *SDLEvent) MouseMotion() *MouseMotionEvent {
    return (*MouseMotionEvent)(cast(e))
}

func (e *SDLEvent) MouseButton() *MouseButtonEvent {
    return (*MouseButtonEvent)(cast(e))
}

func (e *SDLEvent) MouseWheel() *MouseWheelEvent {
    return (*MouseWheelEvent)(cast(e))
}

func (e *SDLEvent) Quit() *QuitEvent {
    return (*QuitEvent)(cast(e))
}

func (e *SDLEvent) User() *UserEvent {
    return (*UserEvent)(cast(e))
}

func (e *SDLEvent) SysWM() *SysWMEvent {
    return (*SysWMEvent)(cast(e))
}

type WindowEvent struct {
    Evtype uint32
    WindowID uint32
    Event uint8
    Padding1 uint8
    Padding2 uint8
    Padding3 uint8
    Data1 int32
    Data2 int32
}

type KeyboardEvent struct {
    Evtype uint32
    WindowID uint32
    State uint8
    Repeat uint8
    Padding2 uint8
    Padding3 uint8
    Keysym uint32
}

type TextEditingEvent struct {
    Evtype uint32
    WindowID uint32
    Text [32]byte
    Start int32
    Length int32
}

type TextInputEvent struct {
    Evtype uint32
    WindowID uint32
    Text [32]byte
}

type MouseMotionEvent struct {
    Evtype uint32
    WindowID uint32
    State uint8
    Padding1 uint8
    Padding2 uint8
    Padding3 uint8
    X int32
    Y int32
    Xrel int32
    Yrel int32
}

type MouseButtonEvent struct {
    Evtype uint32
    WindowID uint32
    Button uint8
    State uint8
    Padding1 uint8
    Padding2 uint8
    X int32
    Y int32
}

type MouseWheelEvent struct {
    Evtype uint32
    WindowID uint32
    X int32
    Y int32
}

type QuitEvent struct {
    Evtype uint32
}

type UserEvent struct {
    Evtype uint32
    WindowID uint32
    Code int32
    Data1 *[0]byte
    Cata2 *[0]byte
}

type SysWMEvent struct {
    Evtype uint32
}

const (
    SDL_INIT_VIDEO = 0x00000020
    SDL_INIT_TIMER = 0x00000001

    SDL_WINDOWPOS_CENTERED = 0x7FFFFFFE
    SDL_WINDOW_SHOWN = 4
    SDL_RENDERER_PRESENTVSYNC = 32
    SDL_RENDERER_ACCELERATED = 64

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
