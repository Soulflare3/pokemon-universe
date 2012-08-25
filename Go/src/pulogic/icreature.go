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
package pulogic

import (
	"pulogic/pokemon"
	pnet "network"
	pos "putools/pos"
)

// CreatureList is map which holds a list of ICreature interfaces
type CreatureList map[uint64]ICreature

// Main interface for all creature objects in the game
type ICreature interface {
	GetUID() uint64
	GetName() string
	GetType() int

	GetPosition() pos.Position
	GetMovement() int

	SetDirection(_dir int)
	GetDirection() int

	SetTile(_tile ITile)
	GetTile() ITile

	GetOutfit() IOutfit

	GetMovementSpeed() int
	GetTimeSinceLastMove() int

	// Methods for all moving creatures
	OnThink(_interval int)
	OnCreatureMove(_creature ICreature, _from ITile, _to ITile, _teleport bool)
	OnCreatureTurn(_creature ICreature)
	OnCreatureAppear(_creature ICreature, _isLogin bool)
	OnCreatureDisappear(_creature ICreature, _isLogout bool)
	
	OnTickCondition(_type int, _interval int, _remove bool)

	// Methods for all creatures who need to see other creatures	
	AddVisibleCreature(_creature ICreature)
	RemoveVisibleCreature(_creature ICreature)
	KnowsVisibleCreature(_creature ICreature) bool
	GetVisibleCreatures() CreatureList
}

type IBattleCreature interface {
	GetName() string
	GetPokemonParty() *pokemon.PokemonParty
	GetType() int
	
	SendBattleMessage(_message pnet.INetMessageWriter)
}