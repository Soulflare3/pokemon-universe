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
package pokemon

type PokemonParty struct {
	Party	[]*PlayerPokemon
}

func NewPokemonParty() *PokemonParty {
	return &PokemonParty{ Party: make([]*PlayerPokemon, 6) }
}

func (p *PokemonParty) Add(_pokemon *PlayerPokemon) {
	for i := 0; i < 6; i++ {
		if p.Party[i] == nil {
			p.Party[i] = _pokemon
			break
		}
	}
}

func (p *PokemonParty) AddSlot(_pokemon *PlayerPokemon, _slot int) {
	if p.Party[_slot] == nil {
		p.Party[_slot] = _pokemon
	}
}

func (p *PokemonParty) GetFromSlot(_slot int) *PlayerPokemon {
	pokemon := p.Party[_slot]
	return pokemon
}

func (p *PokemonParty) HealParty() {
	for i := 0; i < 6; i++ {
		if p.Party[i] != nil {
			p.Party[i].DamagedHp = 0
		}
	}
}