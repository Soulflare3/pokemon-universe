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
	list "container/list"
)

type PU_Map struct {
	tileMap map[int64]*PU_Tile
	creatureList *list.List
}

func NewMap() *PU_Map {
	return &PU_Map{tileMap : make(map[int64]*PU_Tile),
				   creatureList : list.New()}
}

func (m *PU_Map) GetNumTiles() int {
	return len(m.tileMap)
}

func (m *PU_Map) AddTile(_x int, _y int) *PU_Tile {
	var index int64 = pos.Hash(_x, _y, 0)
	tile, present :=  m.tileMap[index]
	if !present {
		tile = NewTile(_x, _y)
		m.tileMap[index] = tile
	}
	return tile
}

func (m *PU_Map) RemoveTile(_tile *PU_Tile) {
	m.tileMap[_tile.GetHash()] = _tile, false
}

func (m *PU_Map) RemoveTileFromPos(_pos *pos.Position) {
	tile, present := m.tileMap[_pos.Hash()]
	if present {
		m.tileMap[_pos.Hash()] = tile, false
	}
}

func (m *PU_Map) GetTile(_x int, _y int) *PU_Tile {
	var index int64 = pos.Hash(_x, _y, 0)
	tile := m.tileMap[index]
	return tile;
}

func (m* PU_Map) AddCreature(_creature ICreature) {
	m.creatureList.PushBack(_creature)
}

func (m *PU_Map) RemoveCreature(_creature ICreature) {
	for e := m.creatureList.Front(); e != nil;  e = e.Next() {
		if e.Value == _creature {
			m.creatureList.Remove(e)
			break
		}
	}
}

func (m *PU_Map) GetCreatureByID(_id uint32) ICreature {
	for e := m.creatureList.Front(); e != nil;  e = e.Next() {
		if e.Value.(ICreature).GetID() == _id {
			return e.Value.(ICreature)
		}
	}
	return nil
}

func (m *PU_Map) GetPlayerByName(_name string) *PU_Player {
	for e := m.creatureList.Front(); e != nil;  e = e.Next() {
		if player, is_player := e.Value.(*PU_Player); is_player {
			if player.name == _name {
				return player
			}
		}
	}
	return nil
}

