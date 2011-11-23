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

type PU_Battle struct {
	battletype int

	eventQueue []IBattleEvent

	fighters     [4]*PU_Fighter
	self         *PU_Fighter
	changeTarget *PU_Fighter

	numPlayers int

	attackList [4]*PU_Attack

	state       int
	sleepTime   int32
	changeValue int
}

func NewBattle(_battletype int) *PU_Battle {
	return &PU_Battle{battletype: _battletype}
}

func (b *PU_Battle) Start() {
	g_game.panel.battleUI.Init()
	b.state = BATTLE_RUNNING
}

func (b *PU_Battle) Stop() {
	g_game.ShowBattleUI(false)
	g_game.ShowGameUI(true)

	g_game.battle = nil
	g_game.state = GAMESTATE_WORLD
}

func (b *PU_Battle) AddEvent(_event IBattleEvent) {
	b.eventQueue = append(b.eventQueue, _event)
}

func (b *PU_Battle) SetPokemon(_slot int) {
	if g_game.self != nil && g_game.self.pokemon[_slot] != nil {
		b.self.pokemon = g_game.self.pokemon[_slot]
	}
}

func (b *PU_Battle) SetPlayer(_id int, _team int, _name string, _pokename string, _pokeid int, _level int, _hp int) {
	fighter := NewFighter(_id)
	if fighter != nil {
		fighter.team = _team
		fighter.SetPokemon(_pokename, _pokeid, _level, _hp)
	}
	b.fighters[_id] = fighter
}

func (b *PU_Battle) SetNPC(_id int, _team int, _pokename string, _pokeid int, _level int) {
	fighter := NewFighter(_id)
	if fighter != nil {
		fighter.team = _team
		fighter.SetPokemon(_pokename, _pokeid, _level, 100)
	}
	b.fighters[_id] = fighter
}

func (b *PU_Battle) SetSelf(_id int, _team int, _starter int) {
	fighter := NewFighter(_id)
	if fighter != nil {
		fighter.SetSelf()
		fighter.team = _team
		fighter.pokemon = g_game.self.pokemon[_starter]
	}
	b.self = fighter
	b.SetPokemon(_starter)
	b.fighters[_id] = fighter
}

func (b *PU_Battle) GetEnemy() *PU_Fighter {
	for _, fighter := range b.fighters {
		if fighter != nil && fighter.team != b.self.team {
			return fighter
		}
	}
	return nil
}

func (b *PU_Battle) Wait(_ticks uint32) {
	b.sleepTime = int32(_ticks)
	b.state = BATTLE_WAITING
}

func (b *PU_Battle) ChangeHP(_fighter int, _hp int) {
	b.changeTarget = b.fighters[_fighter]
	b.changeValue = _hp
	b.state = BATTLE_CHANGEHP
}

func (b *PU_Battle) ChangeExp(_fighter int, _exp int) {
	b.changeTarget = b.fighters[_fighter]
	b.changeValue = _exp
	b.state = BATTLE_CHANGEEXP
}

func (b *PU_Battle) ProcessEvents() {
	switch b.state {
	case BATTLE_WAITING:
		b.sleepTime -= int32(g_frameTime)
		if b.sleepTime <= 0 {
			b.state = BATTLE_RUNNING
		}

	case BATTLE_CHANGEHP:
		oldhp := b.changeTarget.GetHP()
		if oldhp != b.changeValue {
			mod := 0
			if oldhp > b.changeValue {
				mod = -1
			} else {
				mod = 1
			}

			oldhp += mod
			b.changeTarget.SetHP(oldhp)
		} else {
			b.state = BATTLE_RUNNING
		}

	case BATTLE_CHANGEEXP:
		oldexp := b.changeTarget.GetExp()
		if oldexp != b.changeValue {
			mod := 0
			if oldexp > b.changeValue {
				mod = -1
			} else {
				mod = 1
			}

			oldexp += mod
			b.changeTarget.SetExp(oldexp)
		} else {
			b.state = BATTLE_RUNNING
		}

	case BATTLE_RUNNING:
		if len(b.eventQueue) > 0 {
			event := b.eventQueue[0]
			b.eventQueue = append(b.eventQueue[:0], b.eventQueue[1:]...)
			if event != nil {
				event.Execute()
			}
		}
	}
}
