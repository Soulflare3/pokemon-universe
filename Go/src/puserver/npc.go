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
	"container/list"
	"math"
	"math/rand"
	
	"npclib"
	"putools/log"
	pos "putools/pos"
	pnet "network"
	pul "pulogic"
	puh "puhelper"
)

type Npc struct {
	Creature
	
	dbid				int
	script_name			string
	script 				npclib.NpcInteractionInterface
	
	interactingPlayers	PlayerList
	
	moveInterval		int
	moveRadius			int
	moveCenter			pos.Position
	ticksWithoutPlayer	int
}

func NewNpc() *Npc {
	n := Npc{}
	n.uid = puh.GenerateUniqueID()
	n.Outfit = NewOutfit()
	n.moveSpeed = 280
	n.VisibleCreatures = make(pul.CreatureList)
	n.ConditionList = list.New()
	n.script = nil
	
	n.interactingPlayers = make(PlayerList)

	n.moveInterval = 5
	n.moveRadius = 5
	n.ticksWithoutPlayer = 0
		
	return &n
}

func (n *Npc) Load(_data []interface{}) bool {
	id := puh.DBGetInt(_data[0])
	name := puh.DBGetString(_data[1])
	script_name := puh.DBGetString(_data[2])
	position := puh.DBGetInt64(_data[3])

	n.dbid = id
	n.name = name
	n.script_name = script_name
	
	tile, ok := g_map.GetTile(position)
	if !ok {
		logger.Printf("[Error] Could not load position info for npc %s (%d)\n", n.name, n.dbid)
		return false
	}
	n.Position = tile
	n.moveCenter = tile.GetPosition()
	
	return true
}

func (n *Npc) GetType() int {
	return CTYPE_NPC
}

func (n *Npc) OnCreatureMove(_creature pul.ICreature, _from pul.ITile, _to pul.ITile, _teleport bool) {
	// Check if _creature is a Player, otherwise return
	if _creature.GetType() != CTYPE_PLAYER {
		return
	}

	canSeeFromTile := CanSeePosition(n.GetPosition(), _from.GetPosition())
	canSeeToTile := CanSeePosition(n.GetPosition(), _to.GetPosition())

	if canSeeFromTile && !canSeeToTile { // Leaving viewport
		n.RemoveVisibleCreature(_creature)
		_creature.RemoveVisibleCreature(n)
	} else if canSeeToTile && !canSeeFromTile { // Entering viewport
		n.AddVisibleCreature(_creature)
		_creature.RemoveVisibleCreature(n)
	} else { // Moving inside viewport
		n.AddVisibleCreature(_creature)
		_creature.AddVisibleCreature(n)
	}
}

func (n *Npc) OnCreatureTurn(_creature pul.ICreature) {
	// Check if _creature is a Player, otherwise return
	if _creature.GetType() != CTYPE_PLAYER {
		return
	}
}

func (n *Npc) OnCreatureAppear(_creature pul.ICreature, _isLogin bool) {
	// Check if _creature is a Player, otherwise return
	if _creature.GetType() != CTYPE_PLAYER {
		return
	}
	
	canSeeCreature := CanSeeCreature(n, _creature)
	if !canSeeCreature {
		return
	}

	// We're checking the existence of _creature inside the AddVisibleCreature method 
	// so no need to check here
	n.AddVisibleCreature(_creature)
	_creature.AddVisibleCreature(n)
}

func (n *Npc) OnCreatureDisappear(_creature pul.ICreature, _isLogout bool) {
	// Check if _creature is a Player, otherwise return
	if _creature.GetType() != CTYPE_PLAYER {
		return
	}

	n.RemoveVisibleCreature(_creature)
}

func (n *Npc) AddVisibleCreature(_creature pul.ICreature) {
	// Check if _creature is a Player, otherwise return
	if _creature.GetType() != CTYPE_PLAYER {
		return
	}
	
	if _, found := n.VisibleCreatures[_creature.GetUID()]; !found {
		n.VisibleCreatures[_creature.GetUID()] = _creature
	}
}

func (n *Npc) RemoveVisibleCreature(_creature pul.ICreature) {
	// Check if _creature is a Player, otherwise return
	if _creature.GetType() != CTYPE_PLAYER {
		return
	}
	
	// Check if the player was interacting with this npc
	if n.HasInteractingPlayer(_creature.(*Player)) {
		n.RemoveInteractingPlayer(_creature.(*Player))
	}

	// No need to check if the key actually exists because Go is awesome
	// http://golang.org/doc/effective_go.html#maps
	delete(n.VisibleCreatures, _creature.GetUID())
}

// -------------------------------------------------------- //

func (n *Npc) AutoWalk() {
	// Check if this NPC is allowed to move
	if n.moveInterval == 0 || n.GetTimeSinceLastMove() < n.moveInterval || !CreatureCanMove(n) {
		return
	}
	
	// Increment ticker when no players are around
	// We do this to let the NPC move around for a bit when there are no players around
	// to make everything seem more alive
	if len(n.VisibleCreatures) == 0 {
		if n.ticksWithoutPlayer >= 5 {
			return
		} else {
			n.ticksWithoutPlayer++
		}
	} else {
		n.ticksWithoutPlayer = 0
	}
	
	// Don't move when interacting with one or more players
	if len(n.interactingPlayers) > 0 {
		return
	}
	
	rndDirection := rand.Intn(3) + 1
	moveDirection := DIR_NULL
	switch rndDirection {
		case 1:
			moveDirection = DIR_SOUTH
		case 2:
			moveDirection = DIR_WEST
		case 3:
			moveDirection = DIR_NORTH
		case 4:
			moveDirection = DIR_EAST
	}
	
	if n.moveRadius > 0 && n.CanWalk(moveDirection) {
		if g_game.OnCreatureMove(n, moveDirection) != RET_NOTPOSSIBLE {
			n.lastStep = PUSYS_TIME()
		}
	} else { // Can't walk,, just turn
		g_game.OnCreatureTurn(n, moveDirection)
	}
}

func (n *Npc) CanWalk(_direction int) (ret bool) {
	newPosition := n.GetPosition()
	switch _direction {
		case DIR_SOUTH:
			newPosition.Y++
		case DIR_WEST:
			newPosition.X--
		case DIR_NORTH:
			newPosition.Y--
		case DIR_EAST:
			newPosition.X++
	}
	
	ret = false
	if (newPosition.X >= n.moveCenter.X - n.moveRadius) && (newPosition.X <= n.moveCenter.X + n.moveRadius) &&	(newPosition.Y >= n.moveCenter.Y - n.moveRadius) &&	(newPosition.Y <= n.moveCenter.Y + n.moveRadius) {
		ret = true;
	}
	
	return
}

// -------------------------------------------------------- //

func (n *Npc) AddInteractingPlayer(_player *Player) {
	n.interactingPlayers[_player.GetUID()] = _player
	_player.InteractingNpc = n
}

func (n *Npc) HasInteractingPlayer(_player *Player) bool {
	_, found := n.interactingPlayers[_player.GetUID()]
	return found
}

func (n *Npc) RemoveInteractingPlayer(_player *Player) {
	_player.InteractingNpc = nil
	delete(n.interactingPlayers, _player.GetUID())
}

func (n *Npc) SelfSay(_message string) {
	g_game.internalCreatureSay(n, pnet.SPEAK_NORMAL, _message)
}

func (n *Npc) StartDialog(_player *Player) {
	if !n.HasInteractingPlayer(_player) {
		n.AddInteractingPlayer(_player)
		
		// Calculate turn direction
		playerPos := _player.GetPosition()
		npcPos := n.GetPosition()
		diff := playerPos.Sub(npcPos)
		
		var tan float64 = 10
		if diff.X != 0 {
			tan = float64(diff.Y / diff.X)
		}

		direction := DIR_SOUTH // Default		
		if math.Abs(tan) < 1 {
			if diff.X > 0 {
				direction = DIR_WEST
			} else {
				direction = DIR_EAST
			}
		} else if diff.Y > 0 {
			direction = DIR_NORTH
		}
		
		// Turn creature for this player only
		_player.sendCreatureTurnExclusive(n, direction)
		
		// Start conversation
		n.OnDialogueAnswer(_player.GetUID(), 0)
	}
}

func (n *Npc) OnDialogueAnswer(_cid uint64, _answer int) {
	if n.script != nil {
		n.script.OnAnswer(_cid, _answer)
	}
}