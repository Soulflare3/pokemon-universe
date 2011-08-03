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

type BattleStats struct {
	stats []int16
}

func NewBattleStats() BattleStats {
	return BattleStats{stats: make([]int16, 5)}
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
