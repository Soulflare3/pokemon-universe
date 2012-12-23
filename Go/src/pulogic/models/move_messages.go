package models

const (
	MoveMessages_IdmoveMessage string = "move_messages.idmove_message"
	MoveMessages_MoveEffectId string = "move_messages.move_effect_id"
	MoveMessages_Message      string = "move_messages.message"
)

type MoveMessages struct {
	IdmoveMessages int `PK`
	MoveEffectId int
	Message      string
}
