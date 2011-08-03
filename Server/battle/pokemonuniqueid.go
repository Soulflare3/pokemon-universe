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

type PokemonUniqueId struct {
	pokenum uint16
	subnum  uint8
}

func NewPokemonUniqueId() PokemonUniqueId {
	return PokemonUniqueId{pokenum: 0, subnum: 0}
}

func NewPokemonUniqueIdFromNum(_pokenum uint16, _subnum uint8) PokemonUniqueId {
	return PokemonUniqueId{pokenum: _pokenum, subnum: _subnum}
}

func NewPokemonUniqueIdFromRef(pokeRef uint32) PokemonUniqueId {
	unique := PokemonUniqueId{}
	unique.subnum = uint8(pokeRef >> 16)
	unique.pokenum = uint16(pokeRef & 0xFFFF)

	return unique
}

func (p PokemonUniqueId) GetRef() (ref uint32) {
	ref = uint32(p.pokenum) + (uint32(p.subnum) << 16)
	return
}
