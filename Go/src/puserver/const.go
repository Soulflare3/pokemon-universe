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

import (
	"time"
	"putools/log"
	pos "putools/pos"	
)

var (
	CLIENT_VIEWPORT        pos.Position = pos.Position{28, 22, 0}
	CLIENT_VIEWPORT_CENTER pos.Position = pos.Position{14, 11, 0}
)

const (
	RET_NOERROR int = iota
	RET_NOTPOSSIBLE
	RET_PLAYERISTELEPORTED
	RET_YOUAREEXHAUSTED
	RET_PLAYERNOTFOUND
)

const NANOSECONDS_TO_MILLISECONDS = 0.000001

const (
	PlayerFlag_CannotUseCombat uint64	= iota	// 2^0 = 1
	PlayerFlag_CanAlwaysLogin					// 2^1 = 2
	PlayerFlag_CanBroadcast						// 2^2 = 4
	PlayerFlag_CannotBeSeen						// 2^3 = 8
	PlayerFlag_CanSenseInvisibility				// 2^4 = 16
	// Add new flags here
	PlayerFlag_LastFlag
)

const (
	MOVEMENT_WALK = TILEBLOCK_WALK
	MOVEMENT_SURF = TILEBLOCK_SURF
)

const (
	DIR_NULL  = 0
	DIR_SOUTH = 1
	DIR_WEST  = 2
	DIR_NORTH = 3
	DIR_EAST  = 4
)

const (
	CTYPE_CREATURE = 0
	CTYPE_NPC      = 1
	CTYPE_PLAYER   = 2
)

func PUSYS_TIME() int64 {
	timeNano := float64(time.Now().UnixNano())
	return int64(timeNano * NANOSECONDS_TO_MILLISECONDS)
}

func FuncAfter(_d string,  _f func()) (*time.Timer, error) {
	duration, err := time.ParseDuration(_d)
	if err != nil {
		logger.Printf("FuncAfter - Failed to parse parameter (_d) - %v\n", err.Error())
		return nil, err
	}
	
	timer := time.AfterFunc(duration, _f)
	
	return timer, nil
}