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
	
	pul "pulogic"
	pkmn "pulogic/pokemon"
	puh "puhelper"
)

type PlayerList map[uint64]*Player

type Player struct {
	Creature     		// Inherit generic creature data
	dbid     		int // database id

	Conn 			*Connection

	Pokemon			pkmn.PlayerPokemonList
	PokemonParty	*pkmn.PokemonParty
	Friends			FriendList
	
	Backpack *Depot
	Storage	*Depot

	Location       	*Location
	LastPokeCenter 	*Tile
	InteractingNpc	*Npc
	
	Quests			PlayerQuestList

	Money          	int
	TimeoutCounter	int
	GroupFlags		int64
}

func NewPlayer(_name string) *Player {
	p := &Player{}
	p.uid = puh.GenerateUniqueID()
	p.Conn = nil
	p.Outfit = NewOutfit()
	p.name = _name

	p.Pokemon = make(pkmn.PlayerPokemonList)
	p.PokemonParty = pkmn.NewPokemonParty()
	p.Friends = make(FriendList)
	
	p.Backpack = NewDepot(25)
	p.Storage = NewDepot(100)
	
	p.lastStep = PUSYS_TIME()
	p.moveSpeed = 250
	p.VisibleCreatures = make(pul.CreatureList)
	p.ConditionList = list.New()
	p.TimeoutCounter = 0
	
	p.Quests = make(PlayerQuestList)
	
	// Add self to visible creatures
	p.VisibleCreatures[p.GetUID()] = p

	return p
}

// --------------------- INTERFACE ----------------------------//

func (p *Player) GetType() int {
	return CTYPE_PLAYER
}

func (p *Player) SetConnection(_conn *Connection) {
	p.Conn = _conn
	p.Conn.Owner = p
}

// Called by Connection to remove itself from its owner
// when the player disconnects
func (p *Player) removeConnection() {
	if p.Conn == nil || !p.Conn.IsOpen {
		g_game.OnPlayerLoseConnection(p)
	}
}

func (p *Player) SetMoney(_money int) int {
	if p.Money += _money; p.Money < 0 {
		p.Money = 0
	}
	return p.Money
}

func (p *Player) GetMoney() int {
	return p.Money
}

func (p *Player) GetPokemonParty() *pkmn.PokemonParty {
	return p.PokemonParty
}

func (p *Player) OnCreatureMove(_creature pul.ICreature, _from pul.ITile, _to pul.ITile, _teleport bool) {
	if _creature.GetUID() == p.GetUID() {
		p.lastStep = PUSYS_TIME()
		return
	}
	
	from := _from.(*Tile)
	to := _to.(*Tile)

	canSeeFromTile := CanSeePosition(p.GetPosition(), from.Position)
	canSeeToTile := CanSeePosition(p.GetPosition(), to.Position)

	if canSeeFromTile && !canSeeToTile { // Leaving viewport
		p.sendCreatureMove(_creature, from, to)

		p.RemoveVisibleCreature(_creature)
		_creature.RemoveVisibleCreature(p)
	} else if canSeeToTile && !canSeeFromTile { // Entering viewport
		p.AddVisibleCreature(_creature)
		_creature.RemoveVisibleCreature(p)

		p.sendCreatureMove(_creature, from, to)
	} else { // Moving inside viewport
		p.AddVisibleCreature(_creature)
		_creature.AddVisibleCreature(p)

		p.sendCreatureMove(_creature, from, to)
	}
}

func (p *Player) OnCreatureTurn(_creature pul.ICreature) {
	if _creature.GetUID() != p.GetUID() {
		p.sendCreatureTurn(_creature)
	}
}

func (p *Player) OnCreatureAppear(_creature pul.ICreature, _isLogin bool) {
	if _creature.GetUID() == p.GetUID() {
		return
	}
	
	if _isLogin {
		// Check if creature is in friendlist
		p.UpdateFriend(_creature.GetName(), true)
	}
	
	canSeeCreature := CanSeeCreature(p, _creature)
	if !canSeeCreature {
		return
	}

	// We're checking inside the AddVisibleCreature method so no need to check here
	p.AddVisibleCreature(_creature)
	_creature.AddVisibleCreature(p)
}

func (p *Player) OnCreatureDisappear(_creature pul.ICreature, _isLogout bool) {
	if _creature.GetUID() == p.GetUID() {
		return
	}
	
	if _isLogout {
		// Check if creature is in friendlist
		p.UpdateFriend(_creature.GetName(), false)
	}

	p.RemoveVisibleCreature(_creature)
}

func (p *Player) AddVisibleCreature(_creature pul.ICreature) {
	if _, found := p.VisibleCreatures[_creature.GetUID()]; !found {
		p.VisibleCreatures[_creature.GetUID()] = _creature
		p.sendCreatureAdd(_creature)
	}
}

func (p *Player) RemoveVisibleCreature(_creature pul.ICreature) {
	// No need to check if the key actually exists because Go is awesome
	// http://tip.golang.org/doc/effective_go.html#maps
	delete(p.VisibleCreatures, _creature.GetUID())
	p.sendCreatureRemove(_creature)
}

// ------------------------------------------------------ //

func (p *Player) HealParty() {
	p.PokemonParty.HealParty()
	
	// TODO: Send update to client
}

func (p *Player) SetFlags(_flags int64) {
	p.GroupFlags = _flags
}

func (p *Player) HasFlag(_value uint64) bool {
	return (0 != (p.GroupFlags & (1 << _value)))
}

func (p *Player) GetQuestStatus(_questId int64) (status int) {
	status = 0
	if quest, found := p.Quests[_questId]; found {
		status = quest.Status
	}
	
	return
}

func (p *Player) SetQuestStatus(_questId int64, _status int) {
	quest, found := p.Quests[_questId]
	if found {
		quest.Status = _status
	} else {
		quest = NewPlayerQuest(_questId, _status)
		if quest != nil {
			p.Quests[_questId] = quest
		}
	}
	
	p.sendQuestUpdate(quest)
}

func (p *Player) AbandonQuest(_questId int64) {
	if quest, found := p.Quests[_questId]; found {
		quest.Abandon()
	}
}