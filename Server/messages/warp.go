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

func (c *Connection) Send_PlayerWarp(_position pos.Position) {
	msg := pnet.NewData_Warp()
	msg.Warp.X = _position.X
	msg.Warp.Y = _position.Y
	c.SendMessage(msg)
}

func (c *Connection) Send_RefreshComplete() {
	msg := pnet.NewMessage(pnet.HEADER_REFRESHCOMPLETE)
	c.SendMessage(msg)
}

func (c *Connection) Receive_RefreshWorld() {
	// Send whole screen
	c.Send_Tiles(DIR_NULL, c.Owner.GetPosition())

	c.Send_RefreshComplete()
}
