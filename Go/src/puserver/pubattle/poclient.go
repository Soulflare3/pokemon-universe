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
package pubattle

import (
	"fmt"
	pul "pulogic"
	pnet "network"
)

type PlayerInfoList map[int]*PlayerInfo

type POClient struct {
	player 			pul.IBattleCreature
	socket 			*POClientSocket

	players       	PlayerInfoList
	meLoginPlayer 	*FullPlayerInfo
	mePlayer      	*PlayerInfo

	battle 			*Battle

	bID 			int
}

func NewPOClient(_player pul.IBattleCreature) (*POClient, error) {
	poClient := POClient{player: _player,
		players: make(PlayerInfoList)}

	poClient.meLoginPlayer = NewFullPlayerInfoFromPlayer(_player)
	poClient.mePlayer = NewPlayerInfoFromFullPlayerInfo(poClient.meLoginPlayer)

	return &poClient, nil
}

func (c *POClient) UpdatePokemonData() {
	if c.player.GetType() != 2 { // Player
		return
	}

	playerParty := c.player.GetPokemonParty()
	for i := 0; i < 6; i++ {
		if poke := c.battle.myTeam.Pokes[i]; poke != nil {
			if playerPokemon := playerParty.GetFromSlot(i); playerPokemon != nil {
				playerPokemon.DamagedHp = poke.CurrentHP
				
				// Update pokemon moves
				for j := 0; j < 4; j++ {
					if move := poke.Moves[j]; move != nil {
						if playerMove := playerPokemon.Moves[j]; playerMove != nil {
							playerMove.CurrentPP = move.CurrentPP
						} else {
							fmt.Printf("COULD NOT UPDATE MOVE FOR %v IN SLOT %v\n", playerPokemon.GetNickname(), j)
						}
					}
				}
			} else {
				fmt.Printf("COULD NOT UPDATE POKEMON FROM %v IN SLOT %v\n", c.player.GetName(), i)
			}
		}
	}
}

func (c *POClient) Connect(_host string, _port string) {
	c.socket = NewPOClientSocket(c)
	//c.socket.Connect("localhost", "5080")
	c.socket.Connect(_host, _port)
}

func (c *POClient) ProcessPacket(_packet *pnet.QTPacket) {
	header := int(_packet.ReadUint8())
	fmt.Printf("Received header %v\n", header)
	switch header {
	case COMMAND_Login: // 2
		c.login(_packet)
	case COMMAND_PlayersList: // 5
		c.playerList(_packet)
	case COMMAND_ChallengeStuff: // 7
		c.challengeStuff(_packet)
	case COMMAND_EngageBattle: // 8
		c.engageBattle(_packet)
	case COMMAND_BattleMessage: // 10
		c.battleMessage(_packet)
	case COMMAND_KeepAlive: // 12
		c.keepAlive()
	case COMMAND_Register: // 14
		// Do nothing
	case COMMAND_VersionControl: // 33
		// Do nothing
	case COMMAND_TierSelection: // 34
		// Do nothing
	case COMMAND_BattleList: // 43
		// Do nothin
	case COMMAND_ChannelsList: // 44
		// Do nothing
	case COMMAND_ChannelPlayers: // 45
		// Do nothing
	case COMMAND_JoinChannel: // 46
		// Do nothing
	case COMMAND_ChannelMessage: // 51
		// Do nothing
	case COMMAND_HtmlMessage: // 53
		fmt.Printf("[Message] %s\n\r", _packet.ReadString())
	case COMMAND_ServerName: // 55
		// Do nothing
	default:
		fmt.Printf("UNIMPLEMENTED PACKET: %v\n", header)
	}
}

func (c *POClient) SendMessage(_buffer pnet.IPacket, _header int) {
	c.socket.SendMessage(_buffer, _header)
}

// --------------------- Receive Packets ------------------------
func (c *POClient) login(_packet *pnet.QTPacket) {
	c.mePlayer = NewPlayerInfoFromPacket(_packet)
	c.players[c.mePlayer.Id] = c.mePlayer
}

func (c *POClient) playerList(_packet *pnet.QTPacket) {
	p := NewPlayerInfoFromPacket(_packet)
	fmt.Printf("PlayerList: %d - %s\n\r", p.Id, p.Nick)
	if _, found := c.players[p.Id]; !found {
		c.players[p.Id] = p
	}
}

func (c *POClient) challengeStuff(_packet *pnet.QTPacket) {
	challenge := NewIncommingChallengeFromPacket(_packet)
	challenge.SetNick(c.players[challenge.opponent])

	// PU server will handle the challenge stuff
	if challenge.desc == CHALLENGEDESC_SENT {
		if challenge.IsValidChallenge(c.players) {
			response := c.constructChallenge(CHALLENGEDESC_ACCEPTED, challenge.opponent, challenge.clauses, challenge.mode)
			c.SendMessage(response, COMMAND_ChallengeStuff)
		}
	}
}

func (c *POClient) constructChallenge(_desc int, _opp int, _clauses int, _mode int) *pnet.QTPacket {
	packet := pnet.NewQTPacket()
	packet.AddUint8(uint8(_desc))
	packet.AddUint32(uint32(_opp))
	packet.AddUint32(uint32(_clauses))
	packet.AddUint8(uint8(_mode))
	return packet
}

func (c *POClient) engageBattle(_packet *pnet.QTPacket) {
	c.bID = int(_packet.ReadUint32())
	pID1 := int(_packet.ReadUint32())
	pID2 := int(_packet.ReadUint32())
	if pID1 == 0 { // This is us!				
		battleConf := NewBattleConfFromPacket(_packet)
		// Start the battle
		c.battle = NewBattle(c, battleConf, _packet, c.players[battleConf.GetId(0)], c.players[battleConf.GetId(1)], c.mePlayer.Id, c.bID)

		fmt.Printf("Battle between %s and %s started!\n", c.mePlayer.Nick, c.players[pID2].Nick)
	}
}

func (c *POClient) battleMessage(_packet *pnet.QTPacket) {
	if c.battle != nil {
		_packet.ReadUint32() // Supporting only one battle, unneeded
		_packet.ReadUint32() // Discard the size, unneeded
		c.battle.ReceiveCommand(_packet)
	}
}

// --------------------- Send Packets --------------------------- //
func (c *POClient) keepAlive() {
	c.SendMessage(pnet.NewQTPacket(), COMMAND_KeepAlive)
}

// --------------------- Send to PU Client --------------------------- //
func (c *POClient) UpdatePokes(_player int) {
	if _player == c.battle.me {
		c.UpdateMyPoke()
	} else {
		c.UpdateOppPoke()
	}
}

func (c *POClient) UpdateMyPoke() {
	if poke := c.battle.currentPoke(c.battle.me); poke != nil {
		packet := pnet.NewPacketExt(pnet.HEADER_BATTLE_UPDATEMYPOKE)
		packet.AddString(poke.RNick)
		packet.AddUint32(uint32(poke.Level))
		packet.AddUint8(uint8(poke.Gender))
		packet.AddUint32(uint32(poke.GetStatus()))
		
	}
}

func (c *POClient) UpdateOppPoke() {

}
