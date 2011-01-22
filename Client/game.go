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

const (
	NUMTILES_X = 23
	NUMTILES_Y = 17
	
	MID_X = 11
	MID_Y = 8
)

const (
	GAMESTATE_UNLOADING = iota
	GAMESTATE_LOADING
	GAMESTATE_LOGIN
	GAMESTATE_WORLD
	GAMESTATE_BATTLE_INIT
	GAMESTATE_BATTLE
)

type PU_Game struct {
	state int
	tileImageMap map[uint16]*PU_Image
	guiImageMap map[uint16]*PU_Image
	
	screenOffsetX int
	screenOffsetY int
	
	self *PU_Player
}

func NewGame() *PU_Game {
	return &PU_Game{state : GAMESTATE_LOADING,
					tileImageMap : make(map[uint16]*PU_Image),
					guiImageMap : make(map[uint16]*PU_Image)}
}

func (g *PU_Game) SetState(_state int) {
	g.state = _state
}

func (g *PU_Game) Draw() {
	switch g.state {
		case GAMESTATE_LOADING:
			g_engine.GetFont(FONT_PURITANBOLD_14).DrawText("Loading, please wait...", 10, 10)
			
		case GAMESTATE_LOGIN:
			g.GetGuiImage(IMG_GUI_INTROBG).Draw(0, 0)
			g_gui.Draw()
			
		case GAMESTATE_WORLD:
			g.DrawWorld()
	}
}

func (g *PU_Game) DrawWorld() {
	g.screenOffsetX, g.screenOffsetY = g.GetScreenOffset()
			
	layer2tiles := make([]*PU_Tile, (NUMTILES_X*NUMTILES_Y))
	layer2tilesCount := 0
	
	//draw tile layer 0 and 1 
	for x := 0; x < NUMTILES_X; x++ {
		for y := 0; y < NUMTILES_Y; y++ {
			mapx := int(g.self.x-MID_X)+x
			mapy := int(g.self.y-MID_Y)+y
			
			tile := g_map.GetTile(mapx, mapy)
			if tile != nil {
				for i := 0; i < 2; i++ {
					tile.DrawLayer(i, x, y)
				}
				if tile.layers[2] != nil {
					layer2tiles[layer2tilesCount] = tile
					layer2tilesCount++
				}
			}
		}
	}
	
	//draw creatures
	
	//draw tile layer 2
	for i := 0; i < layer2tilesCount; i++ {
		tile := layer2tiles[i]
		
		screenx := MID_X-(int(g.self.x)-tile.position.X)
		screeny := MID_Y-(int(g.self.y)-tile.position.Y)
		
		tile.DrawLayer(2, screenx, screeny)
	}
	
	//draw player names
	
	//draw extra info
	info := fmt.Sprintf("x: %v y: %v", g.self.x, g.self.y)
	if font := g_engine.GetFont(FONT_PURITANBOLD_14); font != nil {
		font.SetColor(255, 242, 0)
		font.DrawText(info, 5, 5)
	}
	
	fmt.Printf("Numtiles: %v\n", len(g_map.tileMap))
}

func (g *PU_Game) GetScreenOffset() (x int, y int) {
	x = 0
	y = 0
	if g.self != nil && g.self.walking {
		switch g.self.direction {
			case DIR_NORTH:
				y = 0-(TILE_HEIGHT-g.self.offset)

			case DIR_EAST:
				x = (TILE_WIDTH-g.self.offset)

			case DIR_SOUTH:
				y = (TILE_HEIGHT-g.self.offset)

			case DIR_WEST:
				x = 0-(TILE_WIDTH-g.self.offset)
		}
	}
	return
}

func (g *PU_Game) LoadFonts () {
	g_engine.LoadFont(FONT_PURITANBOLD_14, GetPath()+"data/font/Puritan2Bold.otf", 14)
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
