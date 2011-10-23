package main

import (
	"os"
	pnet "network"
)

type UniqueId struct {
	PokeNum int
	SubNum int
}

func NewUniqueId() *UniqueId {
	return &UniqueId { PokeNum: 173, SubNum: 0 }
}

func NewUniqueId(_pokeNum, _subNum int) *UniqueId {
	uniqueId := UniqueId { PokeNum: _pokeNum,
							SubNum: _subNum }
	return &uniqueId
}

func NewUniqueIdFromPacket(_packet *pnet.QTPacket) *UniqueId {
	uniqueId := UniqueId { PokeNum: (int)_packet.ReadUint16(),
							SubNum: (int)_packet.ReadUint8() }
	return &uniqueId
}

func (u *UniqueId) WritePacket() (pnet.IPacket, os.Error) {
	packet := NewQTPacket()
	packet.AddUint16(uint16(u.PokeNum))
	packet.AddUint8(uint8(u.SubNum))
	return packet, nil
}