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
	"crypto/sha1"
	"fmt"
	"hash"
	"strings"
)

func CheckAccountInfo(_username string, _password string) bool {
	//_username = g_db.Escape(_username)
	//_password = g_db.Escape(_password)

	var queryString string = "SELECT password, password_salt FROM player WHERE name='" + _username + "'"
	if err := g_db.Query(queryString); err != nil {
		println("Error - CheckAccountInfo: " + err.Error())
		return false
	}

	result, err := g_db.UseResult()
	if err != nil {
		println("Error - CheckAccountInfo: " + err.Error())
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

	var sha1Hash string = strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))
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
	idPlayer := DBGetInt(row["idplayer"])
	name := DBGetString(row["name"])
	result.Free()

	value, found := g_game.GetPlayerByName(name)
	if found {
		p = value
		ret = true
	} else {
		p = NewPlayer(name)
		p.dbid = idPlayer
		ret = p.loadPlayerInfo()
	}

	return
}
