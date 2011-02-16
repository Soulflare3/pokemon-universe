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
	renderer *sdl.Renderer
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

	//Find our available renderers
	openglIndex := 0
	d3dIndex := -1
	numRenderers := sdl.GetNumRenderDrivers()
	for i := 0; i < numRenderers; i++ {
		rendererName := sdl.GetRenderDriverName(i)	
		println(rendererName)
		if rendererName == "opengl" {
			openglIndex = i		
		} else if rendererName == "direct3d" {
			d3dIndex = i
		}
	}
	
	//Default renderer is OpenGL
	rendererIndex := openglIndex
	
	//If we found DirectX (on Windows), use that
	if d3dIndex != -1 {
		rendererIndex = d3dIndex
	}
	
	e.renderer, err = sdl.CreateRenderer(e.window, rendererIndex)
	if err != "" {
		fmt.Printf("Error in CreateRenderer: %v", err) 
		return
	}
	
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
	
	//Destroy the renderer
	e.renderer.Release()

	//Destroy the window
	sdl.DestroyWindow(e.window)
	
	//Quit SDL ttf
	sdl.QuitTTF()
} 

func (e *PU_Engine) DrawFillRect(_rect *PU_Rect, _color *sdl.Color, _alpha uint8) {
	sdl.SetRenderDrawColor(e.renderer, _color.R, _color.G, _color.B, _alpha)
	sdl.SetRenderDrawBlendMode(e.renderer, sdl.SDL_BLENDMODE_BLEND)
	sdl.RenderFillRect(e.renderer, *_rect.ToSDL())
	sdl.SetRenderDrawBlendMode(e.renderer, sdl.SDL_BLENDMODE_NONE)
	sdl.SetRenderDrawColor(e.renderer, 0, 0, 0, 255)
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

