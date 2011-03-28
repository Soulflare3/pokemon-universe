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
	"sdl"
)

type PU_Image struct {
	w, h uint16
	alpha uint8
	blendmode int
	
	surface *sdl.Surface
	texture *sdl.Texture
}

func NewImage(_file string) *PU_Image {
	image := &PU_Image{blendmode : sdl.BLENDMODE_BLEND}
	
	image.surface = sdl.LoadImage(GetPath() + _file)

	image.Reload()
	return image
}

func NewImageFromSurface(_surface *sdl.Surface) *PU_Image {
	image := &PU_Image{surface: _surface,
					   blendmode : sdl.BLENDMODE_BLEND}
	image.Reload()
	return image
}

func (i *PU_Image) Reload() {
	i.texture = i.surface.CreateTexture(g_engine.renderer)
	i.w = uint16(i.texture.W)
	i.h = uint16(i.texture.H)
}

func (i *PU_Image) Release() {
	if i.surface != nil {
		i.surface.Release()
	}
	if i.texture != nil {
		i.texture.Release()
	}
}

func (i *PU_Image) SetBlendMode(_blendmode int) {
	i.blendmode = _blendmode
	i.texture.SetBlendMode(_blendmode)
}

func (i *PU_Image) SetColorMod(_red uint8, _green uint8, _blue uint8) {
	i.texture.SetColorMod(_red, _green, _blue)
}

func (i *PU_Image) SetAlphaMod(_alpha uint8) {
	i.texture.SetAlpha(_alpha)
}

func (i *PU_Image) Render(_src *sdl.Rect, _dst *sdl.Rect) {
	i.texture.SetBlendMode(i.blendmode)
	i.texture.RenderCopy(g_engine.renderer, _src, _dst)
}

func (i *PU_Image) Draw(_x int, _y int) {
	src := &sdl.Rect{0, 0, int32(i.w), int32(i.h)}
	dst := &sdl.Rect{int32(_x), int32(_y), int32(i.w), int32(i.h)}

	i.Render(src, dst)
}

func (i *PU_Image) DrawRect(_rect *PU_Rect) {
	src := &sdl.Rect{0, 0, int32(i.w), int32(i.h)}

	i.Render(src, _rect.ToSDL())
}

func (i *PU_Image) DrawClip(_x int, _y int, _clip *PU_Rect) {
	dst := &sdl.Rect{0, 0, int32(i.w), int32(i.h)}
	
	i.Render(_clip.ToSDL(), dst)
}

func (i *PU_Image) DrawRectClip(_rect *PU_Rect, _clip *PU_Rect) {
	i.Render(_clip.ToSDL(), _rect.ToSDL())
}

func (i *PU_Image) DrawInRect(_x int, _y int, _inrect *PU_Rect) {
	imgRect := NewRect(_x, _y, int(i.w), int(i.h))
	inRect := _inrect.Intersection(imgRect)
	dstRect := NewRect(inRect.x, inRect.y, inRect.width, inRect.height)
	inRect.x -= _x
	inRect.y -= _y
	
	i.Render(inRect.ToSDL(), dstRect.ToSDL())
}

func (i *PU_Image) DrawRectInRect(_rect *PU_Rect, _inrect *PU_Rect) {
	inRect := _inrect.Intersection(_rect)
	dstRect := NewRect(inRect.x, inRect.y, inRect.width, inRect.height)
	inRect.x -= _rect.x
	inRect.y -= _rect.y
	
	i.Render(inRect.ToSDL(), dstRect.ToSDL())
}

