package main

const (
	PlayerInfo_LoggedIn = 1
	PlayerInfo_Battling = 2
	PlayerInfo_Away = 4
)

type PlayerInfo struct {
	id		int32
	team	*BasicInfo
	auth	int8
	flags	uint8
	rating	int16
	pokes	[]*PokemonUniqueId
	avatar	uint16
	tier	string
	color	uint32
	gen		uint8
}

func NewPlayerInfo() *PlayerInfo {
	return &PlayerInfo{ pokes: make([]*PokemonUniqueId, 6) }
}