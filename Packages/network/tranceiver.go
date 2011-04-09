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

import (
	"gob"
	"io"
)

type Tranceiver struct {
	encoder *gob.Encoder
	decoder *gob.Decoder
}

func NewTranceiver(_socket io.ReadWriter) *Tranceiver {
	tranceiver := &Tranceiver{}
	
	types := []interface{}{
		(*Data_Login)(nil),
		(*Data_LoginStatus)(nil),
		(*Data_PlayerData)(nil),
		(*Data_AddCreature)(nil),
		(*Data_Tiles)(nil),
		(*Data_Walk)(nil),
		(*Data_CreatureWalk)(nil),
		(*Data_Turn)(nil),
		(*Data_CreatureTurn)(nil),
		(*Data_Warp)(nil),
	}
	for _, t := range types {
		gob.Register(t)
	}	

	tranceiver.encoder = gob.NewEncoder(_socket)
	tranceiver.decoder = gob.NewDecoder(_socket)
	
	return tranceiver
}

func (t *Tranceiver) Send(_message *Message) {
	if err := t.encoder.Encode(_message); err != nil {
		println("Encode error: " + err.String())
	}
}

func (t *Tranceiver) Receive() (message *Message, received bool) {
	var msg Message
	if t.decoder.Decode(&msg) != nil {
		received = false
		return
	}
	message = &msg
	return
}
