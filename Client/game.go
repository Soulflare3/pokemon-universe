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
	"fmt"
	"sdl"
	"io/ioutil"
	"strings"
	"strconv"
)

type PU_Game struct {
	tileImageMap map[uint16]*PU_Image
}

func NewGame() *PU_Game {
	return &PU_Game{tileImageMap : make(map[uint16]*PU_Image)}
}

func (g *PU_Game) LoadTileImages() {	
	files, err := ioutil.ReadDir("data/tiles/")
	if err != nil {
		fmt.Printf("Couldn't open data/files/ directory. Error: %v\n", err.String())
		return
	}
	
	for i := 0; i < len(files); i++ {
		g.LoadTileImage(files[i].Name)
	}
	
}

func (g *PU_Game) LoadTileImage(_file string) {
	name := strings.Replace(_file, ".png", "", -1)
	id, err := strconv.Atoi(name) 
	if err != nil {
		return
	}
	
	surface := sdl.LoadImage("data/tiles/"+_file)
	if surface == nil {
		return
	}
	
	image := NewImageFromSurface(surface)
	g_engine.AddResource(image)
	
	g.tileImageMap[uint16(id)] = image
}

func (g *PU_Game) GetTileImage(_id uint16) *PU_Image {
	if image, present := g.tileImageMap[_id]; present {
		return image
	}
	return nil
}
