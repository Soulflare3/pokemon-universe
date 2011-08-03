/*Pokemon Universe MMORPG
Copyright (C) 2010 the Pokemon Universe Authors

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.*/
package main

import pnet "network"

const (
	BattleDynamicInfo_Spikes         = 1
	BattleDynamicInfo_SpikesLV2      = 2
	BattleDynamicInfo_SpikesLV3      = 4
	BattleDynamicInfo_StealthRock    = 8
	BattleDynamicInfo_ToxicSpikes    = 16
	BattleDynamicInfo_ToxicSpikesLV2 = 32
)

type BattleDynamicInfo struct {
	boosts []int8
	flags  uint8
}

func NewBattleDynamicInfo() BattleDynamicInfo {
	return BattleDynamicInfo{boosts: make([]int8, 7)}
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
