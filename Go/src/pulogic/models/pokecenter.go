package models

const (
	Pokecenter_Idpokecenter string = "pokecenter.idpokecenter"
	Pokecenter_Position     string = "pokecenter.position"
	Pokecenter_Description  string = "pokecenter.description"
)

type Pokecenter struct {
	Idpokecenter int `PK`
	Position     int64
	Description  string
}
