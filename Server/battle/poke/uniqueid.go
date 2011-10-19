package main

import (
	pnet "network"
)

type UniqueId struct {
	PokeNum int
	SubNum int
}

func NewUniqueId(_pokeNum, _subNum int) *UniqueId {
	uniqueId := UniqueId { PokeNum: _pokeNum,
							SubNum: _subNum }
	return &uniqueId
}

func NewUniqueIdFromPacket(_packet *pnet.QTPacket) *UniqueId {
	uniqueId := UniqueId { PokeNum: (int)_packet.ReadUint16(),
							SubNum: (int)_packet.ReadByte() }
	return &uniqueId
}