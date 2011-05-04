package main

type TeamBattle struct {
	name	string
	info	string
	gen		uint32
	
	pokemons	[]*PokeBattle
	indexes		[]int32
}

func NewTeamBattle() *TeamBattle {
	team := TeamBattle{}
	team.pokemons = make([]*PokeBattle, 6)
	team.indexes = make([]int32, 6)
	
	return &team
}

func (b *TeamBattle) Poke(i int8) *PokeBattle {
	return b.pokemons[b.indexes[i]]
}

func (b *TeamBattle) SetPoke(_index int, _poke *PokeBattle) {
	b.pokemons[_index] = _poke
}

func (b *TeamBattle) SwitchPokemon(_poke1 int8, _poke2 int8) {
	b.indexes[_poke1], b.indexes[_poke2] = b.indexes[_poke2], b.indexes[_poke1]
}

