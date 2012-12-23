package models

const (
	PlayerFriends_IdplayerFriends string = "player_friends.idplayer_friends"
	PlayerFriends_Idplayer        string = "player_friends.idplayer"
	PlayerFriends_Idfriend        string = "player_friends.idfriend"
)

type PlayerFriends struct {
	IdplayerFriends int `PK`
	Idplayer        int
	Idfriend        int
}

type PlayerFriendsJoinPlayer struct {
	IdplayerFriends int `PK`
	Idfriend        int
	Name         	string
	Idlocation   	int
}