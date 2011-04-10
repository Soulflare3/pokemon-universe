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
	"fmt"
)

type IOMapDB struct{}

func (io *IOMapDB) LoadMap(_map *Map) (err os.Error) {
	var query string = "SELECT t.`x`, t.`y`, t.`z`, t.`idlocation`, t.`idmap`, t.`movement`, t.`idteleport`," +
		" tl.`sprite`, tl.`layer`, tp.`x` AS `tp_x`, tp.`y` AS `tp_y`, tp.`z` AS `tp_z`" +
		" FROM tile `t`" +
		" INNER JOIN tile_layer `tl` ON tl.`idtile` = t.`idtile`" +
		" LEFT JOIN teleport `tp` ON tp.`idteleport` = t.`idteleport`"

	if err = g_db.Query(query); err != nil {
		return
	}
	
	var result *mysql.Result
	result, err = g_db.UseResult()
	if err != nil {
		return
	}
	
	defer result.Free()
	g_logger.Printf(" - Processing worldmap data from database")
	count := 0
	for {
		count++
		fmt.Printf("Row %v\r", count)
		row := result.FetchMap()
		if row == nil {
			break
		}
		
		x 			:= row["x"].(int)
		y 			:= row["y"].(int)
		z 			:= row["z"].(int)
		position 	:= pos.NewPositionFrom(x, y, z)
		layer		:= row["layer"].(int)
		sprite		:= row["sprite"].(int)
		blocking	:= row["movement"].(int)
		tp_id 		:= row["idteleport"].(int)
		idlocation	:= row["idlocation"].(int)

		tile, found := _map.GetTile(position.Hash())
		if found == false {
			tile = NewTile(position)
			tile.Blocking = blocking

			// Get location
			location, found := g_game.Locations.GetLocation(idlocation)
			if found {
				tile.Location = location
			}

			// Teleport event
			if tp_id > 0 {
				tp_x := row["tp_x"].(int)
				tp_y := row["tp_y"].(int)
				tp_z := row["tp_z"].(int)
				tp_pos := pos.NewPositionFrom(tp_x, tp_y, tp_z)
				tile.AddEvent(NewWarp(tp_pos))
			}

			_map.addTile(tile)
		}

		tile.AddLayer(layer, sprite)
	}

	return
}

