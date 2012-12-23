package models

const (
	AbilityMessages_IdabilityMessages string = "ability_messages.idability_messages"
	AbilityMessages_AbilityId         string = "ability_messages.ability_id"
	AbilityMessages_Message           string = "ability_messages.message"
)

type AbilityMessages struct {
	IdabilityMessages int `PK`
	AbilityId         int
	Message           string
}
