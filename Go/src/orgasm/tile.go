package main

import (
	"fmt"
	pos "putools/pos"
	puh "puhelper"
)

type LayerMap map[int]*TileLayer
type Tile struct {
	DbId		int64
	Position 	pos.Position
	Blocking 	int
	// Location 	*Location

	Layers    	LayerMap
	Event    	ITileEvent
	
	IsNew		bool
	IsModified	bool
	IsRemoved	bool
}

// NewTile creates a Tile object with Position as parameter
func NewTile(_pos pos.Position) *Tile {
	t := &Tile{Position: _pos}
	t.Blocking = TILEBLOCK_WALK
	t.Layers = make(LayerMap)
	// t.Location = nil
	
	t.IsNew = true;

	return t
}

// NewTileExt creates a Position from _x, _y, _z and then calls NewTile to create a new Tile object
func NewTileExt(_x int, _y int, _z int) *Tile {
	return NewTile(pos.NewPositionFrom(_x, _y, _z))
}

// AddLayer adds a new TileLayer to the tile. 
// If the layer already exists it will return that one otherwise it'll make a new one
func (t *Tile) AddLayer(_layer int, _sprite int) (layer *TileLayer) {
	layer = t.GetLayer(_layer)
	if layer == nil {
		layer = NewTileLayer(_layer, _sprite)
		t.Layers[_layer] = layer
	} else {
		t.Layers[_layer].SetSpriteId(_sprite)
	}

	return
}

func (t *Tile) AddEvent(_event ITileEvent) {
	t.Event = _event
	t.IsModified = true
	
	t.Event.Save()
}

func (t *Tile) RemoveEvent() {
	if t.Event != nil {
		if t.Event.Delete() {
			t.Event = nil
			t.IsModified = true
		}
	}
}

func (t *Tile) SetBlocking(_blocking int) {
	t.Blocking = _blocking
	t.IsModified = true
}

// GetLayer returns a TileLayer object if the layer exists, otherwise nil
func (t *Tile) GetLayer(_layer int) *TileLayer {
	if layer, ok := t.Layers[_layer]; ok {
		return layer
	}

	return nil
}

func (t *Tile) RemoveLayer(_layer *TileLayer) {
	if _layer != nil {
		if _layer.Delete() {
			delete(t.Layers, _layer.Layer)
		}
	}
}

// Save tile (including children) to database
func (t *Tile) Save() bool {
	var eventDbId int64 = 0

	// Check if tile has an event 
	if t.Event != nil {
		t.Event.Save()
		
		eventDbId = t.Event.GetDbId()
	}
	
	var query string
	if t.IsNew {
		query = fmt.Sprintf(QUERY_INSERT_TILE, t.Position.X, t.Position.Y, t.Position.Z, t.Blocking, eventDbId)
	} else if t.IsModified { // Tile is probably changed, update it in the database
		query = fmt.Sprintf(QUERY_UPDATE_TILE, t.Blocking, eventDbId, t.DbId)
	}
	
	if len(query) > 0 {
		if err := puh.DBQuery(query); err != nil {
			return false
		}
		
		if t.IsNew {
			t.DbId = int64(puh.DBGetLastInsertId())
		}
	}
	
	t.IsNew = false
	t.IsModified = false
	
	// Add tile to map
	g_map.AddTile(t)
	
	// Save all layers (if needed)
	for _, tl := range t.Layers {
		if !tl.Save() {
			return false
		}
	}
	
	return true
}

// Remove tile (including children) from database
func (t *Tile) Delete() bool {
	// Delete all layers
	for _, tl := range t.Layers {
		tl.Delete()
	}

	// Check if tile has an event 
	if t.Event != nil {
		t.Event.Delete()
	}
	
	query := fmt.Sprintf(QUERY_DELETE_TILE, t.DbId)
	if err := puh.DBQuery(query); err != nil {
		return false
	}
				
	t.IsRemoved = true
	
	return true
}