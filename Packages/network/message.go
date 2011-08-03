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
package network

//Main message holding pointers to all possible datastructures
type Message struct {
	Header         int
	Login          *Data_Login
	LoginStatus    *Data_LoginStatus
	PlayerData     *Data_PlayerData
	AddCreature    *Data_AddCreature
	RemoveCreature *Data_RemoveCreature
	Tiles          *Data_Tiles
	Walk           *Data_Walk
	CreatureWalk   *Data_CreatureWalk
	Turn           *Data_Turn
	CreatureTurn   *Data_CreatureTurn
	Warp           *Data_Warp
}

func NewMessage(_header int) *Message {
	return &Message{Header: _header}
}
