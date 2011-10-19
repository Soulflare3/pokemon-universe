package main

import (
	pnet "network"
)

type ShallowShownPoke struct {
	Item bool
	UID *UniqueId
	Level int
	Gender int
}

func NewShallowShownPokeFromPacket(_packet *pnet.QTPacket) *ShallowShownPoke {
	shallowShownPoke := S`hallowShownPoke{ UID: NewUniqueIdFromPacket(_packet),
											Level: (int)_packet.ReadByte(),
											Gender: (int)_packet.ReadByte(),
											Item: _packet.ReadBool() }
	return &shallowShownPoke
}