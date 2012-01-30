package main

import (
	list "container/list"
	pos "position"
)

const (
	TILEBLOCK_BLOCK       int = 1
	TILEBLOCK_WALK            = 2
	TILEBLOCK_SURF            = 3
	TILEBLOCK_TOP             = 4
	TILEBLOCK_BOTTOM          = 5
	TILEBLOCK_RIGHT           = 6
	TILEBLOCK_LEFT            = 7
	TILEBLOCK_TOPRIGHT        = 8
	TILEBLOCK_BOTTOMRIGHT     = 9
	TILEBLOCK_BOTTOMLEFT      = 10
	TILEBLOCK_TOPLEFT         = 11
)

type TileLayer struct {
	DbId	 int64
	Layer    int
	SpriteID int
}

type LayerMap map[int]*TileLayer
type Tile struct {
	DbId		int64
	Position 	pos.Position
	Blocking 	int
	// Location 	*Location

	Layers    	LayerMap
	Events    	*list.List
	
	IsRemoved	bool
}

// NewTile creates a Tile object with Position as parameter
func NewTile(_pos pos.Position) *Tile {
	t := &Tile{Position: _pos}
	t.Blocking = TILEBLOCK_WALK
	t.Layers = make(LayerMap)
	// t.Location = nil
	t.Events = list.New()

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
		layer = &TileLayer{Layer: _layer, SpriteID: _sprite}
		t.Layers[_layer] = layer
	}

	return
}

// func (t *Tile) AddEvent(_event ITileEvent) {
//	t.Events.PushBack(_event)
// }

// GetLayer returns a TileLayer object if the layer exists, otherwise nil
func (t *Tile) GetLayer(_layer int) *TileLayer {
	if layer, ok := t.Layers[_layer]; !ok {
		return layer
	}

	return nil
}

func (t *Tile) RemoveLayer(_layer int) {
	delete(t.Layers, _layer)
}

type TileCollection struct {
  Tiles []*Tile
  x, y int
  width, height int
  username, description string
}

func (c *TileCollection) AddTile(_tile *Tile) {
        c.Tiles = append(c.Tiles, _tile)
}

