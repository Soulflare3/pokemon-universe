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
	"hash"
	"crypto/sha1"
	"strings"
	"os"
	"db"
)

func CheckAccountInfo(_username string, _password string) bool {
	//_username = g_db.Escape(_username)
	//_password = g_db.Escape(_password)
	
	var result db.ResultSet
	var err os.Error
	var queryString string = "SELECT password, password_salt FROM player WHERE name='" + _username + "'"
	if result, err = g_db.StoreQuery(queryString); err != nil {
		return false
	}
	
	if !result.Next() {
		return false
	}
	
	password	:= result.GetDataString("password")
	salt		:= result.GetDataString("password_salt")
	_password = _password + salt
	
	var passCheck bool = PasswordTest(_password, password)
	return passCheck
}

func PasswordTest(_plain string, _hash string) bool {
	var h hash.Hash = sha1.New()
	h.Write([]byte(_plain))
	
	var sha1Hash string = strings.ToUpper(fmt.Sprintf("%x", h.Sum()))
	var original string = strings.ToUpper(_hash)
	
	return (sha1Hash == original)
}

func LoadPlayerProfile(_username string) (ret bool, p *Player) {
	p = nil
	ret = false
	//_username = g_db.Escape(_username)
	
	var result db.ResultSet
	var err os.Error
	var queryString string = "SELECT idplayer, name FROM player WHERE name='" + _username +"'"
	if result, err = g_db.StoreQuery(queryString); err != nil {
		if IS_DEBUG {
			g_logger.Printf("[DEBUG] LoadPlayerProfile: %v\n\r", err)
		}
		return
	}
	if !result.Next() {
		return
	}
	
	idPlayer := result.GetDataInt("idplayer")
	name	 := result.GetDataString("name")
	
	if g_game.GetPlayerByName(name) == nil {
		p := NewPlayer(name)
		p.Id = int(idPlayer)
		
		queryString = "SELECT p.`position`, p.`movement`, p.`money`, p.`idlocation`, "
		queryString += "po.`head`, po.`nek`, po.`upper`, po.`lower`, po.`feet`, pc.`position` AS `pc_position` "
		queryString += "FROM `player` as p "
		queryString += "JOIN `player_outfit` AS po ON po.`idplayer`=p.`idplayer` "
		queryString += "JOIN `pokecenter` AS pc ON pc.`idpokecenter`=p.`idpokecenter` "
		queryString += "WHERE p.`idplayer` = '" + string(idPlayer) + "'"
		
		if result, err = g_db.StoreQuery(queryString); err != nil {
			if IS_DEBUG {
				g_logger.Printf("[DEBUG] LoadPlayerProfile: %v\n\r", err)
			}
			return
		}
		if !result.Next() {
			return
		}
				
		positionHash := result.GetDataLong("position")
		movement	:= result.GetDataInt("movement")
		money		:= result.GetDataInt("money")
		idlocation	:= result.GetDataInt("idlocation")
		pcposition	:= result.GetDataLong("pc_position")
		
		var ok bool
		p.Position, ok = g_map.GetTile(positionHash)
		if !ok {
			return
		}
		p.Location, ok = g_game.Locations.GetLocation(int32(idlocation))
		if !ok {
			p.Location = p.Position.Location
		}
		p.LastPokeCenter, ok = g_map.GetTile(pcposition)
		if !ok {
			p.LastPokeCenter, _ = g_map.GetTile(p.Position.Location.PokeCenter.Hash())
		}
		
		p.Movement = uint16(movement)
		p.SetMoney(int32(money))
		
		// Load outfit
		outfitHead	:= result.GetDataInt("head")
		outfitNek	:= result.GetDataInt("nek")
		outfitUpper	:= result.GetDataInt("upper")
		outfitLower	:= result.GetDataInt("lower")
		outfitFeet	:= result.GetDataInt("feet")
		
		p.SetOutfitKey(OUTFIT_HEAD, int(outfitHead))
		p.SetOutfitKey(OUTFIT_NEK, int(outfitNek))
		p.SetOutfitKey(OUTFIT_UPPER, int(outfitUpper))
		p.SetOutfitKey(OUTFIT_LOWER, int(outfitLower))
		p.SetOutfitKey(OUTFIT_FEET, int(outfitFeet))
		
		// Add player object to THE GAME (you've just lost it :3)
		g_game.AddCreature(p)
		ret = true
	}

	return
}
