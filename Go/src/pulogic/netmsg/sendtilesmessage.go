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
package netmsg

import (
	pnet "network"
	pul "pulogic"
)

type SendTilesMessage struct {
	Tiles pul.TilesMap
}

func NewSendTilesMessage() *SendTilesMessage {
	return &SendTilesMessage{Tiles: make(pul.TilesMap)}
}

// GetHeader returns the header value of this message
func (m *SendTilesMessage) GetHeader() uint8 {
	return pnet.HEADER_TILES
}

func (m *SendTilesMessage) AddTile(_tile pul.ITile) {
	m.Tiles[_tile.GetPosition().Hash()] = _tile
}

// WritePacket write the needed object data to a Packet and returns it
func (m *SendTilesMessage) WritePacket() pnet.IPacket {
	totalTiles := uint16(len(m.Tiles))
	packet := pnet.NewPacketExt(m.GetHeader())
	packet.AddUint16(totalTiles)
	
	if totalTiles == 0 {
		return packet
	}
	
	for _, tile := range m.Tiles {
		if tile == nil {
			continue
		}
		
		packet.AddUint16(uint16(tile.GetPosition().X))
		packet.AddUint16(uint16(tile.GetPosition().Y))
		packet.AddUint16(uint16(tile.GetBlocking()))
		
		packet.AddUint16(uint16(len(tile.GetLayers())))
		for id, layer := range tile.GetLayers() {
			if layer == nil {
				continue
			}
			
			packet.AddUint16(uint16(id))
			packet.AddUint32(uint32(layer.SpriteID))
		}
		
		// Tile Location info
		if location := tile.GetLocation(); location != nil {
			packet.AddUint16(uint16(location.GetId()))
			packet.AddString(location.GetName())
		} else {
			packet.AddUint16(0)
			packet.AddString("Unknown")
		}
	}
	
	m.Tiles = make(pul.TilesMap) 
	
	return packet	
}
