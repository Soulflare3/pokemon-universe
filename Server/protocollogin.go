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
	
	var sha1Hash string = strings.ToUpper(string(h.Sum()))
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
		
		g_game.AddPlayer(p)
	}

	return true, p
}
