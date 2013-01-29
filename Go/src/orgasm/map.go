package main

import (
	"sync" 
	"runtime" 
	"time" 
		
	pos "nonamelib/pos"
	"nonamelib/log"
	
	"pulogic/models" 
)

type Map struct {
	tileMap map[int]map[int64]*Tile
	mapNames map[int]string
	
	updateChannel chan *Packet
	
	numOfProcessRoutines int
	processChan chan TileRow
	processExitChan chan bool
	tileMutex	sync.Mutex
}

type TileRow struct {
	Idtile		int64
	IdtileEvent int
	X           int
	Y           int
	Z           int
	Layer       int
	Movement    int
	Sprite      int
	Idlocation  int
	Param1      string
	Param2      string
	Param3      string
	Param4      string
	Param5      string
	Eventtype   int
}

func NewMap() *Map {
	return &Map{ tileMap: make(map[int]map[int64]*Tile), 
				 mapNames: make(map[int]string),
				 updateChannel: make(chan *Packet),
				 numOfProcessRoutines: runtime.NumCPU(),
				 processChan: make(chan TileRow),
				 processExitChan: make(chan bool) }
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

// Gets a tile from the list. If the tile doesnt exists it will be created
// Returns the tile pointer and a boolean, true if the tile is new
// NOTE: Should only be used when loading the map, because of locking
func (m *Map) getOrAddTile(_x, _y, _z int) (*Tile) {
	m.tileMutex.Lock()
	defer m.tileMutex.Unlock()
	
	position := pos.NewPositionFrom(_x, _y, _z)
	tile, ok := m.GetTileFromPosition(position)
	
	if !ok {
		tile = NewTile(position)
		m.AddTile(tile)
	} 
	
	return tile
}

// Waits for all spawned process routines to finish
// This can take a while if there are alot of tiles
func (m *Map) waitForLoadComplete() {
	count := 0
	for {
		select {
			case <-m.processExitChan:
				count++
				if count == m.numOfProcessRoutines {
					return
				}
		}
	}
}

func (m *Map) LoadMapList() (bool, string) {
	var maps []models.Map
	if err := g_orm.FindAll(&maps); err != nil {
		return false, err.Error()
	}

	for _, mapEntity := range maps {
		m.AddMap(mapEntity.Idmap, mapEntity.Name)
	}

	return true, ""
}

func (m *Map) LoadTiles() bool {
	start := time.Now().UnixNano()
	var allTiles []TileRow
	
	err := g_orm.SetTable("tile").Join("INNER", "tile_layer", "tile_layer.tileid = tile.idtile").Join(" LEFT", "tile_events", "tile_events.idtile_events = tile.idtile_event").OrderBy("tile.idtile DESC").FindAll(&allTiles)
	if err != nil {
		log.Error("map", "loadTiles", "Error while loading tiles: %v", err.Error())
		return false
	}
	
	log.Info("Map", "loadTiles", "%d tiles fetched in %dms", len(allTiles), (time.Now().UnixNano()-start)/1e6)
	
	if len(allTiles) > 0 {
		log.Verbose("Map", "loadTiles", "Processing tiles with %d goroutines", m.numOfProcessRoutines)
		
		// Start process goroutine
		for i := 1; i <= m.numOfProcessRoutines; i++ {
			go m.processTiles()
		}

		// Send rows to channel
		for key, row := range(allTiles) {
			//First tile has highest ID
			if (key == 0) {
				g_newTileId = (row.Idtile + 1)	
				log.Verbose("Map", "loadTiles", "Determined next tile ID: %d", g_newTileId)	
			}
			
			m.processChan <- row
		}

		// Close channel so the process goroutine(s) will shutdown
		close(m.processChan)
	}
	return true
}

func (m *Map) processTiles() {
	for {
		row, ok := <-m.processChan
		if !ok {
			break
		}
		
		// Get or create tile
		tile := m.getOrAddTile(row.X, row.Y, row.Z)
		
		// If the tile is new set extra data
		if tile.IsNew {
			//Set tile db id
			tile.DbId = row.Idtile
		
			// Set blocking
			tile.Blocking = row.Movement
			
			// Link location to tile
//			if location, found := g_locations.Get(row.Idlocation); found {
//				tile.Location = location
//			}
	
			// Check if we have a tile event id. If so, do something with it
//			if row.IdtileEvent > 0 {
//				if row.Eventtype == pulogic.EVENTTYPE_TELEPORT {
//					destination_x, _ := strconv.Atoi(row.Param1)
//					destination_y, _ := strconv.Atoi(row.Param2)
//					destination_z, _ := strconv.Atoi(row.Param3)
//					
//					destination := pos.NewPositionFrom(destination_x, destination_y, destination_z)
//					teleport := NewTeleport(destination)
//	
//					tile.AddEvent(teleport)
//				}
//			}
			tile.IsNew = false
		}
	
		// Add layer to tile
		tile.AddLayer(row.Layer, row.Sprite)
	}
	
	m.processExitChan <- true
}