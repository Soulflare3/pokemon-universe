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
	pnet "network"
	pos "position"
)

func (c *Connection) Send_Tiles(_direction int, _centerPosition pos.Position) {
	xMin := 1
	xMax := CLIENT_VIEWPORT.X
	yMin := 1
	yMax := CLIENT_VIEWPORT.Y

	if _direction != DIR_NULL {
		switch _direction {
		case DIR_NORTH:
			yMax = 1
		case DIR_EAST:
			xMin = CLIENT_VIEWPORT.X
		case DIR_SOUTH:
			yMin = CLIENT_VIEWPORT.Y
		case DIR_WEST:
			xMax = 1
		}
	}

	// Top-left coordinates
	positionX := (_centerPosition.X - CLIENT_VIEWPORT_CENTER.X)
	positionY := (_centerPosition.Y - CLIENT_VIEWPORT_CENTER.Y)
	z := _centerPosition.Z

	tilesMessage := pnet.NewData_Tiles()
	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			index := pos.Hash(positionX+x, positionY+y, z)
			tile, ok := g_map.GetTile(index)
			if ok {
				newTile := pnet.NewTile()
				newTile.X = positionX + x
				newTile.Y = positionY + y
				newTile.Blocking = tile.Blocking
				for _, layer := range tile.Layers {
					newTile.AddLayer(layer.Layer, layer.SpriteID)
				}
				tilesMessage.Tiles.Tiles = append(tilesMessage.Tiles.Tiles, newTile)
			}
		}

		if _direction == DIR_NULL {
			c.SendMessage(tilesMessage)
		}
	}

	if _direction != DIR_NULL {
		c.SendMessage(tilesMessage)
	}
}
