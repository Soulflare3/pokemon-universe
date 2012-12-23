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
package pokemon

import "pulogic/models"

type Move struct {
	MoveId				int
	Identifier			string
	TypeId				int
	Power				int
	PP					int
	Accuracy			int
	Priority			int
	TargetId			int
	DamageClassId		int
	EffectId			int
	EffectChance		int
	ContestType			int
	ContestEffect		int
	SuperContestEffect	int
	FlavorText			string
}

func NewMove() *Move {
	return &Move{}
}

func NewMoveFromEntity(_entity models.MovesJoinMoveFlavorText) *Move {
	move := NewMove()
	
	move.MoveId = _entity.Id
	move.Identifier = _entity.Identifier
	move.TypeId = _entity.TypeId
	move.Power = _entity.Power
	move.PP = _entity.Pp
	move.Accuracy = _entity.Accuracy
	move.Priority = _entity.Priority
	move.TargetId = _entity.TargetId
	move.DamageClassId = _entity.DamageClassId
	move.EffectId = _entity.EffectId
	move.EffectChance = _entity.EffectChance
	move.ContestType = _entity.ContestTypeId
	move.ContestEffect = _entity.ContestEffectId
	move.SuperContestEffect = _entity.SuperContestEffectId
	move.FlavorText = _entity.FlavorText
	
	return move
}