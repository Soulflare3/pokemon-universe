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
	"fmt"
	"strings"
	"sync"
	"time"

	pnet "network" // PU Network package
	pul "pulogic"
	pos "putools/pos" // Position package
	"putools/log"
)

type GameState int

const (
	GAME_STATE_STARTUP GameState = iota
	GAME_STATE_INIT
	GAME_STATE_NORMAL
	GAME_STATE_CLOSED
	GAME_STATE_CLOSING
)

type Game struct {
	State         		GameState
	
	Creatures     		pul.CreatureList
	Players       		PlayerList
	PlayersDiscon		PlayerList

	Chat				*Chat
	Locations			*LocationStore
	Items				*ItemStore

	mutexCreatureList   *sync.RWMutex
	mutexPlayerList     *sync.RWMutex
	mutexDisconnectList *sync.RWMutex
}

func NewGame() *Game {
	game := Game{}
	game.State = GAME_STATE_STARTUP
	
	// Initialize maps
	game.Creatures = make(pul.CreatureList)
	game.Players = make(PlayerList)
	game.PlayersDiscon = make(PlayerList)
	
	game.Chat = NewChat()

	// Mutexes
	game.mutexCreatureList = new(sync.RWMutex)
	game.mutexPlayerList = new(sync.RWMutex)
	game.mutexDisconnectList = new(sync.RWMutex)

	return &game
}

func (the *Game) Load() (LostIt bool) {
	LostIt = true // fuck >:(
	g_map = NewMap()
	the.Locations = NewLocationStore()

	logger.Println(" - Loading locations")
	if err := the.Locations.Load(); err != nil {
		logger.Println("[ERROR] Failed to load locations...")
		LostIt = false
	}
	
	logger.Println(" - Loading items")
	if err := the.Items.Load(); !err {
		logger.Println("[ERROR] Failed to load items...")
		LostIt = false
	}

	logger.Println(" - Loading worldmap")
	start := time.Now().UnixNano()
	if err := g_map.Load(); err != nil {
		logger.Println("[ERROR] Failed to load worldmap...")
		LostIt = false
	} else {
		logger.Printf(" - Map loaded in %dms\n", (time.Now().UnixNano()-start)/1e6)
	}
	
	logger.Println(" - Loading NPCs and scripts")
	g_npc = NewNpcManager()
	if !g_npc.Load() {
		logger.Println("[ERROR] Failed to load NPC data...")
		LostIt = false
	}

	return
}

func (g *Game) CheckCreatures() {
	// TODO: Change this to a scheduler
	go func() {
		time.Sleep(1e9)
		g.CheckCreatures()
	}()
	
	g.mutexCreatureList.RLock()
	defer g.mutexCreatureList.RUnlock()
	
	for _, creature := range g.Creatures {
		creature.OnThink(1000)
	}
}

func (g *Game) GetPlayerByGuid(_guid uint64) (*Player, bool) {
	v, ok := g.Players[_guid]
	return v, ok
}

func (g *Game) GetPlayerByName(_name string) (*Player, bool) {
	for _, value := range g.Players {
		if value.GetName() == _name {
			return value, true
		}
	}

	return nil, false
}

func (g *Game) OnPlayerLoseConnection(_player *Player) {
	_player.Conn = nil
	_player.TimeoutCounter = 0
	// g.PlayersDiscon[_player.GetUID()] = _player
	
	g.RemoveCreature(_player.GetUID())
}

func (g *Game) AddCreature(_creature pul.ICreature) {
	// TODO: Maybe only take the creatues from the area the new creature is in. This saves some extra iterating
	// TODO 2: Upgrade this to parallel stuff
	
	g.mutexCreatureList.Lock()
	defer g.mutexCreatureList.Unlock()
	for _, value := range g.Creatures {
		value.OnCreatureAppear(_creature, true)
	}
	
	g.Creatures[_creature.GetUID()] = _creature

	if _creature.GetType() == CTYPE_PLAYER {
		g.mutexPlayerList.Lock()
		defer g.mutexPlayerList.Unlock()
		g.Players[_creature.GetUID()] = _creature.(*Player)
		
		// Join default channels
		g.Chat.AddUserToChannel(_creature.(*Player), pnet.CHANNEL_WORLD)
		g.Chat.AddUserToChannel(_creature.(*Player), pnet.CHANNEL_TRADE)
	}
}

func (g *Game) RemoveCreature(_guid uint64) {
	object, exists := g.Creatures[_guid]
	if exists {
		g.mutexCreatureList.Lock()
		defer g.mutexCreatureList.Unlock()

		delete(g.Creatures, _guid)

		if object.GetType() == CTYPE_PLAYER {			
			g.mutexPlayerList.Lock()
			defer g.mutexPlayerList.Unlock()
			
			// Remove from player list
			delete(g.Players, _guid)
			
			playerObj := object.(*Player)
			
			// Remove from channels
			g.Chat.RemoveUserFromAllChannels(playerObj)
			
			// Save all player data
			playerObj.SaveData()
			
			logger.Printf("[Logout] %d - %v logged out\n", object.GetUID(), object.GetName())
			
		}
		
		// Remove creature from all visible creature lists
		for _, c := range g.Creatures {
			c.RemoveVisibleCreature(object)
		}
	}
}

func (g *Game) OnPlayerMove(_creature pul.ICreature, _direction int, _sendMap bool) {
	ret := g.OnCreatureMove(_creature, _direction)

	player := _creature.(*Player)
	if ret == RET_NOTPOSSIBLE {
		player.sendCreatureMove(_creature, _creature.GetTile(), _creature.GetTile())
	} else if ret == RET_PLAYERISTELEPORTED {
		player.sendPlayerWarp()
		player.sendMapData(DIR_NULL)
	} else {
		player.sendMapData(_direction)
	}
}

func (g *Game) OnPlayerTurn(_creature pul.ICreature, _direction int) {
	if _creature.GetDirection() != _direction {
		g.OnCreatureTurn(_creature, _direction)
	}
}

func (g *Game) OnPlayerSay(_creature *Player, _channelId int, _speakType int, _receiver string, _message string) bool {
	toReturn := false
	switch _speakType {
		case pnet.SPEAK_NORMAL:
			toReturn = g.internalCreatureSay(_creature, pnet.SPEAK_NORMAL, _message)
		case pnet.SPEAK_YELL:
			toReturn = g.playerYell(_creature, _message)
		case pnet.SPEAK_WHISPER:
			toReturn = g.playerWhisper(_creature, _message)
		case pnet.SPEAK_PRIVATE:
			toReturn = g.playerSpeakTo(_creature, _speakType, _receiver, _message)
		case pnet.SPEAK_CHANNEL:
			if g.playerTalkToChannel(_creature, _speakType, _message, _channelId) {
				toReturn = true
			} else if _channelId == 0 {
				// Resend in default channel
				toReturn = g.OnPlayerSay(_creature, 0, pnet.SPEAK_NORMAL, _receiver, _message)
			}
		case pnet.SPEAK_BROADCAST:
			toReturn = g.internalBroadcastMessage(_creature, _message)
	}
	
	return toReturn
}

func (g *Game) OnCreatureMove(_creature pul.ICreature, _direction int) (ret int) {
	ret = RET_NOTPOSSIBLE

	if !CreatureCanMove(_creature) {
		return
	}

	currentTile := _creature.GetTile()
	destinationPosition := currentTile.GetPosition()

	switch _direction {
	case DIR_NORTH:
		destinationPosition.Y -= 1
	case DIR_SOUTH:
		destinationPosition.Y += 1
	case DIR_WEST:
		destinationPosition.X -= 1
	case DIR_EAST:
		destinationPosition.X += 1
	}

	// Check if destination tile exists
	destinationTile, ok := g_map.GetTileFromPosition(destinationPosition)
	if !ok {
		return
	}

	// Check if we can move to the destination tile
	if ret = destinationTile.CheckMovement(_creature, _direction); ret == RET_NOTPOSSIBLE {
		return
	}

	// Update position
	_creature.SetTile(destinationTile)
	
	// fmt.Printf("[%v] From (%v,%v) To (%v,%v)\n", _creature.GetName(), currentTile.Position.X, currentTile.Position.Y, destinationPosition.X, destinationPosition.Y)

	// Tell creatures this creature has moved
	g.mutexCreatureList.RLock()
	defer g.mutexCreatureList.RUnlock()
	for _, value := range g.Creatures {
		if value != nil {
			value.OnCreatureMove(_creature, currentTile, destinationTile, false)
		}
	}

	// Move creature object to destination tile
	if ret = currentTile.RemoveCreature(_creature, true); ret == RET_NOTPOSSIBLE {
		return
	}
	if ret = destinationTile.AddCreature(_creature, true); ret == RET_NOTPOSSIBLE {
		currentTile.AddCreature(_creature, false) // Something went wrong, put creature back on old tile
		return
	}

	// If ICreature is a player type we can check for wild encounter
	g.checkForWildEncounter(_creature)

	return
}

func (g *Game) OnCreatureTurn(_creature pul.ICreature, _direction int) {
	if _creature.GetDirection() != _direction {
		_creature.SetDirection(_direction)

		visibleCreatures := _creature.GetVisibleCreatures()
		for _, value := range visibleCreatures {
			value.OnCreatureTurn(_creature)
		}
	}
}

func (g *Game) internalCreatureTeleport(_creature pul.ICreature, _from pul.ITile, _to pul.ITile) (ret int) {
	ret = RET_NOERROR

	if _from == nil || _to == nil {
		ret = RET_NOTPOSSIBLE
	} else {
		// Move creature object to destination tile
		if ret = _from.RemoveCreature(_creature, true); ret == RET_NOTPOSSIBLE {
			return
		}

		if ret = _to.AddCreature(_creature, false); ret == RET_NOTPOSSIBLE {
			_from.AddCreature(_creature, false) // Something went wrong, put creature back on old tile
			return
		}

		_creature.SetTile(_to)

		// Tell creatures this creature has been teleported
		g.mutexCreatureList.RLock()
		defer g.mutexCreatureList.RUnlock()
		for _, value := range g.Creatures {
			if value != nil {
				value.OnCreatureMove(_creature, _from, _to, true)
			}
		}
	}
	
	// No error, played succesfully teleported
	if ret == RET_NOERROR {
		ret = RET_PLAYERISTELEPORTED
	}

	return
}

func (g *Game) internalCreatureSay(_creature pul.ICreature, _speakType int, _message string) bool {
	list := make(pul.CreatureList)
	if _speakType == pnet.SPEAK_YELL {
		position := _creature.GetPosition()  // Get position of speaker

		g.mutexPlayerList.RLock()
		defer g.mutexPlayerList.RUnlock()
		for _, player := range g.Players {
			if player != nil {
				if position.IsInRange3p(player.GetPosition(), pos.NewPositionFrom(27, 21, 0)) {
					list[player.GetUID()] = player
				}
			}
		}
	} else {
		list = _creature.GetVisibleCreatures()
	}
	
	// Send chat message to all visible players
	for _, creature := range list {
		if creature.GetType() == CTYPE_PLAYER {
			player := creature.(*Player)
			// player.sendCreatureSay(_creature, _speakType, _message)
			player.sendToChannel(_creature, _speakType, _message, pnet.CHANNEL_LOCAL, 0)
		}
	}	
	
	return true
}

func (g *Game) playerYell(_player *Player, _message string) bool {
	addExhaustion := 0
	isExhausted := false
	
	if !_player.HasCondition(CONDITION_EXHAUST_YELL, true) {
		addExhaustion = 10
		g.internalCreatureSay(_player, pnet.SPEAK_YELL, strings.ToUpper(_message))
	} else {
		isExhausted = true
		addExhaustion = 5
		_player.sendCancelMessage(RET_YOUAREEXHAUSTED)
	}
	
	if addExhaustion > 0 {
		condition := CreateCondition(CONDITIONID_DEFAULT, CONDITION_EXHAUST_YELL, addExhaustion, 0)
		_player.AddCondition(condition)
	}
	
	return !isExhausted
}

func (g *Game) playerWhisper(_player *Player, _text string) bool {
	list := _player.GetVisibleCreatures()
	position := _player.GetPosition()
	
	for _, creature := range list {
		if creature.GetType() == CTYPE_PLAYER {
			tmpPlayer := creature.(*Player)
			if position.IsInRange3p(tmpPlayer.GetPosition(), pos.NewPositionFrom(1, 1, 0)) {
				tmpPlayer.sendCreatureSay(_player, pnet.SPEAK_WHISPER, _text)
			} else {
				tmpPlayer.sendCreatureSay(_player, pnet.SPEAK_WHISPER, "pspsps")
			}
		}
	}
	
	return true
}

func (g *Game) playerSpeakTo(_player *Player, _type int, _receiver string, _text string) bool {
	toPlayer, found := g.GetPlayerByName(_receiver)
	if !found {
		_player.sendTextMessage(pnet.MSG_STATUS_SMALL, "A player with this name is not online.")
		return false
	}
	
	toPlayer.sendCreatureSay(_player, _type, _text)
	_player.sendTextMessage(pnet.MSG_STATUS_SMALL, fmt.Sprintf("Message sent to %v", toPlayer.GetName()))
	
	return true
}

func (g *Game) playerTalkToChannel(_player *Player, _type int, _text string, _channelId int) bool {
	return g.Chat.TalkToChannel(_player, _type, _text, _channelId)
}

func (g *Game) internalBroadcastMessage(_creature pul.ICreature, _message string) bool {
	// TOOD: Check if _creature is CanBroadcast player flag

	g.mutexPlayerList.RLock()
	defer g.mutexPlayerList.RUnlock()
	for _, player := range g.Players {
		player.sendCreatureSay(_creature, pnet.SPEAK_BROADCAST, _message)
	}
	
	return true
}

func (g *Game) internalCreatureChangeVisible(_creature pul.ICreature, _visible bool) {
	list := _creature.GetVisibleCreatures()
	for _, tmpCreature := range list {
		if tmpCreature.GetType() == CTYPE_PLAYER {
			tmpCreature.(*Player).sendCreatureChangeVisibility(_creature, _visible)
		}
	}
}

func (g *Game) checkForWildEncounter(_creature pul.ICreature) {
	if _creature.GetType() == CTYPE_PLAYER {
		// Do some checkin'
	}
}