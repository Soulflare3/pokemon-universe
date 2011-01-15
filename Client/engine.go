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
	fonts map[int]*PU_Font
	window *sdl.Window
}

func NewEngine() *PU_Engine {
	return &PU_Engine{resourceList : list.New(),
					  fonts : make(map[int]*PU_Font)}
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

func (e *PU_Engine) AddResource(_res IResource) {
	e.resourceList.PushBack(_res)
}

func (e *PU_Engine) LoadImage(_file string) *PU_Image {
	image := NewImage(_file)
	e.resourceList.PushBack(image)
	return image
}

func (e *PU_Engine) LoadFont(_id int, _file string, _size int) *PU_Font {
	font := NewFont(_file, _size)
	e.resourceList.PushBack(font)
	e.fonts[_id] = font
	return font
}

func (e *PU_Engine) GetFont(_id int) *PU_Font {
	if font, present := e.fonts[_id]; present {
		return font
	}
	return nil
}

