package models

const (
	Location_Idlocation   string = "location.idlocation"
	Location_Name         string = "location.name"
	Location_Idpokecenter string = "location.idpokecenter"
	Location_Idmusic      string = "location.idmusic"
)

type Location struct {
	Idlocation   int `PK`
	Name         string
	Idpokecenter int
	Idmusic      int
}
