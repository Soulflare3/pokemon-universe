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

func NewTeamPoke() *TeamPoke {
	teamPoke := TeamPoke{}
	teamPoke.UID = NewUniqueId()
	teamPoke.Nick = "DERP"
	teamPoke.Item = 71
	teamPoke.Ability = 98
	teamPoke.Nature = 0
	teamPoke.Gender = 1
	teamPoke.Gen = 5
	teamPoke.Shiny = true
	teamPoke.Happiness = 127
	teamPoke.Level = 100
	
	teamPoke.Moves = make([]int, 4)
	teamPoke.Moves[0] = 118
	teamPoke.Moves[1] = 227
	teamPoke.Moves[2] = 150
	teamPoke.Moves[3] = 271
	
	teamPoke.DVs = make([]int, 6)
	for i := 0; i < 6; i++ {
		teamPoke.DVs[i] = 31
	}
	teamPoke.EVs = make([]int, 6)
	for i := 0; i < 6; i++ {
		teamPoke.EVs[i] = 10
	}	
	
	return &teamPoke
}

func NewTeamPokeFromPacket(_packet *pnet.QTPacket) *TeamPoke {
	teamPoke := TeamPoke{}
	teamPoke.UID = NewUniqueIdFromPacket(_packet)
	teamPoke.Nick = _packet.ReadString()
	teamPoke.Item = int(_packet.ReadUint16())
	teamPoke.Ability = int(_packet.ReadUint16())
	teamPoke.Nature = int(_packet.ReadUint8())
	teamPoke.Gender = int(_packet.ReadUint8())
	// teamPoke.Gen = (int)_packet.ReadByte()
	teamPoke.Shiny = _packet.ReadBool()
	teamPoke.Happiness = int(_packet.ReadUint8())
	teamPoke.Level = int(_packet.ReadUint8())
	
	teamPoke.Moves = make([]int, 4)
	for i := 0; i < 4; i++ {
		teamPoke.Moves[i] = int(_packet.ReadUint32())
	}
	teamPoke.DVs = make([]int, 6)
	for i := 0; i < 6; i++ {
		teamPoke.DVs[i] = int(_packet.ReadUint8())
	}
	teamPoke.EVs = make([]int, 6)
	for i := 0; i < 6; i++ {
		teamPoke.EVs[i] = int(_packet.ReadUint8())
	}
	
	return &teamPoke
}

func (t *TeamPoke) WritePacket() pnet.IPacket {
	packet := pnet.NewQTPacket()
	packet.AddBuffer(t.UID.WritePacket().GetBufferSlice())
	packet.AddString(t.Nick)
	packet.AddUint16(uint16(t.Item))
	packet.AddUint16(uint16(t.Ability))
	packet.AddUint8(uint8(t.Nature))
	packet.AddUint8(uint8(t.Gender))
	// packet.AddUint8(uint8(t.Gen)) // XXX Gen would go here
	packet.AddBool(t.Shiny)
	packet.AddUint8(uint8(t.Happiness))
	packet.AddUint8(uint8(t.Level))
	
	for i := 0; i < 4; i++ {
		packet.AddUint32(uint32(t.Moves[i]))
	}
	
	for i := 0; i < 6; i++ {
		packet.AddUint8(uint8(t.DVs[i]))
	}
	
	for i := 0; i < 6; i++ {
		packet.AddUint8(uint8(t.EVs[i]))
	}
	
	return packet
}