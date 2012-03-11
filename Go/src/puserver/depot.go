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

type Depot struct {
	MaxItems	int
	Items		ItemMap
}

func NewDepot(_maxItems int) *Depot {
	return &Depot { MaxItems: _maxItems,
					Items: make(ItemMap) }
}

func CreateItemIndex(_category, _slot int) int64 { 
	index := (int64)((_slot) | ((_category) << 8))
	
	return index 
}

func (d *Depot) SetMaxItems(_maxItems int) {
	d.MaxItems = _maxItems
}

func (d *Depot) GetMaxItems() int {
	return d.MaxItems
}

func (d *Depot) GetStorageCount() int {
	var totalCount int = 0
	for _, item := range(d.Items) {
		totalCount += item.GetCount()
	}
	
	return totalCount
}

func (d *Depot) GetItemCount() int {
	return len(d.Items)
}

func (d *Depot) GetItemList() ItemMap {
	return d.Items
}

func (d *Depot) AddItem(_itemId int64, _count int, _slot int) bool {
	if d.GetMaxItems() != 99999 && ((d.GetItemCount() + _count) > d.GetMaxItems()) {
		return false
	}
	
	item, found := g_game.Items.GetItemByItemId(_itemId)
	if !found {
		return false
	}
	
	// Clone item and set count
	newItem := item.Clone()
	newItem.SetCount(_count)
	newItem.Slot = _slot
	
	itemIndex := CreateItemIndex(item.CategoryId, _slot)
	
	// Check if this index is still available
	if d.GetItemFromIndex(itemIndex) != nil {
		slot := d.GetFreeSlotForCategory(item.CategoryId)
		if slot == -1 {
			return false
		}
		itemIndex = CreateItemIndex(item.CategoryId, slot)
		newItem.Slot = slot
	}
	
	// Add item to depot
	d.Items[itemIndex] = newItem
	
	return true
}

func (d *Depot) AddItemObject(_item *Item, _slot int) bool {
	if d.GetMaxItems() != 99999 && ((d.GetItemCount() + _item.GetCount()) > d.GetMaxItems()) {
		return false
	}
	
	itemIndex := CreateItemIndex(_item.CategoryId, _slot)
	_item.Slot = _slot
	
	// Check if this index is still available
	if d.GetItemFromIndex(itemIndex) != nil {
		slot := d.GetFreeSlotForCategory(_item.CategoryId)
		if slot == -1 {
			return false
		}
		itemIndex = CreateItemIndex(_item.CategoryId, slot)
		_item.Slot = slot
	}
	
	// Add item to depot
	d.Items[itemIndex] = _item
	
	return true
}

func (d *Depot) RemoveItem(_itemIndex int64) {
	delete(d.Items, _itemIndex)
}

func (d *Depot) UpdateItem(_slot, _category, _count int) bool {
	index := CreateItemIndex(_category, _slot)
	return d.UpdateItemByIndex(index, _count)
}

func (d *Depot) UpdateItemByIndex(_itemIndex int64, _count int) bool {
	if d.GetMaxItems() != 99999 && ((d.GetItemCount() + _count) > d.GetMaxItems()) {
		return false
	}
	
	item := d.GetItemFromIndex(_itemIndex)
	if item == nil {
		return false
	}
	
	item.SetCount(item.GetCount() + _count)
	
	if item.GetCount() <= 0 {
		d.RemoveItem(_itemIndex)
	}
	
	return true
}

func (d *Depot) GetItem(_category int, _slot int) *Item {
	index := CreateItemIndex(_category, _slot)
	return d.GetItemFromIndex(index)
}

func (d *Depot) GetItemFromIndex(_itemIndex int64) *Item {
	if item, found := d.Items[_itemIndex]; found {
		return item
	}
	return nil
}

func (d *Depot) GetFreeSlotForCategory(_category int) int {
	for i := 0; i < 54; i++ {
		if d.GetItem(_category, i) == nil {
			return i
		} 
	}
	
	return -1
}

func (d *Depot) SearchItem(_category int, _itemId int64) int {
	for i := 0; i < 54; i++ {
		item := d.GetItem(_category, i)
		if item != nil && item.Id == _itemId {
			return i
		} 
	}
	
	return -1
}

func (d *Depot) SwitchItem(_category, _oldSlot, _newSlot int) {
	indexOld := CreateItemIndex(_category, _oldSlot)
	indexNew := CreateItemIndex(_category, _newSlot)
	
	_, oldFound := d.Items[indexOld]
	_, newFound := d.Items[indexNew]
	
	if oldFound && newFound {
		d.Items[indexOld], d.Items[indexNew] = d.Items[indexNew], d.Items[indexOld]
		d.Items[indexOld].Slot = _newSlot
		d.Items[indexNew].Slot = _oldSlot
		
	} else if oldFound {
		d.Items[indexNew] = d.Items[indexOld]
		d.Items[indexNew].Slot = _newSlot
		delete(d.Items, indexOld)
		
	} else {
		d.Items[indexOld] = d.Items[indexNew]
		d.Items[indexOld].Slot = _oldSlot
		delete(d.Items, indexNew)		
		
	}
}