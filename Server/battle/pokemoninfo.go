package main

import "fmt"

type PokemonInfo struct {
	Names	map[uint32]string
}

func NewPokemonInfo() *PokemonInfo {
	info := PokemonInfo{ Names: make(map[uint32]string) }
	info.init()
	return &info
}

func (p *PokemonInfo) init() {
	p.Names[NewPokemonUniqueIdFromNum(3,0).GetRef()] = "Venusaur"
	p.Names[NewPokemonUniqueIdFromNum(16,0).GetRef()] = "Pidgey"
}

func (p *PokemonInfo) GetPokemonName(_uniqueNumber PokemonUniqueId) string {
	value, _ := p.Names[_uniqueNumber.GetRef()]

	fmt.Printf("PokemonInfo - Number %d,%d (Ref %d)\n", _uniqueNumber.pokenum, _uniqueNumber.subnum, _uniqueNumber.GetRef())
	
	return value
}