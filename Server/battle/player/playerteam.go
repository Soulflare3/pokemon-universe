package main

import (
	pnet "network"
)

type PlayerTeam struct {
	Nick string
	Info string
	LoseMessage string
	WinMessage string
	DefaultTier string
	Tier string
	Team *Team
	
	avatar int
}

func NewPlayerTeamFromPacket(_packet *pnet.QTPacket) *PlayerTeam {
	playerTeam := PlayerTeam{}
	playerTeam.Nick = _packet.ReadString()
	playerTeam.Info = _packet.ReadString()
	playerTeam.LoseMessage = _packet.ReadString()
	playerTeam.WinMessage = _packet.ReadString()
	playerTeam.avatar = (int)_packet.ReadUint16()
	playerTeam.DefaultTier = _packet.ReadString()
	playerTeam.Team = NewTeamFromPacket(_packet)
	
	return &playerTeam
}