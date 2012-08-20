package main

import (
	"fmt"
	puh "puhelper"
	pos "putools/pos"
	pul "pulogic"
)


type NpcList struct {
	Npcs	map[int64]*Npc
}

func NewNpcList() *NpcList{
	return &NpcList { Npcs: make(map[int64]*Npc) }
}

func (m *NpcList) LoadNpcList() (succeed bool, error string) {
	result, err := puh.DBQuerySelect(QUERY_SELECT_NPCS)
	if err != nil {
		fmt.Printf(err.Error())
		return false, err.Error()
	}
	
	defer puh.DBFree()
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		// Create new NPC object with data from database
		npc := NewNpc()
		npc.IsNew = false
		npc.DbId = puh.DBGetInt64(row[0])
		npc.Name = puh.DBGetString(row[1])
		
		npc.SetOutfitPart(pul.OUTFIT_HEAD, puh.DBGetInt(row[2]))
		npc.SetOutfitPart(pul.OUTFIT_NEK, puh.DBGetInt(row[3]))
		npc.SetOutfitPart(pul.OUTFIT_UPPER, puh.DBGetInt(row[4]))
		npc.SetOutfitPart(pul.OUTFIT_LOWER, puh.DBGetInt(row[5]))
		npc.SetOutfitPart(pul.OUTFIT_FEET, puh.DBGetInt(row[6]))
		
		positionHash := puh.DBGetInt64(row[7])
		npc.Position = pos.NewPositionFromHash(positionHash)
		npc.Events = puh.DBGetStringFromArray(row[8])
		npc.EventInitId = puh.DBGetInt(row[9])
		
		// Load NPC pokemon
		npc.LoadPokemon()
		
		// Save in Npc map
		m.Npcs[npc.DbId] = npc
	}
	
	return true, ""
}

func (m *NpcList) GetNpcById(_npcId int64) (*Npc, bool) {
	npc, ok := m.Npcs[_npcId]
	
	return npc, ok
}

func (m *NpcList) GetNumNpcs() int {
	return len(m.Npcs)
}

func (m *NpcList) GetNumPokemons() int {
	var count int
	for _, npc := range(m.Npcs) {
		count += len(npc.Pokemons)
	}
	return count
}

func (m *NpcList) AddNpc(_npc *Npc) {
	if _npc.Save() {
		m.Npcs[_npc.DbId] = _npc
	}
}