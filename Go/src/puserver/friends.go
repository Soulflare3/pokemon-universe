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
	puh "puhelper"
)

type FriendList map[string]*Friend

type Friend struct {
	DbId	int64
	Name	string
	Online	bool
}

func (p *Player) AddFriend(_name string) {
	if _, found := p.GetFriend(_name); found {
		return
	}
	
	isOnline := false
	
	// Check if player is online
	_, found := g_game.GetPlayerByName(_name)
	if found {
		isOnline = true
	} else {
		query := fmt.Sprintf("SELECT idplayer FROM player WHERE name='%s'", _name)
		result, err := puh.DBQuerySelect(query)
		if err == nil {
			defer puh.DBFree()
			found = (result.RowCount() > 0)
		} else {
			found = false
		}
	}
	
	if found {
		friend := &Friend { DbId: 0,
							Name: _name,
							Online: isOnline }
		
		p.Friends[_name] = friend
		
		p.sendFriendUpdate(_name, isOnline)
	} else {
		p.sendCancelMessage(RET_PLAYERNOTFOUND)
	}
}

func (p *Player) RemoveFriend(_name string) {
	delete(p.Friends, _name)
	
	p.sendFriendRemove(_name)
}

func (p *Player) UpdateFriend(_name string, _online bool) {
	if friend, found := p.GetFriend(_name); found {
		friend.Online = _online
	}
	
	p.sendFriendUpdate(_name, _online)
}

func (p *Player) GetFriend(_name string) (friend *Friend, found bool) {
	friend, found = p.Friends[_name]
	
	return
}