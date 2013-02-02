package main

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

func (tl *TileLayer) Delete() bool {
	tl.IsRemoved = true
//	query := fmt.Sprintf(QUERY_DELETE_TILELAYER, tl.DbId)
//	if err := puh.DBQuery(query); err != nil {
//		return false
//	}
	
	return true
}

func (tl *TileLayer) SetSpriteId(_id int) {
	tl.SpriteId = _id
	tl.IsModified = true
} 