/*Pokemon Universe MMORPG
Copyright (C) 2010 the Pokemon Universe Authors

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.*/
package main

import (
	"sync"
	"fmt"
	"time"
	pnet "network"
)

type Battle struct {
	owner *POClient
	players []*PlayerInfo // 0 = you, 1 = opponent
	
	remainingTime []int
	ticking []bool
	startingTime []int64
	
	mode int
	numberOfSlots int	
	me int
	opp int
	bID int
	
	myTeam *BattleTeam
	oppTeam *ShallowShownTeam
	
	gotEnd bool
	allowSwitch bool
	allowAttack bool
	clicked bool
	allowAttacks []bool
	background int
	shouldShowPreview bool
	shouldStruggle bool
	
	displayedMoves []*BattleMove
	conf *BattleConf
	
	pokes [][]*ShallowBattlePoke
	pokeAlive map[int] bool
	
	dynamicInfo []*BattleDynamicInfo
	
	histDelta string
	histMutex *sync.RWMutex
}

func NewBattle(_owner *POClient, _bc *BattleConf, _packet *pnet.QTPacket, _p1 *PlayerInfo, _p2 *PlayerInfo, _meID int, _bID int) *Battle {
	battle := Battle{}
	battle.owner = _owner
	battle.conf = _bc
	battle.bID = _bID
	battle.myTeam = NewBattleTeamFromPacket(_packet)
	
	// Only supporting singles
	battle.numberOfSlots = 2
	battle.players = make([]*PlayerInfo, 2)
	battle.players[0] = _p1
	battle.players[1] = _p2
	
	// Figure out who's who
	if battle.players[0].Id == _meID {
		battle.me = 0
		battle.opp = 1
	} else {
		battle.me = 1
		battle.opp = 0
	}
	
	battle.remainingTime = make([]int, 2)
	battle.remainingTime[0] = 5 * 60
	battle.remainingTime[1] = 5 * 60
	
	battle.ticking = make([]bool, 2)
	battle.ticking[0] = false
	battle.ticking[1] = false
	
	battle.background = 1
	
	battle.pokes = make([][]*ShallowBattlePoke, 2)
	for i := 0; i < 2; i++ {
		battle.pokes[i] = make([]*ShallowBattlePoke, 6)
		for j := 0; j < 6; j++ {
			battle.pokes[i][j] = NewShallowBattlePoke()
		}
	}
	
	battle.displayedMoves = make([]*BattleMove, 4)
	for i := 0; i < 4; i++ {
		battle.displayedMoves[i] = NewBattleMove()
	}
	
	battle.WriteToHist(fmt.Sprintf("Battle between %v and %v started!", battle.players[0].Nick, battle.players[1].Nick))
	
	return &battle
}

func (b *Battle) WriteToHist(_message string) {
	//b.histMutex.Lock()
	//defer b.histMutex.Unlock()
	
	b.histDelta = b.histDelta + _message
}

func (b *Battle) ReceiveCommand(_packet *pnet.QTPacket) {
	bc := int(_packet.ReadUint8())
	player := int(_packet.ReadUint8())
	fmt.Printf("Battle command received: %d | PlayerId: %d\n", bc, player)
	switch bc {
		case BattleCommand_SendOut: // 0
			b.receivedSendOut(_packet, player)
		case BattleCommand_SendBack: // 1
			b.receivedSendBack(player)
		case BattleCommand_UseAttack: // 2
			b.receivedUseAttack(_packet, player)
		case BattleCommand_OfferChoice: // 3
			b.receiveOfferChoice(_packet)
		case BattleCommand_BeginTurn: // 4
			b.receivedBeginTurn(_packet)
		case BattleCommand_ChangePP: // 5
			b.receivedChangePP(_packet)
		case BattleCommand_ChangeHp: // 6
			b.receivedChangeHp(_packet, player)
		case BattleCommand_Ko: // 7
			b.receivedKo(player)
		case BattleCommand_AbsStatusChange: // 25
			b.receivedAbsStatusChange(_packet, player)
		case BattleCommand_BlankMessage: // 28
			fmt.Println("")
		case BattleCommand_DynamicInfo: // 31
			b.receiveDynamicInfo(_packet, player)
		case BattleCommand_DynamicStats: // 32
			b.receiveDynamicStats(_packet, player)
		case BattleCommand_ClockStart: // 37
			b.clockStart(_packet, player)
		case BattleCommand_ClockStop: // 38
			b.clockStop(_packet, player)
		case BattleCommand_Rated: // 39
			b.receivedRated(_packet)
		case BattleCommand_TierSection: // 40
			tier := _packet.ReadString()
			b.WriteToHist("Tier: " + tier)
		case BattleCommand_MakeYourChoice: // 43
			b.receiveMakeYourCoice()
		default:
			fmt.Printf("Battle command unimplemented: %d\n", bc)	
	}
}

func (b *Battle) isOut(_poke int) bool {
	return (_poke < (b.numberOfSlots / 2))
}

func (b *Battle) currentPoke(_player int) *ShallowBattlePoke {
	return b.pokes[_player][0]
}

// -------------------- Received Messages ----------------------
func (b *Battle) receivedSendOut(_packet *pnet.QTPacket, _player int) {
	isSilent := _packet.ReadBool()
	fromSpot := int(_packet.ReadUint8())
	
	if _player == b.me {
		// tmp := b.myTeam.Pokes[0]
		b.myTeam.Pokes[0], b.myTeam.Pokes[fromSpot] = b.myTeam.Pokes[fromSpot], b.myTeam.Pokes[0]
		for i := 0; i < 4; i++ {
			b.displayedMoves[i] = NewBattleMoveFromBattleMove(b.myTeam.Pokes[0].Moves[i])
		}
	}
	
	b.pokes[_player][0], b.pokes[_player][fromSpot] = b.pokes[_player][fromSpot], b.pokes[_player][0]
	if _packet.GetMsgSize() > _packet.GetReadPos() { // this is the first time you've seen it
		b.pokes[_player][0] = NewShallowBattlePokeFromPacket(_packet, (_player == b.me))
	}
	
	// TOOD: Send updatePokes to PU client
	// TODO: Send updatePokeballs to PU client
	
	if !isSilent {
		b.WriteToHist(fmt.Sprintf("%s sent out %s!\n", b.players[_player].Nick, b.currentPoke(_player).RNick))
	}
}

func (b *Battle) receivedSendBack(_player int) {
	b.WriteToHist(fmt.Sprintf("%s called %s back!\n", b.players[_player].Nick, b.currentPoke(_player).RNick))
}

func (b *Battle) receivedUseAttack(_packet *pnet.QTPacket, _player int) {
	attack := int(_packet.ReadUint16())
	b.WriteToHist(fmt.Sprintf("%s used %s!\n", b.currentPoke(_player).Nick, g_PokemonManager.GetMoveById(attack).Identifier))
}

func (b *Battle) receiveOfferChoice(_packet *pnet.QTPacket) {
	_packet.ReadUint8() // We don't need it (numSlot)
	b.allowSwitch = _packet.ReadBool()
	b.allowAttack = _packet.ReadBool()
	canDoAttack := false
	for i := 0; i < 4; i++ {
		b.allowAttacks[i] = _packet.ReadBool()
		if b.allowAttacks[i] {
			canDoAttack = true
		}
	}
	
	if b.allowAttack && !canDoAttack {
		b.shouldStruggle = true
	} else {
		b.shouldStruggle = false
	}
	
	// TODO: Send updateButtons to PU client
}

func (b *Battle) receivedBeginTurn(_packet *pnet.QTPacket) {
	turn := _packet.ReadUint32()
	b.WriteToHist(fmt.Sprintf("Start of turn %d!\n", turn))
}

func (b *Battle) receivedChangePP(_packet *pnet.QTPacket) {
	moveNum := int(_packet.ReadUint8())
	newPP := int(_packet.ReadUint8())
	b.displayedMoves[moveNum].CurrentPP = newPP
	b.myTeam.Pokes[0].Moves[moveNum].CurrentPP = newPP
	
	// TODO: Send updateMovePP to PUClient
}

func (b *Battle) receivedChangeHp(_packet *pnet.QTPacket, _player int) {
	newHp := int(_packet.ReadUint16())
	if _player == b.me {
		b.myTeam.Pokes[0].CurrentHP = newHp;
		b.currentPoke(_player).LastKnownPercent = newHp
		b.currentPoke(_player).LifePercent = (newHp * 100) / b.myTeam.Pokes[0].TotalHP
	} else {
		b.currentPoke(_player).LastKnownPercent = newHp
		b.currentPoke(_player).LifePercent = newHp
	}
	
	// TODO: Send HP update to PU Client
}

func (b *Battle) receivedKo(_player int) {
	b.WriteToHist(fmt.Sprintf("%s fainted!\n", b.currentPoke(_player).Nick))
}

func (b *Battle) receivedAbsStatusChange(_packet *pnet.QTPacket, _player int) {
	poke := int(_packet.ReadUint8())
	status := uint(_packet.ReadUint8())
	
	if (poke >= 0 || poke < 6) && status != STATUS_CONFUSED {
		b.pokes[_player][poke].ChangeStatus(status)
		if _player == b.me {
			b.myTeam.Pokes[poke].ChangeStatus(status)
		}
		
		// if b.isOut(poke) {
			// TODO: Send updatePokes to PU client
		// }
		// TODO: Send updatePokeballs to PU client
	}
}

func (b *Battle) receiveDynamicInfo(_packet *pnet.QTPacket, _player int) {
	b.dynamicInfo[_player] = NewBattleDynamicInfoFromPacket(_packet)
}

func (b *Battle) receiveDynamicStats(_packet *pnet.QTPacket, _player int) {
	for i := 0; i < 5; i++ {
		b.myTeam.Pokes[_player / 2].Stats[i] = int(_packet.ReadUint16())
	}
}

func (b *Battle) clockStart(_packet *pnet.QTPacket, _player int) {
	index := _player % 2
	b.remainingTime[index] = int(_packet.ReadUint16())
	b.startingTime[index] = time.Seconds()
	b.ticking[index] = true
}

func (b *Battle) clockStop(_packet *pnet.QTPacket, _player int) {
	index := _player % 2
	b.remainingTime[index] = int(_packet.ReadUint16())
	b.ticking[index] = false
}

func (b *Battle) receivedRated(_packet *pnet.QTPacket) {
	rated := "Unrated"
	if _packet.ReadBool() {
		rated = "Rated"
	}
	b.WriteToHist("Rule: " + rated)
	
	// TODO: Print clauses
}

func (b *Battle) receiveMakeYourCoice() {
	// TODO: Send updateButtons to PU client
	if b.allowSwitch && !b.allowAttack {
		// TOOD: Send switchToPokeViewer to PU client
	}
}

// -------------------- Send Messages ----------------------
func (b *Battle) sendBattleMessageAttack(_attackSlot int) {
	packet := pnet.NewQTPacket()
	packet.AddUint32(uint32(b.bID))
	ac := NewAttackChoice(_attackSlot, b.opp)
	bc := NewBattleChoiceWithChoice(b.me, ac, CHOICETYPE_ATTACKTYPE)
	packet.AddBuffer(bc.WritePacket().GetBufferSlice())
	
	b.owner.SendMessage(packet, COMMAND_BattleMessage)
}

func (b *Battle) sendBattleMessageSwitch(_toSpot int) {
	packet := pnet.NewQTPacket()
	packet.AddUint32(uint32(b.bID))
	sc := NewSwitchChoice(_toSpot)
	bc := NewBattleChoiceWithChoice(b.me, sc, CHOICETYPE_SWITCHTYPE)
	packet.AddBuffer(bc.WritePacket().GetBufferSlice())
	
	b.owner.SendMessage(packet, COMMAND_BattleMessage)
}

func (b *Battle) sendBattleMessageCancel() {
	packet := pnet.NewQTPacket()
	packet.AddUint32(uint32(b.bID))
	bc := NewBattleChoice(b.me, CHOICETYPE_CANCELTYPE)
	packet.AddBuffer(bc.WritePacket().GetBufferSlice())
	
	b.owner.SendMessage(packet, COMMAND_BattleMessage)
}