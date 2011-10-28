package main

import (
	pnet "network"
)

type BattleConf struct {
	Gen int
	Mode int
	Ids []int
	Clauses int
}

func NewBattleConfFromPacket(_packet *pnet.QTPacket) *BattleConf {
	battleConf := BattleConf{}
	battleConf.Gen = int(_packet.ReadUint8())
	battleConf.Mode = int(_packet.ReadUint8())
	battleConf.Ids = make([]int, 2)
	battleConf.Ids[0] = int(_packet.ReadUint32())
	battleConf.Ids[1] = int(_packet.ReadUint32())
	battleConf.Clauses = int(_packet.ReadUint32())
	
	return &battleConf
}