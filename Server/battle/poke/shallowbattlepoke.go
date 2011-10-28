package main

import (
	pnet "network"
)

type ShallowBattlePoke struct {
	RNick string
	Nick string
	PokeName string
	UID *UniqueId
	Types []int
	Shiny bool
	Gender int
	LifePercent int
	Level int
	LastKnowPercent int
	Sub bool
	
	fullStatus int
}

func NewShallowBattlePoke() *ShallowBattlePoke {
	return &ShallowBattlePoke{ Types: make([]int, 2) }
}

func NewShallowBattlePokeFromPacket(_packet *pnet.QTPacket, _isMe bool) *ShallowBattlePoke {
	shallowPoke := ShallowBattlePoke{ Types: make([]int, 2) }
	shallowPoke.UID = NewUniqueIdFromPacket(_packet)
	shallowPoke.RNick = _packet.ReadString()
	shallowPoke.Nick = shallowPoke.RNick
	if !_isMe {
		shallowPoke.Nick = "the foe's " + shallowPoke.Nick
		
		//getName()
		//getTypes()
	}
	
	shallowPoke.LifePercent = int(_packet.ReadUint8())
	shallowPoke.fullStatus = int(_packet.ReadUint32())
	shallowPoke.Gender = int(_packet.ReadUint8())
	shallowPoke.Shiny = _packet.ReadBool()
	shallowPoke.Level = int(_packet.ReadUint32())
}