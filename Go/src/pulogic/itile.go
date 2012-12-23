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
package pulogic

import (
	list "container/list"
	pos "nonamelib/pos"
)

type TileBlocking	int
const (
	TILEBLOCK_BLOCK       TileBlocking = 1
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

const (
	EVENTTYPE_TELEPORT	int = 1
)

type TilesMap map[int64]ITile
type LayerMap map[int]*TileLayer

type ITile interface {
	GetPosition()	pos.Position
	GetBlocking() 	TileBlocking
	GetCreatures() 	CreatureMap
	GetLayers()		LayerMap
	GetEvents() 	*list.List
	GetLocation()	ILocation
	AddCreature(_creature ICreature, _checkEvents bool) (ret ReturnValue)
	RemoveCreature(_creature ICreature, _checkEvents bool) (ret ReturnValue)
}

type TileLayer struct {
	Layer    int
	SpriteID int
}

type ITileEvent interface {
	OnCreatureEnter(_creature ICreature, _prevRet ReturnValue) ReturnValue
	OnCreatureLeave(_creature ICreature, _prevRet ReturnValue) ReturnValue
}

func GetTileBlockingFromInt(_value int) (ret TileBlocking) {
	switch _value {
		case 1: ret = TILEBLOCK_BLOCK
		case 2:	ret = TILEBLOCK_WALK
		case 3:	ret = TILEBLOCK_SURF
		case 4:	ret = TILEBLOCK_TOP
		case 5:	ret = TILEBLOCK_BOTTOM
		case 6:	ret = TILEBLOCK_RIGHT
		case 7:	ret = TILEBLOCK_LEFT
		case 8:	ret = TILEBLOCK_TOPRIGHT
		case 9:	ret = TILEBLOCK_BOTTOMRIGHT
		case 10: ret = TILEBLOCK_BOTTOMLEFT
		case 11: ret = TILEBLOCK_TOPLEFT
	}
	return
}