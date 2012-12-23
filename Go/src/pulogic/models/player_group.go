package models

const (
	PlayerGroup_PlayerIdplayer string = "player_group.player_idplayer"
	PlayerGroup_GroupIdgroup   string = "player_group.group_idgroup"
)

type PlayerGroup struct {
	PlayerIdplayer int `PK`
	GroupIdgroup   int `PK`
}
