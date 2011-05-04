package main

import pnet "network"

const (
	BattleDynamicInfo_Spikes = 1
	BattleDynamicInfo_SpikesLV2 = 2
	BattleDynamicInfo_SpikesLV3 = 4
	BattleDynamicInfo_StealthRock = 8
	BattleDynamicInfo_ToxicSpikes = 16
	BattleDynamicInfo_ToxicSpikesLV2 = 32
)

type BattleDynamicInfo struct {
	boosts	[]int8
	flags	uint8
}

func NewBattleDynamicInfo() BattleDynamicInfo {
	return BattleDynamicInfo{ boosts: make([]int8, 7) }
}

func NewBattleDynamicInfoFromPacket(_packet *pnet.QTPacket) BattleDynamicInfo {
	info := NewBattleDynamicInfo()
	info.boosts[0] = int8(_packet.ReadUint8())
	info.boosts[1] = int8(_packet.ReadUint8())
	info.boosts[2] = int8(_packet.ReadUint8())
	info.boosts[3] = int8(_packet.ReadUint8())
	info.boosts[4] = int8(_packet.ReadUint8())
	info.boosts[5] = int8(_packet.ReadUint8())
	info.boosts[6] = int8(_packet.ReadUint8())
	info.flags = _packet.ReadUint8()
	
	return info
}