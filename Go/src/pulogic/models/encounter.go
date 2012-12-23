package models

const (
	Encounter_Idencounter          string = "encounter.idencounter"
	Encounter_IdencounterCondition string = "encounter.idencounter_condition"
	Encounter_Rate                 string = "encounter.rate"
)

type Encounter struct {
	Idencounter          int `PK`
	IdencounterCondition int
	Rate                 int
}
