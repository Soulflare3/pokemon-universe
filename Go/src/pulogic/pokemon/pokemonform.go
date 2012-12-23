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

import "pulogic/models"

type PokemonForm struct {
	Id				int
	Identifier		string
	IsDefault		bool
	IsBattleOnly	bool
	//Order			int
}

func NewPokemonForm() *PokemonForm {
	return &PokemonForm{}
}

func NewPokemonFormFromEntity(_entity models.PokemonForms) *PokemonForm {
	form := NewPokemonForm()
	form.Id = _entity.Id
	form.Identifier = _entity.FormIdentifier
	form.IsBattleOnly = _entity.IsBattleOnly
	form.IsDefault = _entity.IsDefault
	//form.Order = _entity.Order
	
	return form
} 