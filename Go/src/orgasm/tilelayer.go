package main

import (
	"fmt"
	puh "puhelper"
)

type TileLayer struct {
	DbId	 int64
	Layer    int
	SpriteId int
	
	IsNew 		bool
	IsModified 	bool
	IsRemoved	bool
}

func NewTileLayer(_layer, _spriteId int) *TileLayer {
	tl := &TileLayer{Layer: _layer, SpriteId: _spriteId}
	tl.IsNew = true
	
	return tl
}

func (tl *TileLayer) Save(_tileId int64) bool {
	var query string
	if tl.IsNew {
		query = fmt.Sprintf(QUERY_INSERT_TILELAYER, _tileId, tl.Layer, tl.SpriteId)
	} else if tl.IsModified {
		query = fmt.Sprintf(QUERY_UPDATE_TILELAYER, tl.SpriteId, tl.DbId)
	}
	
	if query != "" {		
		if err := puh.DBQuery(query); err != nil {
			return false
		}
		
		if tl.IsNew {
			tl.DbId = int64(puh.DBGetLastInsertId())
		}
	}

	tl.IsNew = false
	tl.IsModified = false
	
	return true
}

func (tl *TileLayer) Delete() bool {
	query := fmt.Sprintf(QUERY_DELETE_TILELAYER, tl.DbId)
	if err := puh.DBQuery(query); err != nil {
		return false
	}
	
	return true
}

func (tl *TileLayer) SetSpriteId(_id int) {
	tl.SpriteId = _id
	tl.IsModified = true
} 