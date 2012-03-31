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
	"time"
	
	puh "puhelper"
	"putools/log"
)

type QuestStore struct {
	quests map[int64]*Quest
}

func NewQuestStore() *QuestStore {
	return &QuestStore { quests: make(map[int64]*Quest) }
}

func (s *QuestStore) Load() bool {
	result, err := puh.DBQuerySelect("SELECT idquests, name, description FROM quests")
	if err != nil {
		return false
	}
	
	defer puh.DBFree()
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		quest := &Quest { Dbid: puh.DBGetInt64(row[0]),
						  Name: puh.DBGetString(row[1]),
						  Description: puh.DBGetString(row[2]) }
		s.AddQuest(quest)
	}
	
	return true
}

func (s *QuestStore) AddQuest(_quest *Quest) {
	s.quests[_quest.Dbid] = _quest
}

func (s *QuestStore) GetQuest(_dbid int64) (quest *Quest, found bool) {
	quest, found = s.quests[_dbid]
	return
}

type Quest struct {
	Dbid		int64
	Name		string
	Description	string
}

func NewQuest(_dbid int64, _name string) *Quest {
	return &Quest { Dbid: _dbid,
					Name: _name } 
}

type PlayerQuestList map[int64]*PlayerQuest
type PlayerQuest struct {
	Dbid		int64
	Quest		*Quest
	Status		int
	Created		time.Time
	Finished	time.Time
	IsFinished	bool
	
	// Saving variables
	IsNew		bool
	IsModified	bool
	IsAbandoned bool
}

func NewPlayerQuest(_questId int64, _status int) *PlayerQuest {
	quest, found := g_game.Quests.GetQuest(_questId)
	if !found {
		logger.Printf("NewPlayerQuest - Failed to get quest with id: %d\n", _questId)
		return nil 
	}	
	playerQuest := PlayerQuest { Dbid: 0,
								  Quest: quest,
								  Status: _status,
								  Created: time.Now(),
								  Finished: time.Unix(0, 0),
								  IsFinished: false,
								  IsNew: true,
								  IsModified: false,
								  IsAbandoned: false }
	return &playerQuest
}

func NewPlayerQuestExt(_dbid int64, _questId int64, _status int, _created int64, _finished int64) *PlayerQuest {
	createdTime := time.Unix(_created, 0)
	finishedTime := time.Unix(_finished, 0)
	quest, found := g_game.Quests.GetQuest(_questId)
	if !found {
		logger.Printf("NewPlayerQuestExt - Failed to get quest with id: %d\n", _questId)
		return nil 
	}
	
	playerQuest := PlayerQuest { Dbid: _dbid,
								  Quest: quest,
								  Status: _status,
								  Created: createdTime,
								  Finished: finishedTime,
								  IsFinished: (_finished > 0),
								  IsNew: true,
								  IsModified: false,
								  IsAbandoned: false }
	return &playerQuest
}

func (p *PlayerQuest) UpdateStatus(_newStatus int) {
	if !p.IsAbandoned {
		p.Status = _newStatus
		p.IsModified = true
		
		if p.Status == 100 {
			p.IsFinished = true
			p.Finished = time.Now()
		}
	}
}

func (p *PlayerQuest) Abandon() {
	p.IsAbandoned = true
}