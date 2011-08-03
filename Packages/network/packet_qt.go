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

const (
	QTPACKET_MAXSIZE = 16384
)

type QTPacket struct {
	readPos uint16
	MsgSize uint16

	Buffer [QTPACKET_MAXSIZE]uint8
}

func NewQTPacket() *QTPacket {
	packet := &QTPacket{}
	packet.Reset()

	return packet
}

func NewQTPacketExt(_header uint8) *QTPacket {
	packet := NewQTPacket()
	packet.AddUint8(_header)

	return packet
}

func (p *QTPacket) Reset() {
	p.MsgSize = 0
	p.readPos = 2
}

func (p *QTPacket) CanAdd(_size uint16) bool {
	return (_size+p.readPos < QTPACKET_MAXSIZE-16)
}

func (p *QTPacket) GetHeader() uint16 {
	p.MsgSize = uint16(uint16(p.Buffer[1]) | (uint16(p.Buffer[0]) << 8))
	return p.MsgSize
}

func (p *QTPacket) SetHeader() {
	p.Buffer[0] = uint8(p.MsgSize >> 8)
	p.Buffer[1] = uint8(p.MsgSize)
	p.MsgSize += 2
}


func (p *QTPacket) ReadUint8() uint8 {
	v := p.Buffer[p.readPos]
	p.readPos += 1
	return v
}

func (p *QTPacket) ReadUint16() uint16 {
	v := uint16(uint16(p.Buffer[p.readPos+1]) | (uint16(p.Buffer[p.readPos]) << 8))
	p.readPos += 2
	return v
}

func (p *QTPacket) ReadUint32() uint32 {
	v := uint32((uint32(p.Buffer[p.readPos+3]) | (uint32(p.Buffer[p.readPos+2]) << 8) |
		(uint32(p.Buffer[p.readPos+1]) << 16) | (uint32(p.Buffer[p.readPos]) << 24)))
	p.readPos += 4
	return v
}

func (p *QTPacket) ReadUint64() uint64 {
	v := uint64((uint64(p.Buffer[p.readPos+7]) | (uint64(p.Buffer[p.readPos+6]) << 8) |
		(uint64(p.Buffer[p.readPos+5]) << 16) | (uint64(p.Buffer[p.readPos+4]) << 24) |
		(uint64(p.Buffer[p.readPos+3]) << 32) | (uint64(p.Buffer[p.readPos+2]) << 40) |
		(uint64(p.Buffer[p.readPos+1]) << 48) | (uint64(p.Buffer[p.readPos]) << 56)))
	p.readPos += 8
	return v
}

func (p *QTPacket) ReadString() string {
	stringlen := p.ReadUint32()
	if uint16(stringlen) >= (QTPACKET_MAXSIZE + p.readPos) {
		return ""
	}

	v := ""
	for i := 0; uint32(i) < stringlen/uint32(2); i++ {
		val := uint16(uint16(p.Buffer[p.readPos+1]) | (uint16(p.Buffer[p.readPos]) << 8))
		p.readPos += 2

		v += string(val)
	}
	return v
}

func (p *QTPacket) AddUint8(_value uint8) bool {
	if !p.CanAdd(1) {
		return false
	}

	p.Buffer[p.readPos] = _value
	p.readPos += 1
	p.MsgSize += 1

	return true
}

func (p *QTPacket) AddUint16(_value uint16) bool {
	if !p.CanAdd(2) {
		return false
	}

	p.Buffer[p.readPos] = uint8(_value >> 8)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value)
	p.readPos += 1

	p.MsgSize += 2

	return true
}

func (p *QTPacket) AddUint32(_value uint32) bool {
	if !p.CanAdd(4) {
		return false
	}

	p.Buffer[p.readPos] = uint8(_value >> 24)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 16)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 8)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value)
	p.readPos += 1

	p.MsgSize += 4

	return true
}

func (p *QTPacket) AddUint64(_value uint64) bool {
	if !p.CanAdd(8) {
		return false
	}

	p.Buffer[p.readPos] = uint8(_value >> 56)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 48)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 40)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 32)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 24)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 16)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value >> 8)
	p.readPos += 1
	p.Buffer[p.readPos] = uint8(_value)
	p.readPos += 1

	p.MsgSize += 8

	return true
}

func (p *QTPacket) AddString(_value string) bool {
	stringlen := uint16(len(_value) * 2)
	if !p.CanAdd(stringlen * uint16(2)) {
		return false
	}

	p.AddUint32(uint32(stringlen))
	for i, _ := range _value {
		p.AddUint16(uint16(_value[i]))
	}

	return true
}
