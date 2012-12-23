package models

const (
	PlayerItems_IdplayerItems string = "player_items.idplayer_items"
	PlayerItems_Idplayer      string = "player_items.idplayer"
	PlayerItems_Iditem        string = "player_items.iditem"
	PlayerItems_Count         string = "player_items.count"
	PlayerItems_Slot          string = "player_items.slot"
)

type PlayerItems struct {
	IdplayerItems int `PK`
	Idplayer      int
	Iditem        int
	Count         int
	Slot          int
}
