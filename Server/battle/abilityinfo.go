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