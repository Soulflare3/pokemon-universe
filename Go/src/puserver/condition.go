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
	"putools/math"
)

const (
    CONDITION_NONE				= 0
    CONDITION_INVISIBLE			= 1 << 0
    CONDITION_INFIGHT			= 1 << 1
    CONDITION_DRUNK				= 1 << 2
    CONDITION_EXHAUST_YELL		= 1 << 3
)

const (
	CONDITIONID_DEFAULT int = -1
	CONDITIONID_COMBAT = 0
)

const (
	CONDITIONEND_CLEANUP int = iota
	CONDITIONEND_DIE
	CONDITIONEND_TICKS
	CONDITIONEND_ABORT
)

type ICondition interface {
	StartCondition(creature pul.ICreature) bool
	ExecuteCondition(creature pul.ICreature, interval int) bool
	EndCondition(creature pul.ICreature, reason int)
	AddCondition(creature pul.ICreature, condition ICondition)
	
	GetId() int
	GetSubId() int
	GetType() int
	GetTicks() int
	GetEndTime() int64
}

type Condition struct {
	id				int
	subId			int
	ticks			int
	endTime			int64
	conditionType	int
}

func NewCondition(_id int, _type int, _ticks int) *Condition {
	return &Condition { id: _id,
						subId: 0,
						ticks: _ticks,
						endTime: 0,
						conditionType: _type }
}

func CreateCondition(_id int, _type int, _ticks int, _param int) *Condition {
	switch _type {
		case CONDITION_DRUNK:
			fallthrough
		case CONDITION_INFIGHT:
			fallthrough
		case CONDITION_INVISIBLE:
			fallthrough
		case CONDITION_EXHAUST_YELL:
			return NewCondition(_id, _type, _ticks)
	}
	
	return nil
}

func (c *Condition) GetId() int {
	return c.id
}

func (c *Condition) GetSubId() int {
	return c.subId
}

func (c *Condition) GetType() int {
	return c.conditionType
}

func (c *Condition) GetTicks() int {
	return c.ticks
}

func (c *Condition) SetTicks(_newTicks int) {
	c.ticks = _newTicks
	c.endTime = int64(c.GetTicks()) + PUSYS_TIME()
}

func (c *Condition) GetEndTime() int64 {
	if c.GetTicks() == -1 {
		return 0 
	}
	return c.endTime
}

func (c *Condition) StartCondition(_creature pul.ICreature) bool {
	if c.GetTicks() > 0 {
		c.endTime = int64(c.GetTicks()) + PUSYS_TIME()
	}
	return true
}

func (c *Condition) ExecuteCondition(_creature pul.ICreature, _interval int) bool {
	if _interval > 0 {
		_creature.OnTickCondition(c.conditionType, _interval, false)
	}
	
	if c.GetTicks() != -1 {
		newTicks := putools.Imax(0, c.GetTicks() - _interval)
		// Not using set ticks here since it would reset endTime
		c.ticks = newTicks
		
		return (c.GetEndTime() >= PUSYS_TIME())
	}
	
	return true
}

func (c *Condition) EndCondition(_creature pul.ICreature, _reason int) {
	// ...
}

func (c *Condition) AddCondition(_creature pul.ICreature, _condition ICondition) {
	if c.UpdateCondition(_condition) {
		c.SetTicks(_condition.GetTicks())
	}
}

func (c *Condition) UpdateCondition(_condition ICondition) bool {
	if c.GetType() != _condition.GetType() {
		return false
	}
	
	if c.GetTicks() == -1 && _condition.GetTicks() > 0 {
		return false
	}
	
	if _condition.GetTicks() >= 0 && c.GetEndTime() > (PUSYS_TIME() + int64(_condition.GetTicks())) {
		return false
	}
	
	return true
}