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
	"math"
)

type PU_Fighter struct {
	pokemon *PU_Pokemon
	player  *PU_Player

	name     string
	pokename string

	id     int
	pokeid int
	level  int
	team   int
	hp     int
}

func NewFighter(_id int) *PU_Fighter {
	return &PU_Fighter{id: _id}
}

func (f *PU_Fighter) SetPokemon(_name string, _pokeid int, _level int, _hp int) {
	f.pokename = _name
	f.pokeid = _pokeid
	f.level = _level
	f.hp = _hp
}

func (f *PU_Fighter) SetSelf() {
	f.player = g_game.self
}

func (f *PU_Fighter) IsSelf() bool {
	return f.player == g_game.self
}

func (f *PU_Fighter) IsPlayer() bool {
	return f.player != nil
}

func (f *PU_Fighter) GetPokeName() string {
	if f.pokemon != nil {
		return f.pokemon.name
	}
	return f.pokename
}

func (f *PU_Fighter) GetPokeID() int {
	if f.pokemon != nil {
		return int(f.pokemon.id)
	}
	return f.pokeid
}

func (f *PU_Fighter) GetLevel() int {
	if f.pokemon != nil {
		return int(f.pokemon.level)
	}
	return f.level
}

func (f *PU_Fighter) GetHPPerc() int {
	if f.pokemon != nil {
		return int(math.Floor((float64(f.pokemon.hp) / float64(f.pokemon.hpmax)) * 100.0))
	}
	return f.hp
}

func (f *PU_Fighter) GetHP() int {
	if f.pokemon != nil {
		return int(f.pokemon.hp)
	}
	return f.hp
}

func (f *PU_Fighter) GetExp() int {
	if f.pokemon != nil {
		return int(f.pokemon.expPerc)
	}
	return 0
}

func (f *PU_Fighter) GetHPMax() int {
	if f.pokemon != nil {
		return int(f.pokemon.hpmax)
	}
	return 0
}

func (f *PU_Fighter) SetHP(_hp int) {
	if f.pokemon != nil {
		f.pokemon.hp = int16(_hp)
	} else {
		f.hp = _hp
	}
}

func (f *PU_Fighter) SetLevel(_level int) {
	if f.pokemon != nil {
		f.pokemon.level = int16(_level)
	} else {
		f.level = _level
	}
}

func (f *PU_Fighter) SetExp(_exp int) {
	if f.pokemon != nil {
		f.pokemon.expPerc = int16(_exp)
	}
}
