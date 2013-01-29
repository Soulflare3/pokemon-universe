package models

const (
	PlayerQuests_IdplayerQuests string = "player_quests.idplayer_quests"
	PlayerQuests_Idplayer       string = "player_quests.idplayer"
	PlayerQuests_Idquest        string = "player_quests.idquest"
	PlayerQuests_Status         string = "player_quests.status"
	PlayerQuests_Created        string = "player_quests.created"
	PlayerQuests_Finished       string = "player_quests.finished"
)

type PlayerQuests struct {
	IdplayerQuests int `PK`
	Idplayer       int
	Idquest        int
	Status         int
	Created        int64
	Finished       int64
}
