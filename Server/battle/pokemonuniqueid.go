package main

type PokemonUniqueId struct {
	pokenum	uint16
	subnum	uint8
}

func NewPokemonUniqueId() PokemonUniqueId {
	return PokemonUniqueId{pokenum: 0, subnum: 0}
}

func NewPokemonUniqueIdFromNum(_pokenum uint16, _subnum uint8) PokemonUniqueId {
	return PokemonUniqueId{pokenum: _pokenum, subnum: _subnum }
}

func NewPokemonUniqueIdFromRef(pokeRef uint32) PokemonUniqueId {
	unique := PokemonUniqueId{ }
	unique.subnum = uint8(pokeRef >> 16)
	unique.pokenum = uint16(pokeRef & 0xFFFF)

	return unique
}

func (p PokemonUniqueId) GetRef() (ref uint32) {
	ref = uint32(p.pokenum) + (uint32(p.subnum) << 16)
	return
}