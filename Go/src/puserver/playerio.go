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
	"fmt"
	
	"putools/log"
	pul "pulogic"
	pkmn "pulogic/pokemon"
	puh "puhelper"
)

func (p *Player) LoadData() bool {
	// Load player info
	logger.Println("Loading player info")
	if !p.loadPlayerInfo() {
		return false
	}

	// Load all pokemon player has
	logger.Println("Loading player pokemon")
	if !p.loadPokemon() {
		return false
	}
	
	// Load moves for each pokemon
	for index, pokemon := range p.PokemonParty.Party {
		if pokemon != nil {
			fmt.Printf("Load pokemon moves for: %d - %s\n\r", index, pokemon.GetNickname())
			pokemon.LoadMoves()
		}
	}
	
	// Load player storage items
	logger.Println("Loading player items")
	if !p.loadItems() {
		return false
	}
	
	// Load player backpack
	logger.Println("Loading player backpack")
	if !p.loadBackpack() {
		return false
	}
	
	// Load friends list
	logger.Println("Loading player friends")
	if !p.loadFriends() {
		return false
	}

	return true
}

func (p *Player) loadPlayerInfo() bool {
	var query string = "SELECT p.idplayer, p.name, p.position, p.movement, p.idpokecenter, p.money, p.idlocation," +
		" g.group_idgroup, o.head, o.nek, o.upper, o.lower, o.feet FROM player `p`" +
		" INNER JOIN player_outfit `o` ON o.idplayer = p.idplayer" +
		" INNER JOIN player_group `g` ON g.player_idplayer = p.idplayer" +
		" WHERE p.idplayer=%d"
	result, err := puh.DBQuerySelect(fmt.Sprintf(query, p.dbid))
	if err != nil {
		return false
	}

	defer puh.DBFree()
	row := result.FetchRow()
	if row == nil {
		logger.Printf("[Error] No player data for %s (DB ID: %d)\n", p.name, p.dbid)
		return false
	}

	p.dbid = puh.DBGetInt(row[0])
	p.name = puh.DBGetString(row[1])
	tile, ok := g_map.GetTile(puh.DBGetInt64(row[2]))
	if !ok {
		logger.Printf("[Warning] Could not load position info for player %s (%d)\n", p.name, p.dbid)
		//tile, _ = g_map.GetTileFrom(-510, -236, 0)
		tile, _ = g_map.GetTileFrom(0, 0, 1)
		if tile == nil {
			logger.Println("[Error] Could not load default position")
			return false
		}
	}
	p.Position = tile
	p.SetDirection(DIR_SOUTH)
	p.Movement = puh.DBGetInt(row[3])
	// TODO: p.LastPokeCenter = row[4].(int)
	p.Money = puh.DBGetInt(row[5])
	location, ok := g_game.Locations.GetLocation(puh.DBGetInt(row[6]))
	if !ok {
		logger.Printf("[Error] Could not load location info for player %s (%d)\n", p.name, p.dbid)
		return false
	}
	p.Location = location 

	// Group/Right stuff : row[7].(int)

	p.SetOutfitKey(pul.OUTFIT_HEAD, puh.DBGetInt(row[8]))
	p.SetOutfitKey(pul.OUTFIT_NEK, puh.DBGetInt(row[9]))
	p.SetOutfitKey(pul.OUTFIT_UPPER, puh.DBGetInt(row[10]))
	p.SetOutfitKey(pul.OUTFIT_LOWER, puh.DBGetInt(row[11]))
	p.SetOutfitKey(pul.OUTFIT_FEET, puh.DBGetInt(row[12]))

	return true
}

func (p *Player) loadPokemon() bool {
	var query string = "SELECT idpokemon, nickname, bound, experience, iv_hp, iv_attack, iv_attack_spec, iv_defence, iv_defence_spec," +
		" iv_speed, happiness, gender, in_party, party_slot, idplayer_pokemon, shiny, idability, damaged_hp FROM player_pokemon WHERE idplayer='%d' AND in_party=1"
	result, err := puh.DBQuerySelect(fmt.Sprintf(query, p.dbid))
	if err != nil {
		return false
	}

	logger.Println("Loading player pokemon..")
	defer puh.DBFree()	
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}

		pokemon := pkmn.NewPlayerPokemon(p.dbid)
		pokemon.IdDb = puh.DBGetInt(row[14])
		pokemonId := puh.DBGetInt(row[0])
		pokemon.Base = pkmn.GetInstance().GetPokemon(pokemonId)
		pokemon.Nickname = puh.DBGetString(row[1])
		pokemon.IsBound = puh.DBGetInt(row[2])
		pokemon.Experience = puh.DBGetFloat64(row[3])
		pokemon.Stats[0] = puh.DBGetInt(row[4]) // HP
		pokemon.Stats[1] = puh.DBGetInt(row[5]) // Attack
		pokemon.Stats[2] = puh.DBGetInt(row[7]) // Defence
		pokemon.Stats[3] = puh.DBGetInt(row[6]) // Spec Attack
		pokemon.Stats[4] = puh.DBGetInt(row[8]) // Spec Defence
		pokemon.Stats[5] = puh.DBGetInt(row[9]) // Speed
		pokemon.Happiness = puh.DBGetInt(row[10])
		pokemon.Gender = puh.DBGetInt(row[11])
		pokemon.InParty = puh.DBGetInt(row[12])
		pokemon.Slot = puh.DBGetInt(row[13])
		pokemon.IsShiny = puh.DBGetInt(row[15])
		abilityId := puh.DBGetInt(row[16])
		pokemon.DamagedHp = puh.DBGetInt(row[17])
		
		pokemon.Ability = pkmn.GetInstance().GetAbilityById(abilityId)
		if pokemon.Ability == nil {
			logger.Printf("[Warning] Pokemon (%d) has an invalid abilityId (%d)\n", pokemon.IdDb, abilityId)
			pokemon.Ability = pkmn.GetInstance().GetAbilityById(96)
		}

		// Add to party if needed
		if pokemon.InParty == 1 {
			p.PokemonParty.AddSlot(pokemon, pokemon.Slot)
		}
	}

	return true
}

func (p *Player) loadItems() bool {
	var query string = "SELECT idplayer_items, iditem, count, slot FROM player_items WHERE idplayer=%d"
	result, err := puh.DBQuerySelect(fmt.Sprintf(query, p.dbid))
	if err != nil {
		return false
	}
	
	defer puh.DBFree()
	
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		dbid := puh.DBGetInt64(row[0])
		itemId := puh.DBGetInt64(row[1])
		count := puh.DBGetInt(row[2])
		slot := puh.DBGetInt(row[3])
		
		item, _ := g_game.Items.GetItemByItemId(itemId)
		newItem := item.Clone()
		newItem.DbId = dbid
		newItem.SetCount(count)
		
		p.Storage.AddItemObject(newItem, slot)
	}
	
	return true
}

func (p *Player) loadBackpack() bool {
	var query string = "SELECT idplayer_backpack, iditem, count, slot FROM player_backpack WHERE idplayer=%d"
	result, err := puh.DBQuerySelect(fmt.Sprintf(query, p.dbid))
	if err != nil {
		return false
	}
	
	defer puh.DBFree()
	
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		dbid := puh.DBGetInt64(row[0])
		itemId := puh.DBGetInt64(row[1])
		count := puh.DBGetInt(row[2])
		slot := puh.DBGetInt(row[3])
		
		item, _ := g_game.Items.GetItemByItemId(itemId)
		newItem := item.Clone()
		newItem.DbId = dbid
		newItem.SetCount(count)
		
		p.Backpack.AddItemObject(newItem, slot)
	}
	
	return true	
}

func (p *Player) loadFriends() bool {
	var query string = "SELECT idplayer_friends, friend_name FROM player_friends WHERE idplayer=%d"
	result, err := puh.DBQuerySelect(fmt.Sprintf(query, p.dbid))
	if err != nil {
		return false
	}
	
	defer puh.DBFree()
	
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		dbid := puh.DBGetInt64(row[0])
		name := puh.DBGetString(row[1])
		
		_, online := g_game.GetPlayerByName(name)
		
		friend := &Friend { DbId: dbid,
							Name: name,
							Online: online }
		p.Friends[name] = friend
	}
		
	return true
}

func (p *Player) loadQuestsProgress() bool {
	var query string = "SELECT idplayer_quests, idquest, status, created, finished FROM player_quests WHERE idplayer=%d"
	result, err := puh.DBQuerySelect(fmt.Sprintf(query, p.dbid))
	if err != nil {
		return false
	}
	
	defer puh.DBFree()
	for  {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		dbid := puh.DBGetInt64(row[0])
		questid := puh.DBGetInt64(row[1])
		status := puh.DBGetInt(row[2])
		created := puh.DBGetInt64(row[3]) // Unix seconds
		finished := puh.DBGetInt64(row[4]) // Unix seconds
		
		if playerQuest := NewPlayerQuestExt(dbid, questid, status, created, finished); playerQuest != nil {
			playerQuest.IsNew = false
			p.Quests[questid] = playerQuest
		}
	}
	
	return true
}

// --------------------- SAVING ----------------------------//

func (p *Player) SaveData() {
	p.savePlayerInfo()
	p.savePokemon()
	p.saveItems()
	p.saveBackpack()
	p.saveFriends()
}

func (p *Player) savePlayerInfo() {
	var query string 
	query = fmt.Sprintf("UPDATE player SET position=%d, movement=%d, money=%d, idlocation=%d WHERE idplayer=%d", 
						p.GetPosition().Hash(), 
						p.GetMovement(), 
						p.GetMoney(),
						p.GetTile().GetLocation().GetId(),
						p.dbid)
	puh.DBQuery(query)
	
	// Save outfit
	query = fmt.Sprintf("UPDATE player_outfit SET head=%d, nek=%d, upper=%d, lower=%d, feet=%d WHERE idplayer=%d",
						p.GetOutfit().GetOutfitKey(pul.OUTFIT_HEAD),
						p.GetOutfit().GetOutfitKey(pul.OUTFIT_NEK),
						p.GetOutfit().GetOutfitKey(pul.OUTFIT_UPPER),
						p.GetOutfit().GetOutfitKey(pul.OUTFIT_LOWER),
						p.GetOutfit().GetOutfitKey(pul.OUTFIT_FEET),
						p.dbid)
	puh.DBQuery(query)
}

func (p *Player) savePokemon() {
	for index, pokemon := range p.PokemonParty.Party {
		if pokemon != nil {
			// Save pokemon info
			saveQuery := "UPDATE player_pokemon SET "
			saveQuery += "nickname='%v', bound=%d, experience=%d, iv_hp=%d, iv_attack=%d, iv_attack_spec=%d, iv_defence=%d, iv_defence_spec=%d, iv_speed=%d, happiness=%d, in_party=%d, party_slot=%d, held_item=%d "
			saveQuery += "WHERE idplayer_pokemon=%d"
			puh.DBQuery(fmt.Sprintf(saveQuery,
								pokemon.Nickname,
								pokemon.IsBound,
								int(pokemon.Experience),
								pokemon.Stats[0],
								pokemon.Stats[1],
								pokemon.Stats[3],
								pokemon.Stats[2],
								pokemon.Stats[4],
								pokemon.Stats[5],
								pokemon.Happiness,
								pokemon.InParty,
								index,
								0,
								pokemon.IdDb))
			
			// Save moves
			pokemon.SaveMoves()
		}
	}
}

func (p *Player) saveItems() {
    puh.DBStartTransaction()
	puh.DBCon.SetAutoCommit(false)
        
    // Remove all items from database
	if err := puh.DBQueryNoLock(fmt.Sprintf("DELETE FROM player_items WHERE idplayer=%d", p.dbid)); err != nil {
		logger.Println("Failed to save player items!")
		puh.DBRollback()
	}
        
	// Insert all items in database
	var err error = nil
	for _, item := range(p.Storage.Items) {
		if err = puh.DBQueryNoLock(fmt.Sprintf("INSERT INTO player_items (idplayer, iditem, count, slot) VALUES ('%d','%d','%d','%d')", p.dbid, item.DbId, item.Count, item.Slot)); err != nil {
        	break
    	}
	}
        
	puh.DBCon.SetAutoCommit(true)
        
	if err == nil {
    	puh.DBCommit()
	} else {
		puh.DBRollback()
	}
}

func (p *Player) saveBackpack() {
    puh.DBCon.SetAutoCommit(false)
    puh.DBStartTransaction()
        
	// Remove all items from database
	if err := puh.DBQueryNoLock(fmt.Sprintf("DELETE FROM player_backpack WHERE idplayer=%d", p.dbid)); err != nil {
		logger.Println("Failed to save player items!")
		puh.DBRollback()
    }
        
    // Insert all items in database
    var err error = nil
    for _, item := range(p.Backpack.Items) {
        if err = puh.DBQueryNoLock(fmt.Sprintf("INSERT INTO player_backpack (idplayer, iditem, count, slot) VALUES ('%d','%d','%d','%d')", p.dbid, item.DbId, item.Count, item.Slot)); err != nil {
    		break
    	}
    }
        
    puh.DBCon.SetAutoCommit(true)
        
    if err == nil {
            puh.DBCommit()
    } else {
            puh.DBRollback()
	}
}

func (p *Player) saveFriends() {
	for _, friend := range(p.Friends) {	
		if friend.DbId == 0 && !friend.IsRemoved { // New friend
			puh.DBQuery(fmt.Sprintf("INSERT INTO player_friends (idplayer, friend_name) VALUES (%d, '%s')", p.dbid, friend.Name))
		} else if friend.DbId > 0 && friend.IsRemoved { // Friend removed		
			puh.DBQuery(fmt.Sprintf("DELETE FROM player_friends WHERE idplayer_friends = %d", friend.DbId))
		}
	}
}

func (p *Player) saveQuestProgress() {
	for _, quest := range(p.Quests) {
		var query string = ""
		if quest.IsNew {
			tmpQuery := "INSERT INTO player_quests (idplayer, idquest, status, created, finished) VALUES (%d, %d, %d, %d, %d)"
			query = fmt.Sprintf(tmpQuery, p.dbid, quest.Quest.Dbid, quest.Status, quest.Created.Unix(), quest.Finished.Unix())
		} else if quest.IsModified {
			tmpQuery := "UPDATE player_quests SET status=%d, finished=%d WHERE idplayer_quests=%d"
			query = fmt.Sprintf(tmpQuery, quest.Status, quest.Finished.Unix(), quest.Dbid)
		}
		
		if len(query) > 0 {
			puh.DBQuery(query)
		}
	}
}