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

const (
	NUM_BODYPARTS = 6

	BODY_BASE = 0
	BODY_UPPER = 1
	BODY_NECK = 2
	BODY_HEAD = 3
	BODY_FEET = 4
	BODY_LOWER = 5
)

type PU_Player struct {
	PU_Creature
	
	name string
	
	walkConfirmed bool
	money uint32
	
	bodyParts [NUM_BODYPARTS]*PU_BodyPart
}

func NewPlayer(_id uint32) *PU_Player {
	player := &PU_Player{}
	player.SetDefault(_id)
	
	for i := 0; i < NUM_BODYPARTS; i++ {
		player.bodyParts[i] = NewBodyPart(1)
	}
	
	return player
}

func (p *PU_Player) Turn(_dir int, _send bool) {
	if _dir != p.direction {
		p.direction = _dir
		
		if _send {
			//g_conn.protocol.SendTurn(_dir)
		}
	}
}

type PU_BodyPart struct {
	id int
	
	red uint8
	green uint8
	blue uint8
}

func NewBodyPart(_id int) *PU_BodyPart {
	return &PU_BodyPart{id : _id}
}

func (b *PU_BodyPart) SetColor(_red int, _green int, _blue int) {
	b.red = uint8(_red)
	b.green = uint8(_green)
	b.blue = uint8(_blue)
}
