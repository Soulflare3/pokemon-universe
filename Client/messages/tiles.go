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
	"os"
)

type PU_Message_Tiles struct {
}

func NewTilesMessage(_packet *punet.Packet) *PU_Message_Tiles {
	msg := &PU_Message_Tiles{}
	msg.ReadPacket(_packet)
	return msg
}

func (m *PU_Message_Tiles) ReadPacket(_packet *punet.Packet) os.Error {
	tileCount := int(_packet.ReadUint16())
	if tileCount > 0 {
		for i := 0; i < tileCount; i++ {
			m.ProcessTile(_packet)
		}
	}
	return nil
}

func (m *PU_Message_Tiles) ProcessTile(_packet *punet.Packet) {
	tileExists := true
	tileMovement := TILE_WALK
	layers := [3]int{-1,-1,-1}
	
	x := int16(_packet.ReadUint16())
	y := int16(_packet.ReadUint16())
	
	tile := g_map.GetTile(int(x), int(y))
	if tile == nil {
		tileExists = false
	}	
	
	numLayers := int(_packet.ReadUint16())
	for i := 0; i < numLayers; i++ {
		layer := _packet.ReadUint16()
		id := _packet.ReadUint32()
		movement := _packet.ReadUint16()
		
		layers[layer] = int(id)
		
		if layer == 1 {
			tileMovement = int(movement)
		}
	}
	
	if !tileExists {
		tile = g_map.AddTile(int(x), int(y))
		tile.movement = tileMovement
		
		for i := 0; i < 3; i++ {
			if layers[i] != -1 {
				tile.AddLayer(i, layers[i])
			}
		}
	} else {
		signature := uint64(tileMovement)
		shift := uint16(16)
		for i := 0; i < 3; i++ {
			if layers[i] != -1 {
				signature |= uint64((uint16(layers[i]) & 0xFFFF) << shift);
			}
			shift += 16
		}
		
		//we don't want to remove/add all layers of a tile each time we receive it
		//only when something is different
		if tile.GetSignature() != signature {
			tile.movement = tileMovement
			for i := 0; i < 3; i ++ {
				tile.RemoveLayer(i)
				if layers[i] != -1 {
					tile.AddLayer(i, layers[i])
				}
			}
		}
	}
	
	_packet.ReadUint16() //town id 
	_packet.ReadString() //town name
}
