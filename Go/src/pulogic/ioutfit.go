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
package pulogic

type OutfitPart int
const (
	OUTFIT_BASE OutfitPart = iota
	OUTFIT_UPPER
	OUTFIT_NEK
	OUTFIT_HEAD
	OUTFIT_FEET
	OUTFIT_LOWER
)

type IOutfit interface {
	GetData() []int
	SetOutfitKey(_part OutfitPart, _key int)
	GetOutfitKey(_part OutfitPart) int
	SetOutfitStyle(_part OutfitPart, _style int)
	GetOutfitStyle(_part OutfitPart) int
	SetOutfitColour(_part OutfitPart, _colour int)
	GetOutfitColour(_part OutfitPart) int
}