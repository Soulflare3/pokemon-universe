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


const (
	PACKET_MAXSIZE = 16384
)

type INetMessageWriter interface {
	WritePacket() (*Packet, os.Error)
}

type INetMessageReader interface {
	GetHeader() uint8
	ReadPacket(*Packet) os.Error
}

type Packet struct {
	readPos uint16
	MsgSize uint16	
	
	Buffer [PACKET_MAXSIZE]uint8
}

// NewPacket creates a new Packet with no header
func NewPacket() *Packet {
	packet := &Packet{}
	packet.Reset()
	
	return packet
}

// NewPacketExt creates a new Packet with message header
func NewPacketExt(_header uint8) *Packet {
	packet := NewPacket()
	packet.AddUint8(_header)
	
	return packet
}

func (p *Packet) Reset() {
	p.MsgSize = 0
	p.readPos = 2
}

func (p *Packet) CanAdd(_size uint16) bool {
	return (_size+p.readPos < PACKET_MAXSIZE - 16)
}

func (p *Packet) GetHeader() uint16 {
	p.MsgSize = uint16(uint16(p.Buffer[0]) | (uint16(p.Buffer[1]) << 8))
	return p.MsgSize
}

func (p *Packet) SetHeader() {
	p.Buffer[0] = uint8(p.MsgSize)
	p.Buffer[1] = uint8(p.MsgSize << 8)
	p.MsgSize += 2
}

func (p *Packet) ReadUint8() uint8 {
	v := p.Buffer[p.readPos]
	p.readPos += 1
	return v
}

func (p *Packet) ReadUint16() uint16 {
	v := uint16(uint16(p.Buffer[p.readPos]) | (uint16(p.Buffer[p.readPos+1]) << 8))
	p.readPos += 2
	return v
}

func (p *Packet) ReadUint32() uint32 {
	v := uint32((uint32(p.Buffer[p.readPos]) | (uint32(p.Buffer[p.readPos+1]) << 8) |
				 (uint32(p.Buffer[p.readPos+2]) << 16) | (uint32(p.Buffer[p.readPos+3]) << 24)))
	p.readPos += 4
	return v
}

func (p *Packet) ReadString() string {
	stringlen := p.ReadUint16()
	if uint16(stringlen) >= (PACKET_MAXSIZE+p.readPos) {
		return ""
	}
	
	v := string(p.Buffer[p.readPos:p.readPos+uint16(stringlen)])
	p.readPos += uint16(stringlen)
	return v
}

func (p *Packet) AddUint8(_value uint8) bool {
	if !p.CanAdd(1) {
		return false
	}
	
	p.Buffer[p.readPos] = _value
	p.readPos += 1
	p.MsgSize += 1
	
	return true
}

func (p *Packet) AddUint16(_value uint16) bool {
	if !p.CanAdd(2) {
		return false
	}
	
	p.Buffer[p.readPos] = uint8(_value)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 8)
	p.readPos += 1
	
	p.MsgSize += 2
	
	return true
}

func (p *Packet) AddUint32(_value uint32) bool {
	if !p.CanAdd(4) {
		return false
	}
	
	p.Buffer[p.readPos] = uint8(_value)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 8)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 16)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 24)
	p.readPos += 1
	
	p.MsgSize += 4
	
	return true
}

func (p *Packet) AddString(_value string) bool {
	stringlen := uint16(len(_value))
	if !p.CanAdd(stringlen) {
		return false
	}
	
	p.AddUint16(uint16(stringlen))
	for i, _ := range _value { 
		p.Buffer[p.readPos+uint16(i)] = _value[i]
	}
	
	p.readPos += stringlen
	p.MsgSize += uint16(stringlen)
	
	return true
}