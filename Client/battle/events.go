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

const (
	BATTLEEVENT_STOPBATTLE = 999
	BATTLEEVENT_SLEEP = 0
	BATTLEEVENT_TEXTID = 1
	BATTLEEVENT_TEXT = 2
	BATTLEEVENT_CHANGEHP = 3
	BATTLEEVENT_ANIMATION = 4
	BATTLEEVENT_ALLOWCONTROL = 5
	BATTLEEVENT_CHANGEPOKEMON_SELF = 6
	BATTLEEVENT_CHANGEPOKEMON = 7
	BATTLEEVENT_CHANGESELECTION = 8
	BATTLEEVENT_CHANGEPP = 9
	BATTLEEVENT_CHANGESTATUS = 10
	BATTLEEVENT_CHANGELEVELSELF = 11
	BATTLEEVENT_CHANGELEVEL = 12
	BATTLEEVENT_CHANGEATTACK = 13
	BATTLEEVENT_CHANGESCREEN = 14
	BATTLEEVENT_DIALOGUE = 15
	BATTLEEVENT_REMOVEPLAYER = 16
	BATTLEEVENT_CHANGEEXP = 17
)

const (
	BATTLECONTROL_NONE = 0
	BATTLECONTROL_CHOOSEMOVE = 1
	BATTLECONTROL_CHOOSEPOKEMON = 2
	BATTLECONTROL_CHOOSEPOKEMON_ITEM = 3
	BATTLECONTROL_CHOOSEATTACK_ITEM = 4
)

const (
	BATTLESELECT_MOVE = 0
	BATTLESELECT_POKEMON = 1
)

type IBattleEvent interface {
	Execute()
}

