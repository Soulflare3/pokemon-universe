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
	pnet "network"
)

const (
	ChoiceType_Cancel uint8 = iota
	ChoiceType_Attack
	ChoiceType_Switch
	ChoiceType_Rearrange
	ChoiceType_CenterMove
)

type CancelChoice struct {

}

type AttackChoice struct {
	attackSlot   int8
	attackTarget int8
}

type SwitchChoice struct {
	pokeSlot int8
}

type RearrangeChoice struct {
	pokeIndexes [6]int8
}

type MoveToCenterChoice struct {

}

type ChoiceUnion struct {
	cancel    CancelChoice
	attack    AttackChoice
	switching SwitchChoice
	rearrange RearrangeChoice
	move      MoveToCenterChoice
}

type BattleChoices struct {
	switchAllowed  bool
	attacksAllowed bool
	attackAllowed  []bool

	numSlot uint8
}

func NewBattleChoices() BattleChoices {
	battleChoices := BattleChoices{switchAllowed: true,
		attacksAllowed: true,
		attackAllowed:  make([]bool, 4)}
	for i := 0; i < 4; i++ {
		battleChoices.attackAllowed[i] = true
	}

	return battleChoices
}

func (b *BattleChoices) initFromPackage(_packet *pnet.QTPacket) {
	b.numSlot = _packet.ReadUint8()
	b.switchAllowed = (_packet.ReadUint8() == 1)
	b.attacksAllowed = (_packet.ReadUint8() == 1)
	b.attackAllowed[0] = (_packet.ReadUint8() == 1)
	b.attackAllowed[1] = (_packet.ReadUint8() == 1)
	b.attackAllowed[2] = (_packet.ReadUint8() == 1)
	b.attackAllowed[3] = (_packet.ReadUint8() == 1)
}

func (b *BattleChoices) Struggle() bool {
	var count int = 0
	for i := 0; i < 4; i++ {
		if b.attackAllowed[i] {
			count++
		}
	}

	return (count == 0)
}

type BattleChoice struct {
	choiceType uint8
	playerSlot uint8
	choice     ChoiceUnion
}

func NewBattleChoice() *BattleChoice {
	return &BattleChoice{choice: ChoiceUnion{}}
}

func NewBattleChoiceCancel(_slot uint8, _choice CancelChoice) *BattleChoice {
	bc := NewBattleChoice()
	bc.choice.cancel = _choice
	bc.playerSlot = _slot
	bc.choiceType = ChoiceType_Cancel

	return bc
}

func NewBattleChoiceAttack(_slot uint8, _choice AttackChoice) *BattleChoice {
	bc := NewBattleChoice()
	bc.choice.attack = _choice
	bc.playerSlot = _slot
	bc.choiceType = ChoiceType_Attack

	return bc
}

func NewBattleChoiceSwitch(_slot uint8, _choice SwitchChoice) *BattleChoice {
	bc := NewBattleChoice()
	bc.choice.switching = _choice
	bc.playerSlot = _slot
	bc.choiceType = ChoiceType_Switch

	return bc
}

func NewBattleChoiceRearrange(_slot uint8, _choice RearrangeChoice) *BattleChoice {
	bc := NewBattleChoice()
	bc.choice.rearrange = _choice
	bc.playerSlot = _slot
	bc.choiceType = ChoiceType_Rearrange

	return bc
}

func NewBattleChoiceMoveToCenter(_slot uint8, _choice MoveToCenterChoice) *BattleChoice {
	bc := NewBattleChoice()
	bc.choice.move = _choice
	bc.playerSlot = _slot
	bc.choiceType = ChoiceType_CenterMove

	return bc
}

func (b *BattleChoice) SwitchChoice() bool {
	return (b.choiceType == ChoiceType_Switch)
}

func (b *BattleChoice) AttackingChoice() bool {
	return (b.choiceType == ChoiceType_Attack)
}

func (b *BattleChoice) MoveToCenterChoice() bool {
	return (b.choiceType == ChoiceType_CenterMove)
}

func (b *BattleChoice) Cancelled() bool {
	return (b.choiceType == ChoiceType_Cancel)
}

func (b *BattleChoice) RearrangeChoice() bool {
	return (b.choiceType == ChoiceType_Rearrange)
}

func (b *BattleChoice) setTarget(_target int8) {
	b.choice.attack.attackTarget = _target
}

func (b *BattleChoice) setAttackSlot(_slot int8) {
	b.choice.attack.attackSlot = _slot
}
