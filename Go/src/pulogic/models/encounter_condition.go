package models

const (
	EncounterCondition_IdencounterCondition string = "encounter_condition.idencounter_condition"
	EncounterCondition_Name                 string = "encounter_condition.name"
	EncounterCondition_Default              string = "encounter_condition.default"
)

type EncounterCondition struct {
	IdencounterCondition int `PK`
	Name                 string
	Default              int
}
