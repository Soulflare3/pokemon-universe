package main

import (
	pnet "network"
)

type BattlePoke struct {
	ShallowBattlePoke // Extends
	
	CurrentHP int
	TotalHP int
	ItemString string
	AbilityString string
	TeamNum int
	Stats []int
	Moves []*BattleMove
	
	item int
	ability int
	
	statusCount int
	originalStatusCount int
	nature int
	happiness int
	
	DVs []int
	EVs []int
}

func NewBattlePokeFromPacket(_packet *pnet.QTPacket) *BattlePoke {
	battlePoke := BattlePoke{}
	battlePoke.UID = NewUniqueIdFromPacket(_packet)
	battlePoke.Nick = _packet.ReadString()
	battlePoke.TotalHP = int(_packet.ReadUint16())
	battlePoke.Gender = int(_packet.ReadUint8())
	battlePoke.Shiny = _packet.ReadBool()
	battlePoke.Level = int(_packet.ReadUint8())
	battlePoke.item = int(_packet.ReadUint16())
	battlePoke.ItemString = ""
	battlePoke.ability = int(_packet.ReadUint16())
	battlePoke.happiness = int(_packet.ReadUint8())
	
	//getType()
	//getName()
	
	battlePoke.Stats = make([]int, 5)
	for i := 0; i < 5; i++ {
		battlePoke.Stats[i] = int(_packet.ReadUint16())
	}
	battlePoke.Moves = make([]*BattleMove, 4)
	for i := 0; i < 4; i++ {
		battlePoke.Moves[i] = NewBattleMoveFromPacket(_packet)
	}
	battlePoke.EVs = make([]int, 6)
	for i := 0; i < 6; i++ {
		battlePoke.EVs[i] = int(_packet.ReadUint32())
	}
	battlePoke.DVs = make([]int, 6)
	for i := 0; i < 6; i++ {
		battlePoke.DVs[i] = int(_packet.ReadUint32())
	}
	
	return &battlePoke
}