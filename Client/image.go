package main

import (
	"sdl"
	"os"
	"exec"
	"path"
)

type PU_Image struct {
	w, h uint16
	surface *sdl.Surface
	texture *sdl.Texture 
}

func NewImage(_file string) *PU_Image {
	image := &PU_Image{}

	file, _ := exec.LookPath(os.Args[0])
    dir, _ := path.Split(file)
    os.Chdir(dir)
    path, _ := os.Getwd()
	image.surface = sdl.LoadImage(path+"/"+_file)

	image.Reload()
	return image
}

func NewImageFromSurface(_surface *sdl.Surface) *PU_Image {
	image := &PU_Image{surface : _surface}
	image.Reload()
	return image
}

func (i *PU_Image) Reload() {
	i.texture = i.surface.CreateTexture();
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

func (i* PU_Image) Draw(_x int, _y int) {
	src := sdl.Rect{0, 0, int32(i.w), int32(i.h)}
	dst := sdl.Rect{int32(_x), int32(_y), int32(i.w), int32(i.h)}

	i.texture.RenderCopy(src, dst)
}
