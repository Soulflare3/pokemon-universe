package models

const (
	PlayerOutfit_Idplayer string = "player_outfit.idplayer"
	PlayerOutfit_Head     string = "player_outfit.head"
	PlayerOutfit_Nek      string = "player_outfit.nek"
	PlayerOutfit_Upper    string = "player_outfit.upper"
	PlayerOutfit_Lower    string = "player_outfit.lower"
	PlayerOutfit_Feet     string = "player_outfit.feet"
)

type PlayerOutfit struct {
	Idplayer int `PK`
	Head     int
	Nek      int
	Upper    int
	Lower    int
	Feet     int
}
