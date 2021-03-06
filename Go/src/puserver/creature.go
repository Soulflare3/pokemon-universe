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
	pos "putools/pos"
	pul "pulogic"
)

// CanSeeCreature checks if 2 creatures are near each others viewport
func CanSeeCreature(_self pul.ICreature, _other pul.ICreature) bool {
	return CanSeePosition(_self.GetPosition(), _other.GetPosition())
}

// CanSeePosition checks if 2 positions are near each others viewport
func CanSeePosition(_p1 pos.Position, _p2 pos.Position) bool {
	if _p1.Z != _p2.Z {
		return false
	}

	return _p1.IsInRange2p(_p2, CLIENT_VIEWPORT_CENTER)
}

// Returns true if the passed creature can move
func CreatureCanMove(_creature pul.ICreature) bool {
	canMove := (_creature.GetTimeSinceLastMove() >= _creature.GetMovementSpeed())
	return canMove
}

// Creature struct with generic variables for all creatures
type Creature struct {
	uid  uint64 // Unique ID
	name string
	Id   int // Database ID			

	Position  *Tile
	Direction int

	Movement  int
	lastStep  int64
	moveSpeed int

	Outfit

	VisibleCreatures pul.CreatureMap
	ConditionList	*list.List
}

func (c *Creature) GetUID() uint64 {
	return c.uid
}

func (c *Creature) GetName() string {
	return c.name
}

func (c *Creature) GetType() int {
	return -1
}

func (c *Creature) GetTile() pul.ITile {
	return c.Position
}

func (c *Creature) SetTile(_tile pul.ITile) {
	c.Position = _tile.(*Tile)
}

func (c *Creature) GetPosition() pos.Position {
	return c.Position.GetPosition()
}

func (c *Creature) GetMovement() int {
	return c.Movement
}

func (c *Creature) GetDirection() int {
	return c.Direction
}

func (c *Creature) SetDirection(_dir int) {
	c.Direction = _dir
}

func (c *Creature) GetOutfit() pul.IOutfit {
	return c.Outfit
}

func (c *Creature) GetMovementSpeed() int {
	return c.moveSpeed
}

func (c *Creature) GetTimeSinceLastMove() int {
	return int(PUSYS_TIME() - c.lastStep)
}

func (c *Creature) OnThink(_interval int) {
	c.ExecuteConditions(_interval)
}

func (c *Creature) OnCreatureMove(_creature pul.ICreature, _from pul.ITile, _to pul.ITile, _teleport bool) {
}

func (c *Creature) OnCreatureTurn(_creature pul.ICreature) {
}

func (c *Creature) OnCreatureAppear(_creature pul.ICreature, _isLogin bool) {
}

func (c *Creature) OnCreatureDisappear(_creature pul.ICreature, _isLogout bool) {
}

func (c *Creature) AddVisibleCreature(_creature pul.ICreature) {
	if _, found := c.VisibleCreatures[_creature.GetUID()]; !found {
		c.VisibleCreatures[_creature.GetUID()] = _creature
	}
}

func (c *Creature) RemoveVisibleCreature(_creature pul.ICreature) {
	if _, found := c.VisibleCreatures[_creature.GetUID()]; !found {
		c.VisibleCreatures[_creature.GetUID()] = _creature
	}
}

func (c *Creature) KnowsVisibleCreature(_creature pul.ICreature) (found bool) {
	_, found = c.VisibleCreatures[_creature.GetUID()]
	return
}

func (c *Creature) GetVisibleCreatures() pul.CreatureMap {
	return c.VisibleCreatures
}

// --------------------- CONDITIONS ---------------------------- //
func (c *Creature) OnTickCondition(_type int, _interval int, _remove bool) {
	// ...
}

func (c *Creature) OnAddCondition(_type int, _hadCondition bool) {
	if _type == CONDITION_INVISIBLE && !_hadCondition {
		g_game.internalCreatureChangeVisible(c, false)
	}
}

func (c *Creature) ExecuteConditions(_interval int) {
	for e := c.ConditionList.Front(); e != nil; {
		condition := e.Value.(ICondition)
		var next *list.Element
		if !condition.ExecuteCondition(c, _interval) {
			next = e.Next()
			condition.EndCondition(c, CONDITIONEND_TICKS)
			lastCondition := !c.HasCondition(condition.GetType(), false)
			c.OnEndCondition(condition.GetType(), lastCondition)
			c.ConditionList.Remove(e)
		} else {
			next = e.Next()
		}
		e = next
	}
}

func (c *Creature) AddCondition(_condition ICondition) bool {
	if _condition == nil {
		return false
	}
	
	hadCondition := c.HasCondition(_condition.GetType(), false)
	prevCond := c.GetCondition(_condition.GetType(), _condition.GetId(), _condition.GetSubId())
	
	if prevCond != nil {
		prevCond.AddCondition(c, _condition)
		return true
	}
	
	if _condition.StartCondition(c) {
		c.ConditionList.PushBack(_condition)
		c.OnAddCondition(_condition.GetType(), hadCondition)
		return true
	}
	
	return false
}

func (c *Creature) GetCondition(_type int, _id int, _subId int) ICondition {
	for e := c.ConditionList.Front(); e != nil; e = e.Next() {
		condition := e.Value.(ICondition)
		if condition.GetType() == _type && condition.GetId() == _id && condition.GetSubId() == _subId {
			return condition
		}
	}
	
	return nil
}

func (c *Creature) HasCondition(_type int, _checkTime bool) bool {
	for e := c.ConditionList.Front(); e != nil; e = e.Next() {
		condition := e.Value.(ICondition)
		if condition.GetType() == _type && (!_checkTime || condition.GetEndTime() == 0 || condition.GetEndTime() >= PUSYS_TIME()) {
			return true
		}
	}
	
	return false
}

func (c *Creature) OnEndCondition(_type int, _lastCondition bool) {
	if _type == CONDITION_INVISIBLE && _lastCondition {
		g_game.internalCreatureChangeVisible(c, true);
	}
}