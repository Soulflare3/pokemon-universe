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
	"gomysql"
	puh "puhelper"
	pos "putools/pos"
	"putools/log"
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
		x := puh.DBGetInt(row[0])
		y := puh.DBGetInt(row[1])
		z := puh.DBGetInt(row[2])
		position := pos.NewPositionFrom(x, y, z)
		layer := puh.DBGetInt(row[7])
		sprite := puh.DBGetInt(row[6])
		blocking := puh.DBGetInt(row[4])
		// row `idteleport` may be null sometimes.
		var tp_id = 0
		if row[5] != nil {
			tp_id = puh.DBGetInt(row[5])
		}
		idlocation := puh.DBGetInt(row[3])

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
				tp_x := puh.DBGetInt(row[8])
				tp_y := puh.DBGetInt(row[9])
				tp_z := puh.DBGetInt(row[10])
				tp_pos := pos.NewPositionFrom(tp_x, tp_y, tp_z)
				tile.AddEvent(NewWarp(tp_pos))
			}

			_map.AddTile(tile)
		}

		tile.AddLayer(layer, sprite)
	}
}

func (io *IOMapDB) LoadMap(_map *Map) error {
	// Spawn a row processor, in a different goroutine.
	go processRows(_map)

	// Fetch the rows:
	var query string = "SELECT t.`x`, t.`y`, t.`z`, t.`idlocation`, t.`movement`, t.`idteleport`," +
		" tl.`sprite`, tl.`layer`, tp.`x` AS `tp_x`, tp.`y` AS `tp_y`, tp.`z` AS `tp_z`" +
		" FROM tile `t`" +
		" INNER JOIN tile_layer `tl` ON tl.`idtile` = t.`idtile`" +
		" LEFT JOIN teleport `tp` ON tp.`idteleport` = t.`idteleport`"

	result, err := puh.DBQuerySelect(query)
	if err != nil {
		return err
	}
	defer result.Free()
	
	logger.Printf(" - Processing worldmap data from database")
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
	
	return nil
}
