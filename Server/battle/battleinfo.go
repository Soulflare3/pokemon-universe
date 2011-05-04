package main

import (
	"container/vector"
)

type BattleInfo struct {
	BaseBattleInfo

	myteam		*TeamBattle
	possible	bool
	sent		bool
	
	choices		vector.Vector
	choice		vector.Vector
	available	vector.Vector
	done		vector.Vector
	
	currentSlot	int8
	
	mystats		vector.Vector
	tempPoke	vector.Vector
	
	lastMove	[]int
}

func NewBattleInfo(_team *TeamBattle, _me, _opp *PlayerInfo, _mode uint8, _my, _op int32) *BattleInfo {
	base := &BattleInfo{lastMove: make([]int, 6) }
	base.Init(_me, _opp, _mode, _my, _op)
	
	base.possible = false
	base.myteam = _team
	base.sent = true
	
	base.currentSlot = base.SlotNum(base.myself)
	
	for i := 0; i < base.numberOfSlots/2; i++ {
		base.choices.Push(NewBattleChoices())
		base.choice.Push(NewBattleChoice())
		base.available.Push(false)
		base.done.Push(false)
		
		base.mystats.Push(NewBattleStats())
		base.tempPoke.Push(NewPokeBattle())
	}
	
	for i := 0; i < 6; i++ {
		base.pokemons[base.myself][i] = _team.Poke(int8(i))
	}
	
	return base
}

func (b *BattleInfo) GetTempPoke(_spot int8) *PokeBattle {
	return b.tempPoke.At(int(_spot)).(*PokeBattle)
}

func (b *BattleInfo) SetTempPoke(_spot int8, _poke *PokeBattle) {
	b.tempPoke.Set(int(_spot), _poke)
}

func (b *BattleInfo) SwitchPokeExt(_spot int8, _poke int8, _own bool) {
	b.SwitchPoke(_spot, _poke)
	if _own {
		b.myteam.SwitchPokemon(b.Number(_spot), _poke)
		b.SetCurrentShallow(_spot, b.myteam.Poke(b.Number(_spot)))
		b.SetTempPoke(_spot, b.myteam.Poke(b.Number(_spot)))
	}
}