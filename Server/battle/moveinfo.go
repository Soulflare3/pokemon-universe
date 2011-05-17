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
	"fmt"
)

type MoveMessagesPart []string

type Gen struct {
	gen		uint8
	
	pp		map[uint16]uint8
}

func NewGen(_gen uint8) *Gen {
	gen := &Gen { gen: _gen,
				  pp: make(map[uint16]uint8) }
	return gen
}

type MoveInfo struct {
	Names	map[uint16]string
	MoveMessages map[uint16]MoveMessagesPart
	
	Generations	map[uint8]*Gen
}

func NewMoveInfo() *MoveInfo {
	info := &MoveInfo{ Names: make(map[uint16]string),
				MoveMessages: make(map[uint16]MoveMessagesPart),
				Generations: make(map[uint8]*Gen) }
	info.init()
	return info
}

func (m *MoveInfo) init() {
	m.Names[16] = "Gust"
	m.Names[338] = "Frenzy Plant"
}

func (m *MoveInfo) GetMoveName(_moveNumber uint16) string {
	value, found := m.Names[_moveNumber]
	
	if !found {
		fmt.Printf("ERROR - Could not find move: %d\n", _moveNumber)
		return "Unknown Move"
	}
	
	return value
}

func (m *MoveInfo) GetMoveMessage(_move uint16, _part uint8) string {
	value, found := m.MoveMessages[_move]
	
	if !found {
		fmt.Printf("ERROR - Could not find move message %d part %d\n", _move, _part)
	}
	
	return value[_part]
}

func (m *MoveInfo) GetGeneration(_gen uint8) *Gen {
	return m.Generations[_gen-1]
}

func (m *MoveInfo) PP(_gen uint8, _move uint16) uint8 {
	return m.GetGeneration(_gen).pp[_move]
}