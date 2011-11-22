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
	CHOICETYPE_CANCELTYPE int = iota
	CHOICETYPE_ATTACKTYPE
	CHOICETYPE_SWITCHTYPE
	CHOICETYPE_REARRANGETYPE
	CHOICETYPE_CENTERMOVETYPE
	CHOICETYPE_DRAWTYPE
)

type IChoice interface {
	GetChoiceType() int
	WritePacket() pnet.IPacket
}

//
// Attack Choice
type AttackChoice struct {
	AttackSlot int
	AttackTarget int
}

func NewAttackChoice(_as, _at int) *AttackChoice {
	attackChoice := AttackChoice{ AttackSlot: _as, AttackTarget: _at }
	return &attackChoice
}

func NewAttackChoiceFromPacket(_packet *pnet.QTPacket) *AttackChoice {
	return NewAttackChoice(int(_packet.ReadUint8()), int(_packet.ReadUint8()))
}

func (c *AttackChoice) GetChoiceType() int {
	return CHOICETYPE_ATTACKTYPE
}

func (c *AttackChoice) WritePacket() pnet.IPacket {
	packet := pnet.NewQTPacket()
	packet.AddUint8(uint8(c.AttackSlot))
	packet.AddUint8(uint8(c.AttackTarget))
	return packet
}

//
// Switch Choice
type SwitchChoice struct {
	PokeSlot int
}

func NewSwitchChoice(_slot int) *SwitchChoice {
	switchChoice := SwitchChoice{ PokeSlot: _slot }
	return &switchChoice
}

func NewSwitchChoiceFromPacket(_packet *pnet.QTPacket) *SwitchChoice {
	return NewSwitchChoice(int(_packet.ReadUint8()))
}

func (c *SwitchChoice) GetChoiceType() int {
	return CHOICETYPE_ATTACKTYPE
}

func (c *SwitchChoice) WritePacket() pnet.IPacket {
	packet := pnet.NewQTPacket()
	packet.AddUint8(uint8(c.PokeSlot))
	return packet
}

//
// Rearrange Choice
type RearrangeChoice struct {
	PokeIndexes []int
}

func NewRearrangeChoice(_team *BattleTeam) *RearrangeChoice {
	rearrangeChoice := RearrangeChoice{ }
	rearrangeChoice.PokeIndexes = make([]int, 6)
	for i := 0; i < 6; i++ {
		rearrangeChoice.PokeIndexes[i] = _team.Pokes[i].TeamNum
	}
	
	return &rearrangeChoice
}

func NewRearrangeChoiceFromPacket(_packet *pnet.QTPacket) *RearrangeChoice {
	rearrangeChoice := RearrangeChoice{ }
	rearrangeChoice.PokeIndexes = make([]int, 6)
	for i := 0; i < 6; i++ {
		rearrangeChoice.PokeIndexes[i] = int(_packet.ReadUint8())
	}
	
	return &rearrangeChoice
}

func (c *RearrangeChoice) GetChoiceType() int {
	return CHOICETYPE_REARRANGETYPE
}

func (c *RearrangeChoice) WritePacket() pnet.IPacket {
	packet := pnet.NewQTPacket()
	for i := 0; i < 6; i++ {
		packet.AddUint8(uint8(c.PokeIndexes[i]))
	}
	return packet
}

//
// MovetoCenter Choice
type MoveToCenterChoice struct {
}

func NewMoveToCenterChoice() *MoveToCenterChoice {
	return &MoveToCenterChoice{}
}

func (m *MoveToCenterChoice) GetChoiceType() int {
	return CHOICETYPE_CENTERMOVETYPE
}

func (c *MoveToCenterChoice) WritePacket() pnet.IPacket {
	return pnet.NewQTPacket()
}

//
// Draw Choice
type DrawChoice struct {
}

func NewDrawChoice() *DrawChoice {
	return &DrawChoice{}
}

func (m *DrawChoice) GetChoiceType() int {
	return CHOICETYPE_DRAWTYPE
}

func (c *DrawChoice) WritePacket() pnet.IPacket {
	return pnet.NewQTPacket()
}

//
// Battle Choice
type BattleChoice struct {
	PlayerSlot int
	Choice IChoice
	ChoiceType int
}

func NewBattleChoiceWithChoice(_ps int, _c IChoice, _ct int) *BattleChoice {
	return &BattleChoice { PlayerSlot: _ps,
									Choice: _c,
									ChoiceType: _ct }
}

func NewBattleChoice(_ps int, _ct int) *BattleChoice {
	return &BattleChoice { PlayerSlot: _ps,
									ChoiceType: _ct }
}

func NewBattleChoiceFromPacket(_packet *pnet.QTPacket) *BattleChoice {
	battleChoice := BattleChoice{}
	battleChoice.PlayerSlot = int(_packet.ReadUint8())
	battleChoice.ChoiceType = int(_packet.ReadUint8())
	
	switch battleChoice.ChoiceType {
		case CHOICETYPE_SWITCHTYPE:
			battleChoice.Choice = NewSwitchChoiceFromPacket(_packet)
		case CHOICETYPE_ATTACKTYPE:
			battleChoice.Choice = NewAttackChoiceFromPacket(_packet)
		case CHOICETYPE_REARRANGETYPE:
			battleChoice.Choice = NewRearrangeChoiceFromPacket(_packet)
	}
	
	return &battleChoice
}

func (b *BattleChoice) WritePacket() pnet.IPacket {
	packet := pnet.NewQTPacket()
	packet.AddUint8(uint8(b.PlayerSlot))
	packet.AddUint8(uint8(b.ChoiceType))
	
	switch b.ChoiceType {
		case CHOICETYPE_SWITCHTYPE:
			fallthrough
		case CHOICETYPE_ATTACKTYPE:
			fallthrough
		case CHOICETYPE_REARRANGETYPE:
			packet.AddBuffer(b.Choice.WritePacket().GetBufferSlice())
	}
	
	return packet
}