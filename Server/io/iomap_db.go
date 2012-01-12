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
	"fmt"
	"mysql"
	pos "position"
)

type IOMapDB struct{}

var rowChan = make(chan mysql.Row)

func processRows(_map *Map) {
	for {
		row := <-rowChan
		if row == nil {
			fmt.Println("Map IO Channel closed!");
			break;
		}
		x := DBGetInt(row[0])
		y := DBGetInt(row[1])
		z := DBGetInt(row[2])
		position := pos.NewPositionFrom(x, y, z)
		layer := DBGetInt(row[8])
		sprite := DBGetInt(row[7])
		blocking := DBGetInt(row[5])
		// row `idteleport` may be null sometimes.
		var tp_id = 0
		if row[6] != nil {
			tp_id = DBGetInt(row[6])
		}
		idlocation := DBGetInt(row[3])

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
				tp_x := DBGetInt(row[9])
				tp_y := DBGetInt(row[10])
				tp_z := DBGetInt(row[11])
				tp_pos := pos.NewPositionFrom(tp_x, tp_y, tp_z)
				tile.AddEvent(NewWarp(tp_pos))
			}

			_map.addTile(tile)
		}

		tile.AddLayer(layer, sprite)
	}
}

func (io *IOMapDB) LoadMap(_map *Map) (err error) {
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
	rowChan <- nil
	
	return
}
