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

type PU_BattleEvent_ChangePokemon struct {
	fighter int 
	pokeid int 
	hp int 
	level int 
	name string 
}

func NewBattleEvent_ChangePokemon_Self(_pokeid int) *PU_BattleEvent_ChangePokemon {
	event := &PU_BattleEvent_ChangePokemon{}
	event.fighter = -1
	event.pokeid = _pokeid
	return event
}

func NewBattleEvent_ChangePokemon(_fighter int, _pokeid int, _name string, _hp int, _level int) *PU_BattleEvent_ChangePokemon {
	event := &PU_BattleEvent_ChangePokemon{}
	event.fighter = _fighter
	event.pokeid = _pokeid
	event.name = _name
	event.hp = _hp
	event.level = _level
	return event
}

func (e *PU_BattleEvent_ChangePokemon) Execute() {
	if e.fighter == -1 {
		g_game.battle.SetPokemon(e.pokeid)
		return
	} else {
		fighter := g_game.battle.fighters[e.fighter]
		if fighter != nil {
			fighter.SetPokemon(e.name, e.pokeid, e.level, e.hp)
		}
	}
}
