package main

import (
	pnet "network"
)

type Team struct {
	Gen int
	Pokes []*TeamPoke
}

func NewTeam() *Team {
	team := Team { Pokes: make([]*TeamPoke, 6) }
	for i := 0; i < 6; i++ {
		team.Pokes[i] = NewTeamPoke()
	}
}

func NewTeamFromPacket(_packet *pnet.QTPacket) *Team {
	team := Team { Pokes: make([]*TeamPoke, 6) }
	team.Gen = int(_packet.ReadUint8())
	for i := 0; i < 6; i++ {
		team.Pokes[i] = NewTeamPokeFromPacket(_packet)
	}
	return &team
}

func (t *Team) WritePacket() pnet.IPacket {
	packet := pnet.NewQTPacket()
	packet.AddUint8(uint8(t.Gen))
	for i := 0; i < 6; i++ {
		packet.AddBuffer(t.Pokes[i].WritePacket().GetBufferSlice())
	}
	
	return packet
}