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
	pos "putools/pos"
)

type Warp struct {
	dbid int64
	destination pos.Position
}

func NewWarp(_destination pos.Position) *Warp {
	return &Warp{destination: _destination}
}

func (e *Warp) GetEventType() int {
	return EVENTTYPES_TELEPORT
}

func (e *Warp) ToPacket(_packet *Packet) {
	_packet.AddUint8(uint8(e.GetEventType()))
	_packet.AddUint16(uint16(e.destination.X))
	_packet.AddUint16(uint16(e.destination.Y))
	_packet.AddUint16(uint16(e.destination.Z))
}
