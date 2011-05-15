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
	"fmt"
	pnet "network"
)

type POClient struct {
	connection		*PONetwork
	battleId		uint32
	mid				int32
	myNick			string
	
	myNames			map[string]int32
	myPlayersInfo	map[int32]*PlayerInfo // List of all players on the server
	myBattles		map[int32]*Battle
}

func NewPOClient() (*POClient, os.Error) {
	connection, err := NewPONetwork()
	if err != nil {
		return nil, err
	}
	client := &POClient{}
	client.connection = connection
	client.myPlayersInfo = make(map[int32]*PlayerInfo)
	client.myNames = make(map[string]int32)
	client.myBattles = make(map[int32]*Battle)
	
	connection.owner = client
	return client, nil
}

func (c *POClient) SendLoginInfo() {
	packet := pnet.NewQTPacketExt(Login)
	packet.AddString("HerpDerp") // Name
	packet.AddString("Dark Info") // Info
	packet.AddString("Dark Lose") // Lose text
	packet.AddString("Dark Winrar") // Win text
	packet.AddUint16(0) // Avatar
	packet.AddString("1") // Default Tier
	packet.AddUint8(5) // Generation
	
	// TEAM - Loop Pokemon
	packet.AddUint16(16) // pokemon number
	packet.AddUint8(0) // sub number (alt-forms)
	packet.AddString("Pidgey") // nickname
	packet.AddUint16(0) // item
	packet.AddUint16(65) // ability
	packet.AddUint8(0) // nature
	packet.AddUint8(1) // gender
	packet.AddUint8(0) // shiny	
	packet.AddUint8(0) // happiness
	packet.AddUint8(100) // level

	// Team - Loop Pokemon - Loop Moves
	packet.AddUint32(16) // moveid
	packet.AddUint32(0) // moveid
	packet.AddUint32(0) // moveid
	packet.AddUint32(0) // moveid
	// Team - Loop Pokemon - End Loop Moves
	// Team - Loop Pokemon - Loop DV
	packet.AddUint8(15) // hp
	packet.AddUint8(15) // attack
	packet.AddUint8(15) // defense
	packet.AddUint8(15) // SpAttack
	packet.AddUint8(15) // SpDefence
	packet.AddUint8(15) // Speed
	// Team - Loop Pokemon - End Loop DV
	// Team - Loop Pokemon - Loop EV
	packet.AddUint8(152)
	packet.AddUint8(92)
	packet.AddUint8(60)
	packet.AddUint8(64)
	packet.AddUint8(64)
	packet.AddUint8(76)
	// Team - Loop Pokemon - End Loop EV
	
	// Loop 5 more empty pokemon
	for i := 0; i < 5; i++ {
		packet.AddUint16(0) // pokemon number
		packet.AddUint8(0) // sub number (alt-forms)
		packet.AddString("") // nickname
		packet.AddUint16(0) // item
		packet.AddUint16(0) // ability
		packet.AddUint8(0) // nature
		packet.AddUint8(0) // gender
		packet.AddUint8(0) // shiny
		packet.AddUint8(0) // happiness
		packet.AddUint8(0) // level	
		
		packet.AddUint32(0) // moveid
		packet.AddUint32(0) // moveid
		packet.AddUint32(0) // moveid
		packet.AddUint32(0) // moveid

		packet.AddUint8(0) // hp
		packet.AddUint8(0) // attack
		packet.AddUint8(0) // defense
		packet.AddUint8(0) // SpAttack
		packet.AddUint8(0) // SpDefence
		packet.AddUint8(0) // Speed
		
		packet.AddUint8(0)
		packet.AddUint8(0)
		packet.AddUint8(0)
		packet.AddUint8(0)
		packet.AddUint8(0)
		packet.AddUint8(0)
	}
	
	// Team - End Loop Pokemon
	packet.AddUint8(1) // Ladder
	packet.AddUint8(1) // Show team
	packet.AddUint32(1) // Colour
	
	c.connection.SendMessage(packet)
}

func (c *POClient) PlayerLogin(_p *PlayerInfo) {
	fmt.Printf("Received login %v (%d)\n", _p.team.name, _p.id)

	c.mid = _p.id
	c.myNick = _p.team.name
	c.myPlayersInfo[_p.id] = _p
	c.myNames[_p.team.name] = _p.id
}

func (c *POClient) PlayerReceived(_p *PlayerInfo) {
	if _, found := c.myPlayersInfo[_p.id]; found {
		if id := c.myNames[_p.team.name]; id == _p.id {
			return
		}
	}
	
	fmt.Printf("Received player %v (%d)\n", _p.team.name, _p.id)
	c.myPlayersInfo[_p.id] = _p
	c.myNames[_p.team.name] = _p.id
}

func (c *POClient) GetPlayer(_id int32) (value *PlayerInfo, found bool) {
	value, found = c.myPlayersInfo[_id]
	return
}

func (c *POClient) ChallengeStuff(_info *ChallengeInfo) {
	if(_info.description == ChallengeDesc_Sent) { // We are being challenged
		_info.description = ChallengeDesc_Accepted // Auto accept	
		c.connection.SendChallengeStuff(_info )
	} else { // We challenged someone else (reply)
		// TODO: Handle accepted challenge
	}
}

func (c *POClient) StartBattleSelf(_battleId int32, _id int32, _team *TeamBattle, _conf *BattleConfiguration) {
	c.myBattleStarted(_battleId, c.mid, _id, _team, _conf)
}

func (c *POClient) myBattleStarted(_battleId int32, _id1, _id2 int32, _team *TeamBattle, _conf *BattleConfiguration) {
	c.myPlayersInfo[_id1].flags |= PlayerInfo_Battling
	c.myPlayersInfo[_id2].flags |= PlayerInfo_Battling
	
	p1, _ := c.GetPlayer(_id1)
	p2, _ := c.GetPlayer(_id2)
	c.myBattles[_battleId] = NewBattle(c, _battleId, p1, p2, _team, _conf)
}

func (c *POClient) BattleCommand(_battleId int32, _packet *pnet.QTPacket) {
	battle, found := c.myBattles[_battleId]
	if found {
		battle.ReceiveInfo(_packet)
	}
}

func (c *POClient) SendBattleChoice(_battleId int32, _choice *BattleChoice) {
	c.connection.SendBattleCommandBattleChoice(_battleId, _choice)
}

func (c *POClient) BattleFinished(_battleId int32, _res int8, _winner, _loser int32) {
	if (_res == BattleResult_Close || _res == BattleResult_Forfeit) && (_battleId != 0 || (_winner == c.mid || _loser == c.mid)) {
		// Close battle window
	}
	
	c.myBattles[_battleId] = nil, false
	
	if value, found := c.myPlayersInfo[_winner]; found {
		value.flags &= 0xFF ^ PlayerInfo_Battling
	}
	if value, found := c.myPlayersInfo[_loser]; found {
		value.flags &= 0xFF ^ PlayerInfo_Battling
	}
}