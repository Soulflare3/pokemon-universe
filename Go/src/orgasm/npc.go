package main

import (
	"fmt"
	puh "puhelper"
//	pos "putools/pos"
)

// Simple NPC creature, for test purposes for the editor

type Npc struct {
  NpcName map[int]string
  Head map[int]int
  Nek map[int]int
  Upper map[int]int
  Lower map[int]int
  Feet map[int]int
  
}

func NewNpcList() *Npc{
	return &Npc{ NpcName: make(map[int]string),
				 Head: make(map[int]int),
  				 Nek: make(map[int]int),
  				 Upper: make(map[int]int),
  				 Lower: make(map[int]int),
  			 	 Feet: make(map[int]int) }
}

func (c *Npc) AddNpc(_npcId int, _npcName string, _head int, _nek int, _upper int, _lower int, _feet int) {
        c.NpcName[_npcId] = _npcName
        c.Head[_npcId] = _head
        c.Nek[_npcId] = _nek
        c.Upper[_npcId] = _upper
        c.Lower[_npcId] = _lower
        c.Feet[_npcId] = _feet
}

func (m *Npc) LoadNpcList() (succeed bool, error string) {
	var query string = "SELECT npc.idnpc, npc.name, npc_outfit.head, npc_outfit.nek, npc_outfit.upper, npc_outfit.lower, npc_outfit.feet FROM npc INNER JOIN npc_outfit ON npc.idnpc = npc_outfit.idnpc ORDER BY npc.idnpc"
		
	result, err := puh.DBQuerySelect(query)
	if err != nil {
		fmt.Printf(err.Error())
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
		head := puh.DBGetInt(row[2])
		nek := puh.DBGetInt(row[3])
		upper := puh.DBGetInt(row[4])
		lower := puh.DBGetInt(row[5])
		feet := puh.DBGetInt(row[6])
		
		
		m.AddNpc(idNpc, nameNpc, head, nek, upper, lower, feet )
	}
	
	return true, ""
}

func (m *Npc) GetNumNpcs() int {
	return len(m.NpcName)
}