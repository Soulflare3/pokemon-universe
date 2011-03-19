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
	"os"
	pos "position"
)

// Interface for map loading
type IMapLoader interface {
	LoadMap(_map *Map) os.Error
}

type TilesMap map[int64]*Tile
type Map struct {
	tiles    TilesMap
}

func NewMap() *Map {
	return &Map{ tiles: make(TilesMap) }
}

func (m *Map) Load() os.Error {
	maptype, _ := g_config.GetString("map", "type")
	var loader IMapLoader
	switch maptype {
	case "xml":
		loader = &IOMapXML{}
	case "db":
		loader = &IOMapDB{}
	default:
		return os.NewError("Undefined map format!")
	}

	err := loader.LoadMap(m)
	return err
}

func (m *Map) addTile(_tile *Tile) {
	index := _tile.Position.Hash()
	if _, found := m.GetTile(index); !found {
		m.tiles[index] = _tile
	}
}

func (m *Map) GetTile(_index int64) (tile *Tile, ok bool) {
	tile, ok = m.tiles[_index]
	return
}

func (m *Map) GetTileFrom(_x, _y, _z int) (tile *Tile, ok bool) {
	index := pos.Hash(_x, _y, _z)
	tile, ok = m.GetTile(index)
	return
}

func (m *Map) GetTileFromPosition(_position pos.Position) (tile *Tile, ok bool) {
	tile, ok = m.GetTile(_position.Hash())
	return
}

