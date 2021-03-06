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

type PU_BattleEvent_ChangeLevelSelf struct {
	pokemon int
	level   int
}

func NewBattleEvent_ChangeLevelSelf(_pokemon int, _level int) *PU_BattleEvent_ChangeLevelSelf {
	return &PU_BattleEvent_ChangeLevelSelf{pokemon: _pokemon, level: _level}
}

func (e *PU_BattleEvent_ChangeLevelSelf) Execute() {
	pokemon := g_game.self.pokemon[e.pokemon]
	if pokemon != nil {
		pokemon.level = int16(e.level)
	}
}
