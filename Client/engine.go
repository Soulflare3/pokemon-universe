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

type PU_Engine struct {
	imageList *list.List
	window *sdl.Window
}

func NewEngine() *PU_Engine {
	return &PU_Engine{imageList : list.New()}
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
	//Release all image resources
	for i := e.imageList.Front(); i != nil; i = i.Next() {
		image, valid := i.Value.(*PU_Image)
		if valid {
			image.Release()
		}
	}

	//Destroy the window
	sdl.DestroyWindow(e.window)
	
	//Quit SDL ttf
	sdl.QuitTTF()
} 

func (e *PU_Engine) LoadImage(_file string) *PU_Image {
	image := NewImage(_file)
	e.imageList.PushBack(image)
	return image
}
