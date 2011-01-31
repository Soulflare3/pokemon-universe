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
	pos "position"
)

type Player struct {
	name			string
	uid				uint64 // Unique ID
	Id				int // Database ID			
	
	Position		*Tile
	Direction		int
	Conn			*Connection	
	
	Movement		int
	Location		*Location
	LastPokeCenter	*Tile
	
	Money			int32
	
	Outfit
}

func NewPlayer(_name string) *Player {
	p := Player{ name : _name }
	p.uid 	= GenerateUniqueID()
	p.Conn 	= nil
	p.Outfit = NewOutfit()
	
	return &p
}

func  (p *Player) GetUID() uint64 {
	return p.uid
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetTile() *Tile {
	return p.Position
}

func (p *Player) GetPosition() pos.Position {
	return p.Position.Position
}

func (p *Player) GetMovement() int {
	return p.Movement
}

func (p *Player) SetConnection(_conn *Connection) {
	p.Conn = _conn
	p.Conn.Owner = p
	go _conn.HandleConnection()
}

func (p *Player) SetMoney(_money int32) int32 {
	if p.Money += _money; p.Money < 0 {
		p.Money = 0
	}
	return p.Money
}
