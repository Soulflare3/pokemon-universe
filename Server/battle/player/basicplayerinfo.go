package main

import (
	pnet "network"
)

type BasicPlayerInfo struct {
	Nick string
	Info string
}

func NewBasicPlayerInfo(_nick, _info string) *BasicPlayerInfo {
	return &BasicPlayerInfo { Nick: _nick,
										  Info: _info }
}

func NewBasicPlayerInfoFromPacket(_packet *pnet.QTPacket) *BasicPlayerInfo {
	basicPlayerInfo := &BasicPlayerInfo{}
	basicPlayerInfo.Nick = _packet.ReadString()
	basicPlayerInfo.Info = _packet.ReadString()
	
	return basicPlayerInfo
}