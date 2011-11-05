package main

import (
	pnet "network"
)

type BattleTeam struct {
	nick string
	info string
	gen int
	
	Pokes []*BattlePoke
	indexes []int
}

func NewBattleTeamFromPacket(_packet *pnet.QTPacket) *BattleTeam {
	battleTeam := BattleTeam { Pokes: make([]*BattlePoke, 6),
								indexes: make([]int, 6) }
	for i := 0; i < 6; i++ {
		battleTeam.Pokes[i] = NewBattlePokeFromPacket(_packet)
		battleTeam.Pokes[i].TeamNum = i
	}
	
	return &battleTeam
}