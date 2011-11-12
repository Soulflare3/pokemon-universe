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
	pnet "network"
)

type Battle struct {
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
	histMutex *sync.Mutex
}

func NewBattle(_bc *BattleConf, _packet *pnet.QTPacket, _p1 *PlayerInfo, _p2 *PlayerInfo, _meID int, _bID int) *Battle {
	battle := Battle{}
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
	
	battle.pokes = make([][]*ShallowBattlePoke, 6)
	for i := 0; i < 2; i++ {
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
	b.histMutex.Lock()
	defer b.histMutex.Unlock()
	
	b.histDelta = b.histDelta + _message
}

func (b *Battle) ReceiveCommand(_packet *pnet.QTPacket) {
	bc := int(_packet.ReadUint8())
	player := int(_packet.ReadUint8())
	fmt.Printf("Battle command received: %d | PlayerId: %d", bc, player)
	switch bc {
		default:
			fmt.Printf("Battle command unimplemented: %d", bc)	
	}
}