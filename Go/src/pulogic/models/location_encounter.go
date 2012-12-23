package models

const (
	LocationEncounter_Idencounter       string = "location_encounter.idencounter"
	LocationEncounter_IdlocationSection string = "location_encounter.idlocation_section"
)

type LocationEncounter struct {
	Idencounter       int `PK`
	IdlocationSection int `PK`
}
