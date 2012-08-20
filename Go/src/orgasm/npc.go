package main

import (
	"fmt"
	puh "puhelper"
	pos "putools/pos"
	pul "pulogic"
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
	
	query := fmt.Sprintf(QUERY_SELECT_NPC_POKEMON, m.DbId)
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
		
		pokemon := NewNpcPokemon()
		pokemon.IsNew = false
		pokemon.DbId = puh.DBGetInt64(row[0])
		pokemon.pokId = puh.DBGetInt(row[1])
		pokemon.Hp = puh.DBGetInt(row[2])
		pokemon.Att = puh.DBGetInt(row[3])		
		pokemon.Att_spec = puh.DBGetInt(row[4])
		pokemon.Def = puh.DBGetInt(row[5])
		pokemon.Def_spec = puh.DBGetInt(row[6])
		pokemon.Speed = puh.DBGetInt(row[7])
		pokemon.Gender = puh.DBGetInt(row[8])
		pokemon.Held_item = puh.DBGetInt(row[9])
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

func (m *Npc) SetOutfitPart(_part pul.OutfitPart, _key int) {
	m.Outfit.data[_part] = _key
	m.IsModified = true
}

func (m *Npc) GetOutfitPart(_part pul.OutfitPart) int {
	return m.Outfit.data[_part]
}

func (m *Npc) Save() bool {
	if m.IsNew {
		query := fmt.Sprintf(QUERY_INSERT_NPC, m.Name)
		if puh.DBQuery(query) == nil {
			m.DbId = int64(puh.DBGetLastInsertId())
		} else {
			return false
		}
		
		outfitQuery := fmt.Sprintf(QUERY_INSERT_NPC_OUTFIT, m.DbId, m.GetOutfitPart(pul.OUTFIT_HEAD), m.GetOutfitPart(pul.OUTFIT_NEK), m.GetOutfitPart(pul.OUTFIT_UPPER), m.GetOutfitPart(pul.OUTFIT_LOWER), m.GetOutfitPart(pul.OUTFIT_FEET))
		if puh.DBQuery(outfitQuery) != nil {
			return false
		}
		
		eventQuery := fmt.Sprintf(QUERY_INSERT_NPC_EVENT, m.DbId)
		if puh.DBQuery(eventQuery) != nil {
			return false
		}
	} else if m.IsModified {
		query := fmt.Sprintf(QUERY_UPDATE_NPC, m.Name, m.Position.Hash(), m.DbId)
		if puh.DBQuery(query) != nil {
			return false
		}
		
		outfitQuery := fmt.Sprintf(QUERY_UPDATE_NPC_OUTFIT, m.GetOutfitPart(pul.OUTFIT_HEAD), m.GetOutfitPart(pul.OUTFIT_NEK), m.GetOutfitPart(pul.OUTFIT_UPPER), m.GetOutfitPart(pul.OUTFIT_LOWER), m.GetOutfitPart(pul.OUTFIT_FEET), m.DbId)
		if puh.DBQuery(outfitQuery) != nil {
			return false
		}
		
		eventQuery := fmt.Sprintf(QUERY_UPDATE_NPC_EVENT, m.Events, m.EventInitId, m.DbId)
		if puh.DBQuery(eventQuery) != nil {
			return false
		}
	}
	
	for _, pokemon := range(m.Pokemons) {
		pokemon.Save()
	}
	
	m.IsNew = false
	m.IsModified = false

	return true
}

func (m *Npc) Delete() bool {
	// Delete outfit
	outfitQuery := fmt.Sprintf(QUERY_DELETE_NPC_OUTFIT, m.DbId)
	if puh.DBQuery(outfitQuery) != nil {
		return false
	}
	
	// Delete pokemon
	for _, pokemon := range(m.Pokemons) {
		if !pokemon.Delete() {
			return false
		}
	}
	
	// Delete events
	eventQuery := fmt.Sprintf(QUERY_DELETE_NPC_EVENT, m.DbId)
	if puh.DBQuery(eventQuery) != nil {
		return false
	}
	
	// Delete NPC
	npcQuery := fmt.Sprintf(QUERY_DELETE_NPC, m.DbId)
	if puh.DBQuery(npcQuery) != nil {
		return false
	}
	
	// Remove from npc map
	delete(g_npc.Npcs, m.DbId)

	return true
}