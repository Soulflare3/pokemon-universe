package main

import (
	pnet "network"
)

type BattleMove struct {
	CurrentPP int
	TotalPP int
	Num int
	Name string
	Type int
	
	power string
	accuracy string
	description string
	effect string
}

func NewBattleMove() *BattleMove {
	return &BattleMove{}
}

func NewBattleMoveFromId(_id int) *BattleMove {
	battleMove := BattleMove{}
	
	return &battleMove
}

func NewBattleMoveFromBattleMove(_battleMove *BattleMove) *BattleMove {
	battleMove := &BattleMove{}
	battleMove.CurrentPP = _battleMove.CurrentPP
	battleMove.TotalPP = _battleMove.TotalPP
	battleMove.Num = _battleMove.Num
	battleMove.Name = _battleMove.Name
	battleMove.Type = _battleMove.Type
	battleMove.power = _battleMove.power
	battleMove.accuracy = _battleMove.accuracy
	battleMove.description = _battleMove.description
	battleMove.effect = _battleMove.effect
	
	return battleMove
}

func NewBattleMoveFromPacket(_packet *pnet.QTPacket) *BattleMove {
	battleMove := NewBattleMoveFromId(int(_packet.ReadUint32()))
	battleMove.CurrentPP = int(_packet.ReadUint8())
	battleMove.TotalPP = int(_packet.ReadUint8())
	
	return battleMove
}