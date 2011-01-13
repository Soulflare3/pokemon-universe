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

import (
	pos "position"
)

// Main interface for all creature objects in the game
type ICreature interface {
	GetName() string
	GetPosition() pos.Position
}

// Interface for all moving creatures
type ICreatureMove interface {
	OnCreatureMove(_creature ICreature)
	OnCreatureAppear(_creature ICreature, _isLogin bool)
	OnCreatureDisappear(_creature ICreature, _isLogout bool)
}

// CreatureList is map which holds a list of ICreature interfaces
type CreatureList map[int]ICreature

// Interface for all creatures who need to see other creatures
type ICreatureSee interface {
	AddVisibleCreature(_creature ICreature)
	RemoveVisibleCreature(_creature ICreature)
	KnowsVisibleCreature(_creature ICreature) bool
}

// CanSeeCreature checks if 2 creatures are near each others viewport
func CanSeeCreature(_self ICreature, _other ICreature) bool {
	return CanSeePosition(_self.GetPosition(), _other.GetPosition())
}

// CanSeePosition checks if 2 positions are near each others viewport
func CanSeePosition(_p1 pos.Position, _p2 pos.Position) bool {
	if _p1.Z != _p2.Z {
		return false
	}

	return _p1.IsInRange2p(_p2, CLIENT_VIEWPORT_CENTER)
}

