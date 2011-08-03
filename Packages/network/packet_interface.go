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
package network

import "os"

type INetMessageWriter interface {
	WritePacket() (*Packet, os.Error)
}

type INetMessageReader interface {
	GetHeader() uint8
	ReadPacket(*Packet) os.Error
}

type IPacket interface {
	Reset()
	CanAdd(_size uint16) bool

	GetHeader() uint16
	SetHeader()

	ReadUint8() uint8
	ReadUint16() uint16
	ReadUint32() uint32
	ReadUint64() uint64
	ReadString() string

	AddUint8(_value uint8) bool
	AddUint16(_value uint16) bool
	AddUint32(_value uint32) bool
	AddUint64(_value uint64) bool
	AddString(_value string) bool
}
