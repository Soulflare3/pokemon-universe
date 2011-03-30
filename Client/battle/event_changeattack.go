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

type PU_BattleEvent_ChangeAttack struct {
	pokemon int 
	slot int
	pp int
	ppmax int
	power int
	accuracy int
	
	name string
	description string
	poketype string
	contact string
	category string
	target string
}

func NewBattleEvent_ChangeAttack(_pokemon int, _slot int, _name string, _description string, _poketype string, _pp int, _ppmax int, _power int, _accuracy int, _category string, _target string, _contact string) *PU_BattleEvent_ChangeAttack {
	event := &PU_BattleEvent_ChangeAttack{}
	event.pokemon = _pokemon
	event.slot = _slot
	event.name = _name
	event.description = _description
	event.poketype = _poketype
	event.pp = _pp
	event.ppmax = _ppmax
	event.power = _power
	event.accuracy = _accuracy
	event.category = _category
	event.target = _target
	event.contact = _contact
	return event
}

func (e *PU_BattleEvent_ChangeAttack) Execute() {
	pokemon := g_game.self.pokemon[e.pokemon]
	if pokemon != nil {
		pokemon.SetAttack(e.slot, e.name, e.description, e.poketype, uint16(e.pp), uint16(e.ppmax), uint16(e.power), uint16(e.accuracy), e.category, e.target, e.contact)
	}
}
