package main

import (
	pnet "network"
)

const (
	DYNAMICINFO_SPYKES byte = 1
	DYNAMICINFO_SPYKESL2 = 2
	DYNAMICINFO_SPYKESL3 = 4
	DYNAMICINFO_STEALTHROCK = 8
	DYNAMICINFO_TOXICSPYKES = 16
	DYNAMICINFO_TOXICSPYKES = 32
)

type BattleDynamicInfo struct {
	Boosts []byte
	Flags byte
}

func NewBattleDynamicInfoFromPacket(_packet *pnet.QTPacket) *BattleDynamicInfo {
	battleDynamicInfo := BattleDynamicInfo{}
	battleDynamicInfo.Boosts = make([]byte, 7)
	for i := 0; i < 7; i++ {
		battleDynamicInfo.Boosts[i] = _packet.ReadByte()
	}
	battleDynamicInfo.Flags = _packet.ReadByte()
}