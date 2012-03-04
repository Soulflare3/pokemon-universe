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
	"errors"
	pos "putools/pos"
)

// Interface for map loading
type IMapLoader interface {
	LoadMap(_map *Map) error
}

type TilesMap map[int]map[int64]*Tile
type Map struct {
	tiles TilesMap
}

func NewMap() *Map {
	return &Map{tiles: make(TilesMap)}
}

func (m *Map) Load() error {
	maptype, _ := g_config.GetString("map", "type")
	var loader IMapLoader
	switch maptype {
	case "xml":
		loader = &IOMapXML{}
	case "db":
		loader = &IOMapDB{}
	case "db2":
		loader = &IOMapDB2{}
	default:
		return errors.New("Undefined map format!")
	}

	err := loader.LoadMap(m)
	return err
}

func (m *Map) AddMap(_id int, _name string) {
	m.tiles[_id] = make(map[int64]*Tile)
}

func (m *Map) AddTile(_tile *Tile) {
	if _, found := m.GetTileFromPosition(_tile.Position); !found {
		tiles := m.tiles[_tile.Position.Z]
		tiles[_tile.Position.Hash()] = _tile
	}
}

func (m *Map) GetTileFrom(_x, _y, _z int) (*Tile, bool) {
	position := pos.NewPositionFrom(_x, _y, _z)
	return m.GetTileFromPosition(position)
}

func (m *Map) GetTileFromPosition(_pos pos.Position) (*Tile, bool) {
	tiles := m.tiles[_pos.Z]
	tile, found := tiles[_pos.Hash()]
	return tile, found
}

func (m *Map) GetTile(_hash int64) (*Tile, bool) {
	mapId := pos.NewPositionFromHash(_hash).Z
	tiles := m.tiles[mapId]
	tile, found := tiles[_hash]
	return tile, found
}