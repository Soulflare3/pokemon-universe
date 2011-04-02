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
	"db"
)

type IOMapDB struct{}

func (io *IOMapDB) LoadMap(_map *Map) (err os.Error) {
	var query string = "SELECT t.`x`, t.`y`, t.`z`, t.`idlocation`, t.`idmap`, t.`movement`, t.`idteleport`," +
		" tl.`sprite`, tl.`layer`, tp.`x` AS `tp_x`, tp.`y` AS `tp_y`, tp.`z` AS `tp_z`" +
		" FROM tile `t`" +
		" INNER JOIN tile_layer `tl` ON tl.`idtile` = t.`idtile`" +
		" LEFT JOIN teleport `tp` ON tp.`idteleport` = t.`idteleport`"

	var result db.ResultSet
	if result, err = g_db.StoreQuery(query); err != nil {
		return
	}

	g_logger.Printf(" - Processing worldmap data from database")
	
	for ; result.Next();  {
		
		x 			:= result.GetDataInt("x")
		y 			:= result.GetDataInt("y")
		z 			:= result.GetDataInt("z")
		position 	:= pos.NewPositionFrom(int(x), int(y), int(z))
		layer		:= result.GetDataInt("layer")
		sprite		:= result.GetDataInt("sprite")
		blocking	:= result.GetDataInt("movement")
		tp_id 		:= result.GetDataInt("idteleport")
		idlocation	:= result.GetDataInt("idlocation")

		tile, found := _map.GetTile(position.Hash())
		if found == false {
			tile = NewTile(position)
			tile.Blocking = uint16(blocking)

			// Get location
			location, found := g_game.Locations.GetLocation(idlocation)
			if found {
				tile.Location = location
			}

			// Teleport event
			if tp_id > 0 {
				tp_x := result.GetDataInt("tp_x")
				tp_y := result.GetDataInt("tp_y")
				tp_z := result.GetDataInt("tp_z")
				tp_pos := pos.NewPositionFrom(int(tp_x), int(tp_y), int(tp_z))
				tile.AddEvent(NewWarp(tp_pos))
			}

			_map.addTile(tile)
		}

		tile.AddLayer(int16(layer), int32(sprite))

	}
	result.Free()

	return
}

