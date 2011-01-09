package main

import (
	"sdl"
	"os"
	"exec"
	"path"
	"fmt"
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

	image.texture = image.surface.CreateTexture()
	image.w = uint16(image.texture.W)
	image.h = uint16(image.texture.H)

	fmt.Printf("x:%v h:%v\n", image.w, image.h)

	return image
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
