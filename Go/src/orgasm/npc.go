package main

import (	
	"fmt"
	
	pos "nonamelib/pos"
	"nonamelib/log"
	
	"pulogic"
	"pulogic/models"
)

type Npc struct {
	DbId	int64
  	Name 	string
  	
  	Outfit
  	
  	Position pos.Position
	Pokemons map[int64]*NpcPokemon
	Events string
	EventInitId int
	
	IsNew		bool
	IsModified	bool
	IsRemoved	bool
}

func NewNpc() *Npc {
	return &Npc { IsNew: true,
				  Pokemons: make(map[int64]*NpcPokemon),
				  Outfit: NewOutfit() }
}

func (m *Npc) LoadPokemon() (bool, string) {
	if m.IsNew {
		// If an NPC is new, it isn't saved in the database so it can't have any pokemon in there
		return true, ""
	}
	
	var entities []models.NpcPokemon
	err := g_orm.Where(fmt.Sprintf("%v = '%d'", models.NpcPokemon_Idnpc, m.DbId)).FindAll(&entities)
	if err != nil {
		log.Error("Npc", "LoadPokemon", "Failed to load pokemon from database. %v", err.Error())
		return false, err.Error()
	}
	
	for _, entity := range(entities) {
		pokemon := NewNpcPokemon()
		pokemon.IsNew = false
		pokemon.DbId = int64(entity.IdnpcPokemon)
		pokemon.pokId = entity.Idpokemon
		pokemon.Hp = entity.IvHp
		pokemon.Att = entity.IvAttack
		pokemon.Att_spec = entity.IvAttackSpec
		pokemon.Def = entity.IvDefence
		pokemon.Def_spec = entity.IvDefenceSpec
		pokemon.Speed = entity.IvSpeed
		pokemon.Gender = entity.Gender
		pokemon.Held_item = entity.HeldItem
		pokemon.Name = "Pokemon Name"

		m.Pokemons[pokemon.DbId] = pokemon
	}
	
	return true, ""
}

func (m *Npc) SetEvents(_events string, _eventInitId int) {
	m.Events = _events 
	m.EventInitId = _eventInitId
	
	m.IsModified = true
}

func (m *Npc) SetName(_name string) {
	m.Name = _name
	m.IsModified = true
}

func (m *Npc) SetPositionByCoordinates(_x, _y, _z int) {
	m.Position.X = _x
	m.Position.Y = _y
	m.Position.Z = _z
	m.IsModified = true
}

func (m *Npc) AddPokemon(_pokemon *NpcPokemon) bool {
	_pokemon.NpcId = m.DbId
	return _pokemon.Save()
}

func (m *Npc) DeletePokemon(_pokemon *NpcPokemon) bool {
	if _, ok := m.Pokemons[_pokemon.DbId]; ok {
		return _pokemon.Delete()
	}
	
	return false
}

func (m *Npc) SetOutfitPart(_part pulogic.OutfitPart, _key int) {
	m.Outfit.data[_part] = _key
	m.IsModified = true
}

func (m *Npc) GetOutfitPart(_part pulogic.OutfitPart) int {
	return m.Outfit.data[_part]
}

func (m *Npc) Save() bool {
	// NPC
	npcEntity := models.Npc{ Idnpc: int(m.DbId),
							 Name: m.Name,
					  		 Position: m.Position.Hash() }
	if err := g_orm.Save(&npcEntity); err == nil {
		if m.IsNew {
			m.DbId = int64(npcEntity.Idnpc)
		}
	} else {
		return false
	}
	
	// OUTFIT
	outfitEntity := models.NpcOutfit { Idnpc: int(m.DbId),
									   Head: m.GetOutfitPart(pulogic.OUTFIT_HEAD),
									   Nek: m.GetOutfitPart(pulogic.OUTFIT_NEK),
									   Upper: m.GetOutfitPart(pulogic.OUTFIT_UPPER),
									   Lower: m.GetOutfitPart(pulogic.OUTFIT_LOWER),
									   Feet: m.GetOutfitPart(pulogic.OUTFIT_FEET) }
	if err := g_orm.Save(&outfitEntity); err != nil {
		return false
	}
	
	// EVENTS
	if len(m.Events) > 0 {
		eventEntity := models.NpcEvents { Idnpc: int(m.DbId),
										  Event: m.Events,
										  Initid: m.EventInitId }
		if err := g_orm.Save(&eventEntity); err != nil {
			return false
		}
	}
	
	// POKEMON
	if len(m.Pokemons) > 0 {
		for index, pokemon := range(m.Pokemons) {
			pokemon.Save()
			
			if pokemon.IsRemoved {
				delete(m.Pokemons, index) 
			}
		}
	}
	
	m.IsNew = false
	m.IsModified = false

	return true
}

func (m *Npc) Delete() bool {
	entity := models.Npc { Idnpc: int(m.DbId) }
	if _, err := g_orm.Delete(&entity); err != nil {
		log.Error("Npc", "Delete", err.Error())
		return false
	}
	
	// Remove from npc map
	delete(g_npc.Npcs, m.DbId)

	return true
}