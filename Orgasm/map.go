package main

import (
	"mysql"
	pos "position"
)

type Map struct {
	tileMap map[int]map[int64]*Tile
	mapNames map[int]string
	
	updateChannel chan *Packet
}

func NewMap() *Map {
	return &Map{ tileMap: make(map[int]map[int64]*Tile), 
				 mapNames: make(map[int]string),
				 updateChannel: make(chan *Packet) }
}

func (m *Map) GetNumTiles() int {
		var tiles int = 0
	for _, value := range(m.tileMap) {
		tiles += len(value)
	}
	return tiles
}

func (m *Map) GetNumMaps() int {
	return len(m.mapNames)
}

func (m *Map) AddMap(_id int, _name string) {
	m.mapNames[_id] = _name
	m.tileMap[_id] = make(map[int64]*Tile)
}

func (m *Map) GetMap(_id int) (string, bool) {
	name, ok := m.mapNames[_id]
	return name, ok
}

func (m *Map) DeleteMap(_id int) {
	// Remove all tiles
	delete(m.tileMap, _id)
	
	// Remove map name
	delete(m.mapNames, _id)
}

func (m *Map) AddTile(_tile *Tile) {
	if _, found := m.GetTileFromPosition(_tile.Position); !found {
		tiles := m.tileMap[_tile.Position.Z]
		tiles[_tile.Position.Hash()] = _tile
	}
}

func (m *Map) RemoveTile(_tile *Tile) {
	var index int64 = _tile.Position.Hash()
	tiles := m.tileMap[_tile.Position.Z]
	delete(tiles, index)
}

func (m *Map) GetTileFromCoordinates(_x, _y, _z int) (*Tile, bool) {
	position := pos.NewPositionFrom(_x, _y, _z)
	return m.GetTileFromPosition(position)
}

func (m *Map) GetTileFromPosition(_pos pos.Position) (*Tile, bool) {
	tiles := m.tileMap[_pos.Z]
	tile, found := tiles[_pos.Hash()]
	return tile, found
}

func (m *Map) GetTile(_hash int64) (*Tile, bool) {
	mapId := pos.NewPositionFromHash(_hash).Z
	tiles := m.tileMap[mapId]
	tile, found := tiles[_hash]
	return tile, found
}

func (m *Map) LoadMapList() (succeed bool, error string) {
	var query string = "SELECT idmap, name FROM map ORDER BY name"
	
	if err := g_db.Query(query); err != nil {
		return false, err.Error()
	}
	
	result, err := g_db.UseResult()
	if err != nil {
		return false, err.Error()
	}
	
	defer result.Free()
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		idmap := DBGetInt(row[0])
		name := DBGetString(row[1])
		
		m.AddMap(idmap, name)
	}
	
	return true, ""
}

func (m *Map) LoadTiles() (succeed bool, msg string) {
	var query string = "SELECT t.`x`, t.`y`, t.`z`, t.`idlocation`, t.`movement`, t.`idteleport`," +
		" tl.`sprite`, tl.`layer`, tp.`x` AS `tp_x`, tp.`y` AS `tp_y`, tp.`z` AS `tp_z`," +
		" t.`idtile`, tl.`idtile_layer`" +
		" FROM tile `t`" +
		" INNER JOIN tile_layer `tl` ON tl.`idtile` = t.`idtile`" +
		" LEFT JOIN teleport `tp` ON tp.`idteleport` = t.`idteleport`"

	var err error
	var result *mysql.Result
	if result, err = DBQuerySelect(query); err != nil {
		return false, err.Error()
	}

	defer result.Free()
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}

		x := DBGetInt(row[0])
		y := DBGetInt(row[1])
		z := DBGetInt(row[2])
		position := pos.NewPositionFrom(x, y, z)
		layer := DBGetInt(row[7])
		sprite := DBGetInt(row[6])
		blocking := DBGetInt(row[4])
		// row `idteleport` may be null sometimes.
		var tp_id = 0
		if row[5] != nil {
			tp_id = DBGetInt(row[5])
		}
		// idlocation := DBGetInt(row[3])

		tile, found := m.GetTileFromPosition(position)
		if found == false {
			tile = NewTile(position)
			tile.DbId = row[11].(int64)
			tile.Blocking = blocking

			// Get location
			// location, found := g_game.Locations.GetLocation(idlocation)
			// if found {
			//	tile.Location = location
			// }

			// Teleport event
			if tp_id > 0 {
				tp_x := DBGetInt(row[8])
				tp_y := DBGetInt(row[9])
				tp_z := DBGetInt(row[10])
				tp_pos := pos.NewPositionFrom(tp_x, tp_y, tp_z)
				
				warp := NewWarp(tp_pos)
				warp.dbid = int64(tp_id)
				tile.AddEvent(warp)
			}

			m.AddTile(tile)
		}

		tileLayer := tile.AddLayer(layer, sprite)
		tileLayer.DbId = row[12].(int64)
	}
	return true, ""
}

func (m *Map) ProcessMapChanges() {
	for {
	
	}
}