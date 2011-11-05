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

type QColor struct {
	Spec uint8
	Alpha uint16
	Red uint16
	Green uint16
	Blue uint16
	Pad uint16
	Html string
}

func NewQColor() *QColor {
	return &QColor { Spec: 0, Alpha: uint16(0xFFFF), 
					 Red: 0, Green: 0, Blue: 0,
					 Pad: 0, Html: ">" }
}

func NewQColorFromPacket(_packet *pnet.QTPacket) *QColor {
	color := QColor {}
	color.Spec = _packet.ReadUint8()
	color.Alpha = _packet.ReadUint16()
	color.Red = _packet.ReadUint16()
	color.Green = _packet.ReadUint16()
	color.Blue = _packet.ReadUint16()
	color.Pad = _packet.ReadUint16()
	color.Html = ">"
	
	return &color
}

func (c *QColor) WritePacket() pnet.IPacket {
	packet := pnet.NewQTPacket()
	packet.AddUint8(c.Spec)
	packet.AddUint16(c.Alpha)
	packet.AddUint16(c.Red)
	packet.AddUint16(c.Green)
	packet.AddUint16(c.Blue)
	packet.AddUint16(c.Pad)
	return packet
}