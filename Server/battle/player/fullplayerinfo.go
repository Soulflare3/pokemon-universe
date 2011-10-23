package main

import (
	"os"
	pnet "network"
)

type FullPlayerInfo struct {
	Team *PlayerTeam
	IsDefault bool
	
	ladderEnabled bool
	showTeam bool
	nameColour *QColor
}

func NewFullPlayerInfo(_team *PlayerTeam, _ladderEnabled, _showTeam bool) *FullPlayerInfo {
	fullPlayerInfo := &FullPlayerInfo { Team: _team,
										IsDefault: true,
										ladderEnabled: _ladderEnabled,
										showTeam: _showTeam }
	return fullPlayerInfo
}

func NewFullPlayerInfoFromPacket(_packet *pnet.QTPacket) *FullPlayerInfo {
	fullPlayerInfo := &FullPlayerInfo{}
	fullPlayerInfo.Team = NewPlayerTeamFromPacket(_packet)
	fullPlayerInfo.IsDefault = true
	fullPlayerInfo.ladderEnabled = _packet.ReadBool()
	fullPlayerInfo.nameColor = NewQColorFromPacket(_packet)
	
	return fullPlayerInfo
}

func (p *FullPlayerInfo) Nick() string {
	return p.Team.Nick
}

func (p *FullPlayerInfo) WritePacket() (pnet.IPacket, os.Error) {
	packet := NewQTPacket()
	
	teamPacket, _ := p.Team.WritePacket()
	packet.AddBuffer(teamPacket.GetBuffer(), teamPacket.GetMsgSize)
	packet.AddBool(p.ladderEnabled)
	packet.AddBool(p.showTeam)
	
	colorPacket, _ := p.nameColour.WritePacket()
	packet.AddBuffer(colorPacket.GetBuffer(), colorPacket.GetMsgSize())

	return packet, nil
}