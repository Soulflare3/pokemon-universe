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
package pubattle

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

func (b *BattleConf) GetId(_index int) int {
	return b.Ids[_index]
}