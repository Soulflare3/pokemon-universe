package main

type BattleConfiguration struct {
	gen		uint8
	mode 	uint8
	ids		[]int32
	clauses	uint32
}

func NewBattleConfiguration() *BattleConfiguration {
	return &BattleConfiguration{ ids: make([]int32, 2) }
}

func (b *BattleConfiguration) slot(spot int32, poke int32) int32 {
	return spot + (poke*2)
}

func (b *BattleConfiguration) spot(id int32) (ret int32) {
	ret = 1
	if b.ids[0] == id {
		ret = 0
	}
	return
}