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

type PU_BattleEvent_ChangePP struct {
	pokemon int 
	attack int 
	value int 
}

func NewBattleEvent_ChangePP(_pokemon int, _attack int, _value int) *PU_BattleEvent_ChangePP {
	return &PU_BattleEvent_ChangePP{pokemon : _pokemon, attack : _attack, value : _value}
}

func (e *PU_BattleEvent_ChangePP) Execute() {

}
