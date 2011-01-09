package main

import (
	"sdl"
	"fmt"
	"os"
	"exec"
	"path"
)

const (
    WINDOW_WIDTH = 964
    WINDOW_HEIGHT = 720
)

func InitEngine() {
	//Create the window
   	window, err := sdl.CreateWindow("Pokemon Universe", WINDOW_WIDTH, WINDOW_HEIGHT)
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
	sdl.CreateRenderer(window, rendererIndex)
}

func LoadImage(_file string) *sdl.Surface {
	file, _ := exec.LookPath(os.Args[0])
    dir, _ := path.Split(file)
    os.Chdir(dir)
    path, _ := os.Getwd()
	return sdl.LoadImage(path+"/"+_file)
}
