package main

import (
	"os"
	
	pnet "network"
)

type Team struct {
	Gen int
	Pokes []TeamPoke
}

func NewTeam() *Team {
	team := Team { Pokes: make([]TeamPoke, 6) }
	for i := 0; i < 6; i++ {
		team.Pokes[i] = NewTeamPoke()
	}
}

func (t *Team) WritePacket() (pnet.IPacket, os.Error) {
	packet := NewQTPacket()
	packet.AddUint8(uint8(t.Gen))
	for i := 0; i < 6; i++ {
		pokePacket, _ := team.Pokes[i].WritePacket()
		packet.AddBuffer(pokePacket.GetBuffer(), pokePacket.GetMsgSize())
	}
	
	return packet, nil
}