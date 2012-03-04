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

const (
	MOVEMETHOD_LEVELUP		int = 1
	MOVEMETHOD_EGG				= 2
	MOVEMETHOD_TUTOR		    = 3
	MOVEMETHOD_MACHINE			= 4
	MOVEMETHOD_LIGHTBALLEGG		= 6
	MOVEMETHOD_FORMCHANGE		= 10
)

type PokemonMove struct {
	Pokemon				*Pokemon
	VersionGroup		int
	Move				*Move
	PokemonMoveMethod	int
	Level				int
	Order				int
}

func NewPokemonMove() *PokemonMove {
	return &PokemonMove{}
}