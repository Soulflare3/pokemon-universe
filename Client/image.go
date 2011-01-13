package main

import (
	"sdl"
	"os"
	"exec"
	"path"
)

type PU_Image struct {
	w, h uint16
	alpha uint8
	blendmode int
	
	surface *sdl.Surface
	texture *sdl.Texture
}

func NewImage(_file string) *PU_Image {
	image := &PU_Image{}

	file, _ := exec.LookPath(os.Args[0])
	dir, _ := path.Split(file)
	os.Chdir(dir)
	path, _ := os.Getwd()
	image.surface = sdl.LoadImage(path + "/" + _file)

	image.Reload()
	return image
}

func NewImageFromSurface(_surface *sdl.Surface) *PU_Image {
	image := &PU_Image{surface: _surface}
	image.Reload()
	return image
}

func (i *PU_Image) Reload() {
	i.texture = i.surface.CreateTexture()
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
	i.texture.RenderCopy(_src, _dst)
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

