/*Pokemon Universe MMORPG
Copyright (C) 2010 the Pokemon Universe Authors

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.*/
package main

import (
	"time"
	
	"npclib"
	"putools/log"
	puh "puhelper"
)

type NpcManager struct {
	npcLib			*npclib.NpcLib
	npcList			map[int]*Npc
	
	autoWalkEnable	bool
}

func NewNpcManager() *NpcManager {
	npcManager := NpcManager {}
	npcManager.npcLib = npclib.NewNpcLib()
	npcManager.autoWalkEnable = false
	
	return &npcManager
}

func (n *NpcManager) Load() bool {
	// Fetch NPCs from database
	var query string = "SELECT idnpc, name, script_name, position FROM npc"
	result, err := puh.DBQuerySelect(query)
	if err != nil {
		return false
	}
	
	defer puh.DBFree()
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
	
	n.StartAutoWalkNpc()
	
	return true
}

func (n *NpcManager) StartAutoWalkNpc() {
	logger.Println("Starting NPC ticker...")
	if !n.autoWalkEnable {
		n.autoWalkEnable = true
		go n.autoWalkRoutine()
	}
}

func (n *NpcManager) autoWalkRoutine() {
	for _, npc := range(n.npcList) {
		npc.AutoWalk()
	}
	
	go func() {
		time.Sleep(time.Second)
		
		if n.autoWalkEnable {
			n.autoWalkRoutine()
		}
	}()
}

func (n *NpcManager) StopAutoWalkNpc() {
	logger.Println("Stopping NPC ticker...")
	if n.autoWalkEnable {
		n.autoWalkEnable = false
	}
}