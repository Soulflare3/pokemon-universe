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
	"os"
	"hash"
	"crypto/sha1"
	"strings"
	"mysql"
)

func CheckAccountInfo(_username string, _password string) bool {
	var err os.Error
	var res *mysql.MySQLResult
	var rows map[string]interface{}
	
	_username = g_db.Escape(_username)
	_password = g_db.Escape(_password)
	
	var queryString string = "SELECT password, password_salt FROM player WHERE name='" + _username + "'"
	if res, err = g_db.Query(queryString); err != nil {
		return false
	}
	
	if rows = res.FetchMap(); rows == nil {
		return false
	}
	
	password, _ := rows["password"].(string)
	salt, _ := rows["password_salt"].(string)
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

func LoadPlayerProfile(_username string) (bool, *Player) {
	var err os.Error
	var res *mysql.MySQLResult
	var rows map[string]interface{}
	
	_username = g_db.Escape(_username)
	
	var queryString string = "SELECT idplayer, name FROM player WHERE name='" + _username +"'"
	if res, err = g_db.Query(queryString); err != nil {
		if IS_DEBUG {
			g_logger.Printf("[DEBUG] LoadPlayerProfile: %v\n\r", err)
		}
		return false, nil
	}
	if rows = res.FetchMap(); rows == nil {
		return false, nil
	}
	
	idPlayer, _ := rows["idplayer"].(int)
	name, _ 	:= rows["name"].(string)
	
	var p *Player = g_game.GetPlayerByName(name)
	if p == nil {
		p = NewPlayer(name)
		p.Id = idPlayer
		
		queryString = "SELECT p.`position`, p.`movement`, p.`money`, p.`idlocation`, "
		queryString += "po.`head`, po.`nek`, po.`upper`, po.`lower`, po.`feet`, pc.`position` AS `pc_position` "
		queryString += "FROM `player` as p "
		queryString += "JOIN `player_outfit` AS po ON po.`idplayer`=p.`idplayer` "
		queryString += "JOIN `pokecenter` AS pc ON pc.`idpokecenter`=p.`idpokecenter` "
		queryString += "WHERE p.`idplayer` = '" + string(idPlayer) + "'"
		
		if res, err = g_db.Query(queryString); err != nil {
			if IS_DEBUG {
				g_logger.Printf("[DEBUG] LoadPlayerProfile: %v\n\r", err)
			}
			return false, nil
		}
		if rows = res.FetchMap(); rows == nil {
			return false, nil
		}
		
		positionHash, _ := rows["position"].(int64)
		movement, _		:= rows["movement"].(int)
		money,_			:= rows["money"].(int)
		idlocation, _	:= rows["idlocation"].(int32)
		pcposition, _	:= rows["pc_position"].(int64)
		
		var ok bool
		p.Position, ok = g_game.WorldMap.GetTile(positionHash)
		if !ok {
			return false, nil
		}
		p.Location, ok = g_game.Locations.GetLocation(idlocation)
		if !ok {
			p.Location = p.Position.Location
		}
		p.LastPokeCenter, ok = g_game.WorldMap.GetTile(pcposition)
		if !ok {
			p.LastPokeCenter, _ = g_game.WorldMap.GetTile(p.Position.Location.PokeCenter.Hash())
		}
		
		p.Movement	= movement
		p.SetMoney(money)
		
		// Load outfit
		outfitHead, _	:= rows["head"].(int)
		outfitNek, _	:= rows["nek"].(int)
		outfitUpper, _	:= rows["upper"].(int)
		outfitLower, _	:= rows["lower"].(int)
		outfitFeet, _	:= rows["feet"].(int)
		
		p.NewOutfit()
		p.SetOutfitKey(OUTFIT_HEAD, outfitHead)
		p.SetOutfitKey(OUTFIT_NEK, outfitNek)
		p.SetOutfitKey(OUTFIT_UPPER, outfitUpper)
		p.SetOutfitKey(OUTFIT_LOWER, outfitLower)
		p.SetOutfitKey(OUTFIT_FEET, outfitFeet)
		
		// Add player object to THE GAME (you've just lost it :3)
		g_game.AddPlayer(p)
	}

	return true, p
}
