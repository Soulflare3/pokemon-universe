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
	TILE_WIDTH  = 48
	TILE_HEIGHT = 48
)

const (
	TILE_BLOCKING       = 1
	TILE_WALK           = 2
	TILE_SURF           = 3
	TILE_BLOCKTOP       = 4
	TILE_BLOCKBOTTOM    = 5
	TILE_BLOCKRIGHT     = 6
	TILE_BLOCKLEFT      = 7
	TILE_BLOCKCORNER_TR = 8
	TILE_BLOCKCORNER_BR = 9
	TILE_BLOCKCORNER_BL = 10
	TILE_BLOCKCORNER_TL = 11
)

type PU_Tile struct {
	position pos.Position
	movement int

	layers [3]*PU_Layer
}

func NewTile(_x int, _y int) *PU_Tile {
	tile := &PU_Tile{position: pos.NewPositionFrom(_x, _y, 0)}
	return tile
}

func (t *PU_Tile) DrawLayer(_layer int, _x int, _y int) {
	if t.layers[_layer] == nil {
		return
	}

	drawX := (_x * TILE_WIDTH) - TILE_WIDTH - 22 + g_game.screenOffsetX
	drawY := (_y * TILE_HEIGHT) - TILE_HEIGHT + g_game.screenOffsetY

	t.layers[_layer].Draw(drawX, drawY)
}

func (t *PU_Tile) AddLayer(_layer int, _id int) {
	if t.layers[_layer] == nil {
		t.layers[_layer] = NewLayer(_id)
	} else {
		t.layers[_layer].SetID(_id)
	}
}

func (t *PU_Tile) RemoveLayer(_layer int) {
	if t.layers[_layer] != nil {
		t.layers[_layer] = nil
	}
}

func (t *PU_Tile) GetSignature() uint64 {
	signature := uint64(t.movement)
	shift := uint16(16)
	for i := 0; i < 3; i++ {
		if t.layers[i] != nil {
			signature |= (uint64(t.layers[i].id) << shift)
		}
		shift += 16
	}
	return signature
}

func (t *PU_Tile) GetHash() int64 {
	return t.position.Hash()
}
