package main

import (
	"fmt"
	puh "puhelper"
)

type NpcPokemon struct {
	DbId int64
	NpcId int64
	pokId int
	Name 	string
  	Hp int
  	Att int
  	Att_spec int
  	Def int
  	Def_spec int
  	Speed int
  	Gender int
  	Held_item int
  	
  	IsNew		bool
  	IsModified	bool
  	IsRemoved	bool
}

func NewNpcPokemon() *NpcPokemon {
	return &NpcPokemon{IsNew: true}
}

func (p *NpcPokemon) Save() bool {
	if p.IsNew {
		/*query := fmt.Sprintf(QUERY_INSERT_NPC_POKEMON, p.Name)
		if puh.DBQuery(query) == nil {
			m.DbId = int64(puh.DBGetLastInsertId())
		} else {
			return false
		}*/
	} else if p.IsModified {
		/*query := fmt.Sprintf(QUERY_UPDATE_NPC_POKEMON, p.Name)
		if puh.DBQuery(query) != nil {
			return false
		}*/
	}
	
	return true
}

func (p *NpcPokemon) Delete() bool {
	query := fmt.Sprintf(QUERY_DELETE_NPC_POKEMON, p.DbId)
	if puh.DBQuery(query) != nil {
		return false
	}
	
	p.IsRemoved = true
	return true
}