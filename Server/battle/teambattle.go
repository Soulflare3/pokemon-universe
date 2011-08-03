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

type TeamBattle struct {
	name string
	info string
	gen  uint32

	pokemons []*PokeBattle
	indexes  []int32
}

func NewTeamBattle() *TeamBattle {
	team := TeamBattle{}
	team.pokemons = make([]*PokeBattle, 6)
	team.indexes = make([]int32, 6)

	return &team
}

func (b *TeamBattle) Poke(i int8) *PokeBattle {
	return b.pokemons[b.indexes[i]]
}

func (b *TeamBattle) SetPoke(_index int, _poke *PokeBattle) {
	b.pokemons[_index] = _poke
}

func (b *TeamBattle) SwitchPokemon(_poke1 int8, _poke2 int8) {
	b.indexes[_poke1], b.indexes[_poke2] = b.indexes[_poke2], b.indexes[_poke1]
}
