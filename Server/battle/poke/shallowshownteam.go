package main

import (
	pnet "network"
)

type ShallowShownTeam struct {
	Pokes []*ShallowShownPoke
}

func NewShallowShownTeamFromPacket(_packet *pnet.QTPacket) *ShallowShownTeam {
	shownTeam := ShallowShownTeam { Pokes: make([]*ShallowShownPoke, 6)
	for i := 0; i < 6; i++ {
		shownTeam.Pokes[i] = NewShallowShownPokeFromPacket(_packet)
	}
}

func (s *ShallowShownTeam) Poke(_index int) *ShallowShownPoke {
	return s.Pokes[_index]
}