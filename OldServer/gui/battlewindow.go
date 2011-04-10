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
	BATTLEWINDOW_NONE = -1
	BATTLEWINDOW_ATTACK = 0
	BATTLEWINDOW_ITEMS = 1
	BATTLEWINDOW_POKEMON = 2
)

type IBattleWindow interface {
	SetValue(_value int)
	GetValue() int
	GetType() int
	Close()
}

type PU_BattleWindow struct {
	windowtype int
	value int
}

func (g *PU_BattleWindow) SetValue(_value int) {
	g.value = _value
}

func (g *PU_BattleWindow) GetValue() int {
	return g.value
}

func (g *PU_BattleWindow) GetType() int {
	return g.windowtype
}
