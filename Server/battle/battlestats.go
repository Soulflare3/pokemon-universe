package main

import pnet "network"

type BattleStats struct {
	stats []int16
}

func NewBattleStats() BattleStats {
	return BattleStats{stats: make([]int16, 5) }
}

func NewBattleStatsFromPacket(_packet *pnet.QTPacket) BattleStats {
	battleStats := NewBattleStats()
	battleStats.stats[0] = int16(_packet.ReadUint16())
	battleStats.stats[1] = int16(_packet.ReadUint16())
	battleStats.stats[2] = int16(_packet.ReadUint16())
	battleStats.stats[3] = int16(_packet.ReadUint16())
	battleStats.stats[4] = int16(_packet.ReadUint16())
	
	return battleStats
}