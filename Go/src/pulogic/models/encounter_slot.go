package models

const (
	EncounterSlot_IdencounterSlot string = "encounter_slot.idencounter_slot"
	EncounterSlot_Idencounter     string = "encounter_slot.idencounter"
	EncounterSlot_Idpokemon       string = "encounter_slot.idpokemon"
	EncounterSlot_GenderRate      string = "encounter_slot.gender_rate"
)

type EncounterSlot struct {
	IdencounterSlot int `PK`
	Idencounter     int
	Idpokemon       int
	GenderRate      int
}
