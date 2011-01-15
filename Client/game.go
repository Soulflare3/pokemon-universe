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
	zip "archive/zip"	
	"sdl"
	"io"
	"strings"
	"strconv"
	"time"
)

type PU_Game struct {
	tileImageMap map[uint16]*PU_Image
}

func NewGame() *PU_Game {
	return &PU_Game{tileImageMap : make(map[uint16]*PU_Image)}
}

func (g *PU_Game) LoadTileImages() {	
	zip, err := zip.OpenReader("data/tiles.pu_data")
	if zip == nil {
		fmt.Printf("LoadTileImages zip error: %v\n", err.String())
		return
	}
	
	for i := 0; i < len(zip.File); i++ {
		g.LoadTileImage(zip.File[i])
	}
	
}

func (g *PU_Game) LoadTileImage(_file *zip.File) {
	reader, err := _file.Open()
	if reader == nil {
		fmt.Printf("LoadTileImages file error: %v\n", err.String())
	}
	name := _file.FileHeader.Name
	name = strings.Replace(name, ".png", "", -1)
	id, err := strconv.Atoi(name) 
	if err != nil {
		return
	}
		
	size := _file.FileHeader.UncompressedSize
	data := make([]byte, size)
	numbytes, err := io.ReadFull(reader, data)
	if uint32(numbytes) != size {
		fmt.Printf("LoadTileImages read error: %v\n", err.String())
	} else {
		surface := sdl.LoadImageRW(&data, numbytes)
		image := NewImageFromSurface(surface)
		g_engine.AddResource(image)
		
		g.tileImageMap[uint16(id)] = image
		time.Sleep(1) //Somehow it won't work without this. TODO: Find out why 
	}
}

func (g *PU_Game) GetTileImage(_id uint16) *PU_Image {
	if image, present := g.tileImageMap[_id]; present {
		return image
	}
	return nil
}
