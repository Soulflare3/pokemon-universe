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
)

const (
	DIR_SOUTH = 1
	DIR_WEST = 2
	DIR_NORTH = 3
	DIR_EAST = 4
)

type ICreature interface {
	Draw(_x int, _y int)
	IsWalking() bool
	UpdateWalk()
	GetOffset() int
	GetDirection() int
	GetX() int16
	GetY() int16
}

type PU_Creature struct {
	id uint32
	
	x int16
	y int16
	
	walking bool
	walkEnded bool
	preWalkX int16
	preWalkY int16
	offset int
	walkProgress float32
	speed int
	
	direction int
	
	frame int 
	frames int 
	
	animationRunning bool
	animationInterval int 
	animationLastTicks uint32
}

func (c *PU_Creature) IsWalking() bool {
	return c.walking
}

func (c *PU_Creature) GetOffset() int {
	return c.offset
}

func (c *PU_Creature) GetDirection() int {
	return c.direction
}

func (c *PU_Creature) GetX() int16 {
	if c.walking {
		return c.preWalkX
	}
	return c.x
}

func (c *PU_Creature) GetY() int16 {
	if c.walking {
		return c.preWalkY
	}
	return c.y
}

func (c *PU_Creature) SetDefault(_id uint32) {
	c.speed = 300
	c.direction = DIR_SOUTH
	c.frames = 3
	c.animationInterval = 150
	c.animationLastTicks = sdl.GetTicks()
}

func (c *PU_Creature) SetPosition(_x int, _y int) {
	c.x = int16(_x)
	c.y = int16(_y)
}

func (c *PU_Creature) StartAnimation() {
	c.animationRunning = true
}

func (c *PU_Creature) StopAnimation() {
	c.animationRunning = false
	c.frame = 0
}

func (c *PU_Creature) UpdateAnimation() {
	if c.animationRunning {
		passedTicks := sdl.GetTicks()-c.animationLastTicks
		if passedTicks >= uint32(c.animationInterval) {
			c.frame++
			if c.frame > c.frames {
				c.frame = 0
			}
			
			c.animationLastTicks = sdl.GetTicks()
		}
	}
}

