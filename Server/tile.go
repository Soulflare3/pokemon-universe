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
	TILEBLOCK_BLOCK = 1
	TILEBLOCK_WALK = 2
	TILEBLOCK_SURF = 3
	TILEBLOCK_TOP = 4
	TILEBLOCK_BOTTOM = 5
	TILEBLOCK_RIGHT = 6
	TILEBLOCK_LEFT = 7
	TILEBLOCK_TOPRIGHT = 8
	TILEBLOCK_BOTTOMRIGHT = 9
	TILEBLOCK_BOTTOMLEFT = 10
	TILEBLOCK_TOPLEFT = 11
)

type TileLayer struct {
	Layer		int32
	SpriteID	int32
}

type LayerMap map[int32]*TileLayer
type Tile struct {
	Position	pos.Position
	Blocking	int32
	
	Layers		LayerMap
	Creatures	CreatureList // List of creatures who are active on this tile
}

// NewTile creates a Tile object with Position as parameter
func NewTile(_pos pos.Position) *Tile {
	t := &Tile { Position: _pos }
	t.Blocking = TILEBLOCK_WALK
	t.Layers = make(LayerMap)
	t.Creatures = make(CreatureList)
	
	return t
}

// NewTileExt creates a Position from _x, _y, _z and then calls NewTile to create a new Tile object
func NewTileExt(_x int, _y int, _z int) *Tile {
	return NewTile(pos.NewPositionFrom(_x, _y, _z))
}

// AddLayer adds a new TileLayer to the tile. 
// If the layer already exists it will return that one otherwise it'll make a new one
func (t *Tile) AddLayer(_layer int32, _sprite int32) (layer *TileLayer) {
	layer = t.GetLayer(_layer)
	if layer == nil {
		layer = &TileLayer{Layer: _layer, SpriteID: _sprite}
		t.Layers[_layer] = layer
	}
	
	return
}

// GetLayer returns a TileLayer object if the layer exists, otherwise nil
func (t *Tile) GetLayer(_layer int32) *TileLayer {
	if layer, ok := t.Layers[_layer]; !ok {
		return layer
	}
	
	return nil
}

// CheckMovement checks if a creature can move to this tile
func (t *Tile) CheckMovement(_creature ICreature, _dir int) ReturnValue {
	movement := _creature.GetMovement()
	blocking := t.Blocking
	
	if blocking != TILEBLOCK_WALK {
		if blocking == TILEBLOCK_BLOCK ||
			(blocking == TILEBLOCK_SURF		&& movement != MOVEMENT_SURF) ||
			(blocking == TILEBLOCK_TOP		&& _dir == DIR_SOUTH) ||
			(blocking == TILEBLOCK_BOTTOM	&& _dir == DIR_NORTH) ||
			(blocking == TILEBLOCK_LEFT		&& _dir == DIR_EAST) ||
			(blocking == TILEBLOCK_RIGHT	&& _dir == DIR_WEST) ||
			(blocking == TILEBLOCK_TOPLEFT	&& (_dir == DIR_EAST || _dir == DIR_SOUTH)) ||
			(blocking == TILEBLOCK_TOPRIGHT && (_dir == DIR_WEST || _dir == DIR_SOUTH)) ||
			(blocking == TILEBLOCK_BOTTOMLEFT  && (_dir == DIR_EAST || _dir == DIR_NORTH)) ||
			(blocking == TILEBLOCK_BOTTOMRIGHT && (_dir == DIR_WEST || _dir == DIR_NORTH)) {
			return RET_NOTPOSSIBLE
		}
	}
	
	return RET_NOERROR
}

// AddCreature adds a new active creature to this tile
func (t *Tile) AddCreature(_creature ICreature) (ret ReturnValue) {
	ret = RET_NOERROR
	
	_, found := t.Creatures[_creature.GetUID()]
	if !found {
		t.Creatures[_creature.GetUID()] = _creature
	}
	
	return
}

// RemoveCreature removes an active creature from this tile
func (t *Tile) RemoveCreature(_creature ICreature) (ret ReturnValue) {
	ret = RET_NOERROR
	
	_, found := t.Creatures[_creature.GetUID()]
	if found {
		t.Creatures[_creature.GetUID()] = nil, false
	}
	
	return
}
