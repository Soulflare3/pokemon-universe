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

var rowChan = make(chan mysql.Row)

func processRows(_map *Map) {
	for {
		row := <-rowChan
		x := row[0].(int)
		y := row[1].(int)
		z := row[2].(int)
		position := pos.NewPositionFrom(x, y, z)
		layer := row[8].(int)
		sprite := row[7].(int)
		blocking := row[5].(int)
		// row `idteleport` may be null sometimes.
		var tp_id = 0
		if row[6] != nil {
			tp_id = row[6].(int)
		}
		idlocation := row[3].(int)

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
				tp_x := row[9].(int)
				tp_y := row[10].(int)
				tp_z := row[11].(int)
				tp_pos := pos.NewPositionFrom(tp_x, tp_y, tp_z)
				tile.AddEvent(NewWarp(tp_pos))
			}

			_map.addTile(tile)
		}

		tile.AddLayer(layer, sprite)
	}
}

func (io *IOMapDB) LoadMap(_map *Map) (err os.Error) {
	// Spawn a row processor, in a different goroutine.
	go processRows(_map)

	// Fetch the rows:
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
		if count%1000 == 0 {
			fmt.Printf("Row %v\r", count)
		}
		row := result.FetchRow()
		if row == nil {
			break
		}
		// Send the row to the processor!
		rowChan <- row
	}
	return
}
