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
	PlayerInfo_LoggedIn = 1
	PlayerInfo_Battling = 2
	PlayerInfo_Away = 4
)

type PlayerInfo struct {
	id		int32
	team	*BasicInfo
	auth	int8
	flags	uint8
	rating	int16
	pokes	[]PokemonUniqueId
	avatar	uint16
	tier	string
	color	uint32
	gen		uint8
}

func NewPlayerInfo() *PlayerInfo {
	return &PlayerInfo{ pokes: make([]PokemonUniqueId, 6) }
}