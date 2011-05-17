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

import "fmt"

type AbilityInfo struct {
	Names	map[uint16]string
}

func NewAbilityInfo() *AbilityInfo {
	info := &AbilityInfo{ Names: make(map[uint16]string) }
	info.init()
	return info
}

func (m *AbilityInfo) init() {
	
}

func (m *AbilityInfo) GetAbilityName(_abilityId uint16) string {
	value, found := m.Names[_abilityId]
	
	if !found {
		fmt.Printf("ERROR - Could not find ability: %d\n", _abilityId)
		return "Unknown Item"
	}
	
	return value
}