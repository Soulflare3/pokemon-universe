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
import (
	"unsafe"
	"fmt"
	"os"
	"bufio"
	"image"
	"image/png"
)

type cast unsafe.Pointer

func GetTicks() uint32 { 
    return uint32(C.SDL_GetTicks())
}   

func Quit() {
	C.SDL_Quit()
}

func Init() (error string) {
    if C.SDL_Init(SDL_INIT_VIDEO|SDL_INIT_TIMER) != 0 {
        error = C.GoString(C.SDL_GetError())
        return
    }
    return ""
}

func CreateWindow(_title string, _width int, _height int) (error string) {
	ctitle := C.CString(_title)
    var window *C.struct_SDL_Window = C.SDL_CreateWindow(ctitle, 
							 							 SDL_WINDOWPOS_CENTERED, 
														 SDL_WINDOWPOS_CENTERED,
														 C.int(_width), 
														 C.int(_height), 
														 SDL_WINDOW_SHOWN)
	C.free(unsafe.Pointer(ctitle))
    if window == nil {
        error = C.GoString(C.SDL_GetError())
        return
    }
	
	numRenderers := int(C.SDL_GetNumRenderDrivers())
	for i := 0; i < numRenderers; i++ {
		var rendererInfo *RendererInfo = &RendererInfo{}
		C.SDL_GetRenderDriverInfo(C.int(i), (*C.SDL_RendererInfo)(cast(rendererInfo)));

		strname := ""
		for c := 0;; c++ { 
			var name = uintptr(unsafe.Pointer(rendererInfo.name))+uintptr(c)
			ch := (*uint8)(cast(name))
			if *ch == uint8(0) {
				break
			}
			strname += string(*ch)	
		}
		fmt.Printf("Renderer: %v\n", strname)
	}
	
    if C.SDL_CreateRenderer(window, 1, SDL_RENDERER_PRESENTVSYNC | SDL_RENDERER_ACCELERATED) != 0 {
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
    texture_formats [50]uint32
    max_texture_width int32
    max_texture_height int32
}

func PollEvent() (*SDLEvent, bool) {
	var ev *SDLEvent = &SDLEvent{}
    if C.SDL_PollEvent((*C.SDL_Event)(cast(ev))) != 0 {
		return ev, true
    }
	return nil, false
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

type Surface struct {
	Flags uint32
	Format *PixelFormat
	W int32
	H int32
	Pitch uint16
	Pad0 [2]byte
	Pixels *byte
	Offset int32
	Hwdata *[0]byte
	Clip_rect Rect
	Unused1 uint32
	Locked uint32
	Map *[0]byte
	Format_version uint32
	Refcount int32
}

func (s *Surface) Get() *C.SDL_Surface {
	return (*C.SDL_Surface)(cast(s))
}

func (s *Surface) Release() {
	C.SDL_FreeSurface(s.Get())
}

func (s *Surface) CreateTexture() *Texture {
	return &Texture{texture : C.SDL_CreateTextureFromSurface(C.Uint32(0), s.Get())}
}

func (s *Surface) DisplayFormatAlpha(_surface *Surface) {
	surface := (*Surface)(cast(C.SDL_DisplayFormatAlpha(_surface.Get())))
	s = surface
}

func (s *Surface) SaveBMP(_file string) {
	cfile := C.CString(_file); defer C.free(unsafe.Pointer(cfile))
	cparams := C.CString("wb"); defer C.free(unsafe.Pointer(cparams))
	C.SDL_SaveBMP_RW(s.Get(), C.SDL_RWFromFile(cfile, cparams), C.int(1))  
}

type Texture struct {
	texture *C.SDL_Texture
	Alpha uint8
}

func (t *Texture) Get() *C.SDL_Texture {
	return t.texture 
}

func (t *Texture) Release() {
	C.SDL_DestroyTexture(t.texture)
} 

func (t *Texture) SetAlpha(_alpha uint8) {
	t.Alpha = _alpha
	C.SDL_SetTextureAlphaMod(t.texture, C.Uint8(_alpha))
}

func (t *Texture) SetScaleMode(_mode int) {
	C.SDL_SetTextureScaleMode(t.texture, C.SDL_ScaleMode(_mode))
}

func (t *Texture) RenderCopy(_srcrect Rect, _dstrect Rect) {
	src := (*C.SDL_Rect)(cast(&_srcrect))
	dst := (*C.SDL_Rect)(cast(&_dstrect))
	C.SDL_RenderCopy(t.texture, src, dst)
}

func RenderClear() {
	C.SDL_RenderClear()
}

func RenderPresent() {
	C.SDL_RenderPresent()
}

func LoadImage(_file string) *Surface {
	img := LoadPNG(_file)
	if img != nil {
		bpp := 0
		if _, is_type := img.(*image.RGBA); is_type {
			bpp = 3
		} else if _, is_type := img.(*image.NRGBA); is_type {
			bpp = 4
		}
		fmt.Printf("Bpp: %v", bpp)
		
		width := img.Bounds().Size().X
		height := img.Bounds().Size().Y
		depth := bpp*8

		var sf = C.SDL_CreateRGBSurface(C.SDL_HWSURFACE, C.int(width), C.int(height), C.int(depth), C.Uint32(0), C.Uint32(0), C.Uint32(0), C.Uint32(0))
		var surface *Surface = (*Surface)(cast(sf))

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				var pixels = uintptr(unsafe.Pointer(surface.Pixels))
				r, g, b, a := img.At(x, y).RGBA()

				color := uint32(C.SDL_MapRGBA(surface.Get().format, C.Uint8(r >> 8), C.Uint8(g >> 8), C.Uint8(b >> 8), C.Uint8(a >> 8)))
				*(*uint32)(unsafe.Pointer(pixels + uintptr(y*int(surface.Pitch)+x<<2))) = color
			}		
		}
		return surface
	}
	return nil
}

func LoadPNG(_file string) image.Image {
	file, err := os.Open(_file, 0, 0) 
	if file == nil {
		fmt.Printf("LoadPNG error: ", err.String())
		return nil
	}
	defer file.Close()
	data := bufio.NewReader(file)
	img, err := png.Decode(data)
	if err != nil {
		fmt.Printf("LoadPNG error: ", err.String())
		return nil
	}
	return img
}

type Rect struct {
	X int16
	Y int16
	W uint16
	H uint16
}

type PixelFormat struct {
	Palette *Palette
	BitsPerPixel uint8
	BytesPerPixel uint8
	Rloss uint8
	Gloss uint8
	Bloss uint8
	Aloss uint8
	Rshift uint8
	Gshift uint8
	Bshift uint8
	Ashift uint8
	Pad0 [2]byte
	Rmask uint32
	Gmask uint32
	Bmask uint32
	Amask uint32
	Colorkey uint32
	Alpha uint8
	Pad1 [3]byte
}

type Palette struct {
	Ncolors int32
	Colors *Color
}

type Color struct {
	R uint8
	G uint8
	B uint8
	Unused uint8
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
