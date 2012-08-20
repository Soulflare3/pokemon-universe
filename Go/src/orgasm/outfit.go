package main

import (
	pul "pulogic"
)

type Outfit struct {
	data []int
}

func NewOutfit() Outfit {
	return Outfit{data: make([]int, 6)}
}

func NewOutfitExt(_head, _nek, _upper, _lower, _feet int) Outfit {
	outfit := NewOutfit()
	outfit.data[pul.OUTFIT_HEAD] = _head
	outfit.data[pul.OUTFIT_NEK] = _nek
	outfit.data[pul.OUTFIT_UPPER] = _upper
	outfit.data[pul.OUTFIT_LOWER] = _lower
	outfit.data[pul.OUTFIT_FEET] = _feet

	return outfit
}