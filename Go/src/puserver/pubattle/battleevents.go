package pubattle

import (
	pnet "network"
	pul "pulogic"
	"pulogic/netmsg"
)

func SendBattleEvent_ChangePP(_player pul.IBattleCreature, _pokemonId int, _moveSlotId int, _newPP int) {
	message := netmsg.NewBattleEventMessage(pnet.BATTLEEVENT_CHANGEPP)
	message.PokemonId = uint32(_pokemonId)
	message.MoveSlotId = uint32(_moveSlotId)
	message.NewPP = uint32(_newPP)
	
	_player.SendBattleMessage(message)
}

func SendBattleEvent_ChangeHP(_player pul.IBattleCreature, _pokemonId int, _hp int) {
	message := netmsg.NewBattleEventMessage(pnet.BATTLEEVENT_CHANGEHP)
	message.PokemonId = uint32(_pokemonId)
	message.NewHP = uint16(_hp)
	
	_player.SendBattleMessage(message)
}

func SendBattleEvent_Message(_player pul.IBattleCreature, _message string) {
	message := netmsg.NewBattleEventMessage(pnet.BATTLEEVENT_TEXT)
	message.Text = _message
}