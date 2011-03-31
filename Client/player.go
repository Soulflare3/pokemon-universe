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
	"math"
)

const (
	NUM_BODYPARTS = 6
	NUM_POKEMON = 6

	BODY_BASE = 0
	BODY_UPPER = 1
	BODY_NECK = 2
	BODY_HEAD = 3
	BODY_FEET = 4
	BODY_LOWER = 5
)

type PU_Player struct {
	PU_Creature
	
	name string
	
	walkConfirmed bool
	money uint32
	
	bodyParts [NUM_BODYPARTS]*PU_BodyPart
	pokemon [NUM_POKEMON]*PU_Pokemon
}

func NewPlayer(_id uint32) *PU_Player {
	player := &PU_Player{}
	player.SetDefault(_id)
	
	for i := 0; i < NUM_BODYPARTS; i++ {
		player.bodyParts[i] = NewBodyPart(1)
	}
	
	return player
}

func (p *PU_Player) Turn(_dir int, _send bool) {
	if _dir != p.direction {
		p.direction = _dir
		
		if _send {
			//g_conn.protocol.SendTurn(_dir)
		}
	}
}

func (p *PU_Player) Draw(_x int, _y int) {
	for part := BODY_BASE; part < BODY_LOWER; part++ {
		image := g_game.GetCreatureImage(part, p.bodyParts[part].id, p.direction, p.frame)
		if image != nil {
			if part != BODY_BASE {
				image.SetColorMod(p.bodyParts[part].red, p.bodyParts[part].green, p.bodyParts[part].blue)
			}
			
			image.Draw(_x, _y)
		}
	}
}

func (p *PU_Player) Walk(_direction int) {
	if !p.walking {
		if p.PreWalk(_direction) {
			g_conn.Game().SendMove(_direction, true)
		} else {
			p.CancelWalk()
		}
	}
}

func (p *PU_Player) PreWalk(_direction int) bool {
	var toTile *PU_Tile
	switch _direction {
		case DIR_NORTH:
			toTile = g_map.GetTile(int(p.x), int(p.y-1))
			
		case DIR_EAST:
			toTile = g_map.GetTile(int(p.x+1), int(p.y))
			
		case DIR_SOUTH:
			toTile = g_map.GetTile(int(p.x), int(p.y+1))
			
		case DIR_WEST:
			toTile = g_map.GetTile(int(p.x-1), int(p.y))
	}
	if p.CanWalkTo(_direction, toTile) {
		p.preWalkX = int16(toTile.position.X)
		p.preWalkY = int16(toTile.position.Y)
		
		p.Turn(_direction, false)
		
		if !p.animationRunning {
			p.StartAnimation()
		}
		
		p.walkProgress = 0.0
		p.offset = 0
		
		p.walkConfirmed = false
		p.walking = true
		
		return true
	}
	return false
}

func (p *PU_Player) CanWalkTo(_direction int, _tile *PU_Tile) bool {
	if _tile == nil {
		return false
	}
	
	tileMovement := _tile.movement
	if tileMovement != TILE_WALK {
		if (tileMovement == TILE_BLOCKING) ||
			(tileMovement == TILE_SURF) ||
			(tileMovement == TILE_BLOCKTOP && _direction == DIR_SOUTH) ||
			(tileMovement == TILE_BLOCKBOTTOM && _direction == DIR_NORTH) ||
			(tileMovement == TILE_BLOCKLEFT && _direction == DIR_EAST) ||
			(tileMovement == TILE_BLOCKRIGHT && _direction == DIR_WEST) ||
			(tileMovement == TILE_BLOCKCORNER_TL && (_direction == DIR_EAST || _direction == DIR_SOUTH)) ||
			(tileMovement == TILE_BLOCKCORNER_TR && (_direction == DIR_WEST || _direction == DIR_SOUTH)) ||
			(tileMovement == TILE_BLOCKCORNER_BL && (_direction == DIR_EAST || _direction == DIR_NORTH)) ||
			(tileMovement == TILE_BLOCKCORNER_BR && (_direction == DIR_WEST || _direction == DIR_NORTH)) {
			return false
		}
	}
	return true
}

func (p *PU_Player) CancelWalk() {
	p.walking = false
	p.walkProgress = 0.0
	p.offset = 0
	p.StopAnimation()
}

func (p *PU_Player) ReceiveWalk(_fromTile *PU_Tile, _toTile *PU_Tile) {
	if _toTile == nil || _fromTile == nil {
		p.CancelWalk()
		return
	}
	
	if p.x != int16(_fromTile.position.X) || p.y != int16(_fromTile.position.Y) {
		p.x = int16(_fromTile.position.X)
		p.x = int16(_fromTile.position.Y)
		
		p.preWalkX = int16(_toTile.position.X)
		p.preWalkY = int16(_toTile.position.Y)
	}
	
	if p != g_game.self {
		p.preWalkX = int16(_toTile.position.X)
		p.preWalkY = int16(_toTile.position.Y)
		
		if p.preWalkY > p.y {
			p.Turn(DIR_SOUTH, false)
		} else if (p.preWalkY < p.y) {
			p.Turn(DIR_NORTH, false)
		} else if (p.preWalkX > p.x) {
			p.Turn(DIR_EAST, false) 
		} else if(p.preWalkX < p.x) {
			p.Turn(DIR_WEST, false)
		}
		
		p.walkProgress = 0.0
		p.offset = 0
		p.StartAnimation()
		p.walking = true
		
	} else {
		p.x = p.preWalkX
		p.y = p.preWalkY
		
		p.CancelWalk()
	}
}

func (p *PU_Player) UpdateWalk() {
	if p.walking {
		p.walkProgress += (float32(1000.0)/float32(p.speed))*(float32(g_frameTime)/float32(1000.0))
		if p.walkProgress >= 1.0 {
			p.offset = TILE_WIDTH
			p.EndWalk()
		} else {
			p.offset = int(math.Ceil(float64(p.walkProgress*float32(TILE_WIDTH))))
		}
		p.UpdateAnimation()
	}
}

func (p *PU_Player) EndWalk() {
	if g_game.self != nil && g_game.self == p {
		p.walkEnded = true
		
		p.x = p.preWalkX
		p.y = p.preWalkY
		
		p.walking = false
		if !p.ContinueWalk() {
			p.StopAnimation()
		}
	} else {
		p.x = p.preWalkX
		p.y = p.preWalkY
		
		p.walking = false
		p.StopAnimation()
	}
}

func (p *PU_Player) ContinueWalk() bool {
	if g_game.state != GAMESTATE_WORLD {
		return false
	}
	
	if p == g_game.self {
		if g_game.lastDirKey != 0 {
			isDown := sdl.KeyDown(g_game.lastDirKey)
			if isDown {
				switch g_game.lastDirKey {
					case sdl.KEY_UP:
						p.Walk(DIR_NORTH)
						
					case sdl.KEY_DOWN:
						p.Walk(DIR_SOUTH)
						
					case sdl.KEY_LEFT:
						p.Walk(DIR_WEST)
						
					case sdl.KEY_RIGHT:
						p.Walk(DIR_EAST)
				}
				return true
			}
		}
	}
	return false
}

func (p *PU_Player) GetPokemonCount() int {
	pokecount := 0
	for i := 0; i < NUM_POKEMON; i++ {
		if p.pokemon[i] != nil {
			pokecount++
		} else {
			break
		}
	}
	return pokecount
}

type PU_BodyPart struct {
	id int
	
	red uint8
	green uint8
	blue uint8
}

func NewBodyPart(_id int) *PU_BodyPart {
	return &PU_BodyPart{id : _id}
}

func (b *PU_BodyPart) SetColor(_red int, _green int, _blue int) {
	b.red = uint8(_red)
	b.green = uint8(_green)
	b.blue = uint8(_blue)
}

//Used by Game's drawing procedure
type PU_PlayerName struct {
	name string
	x int 
	y int 
}

