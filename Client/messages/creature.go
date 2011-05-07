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
	"fmt"
)

func (p *PU_GameProtocol) Receive_AddCreature(_message *punet.Message) {
	data := _message.AddCreature
	player := NewPlayer(data.UID)
	player.name = data.Name
	player.x = data.X
	player.y = data.Y
	player.direction = data.Direction
	
	for part := BODY_UPPER; part <= BODY_LOWER; part++ {
		player.bodyParts[part].id = data.Outfit[part].ID
		color := data.Outfit[part].Color
		red := uint8(color >> 16)
		green := uint8(color >> 8)
		blue := uint8(color)
		player.bodyParts[part].SetColor(int(red), int(green), int(blue))
	}
	
	fmt.Printf("%v\n", player)
	
	g_map.AddCreature(player)
}

func (p *PU_GameProtocol) Receive_RemoveCreature(_message *punet.Message) {
	data := _message.RemoveCreature
	creature := g_map.GetCreatureByID(data.UID)
	if creature != nil {
		g_map.RemoveCreature(creature)
	}
}