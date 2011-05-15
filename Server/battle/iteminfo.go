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