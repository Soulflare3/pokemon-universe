package models

const (
	PlayerBackpack_IdplayerBackpack string = "player_backpack.idplayer_backpack"
	PlayerBackpack_Idplayer         string = "player_backpack.idplayer"
	PlayerBackpack_Iditem           string = "player_backpack.iditem"
	PlayerBackpack_Count            string = "player_backpack.count"
	PlayerBackpack_Slot             string = "player_backpack.slot"
)

type PlayerBackpack struct {
	IdplayerBackpack int `PK`
	Idplayer         int
	Iditem           int
	Count            int
	Slot             int
}
