package main

type PokemonUniqueId struct {
	pokenum	uint16
	subnum	uint8
}

func NewPokemonUniqueId() *PokemonUniqueId {
	return &PokemonUniqueId{pokenum: 0, subnum: 0}
}

func NewPokemonUniqueIdFromNum(_pokenum uint16, _subnum uint8) *PokemonUniqueId {
	unique := PokemonUniqueId{pokenum: _pokenum, subnum: _subnum }
	return &unique
}

func NewPokemonUniqueIdFromRef(pokeRef uint32) *PokemonUniqueId {
	unique := PokemonUniqueId{ }
	unique.subnum = uint8(pokeRef >> 16)
	unique.pokenum = uint16(pokeRef & 0xFFFF)
	return &unique
}