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

type PU_BattleEvent_AllowControl struct {
	state int
}

func NewBattleEvent_AllowControl(_state int) *PU_BattleEvent_AllowControl {
	return &PU_BattleEvent_AllowControl{state : _state}
}

func (e *PU_BattleEvent_AllowControl) Execute() {
	switch e.state {
	case BATTLECONTROL_CHOOSEMOVE:
		g_game.panel.battleUI.Reset()
		
	case BATTLECONTROL_CHOOSEPOKEMON:
		g_game.panel.battleUI.Reset()
		g_game.panel.battleUI.moveState = BATTLEUI_CHOOSEPOKEMON
		g_game.panel.battleUI.OpenWindow(BATTLEWINDOW_POKEMON)
		
	case BATTLECONTROL_CHOOSEPOKEMON_ITEM:
		g_game.panel.battleUI.moveState = BATTLEUI_CHOOSEPOKEMON_ITEM
		g_game.panel.battleUI.OpenWindow(BATTLEWINDOW_POKEMON)
		
	case BATTLECONTROL_CHOOSEATTACK_ITEM:
		g_game.panel.battleUI.moveState = BATTLEUI_CHOOSEATTACK_ITEM
		g_game.panel.battleUI.OpenWindow(BATTLEWINDOW_POKEMON)
	}
}
