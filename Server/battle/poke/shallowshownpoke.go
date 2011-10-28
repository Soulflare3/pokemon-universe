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
	shallowShownPoke := ShallowShownPoke{ UID: NewUniqueIdFromPacket(_packet),
											Level: int(_packet.ReadUint8()),
											Gender: int(_packet.ReadUint8()),
											Item: _packet.ReadBool() }
	return &shallowShownPoke
}