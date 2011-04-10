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
	punet "network"
	"os"
)

type PU_Message_Battle struct {

}

func NewBattleMessage(_packet *punet.Packet) *PU_Message_Battle {
	msg := &PU_Message_Battle{}
	msg.ReadPacket(_packet)
	return msg
}

func (m *PU_Message_Battle) ReadPacket(_packet *punet.Packet) os.Error {
	battletype := int(_packet.ReadUint16())
	
	g_game.SetBattle(battletype)
	
	numPlayers := int(_packet.ReadUint16())
	for i := 0; i < numPlayers; i++ {
		fightertype := int(_packet.ReadUint16())
		switch fightertype {
		case SELF:
			id := int(_packet.ReadUint16())
			team := int(_packet.ReadUint16())
			starter := int(_packet.ReadUint16())
			g_game.battle.SetSelf(id, team, starter)
			
		case NPC, POKEMON:
			id := int(_packet.ReadUint16())
			team := int(_packet.ReadUint16())
			pokeid := int(_packet.ReadUint16())
			pokename := _packet.ReadString()
			level := int(_packet.ReadUint16())
			g_game.battle.SetNPC(id, team, pokename, pokeid, level)
			
		case PLAYER:
			id := int(_packet.ReadUint16())
			team := int(_packet.ReadUint16())
			name := _packet.ReadString()
			pokeid := int(_packet.ReadUint16())
			pokename := _packet.ReadString()
			level := int(_packet.ReadUint16())
			hp := int(_packet.ReadUint16())
			g_game.battle.SetPlayer(id, team, name, pokename, pokeid, level, hp)
		}
	}
	
	g_game.battle.Start()

	return nil
}

type PU_Message_BattleMove struct {
	movetype int
	param1 int 
	param2 int
}

func NewBattleMoveMessage() *PU_Message_BattleMove {
	return &PU_Message_BattleMove{}
}

func (m *PU_Message_BattleMove) WritePacket() (*punet.Packet, os.Error) {
	packet := punet.NewPacketExt(0xD3)
	packet.AddUint16(uint16(m.movetype))
	packet.AddUint16(uint16(m.param1))
	packet.AddUint16(uint16(m.param2))
	return packet, nil
}


