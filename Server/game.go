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

type GameState int
const (
	GAME_STATE_STARTUP GameState = iota
	GAME_STATE_INIT
	GAME_STATE_NORMAL
	GAME_STATE_CLOSED
	GAME_STATE_CLOSING
)

type PlayerList map[uint64]*Player

type Game struct {
	State		GameState
	Creatures	CreatureList
	Players		PlayerList
	
	WorldMap	*Map
	Locations	*LocationStore
}

func NewGame() *Game {
	game := Game{}
	game.State = GAME_STATE_STARTUP
	// Initialize maps
	game.Creatures = make(CreatureList)
	game.Players = make(PlayerList)
	
	return &game
}

func (the *Game) Load() (LostIt bool) {
	LostIt = true // fuck >:(
	the.WorldMap = NewMap()
	the.Locations = NewLocationStore()
	
	g_logger.Println(" - Loading locations")
	// Load locations
	if err := the.Locations.Load(); err != nil {
		g_logger.Println(err)
		LostIt = false
	}
	
	// Load worldmap
	g_logger.Println(" - Loading worldmap")
	if err := the.WorldMap.Load(); err != nil {
		g_logger.Println(err)
		LostIt = false
	}
	
	return
}

func (g *Game) AddPlayer(_player *Player) {
	uid := _player.GetUID()
	
	if g.Players[uid] == nil {
		g.Players[uid] = _player
	}
	
	if g.Creatures[uid] == nil {
		g.Creatures[uid] = _player
	}
}

func (g *Game) GetPlayerByName(_name string) *Player {
	for _, value := range g.Players {
		if value.GetName() == _name {
			return value
		}
	}
	
	return nil
}
