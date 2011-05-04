package main

import "container/vector"

type BaseBattleInfo struct {
	// name [0] = mine, name[1] = other //
	pInfo					[]*PlayerInfo
	sub						vector.Vector
	specialSprite			vector.Vector
	lastSeenSpecialSprite	vector.Vector
	
	time			[]uint16
	ticking			[]bool
	startingTime	[]int64
	
	mode			uint8
	numberOfSlots	int
	
	myself			int8
	opponent		int8
	
	gen				uint8
	
	pokemons		[][]*PokeBattle
	pokeAlive		vector.Vector
	
	statChanges		vector.Vector
}

func (b *BaseBattleInfo) Init(_me *PlayerInfo, _opp *PlayerInfo, _mode uint8, _myself int32, _opponent int32) {
	b.pInfo = make([]*PlayerInfo, 2)
	b.time = make([]uint16, 2)
	b.ticking = make([]bool, 2)
	b.startingTime = make([]int64, 2)
	b.mode = _mode
	b.myself = int8(_myself)
	b.opponent = int8(_opponent)
	
	b.pokemons = make([]([]*PokeBattle), 2)
	for i := 0; i < 2; i++ {
		b.pokemons[i] = make([]*PokeBattle, 6)
	}
							
	if _mode == ChallengeInfo_Doubles {
		b.numberOfSlots = 4
	} else if _mode == ChallengeInfo_Triples {
		b.numberOfSlots = 6
	} else {
		b.numberOfSlots = 2
	}
	
	for i := 0; i < b.numberOfSlots; i++ {
		b.sub.Push(false)
		b.pokeAlive.Push(false)
		b.specialSprite.Push(0)
		b.lastSeenSpecialSprite.Push(0)
		b.statChanges.Push(NewBattleDynamicInfo())
	}
	
	b.pInfo[_myself] = _me
	b.pInfo[_opponent] = _opp
	b.time[_myself] = 5 * 60
	b.time[_opponent] = 5 * 60
	b.ticking[_myself] = false
	b.ticking[_opponent] = false
}

func NewBaseBattleInfo() *BaseBattleInfo {
	return &BaseBattleInfo{ }
}

func NewBaseBattleInfoDefault(_me *PlayerInfo, _opp *PlayerInfo, _mode uint8) *BaseBattleInfo {
	base := NewBaseBattleInfo()
	base.Init(_me, _opp, _mode, 0, 1)
	return base
}

func (i *BaseBattleInfo) SwitchPoke(spot, poke int8) {
	i.pokemons[i.Player(spot)][i.SlotNum(poke)], i.pokemons[i.Player(spot)][i.SlotNum(spot)] = i.CurrentShallow(spot), i.pokemons[i.Player(spot)][i.SlotNum(poke)]
	i.pokeAlive[spot] = true;
}

func (i *BaseBattleInfo) CurrentShallow(spot int8) *PokeBattle {
	return i.pokemons[i.Player(spot)][i.SlotNum(spot)]
}

func (i *BaseBattleInfo) SetCurrentShallow(spot int8, shallow *PokeBattle) {
	i.pokemons[i.Player(spot)][i.SlotNum(spot)] = shallow
}

func (i *BaseBattleInfo) Number( _spot int8) int8 {
	return (_spot / 2)
}

func (i *BaseBattleInfo) Name(_x int8) string {
	return i.pInfo[_x].team.name
}

func (i *BaseBattleInfo) Player(_slot int8) int8 {
	return _slot % 2
}

func (i *BaseBattleInfo) Slot(_player int8, _poke int8) int8 {
	return (_player + (_poke * 2))
}

func (i *BaseBattleInfo) SlotNum(_slot int8) int8 {
	return _slot / 2
}

func (i *BaseBattleInfo) IsOut(_poke int8) bool {
	return (_poke < int8(i.numberOfSlots / 2))
}

func (i *BaseBattleInfo) Multiples() bool {
	return ((i.mode == ChallengeInfo_Doubles) || (i.mode == ChallengeInfo_Triples))
}