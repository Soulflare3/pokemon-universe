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

type PU_Creature struct {
	id uint32
	
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

func (c *PU_Creature) SetDefault(_id uint32) {
	c.speed = 300
	c.direction = DIR_SOUTH
	c.frames = 3
	c.animationInterval = 150
	c.animationLastTicks = sdl.GetTicks()
}
