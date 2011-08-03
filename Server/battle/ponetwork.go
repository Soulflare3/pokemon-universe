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
	"net"
	"os"
	"fmt"
	pnet "network" // PU Network packet
)

type PONetwork struct {
	socket net.Conn
	IsOpen bool
	owner  *POClient
}

func NewPONetwork() (*PONetwork, os.Error) {
	network := &PONetwork{}
	err := network.Connect()

	return network, err
}

func (c *PONetwork) Connect() (err os.Error) {
	c.socket, err = net.Dial("tcp", BATTLESERVER_IP)
	if err == nil {
		go c.HandleConnection()
	}
	return
}

func (c *PONetwork) HandleConnection() {
	c.IsOpen = true

	for {
		var headerbuffer [2]uint8
		recv, err := c.socket.Read(headerbuffer[0:])
		if err != nil || recv == 0 {
			fmt.Printf("Error while reading socket: %v", err)
			break
		}

		packet := pnet.NewQTPacket()
		copy(packet.Buffer[0:2], headerbuffer[0:2])
		packet.GetHeader()

		databuffer := make([]uint8, packet.MsgSize)
		recv, err = c.socket.Read(databuffer[0:])
		if recv == 0 {
			continue
		} else if err != nil {
			fmt.Printf("Connection read error: %v", err)
			continue
		}

		copy(packet.Buffer[2:], databuffer[:])
		c.ProcessPacket(packet)
		//c.TestRead(packet)
	}

	c.IsOpen = false
}

func (c *PONetwork) TestRead(_packet *pnet.QTPacket) {
	command := _packet.ReadUint32()
	charCmd := _packet.ReadUint8()
	player := _packet.ReadUint8()

	fmt.Printf("Command %v | charCmd %v | Player %v", command, charCmd, player)
}

func (c *PONetwork) ProcessPacket(_packet *pnet.QTPacket) {
	header := _packet.ReadUint8()
	switch header {
	case Register: // 14
		return // do nothing
	case VersionControl, TierSelection: // 33, 34
		return
	case BattleList, ChannelsList, ChannelPlayers, JoinChannel: // 43, 44, 45, 46
		return
	case ChannelMessage, HtmlMessage: // 51, 53
		return
	case Login: // 2
		c.ReceiveLogin(_packet)
	case SendMessage: // 4
		c.ReceivedMessage(_packet)
	case PlayersList: // 5
		c.ReceivePlayersList(_packet)
	case ChallengeStuff: // 7
		c.ReceiveChallengeStuff(_packet)
	case EngageBattle: // 8
		c.ReceiveEngageBattle(_packet)
	case BattleFinished: // 9
		c.ReceiveBattleFinished(_packet)
	case BattleMessage: // 10
		c.ReceiveBattleMessage(_packet)
	case KeepAlive: // 12
		c.KeepAlive()

		//		case BattleFinished: // 9
		// TODO

	default:
		fmt.Printf("[Warning] Received unknown header %v\n\r", header)
	}
}

/*******************************************************/
//		Receive Messages
/*******************************************************/

func (c *PONetwork) ReceiveLogin(_packet *pnet.QTPacket) {
	playerInfo := NewPlayerInfo()
	playerInfo.id = int32(_packet.ReadUint32())

	basicInfo := NewBasicInfo()
	basicInfo.name = _packet.ReadString()
	basicInfo.info = _packet.ReadString()
	playerInfo.team = basicInfo

	playerInfo.auth = int8(_packet.ReadUint8())
	playerInfo.flags = _packet.ReadUint8()
	playerInfo.rating = int16(_packet.ReadUint16())

	for i := 0; i < 6; i++ {
		playerInfo.pokes[i] = NewPokemonUniqueIdFromNum(_packet.ReadUint16(), _packet.ReadUint8())
	}

	playerInfo.avatar = _packet.ReadUint16()
	playerInfo.tier = _packet.ReadString()
	playerInfo.color = _packet.ReadUint32()
	playerInfo.gen = _packet.ReadUint8()

	c.owner.PlayerLogin(playerInfo)
}

func (c *PONetwork) ReceivedMessage(_packet *pnet.QTPacket) {
	message := _packet.ReadString()
	fmt.Printf("Recv Message: %v\n\r", message)
}

func (c *PONetwork) ReceiveBattleMessage(_packet *pnet.QTPacket) {
	battleId := int32(_packet.ReadUint32())
	_packet.ReadUint32() // Dummy var, without this everything goes wrong >.>
	c.owner.BattleCommand(battleId, _packet)
}

func (c *PONetwork) ReceivePlayersList(_packet *pnet.QTPacket) {
	for _packet.CanRead() {
		playerInfo := NewPlayerInfo()
		playerInfo.id = int32(_packet.ReadUint32())
		if playerInfo.id == 0 {
			break
		}

		basicInfo := NewBasicInfo()
		basicInfo.name = _packet.ReadString()
		basicInfo.info = _packet.ReadString()
		playerInfo.team = basicInfo

		playerInfo.auth = int8(_packet.ReadUint8())
		playerInfo.flags = _packet.ReadUint8()
		playerInfo.rating = int16(_packet.ReadUint16())

		for i := 0; i < 6; i++ {
			playerInfo.pokes[i] = NewPokemonUniqueIdFromNum(_packet.ReadUint16(), _packet.ReadUint8())
		}

		playerInfo.avatar = _packet.ReadUint16()
		playerInfo.tier = _packet.ReadString()
		playerInfo.color = _packet.ReadUint32()
		playerInfo.gen = _packet.ReadUint8()

		c.owner.PlayerReceived(playerInfo)
	}
}

func (c *PONetwork) ReceiveChallengeStuff(_packet *pnet.QTPacket) {
	fmt.Println("<- Receive Challenge Stuff")
	challengeInfo := NewChallengeInfo()
	challengeInfo.description = _packet.ReadUint8()
	challengeInfo.opponent = _packet.ReadUint32()
	challengeInfo.clauses = _packet.ReadUint32()
	challengeInfo.mode = _packet.ReadUint8()

	c.owner.ChallengeStuff(challengeInfo)
}

func (c *PONetwork) ReceiveEngageBattle(_packet *pnet.QTPacket) {
	fmt.Println("<- Receive Engage Battle")
	battleid := int32(_packet.ReadUint32())
	id1 := int32(_packet.ReadUint32())
	id2 := int32(_packet.ReadUint32())

	if id1 == 0 {
		// This is a battle we take part in
		conf := NewBattleConfiguration()
		conf.gen = _packet.ReadUint8()
		conf.mode = _packet.ReadUint8()
		conf.ids[0] = int32(_packet.ReadUint32())
		conf.ids[1] = int32(_packet.ReadUint32())
		conf.clauses = _packet.ReadUint32()

		team := NewTeamBattle()
		for i := 0; i < 6; i++ {
			poke := NewPokeBattle()
			poke.num = NewPokemonUniqueIdFromRef(_packet.ReadUint32())
			poke.nick = _packet.ReadString()
			poke.totalLifePoints = _packet.ReadUint16()
			poke.lifePoints = _packet.ReadUint16()
			poke.gender = _packet.ReadUint8()
			poke.shiny = (_packet.ReadUint8() == 1)
			poke.level = _packet.ReadUint8()
			poke.item = _packet.ReadUint16()
			poke.ability = _packet.ReadUint16()
			poke.happiness = _packet.ReadUint8()

			for i := 0; i < 5; i++ {
				poke.normal_stats[i] = _packet.ReadUint16()
			}

			for i := 0; i < 4; i++ {
				move := NewBattleMove()
				move.num = _packet.ReadUint16()
				move.pp = _packet.ReadUint8()
				move.totalPP = _packet.ReadUint8()
				poke.moves[i] = move
			}

			for i := 0; i < 6; i++ {
				poke.evs[i] = _packet.ReadUint8()
			}

			for i := 0; i < 6; i++ {
				poke.dvs[i] = _packet.ReadUint8()
			}

			team.SetPoke(i, poke)
		}

		c.owner.StartBattleSelf(battleid, id2, team, conf)
	} else {
		// This is a battle of strangers
	}
}

func (c *PONetwork) ReceiveBattleFinished(_packet *pnet.QTPacket) {
	battleId := int32(_packet.ReadUint32())
	desc := int8(_packet.ReadUint8())
	id1 := int32(_packet.ReadUint32())
	id2 := int32(_packet.ReadUint32())

	c.owner.BattleFinished(battleId, desc, id1, id2)
}

/*******************************************************/
//		Send Messages
/*******************************************************/

func (c *PONetwork) KeepAlive() {
	packet := pnet.NewQTPacketExt(KeepAlive)
	c.SendMessage(packet)
}

func (c *PONetwork) SendMessage(_packet *pnet.QTPacket) {
	_packet.SetHeader()
	c.socket.Write(_packet.Buffer[0:_packet.MsgSize])
}

func (c *PONetwork) SendChallengeStuff(_info *ChallengeInfo) {
	fmt.Println("-> Send Challenge Stuff")
	packet := pnet.NewQTPacketExt(ChallengeStuff)
	packet.AddUint8(_info.description)
	packet.AddUint32(_info.opponent)
	packet.AddUint32(_info.clauses)
	packet.AddUint8(_info.mode)

	c.SendMessage(packet)
}

func (c *PONetwork) SendBattleCommandBattleChoice(_battleId int32, _choice *BattleChoice) {
	packet := pnet.NewQTPacketExt(BattleMessage)
	packet.AddUint32(uint32(_battleId))
	packet.AddUint8(_choice.playerSlot)
	packet.AddUint8(_choice.choiceType)

	if _choice.choiceType == ChoiceType_Switch {
		packet.AddUint8(uint8(_choice.choice.switching.pokeSlot))
	} else if _choice.choiceType == ChoiceType_Attack {
		packet.AddUint8(uint8(_choice.choice.attack.attackSlot))
		packet.AddUint8(uint8(_choice.choice.attack.attackTarget))
	} else if _choice.choiceType == ChoiceType_Rearrange {
		for i := 0; i < 6; i++ {
			packet.AddUint8(uint8(_choice.choice.rearrange.pokeIndexes[i]))
		}
	}

	c.SendMessage(packet)
}
