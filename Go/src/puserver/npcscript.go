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
	pnet "network"
)

type NpcScript struct {
	self *Npc
}

func NewNpcScript(_npc *Npc) *NpcScript {
	return &NpcScript{ self: _npc }
}

func (s *NpcScript) GetCreatureName(cid uint64) (name string) {
	player, ok := g_game.GetPlayerByGuid(cid)
	if ok {
		name = player.GetName()
	} else {
		name = "Unknown"
	}
	return
}

func (s *NpcScript) SelfSay(message string) {
	s.self.SelfSay(message)
}

// Dialogue
func (s *NpcScript) SendDialogue(_cid uint64, _title string, _options ...string) {
	if player, ok := g_game.GetPlayerByGuid(_cid); ok {
		if len(_options) > 0 {
			player.sendDialog(pnet.DIALOG_NPC, s.self.GetUID(), _title, _options)
		} else {
			player.sendDialog(pnet.DIALOG_NPCTEXT, s.self.GetUID(), _title, _options)
		}	
	}
}

func (s *NpcScript) HideDialogue(_cid uint64) {
	if player, ok := g_game.GetPlayerByGuid(_cid); ok {
		player.sendDialog(pnet.DIALOG_CLOSE, 0, "", nil)
	}
}

func (s *NpcScript) EndDialogue(_cid uint64) {
	if player, ok := g_game.GetPlayerByGuid(_cid); ok {
		s.self.RemoveInteractingPlayer(player)
		player.sendDialog(pnet.DIALOG_CLOSE, 0, "", nil)
	}
}

// Pokecenter
func (s *NpcScript) HealParty(cid uint64) {
	player, ok := g_game.GetPlayerByGuid(cid)
	if ok {
		player.HealParty()
	}
}

// Market
func (s *NpcScript) OpenShopWindow(_cid uint64) {
}

func (s *NpcScript) CloseShopWindow(_cid uint64) {
}

// Quest
func (s *NpcScript) GetQuestProgress(_cid uint64, _questId int) (status int) {
	status = 0
	if player, found := g_game.GetPlayerByGuid(_cid); found {
		status = player.GetQuestStatus(int64(_questId))
	}
	
	return
}

func (s *NpcScript) SetQuestProgress(_cid uint64, _questId int, _progress int) {
	if player, found := g_game.GetPlayerByGuid(_cid); found {
		player.SetQuestStatus(int64(_questId), _progress)
	}
}

// Items
func (s *NpcScript) AddItem(_cid uint64, _itemId int64, _amount int) (ret bool) {
	ret = false
	
	// Get creature
	if player, found := g_game.GetPlayerByGuid(_cid); found {
		// Get item
		if item, ok := g_game.Items.GetItemByItemId(_itemId); ok {
			// Find free slot
			slot := player.Backpack.GetFreeSlotForCategory(item.CategoryId)
			// Add to backpack
			ret = player.Backpack.AddItem(_itemId, _amount, slot)
		}
	}
	
	return
}

func (s *NpcScript) CheckItem(cid uint64, itemId int64, amount int) (ret bool) {
	ret = false
	
	// Get creature
	if player, found := g_game.GetPlayerByGuid(cid); found {
		if item := player.Backpack.GetItemFromIndex(itemId); item != nil {
			ret = (item.GetCount() >= amount)
		}
	}
	
	return
}

func (s *NpcScript) RemoveItem(cid uint64, itemId int64, amount int) (ret bool) {
	ret = false
	
	// Get creature
	if player, found := g_game.GetPlayerByGuid(cid); found {
		ret = player.Backpack.UpdateItemByIndex(itemId, amount)
	}
	
	return
}

// Golds
func (s *NpcScript) AddMoney(cid uint64, amount int) {
	player, ok := g_game.GetPlayerByGuid(cid)
	if ok {
		player.SetMoney(amount)
	}
}

func (s *NpcScript) CheckMoney(cid uint64, amount int) bool {
	player, ok := g_game.GetPlayerByGuid(cid)
	if ok && player.Money >= amount {
		return true
	}
	
	return false
}

func (s *NpcScript) RemoveMoney(cid uint64, amount int) {
	player, ok := g_game.GetPlayerByGuid(cid)
	if ok {
		player.SetMoney(-amount)
	}
}