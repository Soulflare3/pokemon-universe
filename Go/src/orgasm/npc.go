package main

import (
	"fmt"
	puh "puhelper"
	pos "putools/pos"
)

// Simple NPC creature, for test purposes for the editor
type Npc struct {
	Id		int
  	Name 	string
  	Head 	int
  	Nek 	int
  	Upper 	int
	Lower	int
  	Feet 	int
  	Position pos.Position
}

type NpcList struct {
	Npcs	map[int]*Npc
}

func NewNpcList() *NpcList{
	return &NpcList { Npcs: make(map[int]*Npc) }
}

func (m *NpcList) LoadNpcList() (succeed bool, error string) {
	var query string = "SELECT npc.idnpc, npc.name, npc_outfit.head, npc_outfit.nek, npc_outfit.upper, npc_outfit.lower, npc_outfit.feet, npc.position FROM npc INNER JOIN npc_outfit ON npc.idnpc = npc_outfit.idnpc ORDER BY npc.idnpc"
		
	result, err := puh.DBQuerySelect(query)
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
		
		idNpc := puh.DBGetInt(row[0])
		nameNpc := puh.DBGetString(row[1])		
		head := puh.DBGetInt(row[2])
		nek := puh.DBGetInt(row[3])
		upper := puh.DBGetInt(row[4])
		lower := puh.DBGetInt(row[5])
		feet := puh.DBGetInt(row[6])
		positionHash := puh.DBGetInt64(row[7])
		
		m.AddNpc(idNpc, nameNpc, head, nek, upper, lower, feet, pos.NewPositionFromHash(positionHash))
	}
	
	return true, ""
}

func (m *NpcList) GetNumNpcs() int {
	return len(m.Npcs)
}

func (m *NpcList) AddNpc(_npcId int, _npcName string, _head int, _nek int, _upper int, _lower int, _feet int, _position pos.Position) {
	npc := &Npc { Id: _npcId,
				  Name: _npcName,
				  Head: _head,
				  Nek: _nek,
				  Upper: _upper,
				  Lower: _lower,
				  Feet: _feet, 
				  Position: _position }
	
	m.Npcs[_npcId] = npc
}

func (m *NpcList) UpdateNpcAppearance(_npcId int, _npcName string, _head int, _nek int, _upper int, _lower int, _feet int) {
	// Get Npc from list
	npc, found := m.Npcs[_npcId]
	if !found {
		m.AddNpc(_npcId, _npcName, _head, _nek, _upper, _lower, _feet, pos.NewPosition())
	} else {
		npc.Name = _npcName
		npc.Head = _head
		npc.Nek = _nek
		npc.Upper = _upper
		npc.Lower = _lower
		npc.Feet = _feet
	}
}

func (m *NpcList) UpdateNpcPosition(_npcId int, _position pos.Position) {
	// Get Npc from list
	npc, found := m.Npcs[_npcId]
	if found {
		npc.Position = _position;
	} 
}

func (m *NpcList) DeleteNpc(_npcId int) {

// Delete NPC
delete(m.Npcs, _npcId)

}

