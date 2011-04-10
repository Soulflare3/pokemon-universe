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
	"math"
)

func (p *PU_GameProtocol) Send_Walk(_direction int, _requestTiles bool) {
	message := punet.NewData_Walk()
	message.Walk.Direction = _direction
	message.Walk.RequestTiles = _requestTiles
	g_conn.SendMessage(message)
}

func (p *PU_GameProtocol) Receive_CancelWalk() {
	if g_game.self != nil {
		g_game.self.CancelWalk()
	}
}

func (p *PU_GameProtocol) Receive_CreatureWalk(_message *punet.Message) {
	data := _message.CreatureWalk
	creature := g_map.GetCreatureByID(data.UID)
	fromTile := g_map.GetTile(data.FromX, data.FromY)
	toTile := g_map.GetTile(data.ToX, data.ToY)
	if creature != nil {
		if int(math.Fabs(float64(data.FromX)-float64(data.ToX))) > 1 || int(math.Fabs(float64(data.FromY)-float64(data.ToY))) > 1 {
			creature.SetPosition(data.ToX, data.ToY)
		} else {
			creature.ReceiveWalk(fromTile, toTile)
		}
	}
}
