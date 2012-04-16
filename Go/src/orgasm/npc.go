package main

import (
	"fmt"
	puh "puhelper"
	pos "putools/pos"
)

type Npc struct {
  	Name 	string
  	Head 	int
  	Nek 	int
  	Upper 	int
	Lower	int
  	Feet 	int
  	Position pos.Position
	Pokemons map[int]*NpcPokemon
}

type NpcPokemon struct {
	Name 	string
  	Hp int
  	Att int
  	Att_spec int
  	Def int
  	Def_spec int
  	Speed int
  	Gender int
  	Held_item int
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

func (m *NpcList) LoadNpcPokemon() (succeed bool, error string) {
		var query string = "SELECT idnpc, idnpc_pokemon, iv_hp, iv_attack, iv_attack_spec, iv_defence, iv_defence_spec, iv_speed, gender, held_item FROM npc_pokemon"
		
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
			idPokemon := puh.DBGetInt(row[1])
			hp := puh.DBGetInt(row[2])
			attack := puh.DBGetInt(row[3])		
			attack_spec := puh.DBGetInt(row[4])
			defence := puh.DBGetInt(row[5])
			defence_spec := puh.DBGetInt(row[6])
			speed := puh.DBGetInt(row[7])
			gender := puh.DBGetInt(row[8])
			held_item := puh.DBGetInt(row[9])
			
			//TODO Pokemon Names
			m.Npcs[idNpc].AddPokemon(idPokemon, "Pokemon Name", hp, attack, attack_spec, defence, defence_spec, speed, gender, held_item)
		}
	return true, ""
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

func (m *NpcList) AddNpc(_npcId int, _npcName string, _head int, _nek int, _upper int, _lower int, _feet int, _position pos.Position) {
	npc := &Npc { Name: _npcName,
				  Head: _head,
				  Nek: _nek,
				  Upper: _upper,
				  Lower: _lower,
				  Feet: _feet, 
				  Position: _position,
				  Pokemons: make(map[int]*NpcPokemon) }
			
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

func (m *Npc) AddPokemon(_id int, _name string, _hp int, _att int, _att_spec int, _def int, _def_spec int, _speed int, _gender int, _held_item int) {
	npcPokemon := &NpcPokemon { 	Name : _name,
									Hp : _hp,
									Att : _att,
									Att_spec : _att_spec,
									Def : _def,
									Def_spec : _def_spec,
									Speed : _speed,
									Gender : _gender,
									Held_item : _held_item }
	m.Pokemons[_id] = npcPokemon
}
