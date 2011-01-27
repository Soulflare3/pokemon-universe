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
	"os"
	pos "position"
	"mysql"
)

type IOMapDB struct {

}

func (io *IOMapDB) LoadMap(_map *Map) (err os.Error) {
	var res *mysql.MySQLResult
	var row map[string]interface{}
	var	query string = "SELECT t.`x`, t.`y`, t.`z`, t.`idlocation`, t.`idmap`, t.`movement`, t.`idteleport`," +
					   	" tl.`sprite`, tl.`layer`, tp.`x` AS `tp_x`, tp.`y` AS `tp_y`, tp.`z` AS `tp_z`" +
					   	" FROM tile `t`" +
						" INNER JOIN tile_layer `tl` ON tl.`idtile` = t.`idtile`" +
						" LEFT JOIN teleport `tp` ON tp.`idteleport` = t.`idteleport`"
	
	if res, err = g_db.Query(query); err != nil {
		return
	}
	
	for {
		if row = res.FetchMap(); row == nil {
			break
		}
		
		// DB vars
		x			:=	row["x"].(int)
		y			:=	row["y"].(int)
		z			:=	row["z"].(int)
		position	:=	pos.NewPositionFrom(x, y, z)
		layer		:=	row["layer"].(int32)
		sprite		:=	row["sprite"].(int32)
		blocking	:=	row["movement"].(int32)
		//tp_id		:=	row["idteleport"].(int32)
		//location	:=	row["idlocation"].(int32)
		
		tile, found := _map.GetTile(position.Hash())
		if found == false {
			tile = NewTile(position)
			tile.Blocking = blocking
			
			// ToDo: Add teleport stuff
			// ---
			
			_map.addTile(tile)
		}
		
		tile.AddLayer(layer, sprite)
	}	
						
	return
}
