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
// Client -> Server

//Send turn action to server (HEADER_TURN)
type Data_Turn struct {
	Direction int
}

func NewData_Turn() (msg *Message) {
	msg = NewMessage(HEADER_TURN)
	msg.Turn = &Data_Turn{}
	return
}

//===============================================
// Server -> Client

//Notify client of a turning creature (HEADER_TURN)
type Data_CreatureTurn struct {
	UID       uint64
	Direction int
}

func NewData_CreatureTurn() (msg *Message) {
	msg = NewMessage(HEADER_TURN)
	msg.CreatureTurn = &Data_CreatureTurn{}
	return
}
