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

type PokemonInfo struct {
	Names	map[uint32]string
}

func NewPokemonInfo() *PokemonInfo {
	info := &PokemonInfo{ Names: make(map[uint32]string) }
	info.init()
	return info
}

func (p *PokemonInfo) init() {
	p.Names[NewPokemonUniqueIdFromNum(3,0).GetRef()] = "Venusaur"
	p.Names[NewPokemonUniqueIdFromNum(16,0).GetRef()] = "Pidgey"
}

func (p *PokemonInfo) GetPokemonName(_uniqueNumber PokemonUniqueId) (value string) {
	value, _ = p.Names[_uniqueNumber.GetRef()]	
	return
}