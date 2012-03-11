package main

import (
	puh "puhelper"
//	pos "putools/pos"
)

// Simple NPC creature, for test purposes for the editor

type Npc struct {
  NpcName map[int]string
}

func NewNpcList() *Npc{
	return &Npc{ NpcName: make(map[int]string)  }
}

func (c *Npc) AddNpc(_npcId int, _npcName string) {
        c.NpcName[_npcId] = _npcName;
}

func (m *Npc) LoadNpcList() (succeed bool, error string) {
	var query string = "SELECT idnpc, name FROM npc ORDER BY idnpc"
		
	result, err := puh.DBQuerySelect(query)
	if err != nil {
		return false, err.Error()
	}
	
	defer result.Free()
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		idNpc := puh.DBGetInt(row[0])
		nameNpc := puh.DBGetString(row[1])
		
		m.AddNpc(idNpc, nameNpc)
	}
	
	return true, ""
}

func (m *Npc) GetNumNpcs() int {
	return len(m.NpcName)
}