package main

import (
	"pu_npclib"
)

type NpcManager struct {
	npcLib		*pu_npclib.NpcLib
	npcList		map[int]*Npc
}

func NewNpcManager() *NpcManager {
	npcManager := NpcManager {}
	npcManager.npcLib = pu_npclib.NewNpcLib()
	
	return &npcManager
}

func (n *NpcManager) Load() bool {
	// Fetch NPCs from database
	var query string = "SELECT idnpc, name, script_name, position FROM npc"
	result, err := DBQuerySelect(query)
	if err != nil {
		return false
	}
	
	defer result.Free()
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
				
		npc := NewNpc()
		if npc.Load(row) {
			if len(npc.script_name) > 0 {
				// Assign script to npc if exists
				npcScript, found := n.npcLib.GetNpcScript(npc.script_name)
				if found {
					npcScript.SetScriptInterface(NewNpcScript(npc))
					npc.script = npcScript
				}
			}
			
			n.npcList[npc.dbid] = npc
			g_game.AddCreature(npc)
		}
	}
	
	return true
}