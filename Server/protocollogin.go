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
)

const (
	LOGINSTATUS_IDLE            = 0
	LOGINSTATUS_WRONGACCOUNT    = 1
	LOGINSTATUS_SERVERERROR     = 2
	LOGINSTATUS_DATABASEERROR   = 3
	LOGINSTATUS_ALREADYLOGGEDIN = 4
	LOGINSTATUS_READY           = 5
	LOGINSTATUS_CHARBANNED      = 6
	LOGINSTATUS_SERVERCLOSED    = 7
	LOGINSTATUS_WRONGVERSION    = 8
	LOGINSTATUS_FAILPROFILELOAD = 9
)

func CheckAccountInfo(_username string, _password string) bool {
	//_username = g_db.Escape(_username)
	//_password = g_db.Escape(_password)

	var queryString string = "SELECT password, password_salt FROM player WHERE name='" + _username + "'"
	if err := g_db.Query(queryString); err != nil {
		println("Error - CheckAccountInfo: " + err.String())
		return false
	}

	result, err := g_db.UseResult()
	if err != nil {
		println("Error - CheckAccountInfo: " + err.String())
		return false
	}

	row := result.FetchMap()
	defer result.Free()
	if row == nil {
		return false
	}

	password := row["password"].(string)
	salt := row["password_salt"].(string)
	_password = _password + salt

	passCheck := PasswordTest(_password, password)
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

	var queryString string = "SELECT idplayer, name FROM player WHERE name='%v'"
	if err := g_db.Query(fmt.Sprintf(queryString, _username)); err != nil {
		if IS_DEBUG {
			g_logger.Printf("[DEBUG] LoadPlayerProfile: %v\n\r", err)
		}
		return
	}

	result, err := g_db.UseResult()
	if err != nil {
		return
	}

	row := result.FetchMap()
	if row == nil {
		return
	}
	idPlayer := row["idplayer"].(int)
	name := row["name"].(string)
	result.Free()

	value, found := g_game.GetPlayerByName(name)
	if found {
		p = value.(*Player)
		ret = true
	} else {
		p = NewPlayer(name)
		p.Id = idPlayer

		queryString = "SELECT p.`position`, p.`movement`, p.`money`, p.`idlocation`, "
		queryString += "po.`head`, po.`nek`, po.`upper`, po.`lower`, po.`feet`, pc.`position` AS `pc_position` "
		queryString += "FROM `player` as p "
		queryString += "JOIN `player_outfit` AS po ON po.`idplayer`=p.`idplayer` "
		queryString += "JOIN `pokecenter` AS pc ON pc.`idpokecenter`=p.`idpokecenter` "
		queryString += "WHERE p.`idplayer` = '%d'"

		if err := g_db.Query(fmt.Sprintf(queryString, idPlayer)); err != nil {
			if IS_DEBUG {
				g_logger.Printf("[DEBUG] LoadPlayerProfile 1: %v\n\r", err)
			}
			return
		}
		result, err := g_db.UseResult()
		if err != nil {
			if IS_DEBUG {
				g_logger.Printf("[DEBUG] LoadPlayerProfile 2: %v\n\r", err)
			}
			return
		}

		defer result.Free()
		row := result.FetchMap()
		if row == nil {
			if IS_DEBUG {
				g_logger.Printf("[DEBUG] LoadPlayerProfile 3: no row for this id %d\n\r", string(idPlayer))
			}
			return
		}

		positionHash := row["position"].(int64)
		movement := row["movement"].(int)
		money := row["money"].(int)
		idlocation := row["idlocation"].(int)
		pcposition := row["pc_position"].(int64)

		var ok bool
		p.Position, ok = g_map.GetTile(positionHash)
		if !ok {
			if IS_DEBUG {
				g_logger.Println("[DEBUG] LoadPlayerProfile: Could not find the tile for this position", err)
			}
			return
		}
		p.Location, ok = g_game.Locations.GetLocation(idlocation)
		if !ok {
			if IS_DEBUG {
				g_logger.Println("[DEBUG] LoadPlayerProfile: Could not find the location", err)
			}
			p.Location = p.Position.Location
		}
		p.LastPokeCenter, ok = g_map.GetTile(pcposition)
		if !ok {
			p.LastPokeCenter, _ = g_map.GetTile(p.Position.Location.PokeCenter.Hash())
		}

		p.Movement = movement
		p.SetMoney(money)

		// Load outfit
		outfitHead := row["head"].(int)
		outfitNek := row["nek"].(int)
		outfitUpper := row["upper"].(int)
		outfitLower := row["lower"].(int)
		outfitFeet := row["feet"].(int)

		p.SetOutfitKey(OUTFIT_HEAD, outfitHead)
		p.SetOutfitKey(OUTFIT_NEK, outfitNek)
		p.SetOutfitKey(OUTFIT_UPPER, outfitUpper)
		p.SetOutfitKey(OUTFIT_LOWER, outfitLower)
		p.SetOutfitKey(OUTFIT_FEET, outfitFeet)

		println("- Loaded all data, adding to Game")
		// Add player object to THE GAME (you've just lost it :3)
		g_game.AddCreature(p)
		ret = true
	}

	return
}
