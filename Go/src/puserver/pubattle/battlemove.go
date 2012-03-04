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
	"pulogic/pokemon"
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
	move := pokemon.GetInstance().GetMoveById(_id)
	battleMove := BattleMove{}
	battleMove.CurrentPP = move.PP
	battleMove.TotalPP = move.PP
	battleMove.Num = move.MoveId
	battleMove.Name = move.Identifier
	battleMove.Type = move.TypeId
	battleMove.power = string(move.Power)
	battleMove.accuracy = string(move.Accuracy)
	battleMove.description = ""
	battleMove.effect = ""
	
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
	battleMove := NewBattleMoveFromId(int(_packet.ReadUint16()))
	battleMove.CurrentPP = int(_packet.ReadUint8())
	battleMove.TotalPP = int(_packet.ReadUint8())
	
	return battleMove
}