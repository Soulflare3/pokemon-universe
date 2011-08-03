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
	punet "network"
)

func (p *PU_GameProtocol) Receive_Tiles(_message *punet.Message) {
	data := _message.Tiles
	for _, t := range data.Tiles {
		x := t.X
		y := t.Y

		tile := g_map.GetTile(x, y)
		if tile == nil {
			tile = g_map.AddTile(x, y)
			tile.movement = t.Blocking

			for _, l := range t.Layers {
				tile.AddLayer(l.Index, l.Sprite)
			}

			continue
		}
		//else

		layers := [3]int{-1, -1, -1}
		for _, l := range t.Layers {
			layers[l.Index] = l.Sprite
		}
		signature := uint64(t.Blocking)
		shift := uint16(16)
		for i := 0; i < 3; i++ {
			if layers[i] != -1 {
				signature |= (uint64(layers[i]) << shift)
			}
			shift += 16
		}

		//we don't want to remove/add all layers of a tile each time we receive it
		//only when something is different
		if tile.GetSignature() != signature {
			tile.movement = t.Blocking
			for i := 0; i < 3; i++ {
				tile.RemoveLayer(i)
				if layers[i] != -1 {
					tile.AddLayer(i, layers[i])
				}
			}
		}
	}
}
