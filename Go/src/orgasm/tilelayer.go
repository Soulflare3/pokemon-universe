package main

import (
	"fmt"
	puh "puhelper"
)

type TileLayer struct {
	DbId	 int64
	Layer    int
	SpriteId int
	TileId int64
	
	IsNew 		bool
	IsModified 	bool
	IsRemoved	bool
}

func NewTileLayer(_layer, _spriteId int, _tileId int64) *TileLayer {
	tl := &TileLayer{Layer: _layer, SpriteId: _spriteId, TileId: _tileId}
	tl.IsNew = true
	
	return tl
}

func (tl *TileLayer) Save() bool {
	var query string
	if tl.IsNew {
		query = fmt.Sprintf(QUERY_INSERT_TILELAYER, tl.TileId, tl.Layer, tl.SpriteId)
	} else if tl.IsModified {
		query = fmt.Sprintf(QUERY_UPDATE_TILELAYER, tl.SpriteId, tl.DbId)
	}
	
	if query != "" {		
		if err := puh.DBQuery(query); err != nil {
			return false
		}
		
		if tl.IsNew {
			tl.DbId = int64(puh.DBGetLastInsertId())
			if IS_DEBUG {
				fmt.Printf("Added New tilelayer to DB - DbId: %d\n", tl.DbId) 
			}
		} else if tl.IsModified {
			if IS_DEBUG {
				fmt.Printf("Updated tilelayer - DbId: %d\n", tl.DbId) 
			}
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