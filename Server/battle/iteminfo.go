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

type ItemInfo struct {
	Names	map[uint16]string
}

func NewItemInfo() *ItemInfo {
	info := &ItemInfo{ Names: make(map[uint16]string) }
	info.init()
	return info
}

func (m *ItemInfo) init() {
	
}

func (m *ItemInfo) GetItemName(_itemId uint16) string {
	value, found := m.Names[_itemId]
	
	if !found {
		fmt.Printf("ERROR - Could not find item: %d\n", _itemId)
		return "Unknown Item"
	}
	
	return value
}