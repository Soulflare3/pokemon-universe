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
	pul "pulogic"
	pos "putools/pos"
)

type Warp struct {
	destination pos.Position
}

func NewWarp(_destination pos.Position) *Warp {
	return &Warp{destination: _destination}
}

func (e *Warp) OnCreatureEnter(_creature pul.ICreature, _prevRet int) (ret int) {
	currentTile := _creature.GetTile()
	destinationTile, found := g_map.GetTileFromPosition(e.destination)

	if found {
		ret = g_game.internalCreatureTeleport(_creature, currentTile, destinationTile)
	} else {
		ret = RET_NOTPOSSIBLE
	}

	return
}

func (e *Warp) OnCreatureLeave(_creature pul.ICreature, _prevRet int) int {
	return RET_NOERROR
}
