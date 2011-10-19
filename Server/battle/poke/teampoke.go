package main

import (
	pnet "network"
)

type TeamPoke struct {
	UID *UniqueId
	Nick string
	Item int
	Ability int
	Nature int
	Gender int
	Gen int
	Shiny bool
	Happiness int
	Level int
	
	Moves []int
	DVs []int
	EVs []int
}

func NewTeamPokeFromPacket(_packet *pnet.QTPacket) *TeamPoke {
	teamPoke := TeamPoke{}
	teamPoke.UID = NewUniqueIdFromPacket(_packet)
	teamPoke.Nick = _packet.ReadString()
	teamPoke.Item = (int)_packet.ReadUint16()
	teamPoke.Ability = (int)_packet.ReadUint16()
	teamPoke.Nature = (int)_packet.ReadByte()
	teamPoke.Gender = (int)_packet.ReadByte()
	// teamPoke.Gen = (int)_packet.ReadByte()
	teamPoke.Shiny = _packet.ReadBool()
	teamPoke.Happiness = (int)_packet.ReadByte()
	teamPoke.Level = (int)_packet.ReadByte()
	
	teamPoke.Moves = make([]int, 4)
	for i := 0; i < 4; i++ {
		teamPoke.Moves[i] = (int)_packet.ReadUint32()
	}
	teamPoke.DVs = make([]int, 6)
	for i := 0; i < 6; i++ {
		teamPoke.DVs[i] = (int)_packet.ReadByte()
	}
	teamPoke.EVs = make([]int, 6)
	for i := 0; i < 6; i++ {
		teamPoke.EVs[i] = (int)_packet.ReadByte()
	}
}