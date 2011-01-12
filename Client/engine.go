package main

import (
	"sdl"
	"fmt"
	list "container/list"
)

const (
    WINDOW_WIDTH = 964
    WINDOW_HEIGHT = 720
)

type IResource interface {
	Release()
}

type PU_Engine struct {
	resourceList *list.List
	window *sdl.Window
}

func NewEngine() *PU_Engine {
	return &PU_Engine{resourceList : list.New()}
}

func (e *PU_Engine) Init() {
	//Create the window
	var err string
   	e.window, err = sdl.CreateWindow("Pokemon Universe", WINDOW_WIDTH, WINDOW_HEIGHT)
    if err != "" {
        fmt.Printf("Error in CreateWindow: %v", err) 
        return
    }

	//Try to find and use OpenGL
	rendererIndex := 0
	numRenderers := sdl.GetNumRenderDrivers()
	for i := 0; i < numRenderers; i++ {
		rendererName := sdl.GetRenderDriverName(i)	
		if rendererName == "opengl" {
			rendererIndex = i		
		}
	}
	sdl.CreateRenderer(e.window, rendererIndex)
	sdl.SelectRenderer(e.window)
	
	sdl.InitTTF();
}

func (e *PU_Engine) Exit() {
	//Release all resources
	for i := e.resourceList.Front(); i != nil; i = i.Next() {
		res, valid := i.Value.(IResource)
		if valid {
			res.Release()
		}
	}

	//Destroy the window
	sdl.DestroyWindow(e.window)
	
	//Quit SDL ttf
	sdl.QuitTTF()
} 

func (e *PU_Engine) LoadImage(_file string) *PU_Image {
	image := NewImage(_file)
	e.resourceList.PushBack(image)
	return image
}

func (e *PU_Engine) LoadFont(_file string, _size int) *PU_Font {
	font := NewFont(_file, _size)
	e.resourceList.PushBack(font)
	return font
}

