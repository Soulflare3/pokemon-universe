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
	POKESTAT_ATTACK = iota
	POKESTAT_DEFENSE
	POKESTAT_SPECIALATTACK
	POKESTAT_SPECIALDEFENSE
	POKESTAT_SPEED
)

type PU_Pokemon struct {
	uid   uint32
	id    int16
	level int16
	hp    int16
	hpmax int16
	sex   int16

	expPerc    int16
	expCurrent int32
	expTnl     int32

	name   string
	flavor string

	type1 string
	type2 string

	stats [5]int16

	attacks [4]*PU_Attack
}

func NewPokemon() *PU_Pokemon {
	return &PU_Pokemon{}
}

func (p *PU_Pokemon) SetAttack(_num int, _name string, _description string, _type string, _pp uint16, _ppmax uint16, _power uint16, _accuracy uint16, _category string, _target string, _contact string) {
	if p.attacks[_num] == nil {
		p.attacks[_num] = NewAttack()
	}

	attack := p.attacks[_num]
	attack.name = _name
	attack.description = _description
	attack.poketype = _type
	attack.pp = _pp
	attack.ppmax = _ppmax
	attack.power = _power
	attack.accuracy = _accuracy
	attack.category = _category
	attack.target = _target
	attack.contact = _contact
}

func (p *PU_Pokemon) GetAttack(_num int) *PU_Attack {
	return p.attacks[_num]
}

func GetTypeByName(_type string) uint16 {
	switch _type {
	case "ground":
		return 100

	case "water":
		return 101

	case "ghost":
		return 102

	case "bug":
		return 103

	case "fighting":
		return 104

	case "psychic":
		return 105

	case "grass":
		return 106

	case "dark":
		return 107

	case "normal":
		return 108

	case "poison":
		return 109

	case "electric":
		return 110

	case "unknown":
		return 111

	case "steel":
		return 112

	case "rock":
		return 113

	case "dragon":
		return 114

	case "flying":
		return 115

	case "fire":
		return 116

	case "ice":
		return 117
	}
	return 111
}
