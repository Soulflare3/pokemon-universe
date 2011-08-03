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
	pos "position"
)

const (
	MOVEMENT_WALK = TILEBLOCK_WALK
	MOVEMENT_SURF = TILEBLOCK_SURF
)

const (
	DIR_NULL  = 0
	DIR_SOUTH = 1
	DIR_WEST  = 2
	DIR_NORTH = 3
	DIR_EAST  = 4
)

const (
	CTYPE_CREATURE = 0
	CTYPE_NPC      = 1
	CTYPE_PLAYER   = 2
)

// CreatureList is map which holds a list of ICreature interfaces
type CreatureList map[uint64]ICreature

// Main interface for all creature objects in the game
type ICreature interface {
	GetUID() uint64
	GetName() string
	GetType() int

	GetPosition() pos.Position
	GetMovement() int

	SetDirection(_dir int)
	GetDirection() int

	SetTile(_tile *Tile)
	GetTile() *Tile

	GetOutfit() Outfit

	GetMovementSpeed() int
	GetTimeSinceLastMove() int

	// Methods for all moving creatures
	OnCreatureMove(_creature ICreature, _from *Tile, _to *Tile, _teleport bool)
	OnCreatureTurn(_creature ICreature)
	OnCreatureAppear(_creature ICreature, _isLogin bool)
	OnCreatureDisappear(_creature ICreature, _isLogout bool)

	// Methods for all creatures who need to see other creatures	
	AddVisibleCreature(_creature ICreature)
	RemoveVisibleCreature(_creature ICreature)
	KnowsVisibleCreature(_creature ICreature) bool
	GetVisibleCreatures() CreatureList
}

// CanSeeCreature checks if 2 creatures are near each others viewport
func CanSeeCreature(_self ICreature, _other ICreature) bool {
	return CanSeePosition(_self.GetPosition(), _other.GetPosition())
}

// CanSeePosition checks if 2 positions are near each others viewport
func CanSeePosition(_p1 pos.Position, _p2 pos.Position) bool {
	if _p1.Z != _p2.Z {
		return false
	}

	return _p1.IsInRange2p(_p2, CLIENT_VIEWPORT_CENTER)
}

// Returns true if the passed creature can move
func CreatureCanMove(_creature ICreature) bool {
	canMove := (_creature.GetTimeSinceLastMove() >= _creature.GetMovementSpeed())
	return canMove
}

// Creature struct with generic variables for all creatures
type Creature struct {
	uid  uint64 // Unique ID
	name string
	Id   int // Database ID			

	Position  *Tile
	Direction int

	Movement  int
	lastStep  int64
	moveSpeed int

	Outfit

	VisibleCreatures CreatureList
}

func (c *Creature) GetUID() uint64 {
	return c.uid
}

func (c *Creature) GetName() string {
	return c.name
}

func (c *Creature) GetTile() *Tile {
	return c.Position
}

func (c *Creature) SetTile(_tile *Tile) {
	c.Position = _tile
}

func (c *Creature) GetPosition() pos.Position {
	return c.Position.Position
}

func (c *Creature) GetMovement() int {
	return c.Movement
}

func (c *Creature) GetDirection() int {
	return c.Direction
}

func (c *Creature) SetDirection(_dir int) {
	c.Direction = _dir
}

func (c *Creature) GetOutfit() Outfit {
	return c.Outfit
}

func (c *Creature) GetMovementSpeed() int {
	return c.moveSpeed
}

func (c *Creature) GetTimeSinceLastMove() int {
	return int(PUSYS_TIME() - c.lastStep)
}

func (c *Creature) KnowsVisibleCreature(_creature ICreature) (found bool) {
	_, found = c.VisibleCreatures[_creature.GetUID()]
	return
}

func (c *Creature) GetVisibleCreatures() CreatureList {
	return c.VisibleCreatures
}
