package main

import (
	"pu_npclib"
)

type Npc struct {
	Creature
	
	dbid		int
	script_name	string
	script 		pu_npclib.NpcInteractionInterface
}

func NewNpc() *Npc {
	n := Npc{}
	n.uid = GenerateUniqueID()
	n.Outfit = NewOutfit()
	n.moveSpeed = 280
	n.VisibleCreatures = make(CreatureList)
	n.script = nil
	
	return &n
}

func (n *Npc) Load(_data []interface{}) bool {
	id := DBGetInt(_data[0])
	name := DBGetString(_data[1])
	script_name := DBGetString(_data[2])
	position := _data[3].(int64)

	n.dbid = id
	n.name = name
	n.script_name = script_name
	
	tile, ok := g_map.GetTile(position)
	if !ok {
		g_logger.Printf("[Error] Could not load position info for npc %s (%d)", n.name, n.dbid)
		return false
	}
	n.Position = tile
	
	return true
}

func (n *Npc) GetType() int {
	return CTYPE_NPC
}

func (n *Npc) OnCreatureMove(_creature ICreature, _from *Tile, _to *Tile, _teleport bool) {
	// Check if _creature is a Player, otherwise return
	if _creature.GetType() != CTYPE_PLAYER {
		return
	}

	canSeeFromTile := CanSeePosition(n.GetPosition(), _from.Position)
	canSeeToTile := CanSeePosition(n.GetPosition(), _to.Position)

	if canSeeFromTile && !canSeeToTile { // Leaving viewport
		n.RemoveVisibleCreature(_creature)
		_creature.RemoveVisibleCreature(n)
	} else if canSeeToTile && !canSeeFromTile { // Entering viewport
		n.AddVisibleCreature(_creature)
		_creature.RemoveVisibleCreature(n)
	} else { // Moving inside viewport
		n.AddVisibleCreature(_creature)
		_creature.AddVisibleCreature(n)
	}
}

func (n *Npc) OnCreatureTurn(_creature ICreature) {
	// Check if _creature is a Player, otherwise return
	if _creature.GetType() != CTYPE_PLAYER {
		return
	}
}

func (n *Npc) OnCreatureAppear(_creature ICreature, _isLogin bool) {
	// Check if _creature is a Player, otherwise return
	if _creature.GetType() != CTYPE_PLAYER {
		return
	}
	
	canSeeCreature := CanSeeCreature(n, _creature)
	if !canSeeCreature {
		return
	}

	// We're checking inside the AddVisibleCreature method so no need to check here
	n.AddVisibleCreature(_creature)
	_creature.AddVisibleCreature(n)
}

func (n *Npc) OnCreatureDisappear(_creature ICreature, _isLogout bool) {
	// Check if _creature is a Player, otherwise return
	if _creature.GetType() != CTYPE_PLAYER {
		return
	}
	
	// TODO: Have to do something here with _isLogout

	n.RemoveVisibleCreature(_creature)
}

func (n *Npc) AddVisibleCreature(_creature ICreature) {
	// Check if _creature is a Player, otherwise return
	if _creature.GetType() != CTYPE_PLAYER {
		return
	}
	
	if _, found := n.VisibleCreatures[_creature.GetUID()]; !found {
		n.VisibleCreatures[_creature.GetUID()] = _creature
	}
}

func (n *Npc) RemoveVisibleCreature(_creature ICreature) {
	// Check if _creature is a Player, otherwise return
	if _creature.GetType() != CTYPE_PLAYER {
		return
	}

	// No need to check if the key actually exists because Go is awesome
	// http://golang.org/doc/effective_go.html#maps
	delete(n.VisibleCreatures, _creature.GetUID())
}

// -------------------------------------------------------- //

func (n *Npc) SelfSay(_message string) {

}

func (n *Npc) OnDialogueAnswer(_cid uint64, _answer int) {
	if n.script != nil {
		n.script.OnAnswer(_cid, _answer)
	}
}