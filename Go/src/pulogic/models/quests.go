package models

const (
	Quests_Idquests    string = "quests.idquests"
	Quests_Name        string = "quests.name"
	Quests_Description string = "quests.description"
)

type Quests struct {
	Idquests    int `PK`
	Name        string
	Description string
}
