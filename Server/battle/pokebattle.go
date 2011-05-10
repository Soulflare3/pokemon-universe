package main

import (
	pnet "network"
	"fmt"
)

type PokeBattle struct {
	nick			string
	num				PokemonUniqueId
	
	dvs				[]uint8
	evs				[]uint8
	
	normal_stats	[]uint16
	moves			[]*BattleMove
	
	lifePoints		uint16
	totalLifePoints uint16
	lifePercent		uint8
	
	item			uint16
	ability			uint16
	fullStatus		uint32
	statusCount		int8
	oriStatusCount	int8
	gender			uint8
	level			uint8
	nature			uint8
	happiness		uint8
	shiny			bool
}

func NewPokeBattle() *PokeBattle {
	return &PokeBattle{ dvs: make([]uint8, 6),
					evs: make([]uint8, 6),
					normal_stats: make([]uint16, 5),
					moves: make([]*BattleMove, 4) }
}

func NewPokeBattleFromPacket(_packet *pnet.QTPacket) *PokeBattle {
	poke := NewPokeBattle()

	pokeNum := _packet.ReadUint16()
	// subNum := _packet.ReadUint8()
	// derp := _packet.ReadUint8()
	
	poke.num = NewPokemonUniqueIdFromRef(uint32(pokeNum))
//	if poke.num.pokenum > 0 && poke.num.subnum >= 0 {
		poke.nick = _packet.ReadString()
		poke.lifePercent = _packet.ReadUint8()
		poke.fullStatus = _packet.ReadUint32()
		poke.gender = _packet.ReadUint8()
		poke.shiny = (_packet.ReadUint8() == 1) 
		poke.level = _packet.ReadUint8()
//	}
	
	fmt.Printf("POKEBATTLE - pokeNum: %d | subNum %d | %d | %d\n", pokeNum, 0, 0, poke.level)
	
	return poke	
}

func (p *PokeBattle) Init() {
	if p.totalLifePoints > 0 {
		p.lifePercent = uint8((p.lifePoints * 100) / p.totalLifePoints)
		if p.lifePercent == 0 && p.lifePoints > 0 {
			p.lifePercent = 1
		}
	}
}

func (p *PokeBattle) ChangeStatus(_status uint8) {
	// Clear past status
	p.fullStatus = p.fullStatus & ^(uint32(1 << PokemonStatus_Koed) | 0x3F)	
	// Add new status
	p.fullStatus = p.fullStatus | (1 << _status)
}