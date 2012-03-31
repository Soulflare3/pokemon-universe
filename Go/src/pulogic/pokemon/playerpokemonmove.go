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

import "fmt"

type PlayerPokemonMove struct {
	Move *Move
	
	DbId		int64
	CurrentPP 	int
}

func NewPlayerPokemonMove(_dbid int64, _moveId int, _currentPP int) *PlayerPokemonMove {
	pMove := PlayerPokemonMove{ DbId: _dbid, CurrentPP: _currentPP }
	pMove.Move = GetInstance().GetMoveById(_moveId)
	
	if pMove.Move == nil {
		fmt.Printf("MOVE IS NIL - %d\n", _moveId) 
	}
	
	return &pMove
}