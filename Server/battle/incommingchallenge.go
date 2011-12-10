package main

import (
	pnet "network"
)

type IncommingChallenge struct {
	desc		int
	mode		int
	opponent	int
	clauses		int
	oppName	 	string
	note 		int
}

func NewIncommingChallengeFromPacket(_packet *pnet.QTPacket) *IncommingChallenge {
	return &IncommingChallenge {  desc: int(_packet.ReadUint8()),
								opponent: int(_packet.ReadUint32()),
								clauses: int(_packet.ReadUint32()),
								mode: int(_packet.ReadUint8()) }
}

func (i *IncommingChallenge) SetNick(_playerInfo *PlayerInfo) {
	if _playerInfo != nil {
		i.oppName = _playerInfo.Nick
	}
}

func (i *IncommingChallenge) IsValidChallenge(_players PlayerInfoList) bool {
	i.SetNick(_players[i.opponent])
	return (i.desc == CHALLENGEDESC_SENT && i.oppName != "")
}