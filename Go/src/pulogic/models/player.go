package models

const (
	Player_Idplayer     string = "player.idplayer"
	Player_Idaccount    string = "player.idaccount"
	Player_Name         string = "player.name"
	Player_Password     string = "player.password"
	Player_PasswordSalt string = "player.password_salt"
	Player_Position     string = "player.position"
	Player_Movement     string = "player.movement"
	Player_Idpokecenter string = "player.idpokecenter"
	Player_Money        string = "player.money"
	Player_Idlocation   string = "player.idlocation"
)

type Player struct {
	Idplayer     int `PK`
	Idaccount    int
	Name         string
	Password     string
	PasswordSalt string
	Position     int64
	Movement     int
	Idpokecenter int
	Money        int
	Idlocation   int
}

type PlayerJoinOutfitJoinGroup struct {
	Idaccount    int
	Name         string
	Position     int64
	Movement     int
	Idpokecenter int
	Money        int
	Idlocation   int
	
	Head     int
	Nek      int
	Upper    int
	Lower    int
	Feet     int
	
	GroupIdgroup   int
}