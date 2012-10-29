package main

import (
	 "fmt"
	puh "puhelper"
	pos "putools/pos"
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
		if tiles == nil {
			tiles = make(map[int64]*Tile)
		}
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
	if tiles == nil {
		tiles = make(map[int64]*Tile)
	}
	tile, found := tiles[_hash]
	return tile, found
}

func (m *Map) LoadMapList() (succeed bool, error string) {
	var query string = "SELECT idmap, name FROM map ORDER BY name"
		
	result, err := puh.DBQuerySelect(query)
	if err != nil {
		return false, err.Error()
	}
	
	defer puh.DBFree()
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		idmap := puh.DBGetInt(row[0])
		name := puh.DBGetString(row[1])
		
		m.AddMap(idmap, name)
	}
	
	return true, ""
}

func (m *Map) LoadTiles() (succeed bool) {

	result, err := puh.DBQuerySelect(QUERY_LOAD_TILES)
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err.Error())
		return false
	}
	
	count := 0
	
	defer puh.DBFree()
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		count++
		fmt.Printf("\rRetrieving tiles... %d", count)

		x := puh.DBGetInt(row[0])
		y := puh.DBGetInt(row[1])
		z := puh.DBGetInt(row[2])
		position := pos.NewPositionFrom(x, y, z)
		layer := puh.DBGetInt(row[7])
		sprite := puh.DBGetInt(row[6])
		blocking := puh.DBGetInt(row[4])
		// row `idtile_event` may be null sometimes.
		var te_id = 0
		if row[5] != nil {
			te_id = puh.DBGetInt(row[5])
		}
		// idlocation := DBGetInt(row[3])

		tile, found := m.GetTileFromPosition(position)
		
		if found == false {
			tile = NewTile(position)
			tile.IsNew = false
			tile.DbId = int64(puh.DBGetInt64(row[8]))
			tile.Blocking = blocking

			// Get location
			// location, found := g_game.Locations.GetLocation(idlocation)
			// if found {
			//	tile.Location = location
			// }

			//TODO tile events
			// Event
			if te_id > 0 {
				//Do some tile event stuff, like teleports
			}
			
			m.AddTile(tile)
		}
		tileLayer := tile.AddLayer(layer, sprite)
		tileLayer.DbId = puh.DBGetInt64(row[9])
		tileLayer.IsNew = false

	}
	
	fmt.Printf("\rRetrieving tiles... Done")
	return true
}

func (m *Map) ProcessMapChanges() {
	for {
	
	}
}