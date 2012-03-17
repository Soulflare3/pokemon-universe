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
	puh "puhelper"
)

type Item struct {
	DbId			int64
	Id				int64
	Identifier		string
	CategoryId		int
	Cost			int
	FlingPower		int
	FlingEffectId	int
	MaxStack		int
	
	CanBeSold		bool
	CanBeTraded		bool
	
	Count			int
	Slot			int
}

func (i *Item) Clone() *Item {
	newItem := &Item{DbId: i.DbId,
					 Id: i.Id,
					 Identifier: i.Identifier,
					 CategoryId: i.CategoryId,
					 Cost: i.Cost,
			 		 FlingPower: i.FlingPower,
					 FlingEffectId: i.FlingEffectId,
					 MaxStack: i.MaxStack,
			
					 CanBeSold: i.CanBeSold,
					 CanBeTraded: i.CanBeTraded,
					 Count: i.Count }
					 
	return newItem
}

func (i *Item) GetCount() int {
	return i.Count
}

func (i *Item) SetCount(_count int) {
	i.Count = _count
}

/************************
		ITEM STORE
*************************/
type ItemMap map[int64]*Item
type ItemStore struct {
	Items ItemMap
}

func NewItemStore() *ItemStore {
	return &ItemStore { Items: make(ItemMap) }
}

func (s *ItemStore) Load() bool {
	var query string = "SELECT id, identifier, category_id, cost, fling_power, fling_effect_id FROM items"
	result, err := puh.DBQuerySelect(query)
	if err != nil {
		return false
	}
	
	defer puh.DBFree()
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		itemId := puh.DBGetInt64(row[0])
		identifier := puh.DBGetString(row[1])
		category := puh.DBGetInt(row[2])
		cost := puh.DBGetInt(row[3])
		flingPower := puh.DBGetInt(row[4])
		flingEffectId := puh.DBGetInt(row[5])
		
		item := &Item { Id: itemId,
						Identifier: identifier,
						CategoryId: category,
						Cost: cost,
						FlingPower: flingPower,
						FlingEffectId: flingEffectId,
						MaxStack: 0,
						CanBeSold: true,
						CanBeTraded: true }
		
		// Add item to store
		s.Items[itemId] = item
	}
	
	return true;
}

func (s *ItemStore) GetItemByItemId(_itemId int64) (*Item, bool) {
	value, ok := s.Items[_itemId]
	return value, ok
}