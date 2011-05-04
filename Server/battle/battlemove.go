package main

type BattleMove struct {
	pp		uint8
	totalPP	uint8
	num		uint16
}

func NewBattleMove() *BattleMove {
	return &BattleMove{ }
}