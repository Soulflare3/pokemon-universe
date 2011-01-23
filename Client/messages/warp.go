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
	"sdl"
)

type PU_Message_Warp struct {
	x int16
	y int16	
}

func NewWarpMessage(_packet *punet.Packet) *PU_Message_Warp {
	msg := &PU_Message_Warp{}
	msg.ReadPacket(_packet)
	return msg
}

func (m *PU_Message_Warp) ReadPacket(_packet *punet.Packet) os.Error {
	m.x = int16(_packet.ReadUint16())
	m.y = int16(_packet.ReadUint16())	
	
	g_game.state = GAMESTATE_LOADING
	sdl.Delay(10)
	
	if g_game.self != nil {
		g_game.self.CancelWalk()
		g_game.self.SetPosition(m.x, m.y)
	}
	
	//request the tiles of the new area we just arrived at 
	message := NewTileRefreshMessage()
	g_conn.SendMessage(message)
	
	return nil
}

//request the tiles of the new area we warped to
type PU_Message_TileRefresh struct {
}

func NewTileRefreshMessage() *PU_Message_TileRefresh {
	return &PU_Message_TileRefresh{}
}

func (m *PU_Message_TileRefresh) WritePacket() (*punet.Packet, os.Error) {
	packet := punet.NewPacket()
	packet.AddUint8(0xC4)
	return packet, nil
}

