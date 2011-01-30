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

func (g *PU_Game) LoadFonts () {
	g_engine.LoadFont(FONT_PURITANBOLD_14, GetPath()+"data/font/Puritan2Bold.otf", 14)
	g_engine.LoadFont(FONT_PURITANBOLD_12, GetPath()+"data/font/Puritan2Bold.otf", 12)
	g_engine.LoadFont(FONT_ARIALBLACK_10, GetPath()+"data/font/ariblk.ttf", 10)
}

func (g* PU_Game) LoadTileImages() {
	g.LoadGameImages(GetPath()+"data/tiles/", g.tileImageMap)
}

func (g* PU_Game) LoadGuiImages() {
	g.LoadGameImages(GetPath()+"data/gui/", g.guiImageMap)
}

func (g *PU_Game) GetTileImage(_id uint16) *PU_Image {
	if image, present := g.tileImageMap[_id]; present {
		return image
	}
	return nil
}

func (g *PU_Game) GetGuiImage(_id uint16) *PU_Image {
	if image, present := g.guiImageMap[_id]; present {
		return image
	}
	return nil
}

func (g *PU_Game) GetCreatureImage(_bodypart int, _id int, _dir int, _frame int) *PU_Image {
	key := (uint32(_bodypart) | (uint32(_id) << 8) | (uint32(_dir) << 16) | (uint32(_frame) << 24))
	if image, present := g.creatureImageMap[key]; present {
		return image
	}
	return nil
}

func (g *PU_Game) LoadGameImages(_dir string, _map map[uint16]*PU_Image) {	
	files, err := ioutil.ReadDir(_dir)
	if err != nil {
		fmt.Printf("Couldn't open directory: %v. Error: %v\n", _dir, err.String())
		return
	}
	
	for i := 0; i < len(files); i++ {
		img, id := g.LoadGameImage(files[i].Name, _dir)
		if img != nil {
			_map[uint16(id)] = img
		}
	}
}

func (g *PU_Game) LoadGameImage(_file string, _dir string) (*PU_Image, int) {
	name := strings.Replace(_file, ".png", "", -1)
	id, err := strconv.Atoi(name) 
	if err != nil {
		return nil, 0
	}
	
	surface := sdl.LoadImage(_dir+_file)
	if surface == nil {
		return nil, 0
	}
	
	image := NewImageFromSurface(surface)
	g_engine.AddResource(image)
	
	return image, id	
}

func (g *PU_Game) LoadCreatureImages() {
	dir := GetPath()+"data/creatures/"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("Couldn't open directory: %v. Error: %v\n", dir, err.String())
		return
	}
	
	for i := 0; i < len(files); i++ {
		surface := sdl.LoadImage(dir+files[i].Name)
		if surface == nil {
			return
		}
	
		image := NewImageFromSurface(surface)
		g_engine.AddResource(image)
		if image != nil {
			file := strings.Replace(files[i].Name, ".png", "", -1)
			parts := strings.Split(file, "_", -1)
		
			bodypart, err := strconv.Atoi(parts[0])
			if err != nil {
				continue
			}
			
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				continue
			}
			
			dir, err := strconv.Atoi(parts[2])
			if err != nil {
				continue
			}
			
			frame, err := strconv.Atoi(parts[3])
			if err != nil {
				continue
			}
		
			key := (uint32(bodypart) | (uint32(id) << 8) | (uint32(dir) << 16) | (uint32(frame) << 24))
			g.creatureImageMap[key] = image
		}
	}
}
