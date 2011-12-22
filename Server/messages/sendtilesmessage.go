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
)

type SendTilesMessage struct {
	tiles TilesMap
}

func NewSendTilesMessage() *SendTilesMessage {
	return &SendTilesMessage{tiles: make(TilesMap)}
}

// GetHeader returns the header value of this message
func (m *SendTilesMessage) GetHeader() uint8 {
	return pnet.HEADER_TILES
}

func (m *SendTilesMessage) AddTile(_index int64) {
	tile, ok := g_map.GetTile(_index)
	if ok {
		m.tiles[_index] = tile
	}
}

// WritePacket write the needed object data to a Packet and returns it
func (m *SendTilesMessage) WritePacket() pnet.IPacket {
	totalTiles := uint16(len(m.tiles))
	packet := pnet.NewPacketExt(m.GetHeader())
	packet.AddUint16(totalTiles)
	
	if totalTiles == 0 {
		return packet
	}
	
	for _, tile := range m.tiles {
		if tile == nil {
			continue
		}
		
		packet.AddUint16(uint16(tile.Position.X))
		packet.AddUint16(uint16(tile.Position.Y))
		packet.AddUint16(uint16(tile.Blocking))
		
		for id, layer := range tile.Layers {
			if layer == nil {
				continue
			}
			
			packet.AddUint16(uint16(id))
			packet.AddUint32(uint32(layer.SpriteID))
		}
		
		// Tile Location info
		if location := tile.Location; location != nil {
			packet.AddUint16(uint16(tile.Location.ID))
			packet.AddString(tile.Location.Name)
		} else {
			packet.AddUint16(0)
			packet.AddString("Unknown")
		}
	}
	
	m.tiles = make(TilesMap) 
	
	return packet	
}
