package main

import (
	"github.com/astaxie/beedb"
	
	pos "nonamelib/pos"
	pul "pulogic"
	"pulogic/models"
)


type NpcList struct {
	Npcs	map[int64]*Npc
}

func NewNpcList() *NpcList{
	return &NpcList { Npcs: make(map[int64]*Npc) }
}

func (m *NpcList) LoadNpcList() (bool, string) {
	var entities []models.NpcJoinOutfitJoinEvent
	err := g_orm.SetTable("npc").Join("INNER", "npc_outfit", "npc_outfit.Idnpc = npc.Idnpc").Join(" INNER", "npc_events", "npc_events.Idnpc = npc.Idnpc").FindAll(&entities)
	if err != nil {
		return false, err.Error()
	}
	
	for _, entity := range(entities) {
		// Create new NPC object with data from database
		npc := NewNpc()
		npc.IsNew = false
		npc.DbId = int64(entity.Idnpc)
		npc.Name = entity.Name
		
		npc.SetOutfitPart(pul.OUTFIT_HEAD, entity.Head)
		npc.SetOutfitPart(pul.OUTFIT_NEK, entity.Nek)
		npc.SetOutfitPart(pul.OUTFIT_UPPER, entity.Upper)
		npc.SetOutfitPart(pul.OUTFIT_LOWER, entity.Lower)
		npc.SetOutfitPart(pul.OUTFIT_FEET, entity.Feet)
		
		positionHash := entity.Position
		npc.Position = pos.NewPositionFromHash(positionHash)
		npc.Events = entity.Event
		npc.EventInitId = entity.Initid
		
		// Load NPC pokemon
		npc.LoadPokemon()
		
		// Save in Npc map
		m.Npcs[npc.DbId] = npc
	}
	
	return true,  ""
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