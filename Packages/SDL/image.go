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

/*
#include "SDL.h"
#include "SDL_image.h" 
*/
import "C"
import (
	"unsafe"
	"fmt"
)

type Surface C.SDL_Surface

func (s *Surface) Get() *C.SDL_Surface {
	return (*C.SDL_Surface)(s)
}

func (s *Surface) Release() {
	C.SDL_FreeSurface(s.Get())
}

func (s *Surface) CreateTexture() *Texture {
	var tex = C.SDL_CreateTextureFromSurface(C.Uint32(0), s.Get())
	return (*Texture)(cast(tex))
}

func (s *Surface) DisplayFormatAlpha() *Surface {
	return (*Surface)(C.SDL_DisplayFormatAlpha(s.Get()))
}

func (s *Surface) SaveBMP(_file string) {
	cfile := C.CString(_file); defer C.free(unsafe.Pointer(cfile))
	cparams := C.CString("wb"); defer C.free(unsafe.Pointer(cparams))
	C.SDL_SaveBMP_RW(s.Get(), C.SDL_RWFromFile(cfile, cparams), C.int(1))  
}

func LoadBMP(_file string) *Surface {
	cfile := C.CString(_file); defer C.free(unsafe.Pointer(cfile))
	cparams := C.CString("rb"); defer C.free(unsafe.Pointer(cparams))
	return (*Surface)(C.SDL_LoadBMP_RW(C.SDL_RWFromFile(cfile, cparams), C.int(1)))
}

func LoadImage(_file string) *Surface {
	cfile := C.CString(_file); defer C.free(unsafe.Pointer(cfile))
	img := C.IMG_Load(cfile)
	if img == nil {
		fmt.Printf("Image load error: %v", C.GoString(C.IMG_GetError()))
	}
	return (*Surface)(cast(img))
}

func LoadImageRW(_data *[]byte, _size int) *Surface {
	rawImage := C.SDL_RWFromMem(unsafe.Pointer(&((*_data)[0])), C.int(_size));
	img := C.IMG_Load_RW(rawImage, C.int(0))
	if img == nil {
		fmt.Printf("ImageRW load error: %v", C.GoString(C.IMG_GetError()))
	}
	return (*Surface)(cast(img))
}

type Rect struct {
	X int32
	Y int32
	W int32
	H int32
}

type Color struct {
	R uint8
	G uint8
	B uint8
	Unused uint8
}
