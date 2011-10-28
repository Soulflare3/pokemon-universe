package main

import (
	pnet "network"
)

type PlayerInfo struct {
	Id int
	Nick string
	Info string
	Auth int
	Rating int
	Pokes []*UniqueId
	Tier string
	
	flags int
	rating int
	avatar int
	color *QColor
	gen int
}

func NewPlayerInfo() *PlayerInfo {
	return &PlayerInfo{}
}

func NewPlayerInfoFromPacket(_packet *pnet.QTPacket) *PlayerInfo {
	playerInfo := &PlayerInfo{};
	playerInfo.Id = int(_packet.ReadUint32())
	playerInfo.Nick = _packet.ReadString()
	playerInfo.Info = _packet.ReadString()
	playerInfo.Auth = int(_packet.ReadUint8())
	playerInfo.flags = int(_packet.ReadUint8())
	playerInfo.rating = int(_packet.ReadUint16())
	
	for i := 0; i < 6; i++ {
		playerInfo.Pokes[i] = NewUniqueIdFromPacket(_packet)
	}
	
	playerInfo.avatar = int(_packet.ReadUint8())
	playerInfo.Tier = _packet.ReadString()
	playerInfo.color = NewQColorFromPacket(_packet)
	playerInfo.gen = int(_packet.ReadUint8())
	
	return playerInfo
}