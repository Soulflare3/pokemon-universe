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

func (c *Connection) Send_CreatureWalk(_creature ICreature, _from *Tile, _to *Tile) {
	msg := pnet.NewData_CreatureWalk()
	msg.CreatureWalk.UID = _creature.GetUID()
	msg.CreatureWalk.FromX = _from.Position.X
	msg.CreatureWalk.FromY = _from.Position.Y
	msg.CreatureWalk.ToX = _to.Position.X
	msg.CreatureWalk.ToY = _to.Position.Y
	c.SendMessage(msg)
}

func (c *Connection) Send_CancelWalk() {
	msg := pnet.NewMessage(pnet.HEADER_CANCELWALK)
	c.SendMessage(msg)
}

func (c *Connection) Receive_Walk(_message *pnet.Message) {
	data := _message.Walk
	g_game.OnPlayerMove(c.Owner, data.Direction, data.RequestTiles)
}
