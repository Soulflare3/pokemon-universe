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

type PU_BattleEvent_ChangeHP struct {
	fighter int
	hp      int
}

func NewBattleEvent_ChangeHP(_fighter int, _hp int) *PU_BattleEvent_ChangeHP {
	return &PU_BattleEvent_ChangeHP{fighter: _fighter, hp: _hp}
}

func (e *PU_BattleEvent_ChangeHP) Execute() {
	g_game.battle.ChangeHP(e.fighter, e.hp)
}
