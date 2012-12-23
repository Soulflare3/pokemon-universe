package main

import (
	"nonamelib/log" 
	
	"pulogic/models"
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
	entity := models.NpcPokemon { IdnpcPokemon: int(p.DbId),
								  Idpokemon: p.pokId,
								  Idnpc: int(p.NpcId),
								  IvHp: p.Hp,
								  IvAttack: p.Att,
								  IvAttackSpec: p.Att_spec,
								  IvDefence: p.Def,
								  IvDefenceSpec: p.Def_spec,
								  IvSpeed: p.Speed,
								  Gender: p.Gender,
								  HeldItem: p.Held_item }
	
	if err := g_orm.Save(&entity); err == nil {
		if p.IsNew {
			p.DbId = int64(entity.IdnpcPokemon)
		}
	}
	
	p.IsNew = false
	p.IsModified = false
	
	return true
}

func (p *NpcPokemon) Delete() bool {
	entity := models.NpcPokemon { IdnpcPokemon: int(p.DbId) }
	if _, err := g_orm.Delete(&entity); err != nil {
		log.Error("NpcPokemon", "Delete", "Failed to remove pokemon: %v", err.Error())
		return false
	}
	
	p.IsRemoved = true
	return true
}