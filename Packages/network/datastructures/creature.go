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
package network
//Datastructures to be sent between server and client

//===============================================
// Server -> Client

type BodyPart struct {
	ID int
	Color uint32
}

//Add creature (HEADER_ADDCREATURE)
type Data_AddCreature struct {
	UID 		uint64
	Name 		string
	X			int
	Y			int
	Direction	int
	Outfit		[5]*BodyPart
}

func NewData_AddCreature() (msg *Message) {
	msg = NewMessage(HEADER_ADDCREATURE)
	msg.AddCreature = &Data_AddCreature{}
	return
}

//Remove creature (HEADER_REMOVECREATURE)
type Data_RemoveCreature struct {
	UID 	uint64
}

func NewData_RemoveCreature() (msg *Message) {
	msg = NewMessage(HEADER_REMOVECREATURE)
	msg.RemoveCreature = &Data_RemoveCreature{}
	return
}
