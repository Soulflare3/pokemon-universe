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
	pul "pulogic"
)
	
func (p *Player) sendCancelMessage(_message int) {
	switch _message {		
		case RET_YOUAREEXHAUSTED:
			p.Conn.SendCancel("You are exhausted.")
		case RET_PLAYERNOTFOUND:
			p.Conn.SendCancel("Player with this name could not be found.")
		case RET_NOTPOSSIBLE:
			fallthrough
		default:
			p.Conn.SendCancel("Sorry, not possible.")
	}
}

func (p *Player) sendTextMessage(_mclass int, _message string) {
	if p.Conn != nil {
		// p.Conn.sendTextMessage(_mclass, _message)
	}
}

func (p *Player) sendMapData(_dir int) {
	if p.Conn != nil {
		p.Conn.SendMapData(_dir, p.GetPosition())
	}
}

func (p *Player) sendCreatureMove(_creature pul.ICreature, _from, _to pul.ITile) {
	if p.Conn != nil {
		p.Conn.SendCreatureMove(_creature, _from.(*Tile), _to.(*Tile))
	}
}

func (p *Player) sendCreatureTurn(_creature pul.ICreature) {
	if p.Conn != nil {
		p.Conn.SendCreatureTurn(_creature, p.GetDirection())
	}
}

func (p *Player) sendCreatureAdd(_creature pul.ICreature) {
	if p.Conn != nil {
		p.Conn.SendCreatureAdd(_creature)
	}
}

func (p *Player) sendCreatureRemove(_creature pul.ICreature) {
	if p.Conn != nil {
		p.Conn.SendCreatureRemove(_creature)
	}
}

func (p *Player) sendPlayerWarp() {
	if p.Conn != nil {
		p.Conn.SendPlayerWarp(p.GetPosition())
	}
}

func (p *Player) sendCreatureSay(_creature pul.ICreature, _speakType int, _message string) {
	if p.Conn != nil {
		//p.Conn.SendCreatureSay(_creature, _speakType, _message)
	}
}

func (p *Player) sendCreatureChangeVisibility(_creature pul.ICreature, _visible bool) {
	if _creature.GetUID() != p.GetUID() {
		if _visible {
			p.AddVisibleCreature(_creature)
		} else if !p.hasFlag(PlayerFlag_CanSenseInvisibility) {
			p.RemoveVisibleCreature(_creature)
		}
	}
}

// --------------------- CHAT ----------------------------//
func (p *Player) sendClosePrivateChat(_channelId int) {

}

func (p *Player) sendToChannel(_fromPlayer pul.ICreature, _type int, _text string, _channelId int, _time int) {
	p.Conn.SendCreatureSay(_fromPlayer, _type, _text, _channelId, _time)
}

// --------------------- FRIEND ----------------------------//
func (p *Player) sendFriendList(_friends FriendList) {
	p.Conn.SendFriendList(_friends)
}

func (p *Player) sendFriendUpdate(_name string, _online bool) {
	p.Conn.SendFriendUpdate(_name, _online)
}

func (p *Player) sendFriendRemove(_name string) {
	p.Conn.SendFriendRemove(_name)
}