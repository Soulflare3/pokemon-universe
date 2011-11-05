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
)

type PlayerList map[uint64]*Player

type Player struct {
	Creature // Inherit generic creature data
	dbid			int // database id

	Conn			*Connection

	Pokemon			PlayerPokemonList
	PokemonParty	*PokemonParty
	
	Location		*Location
	LastPokeCenter	*Tile

	Money          int
	TimeoutCounter int
}

func NewPlayer(_name string) *Player {
	p := Player{}
	p.uid = GenerateUniqueID()
	p.Conn = nil
	p.Outfit = NewOutfit()
	p.name = _name

	p.Pokemon = make(PlayerPokemonList)
	p.PokemonParty = NewPokemonParty()
	p.lastStep = PUSYS_TIME()
	p.moveSpeed = 280
	p.VisibleCreatures = make(CreatureList)
	p.TimeoutCounter = 0

	return &p
}

func (p *Player) LoadData() {
	// Load player info
	if !p.loadPlayerInfo() {
		// TODO: Unload player and disconnect
	}
	
	// Load all pokemon player has
	if !p.loadPokemon() {
		// TODO: Unload player and disconnect
	}
}

func (p *Player) loadPlayerInfo() bool {
	var query string = "SELECT p.idplayer, p.name, p.position, p.movement, p.idpokecenter, p.money, p.idlocation," +
						" g.group_idgroup, o.head, o.nek, o.upper, o.lower, o.feet FROM player `p`" +
						" INNER JOIN player_outfit `o` ON o.idplayer = p.idplayer" +
						" INNER JOIN player_group `g` ON g.player_idplayer = p.idplayer" +
						" WHERE p.name='%s'"
	result, err := DBQuerySelect(fmt.Sprintf(query, p.name))
	if err != nil {
		return false
	}
	
	defer result.Free()
	row := result.FetchRow()
	if row == nil {
		g_logger.Printf("[Error] No data for player %s (%d)", p.name, p.dbid)
		return false;
	}
		
	p.dbid = row[0].(int)
	p.name = row[1].(string)
	tile, ok := g_map.GetTile(row[2].(int64))
	if !ok {
		g_logger.Printf("[Error] Could not load position info for player %s (%d)", p.name, p.dbid)
		return false
	}
	p.Position = tile
	p.Movement = row[3].(int)
	// TODO: p.LastPokeCenter = row[4].(int)
	p.Money = row[5].(int)
	location, ok := g_game.Locations.GetLocation(row[6].(int))
	if !ok {
		g_logger.Printf("[Error] Could not load location info for player %s (%d)", p.name, p.dbid)
		return false
	}
	p.Location = location
	
	// Group/Right stuff : row[7].(int)
	
	p.SetOutfitKey(OUTFIT_HEAD, row[8].(int))
	p.SetOutfitKey(OUTFIT_NEK, row[9].(int))
	p.SetOutfitKey(OUTFIT_UPPER, row[10].(int))
	p.SetOutfitKey(OUTFIT_LOWER, row[11].(int))
	p.SetOutfitKey(OUTFIT_FEET, row[12].(int))
	
	return true
}

func (p *Player) loadPokemon() bool {
	var query string = "SELECT idpokemon, nickname, bound, experience, iv_hp, iv_attack, iv_attack_spec, iv_defence, iv_defence_spec," +
						" iv_speed, happiness, gender, in_party, party_slot WHERE idplayer='%d'"
	result, err := DBQuerySelect(fmt.Sprintf(query, p.dbid))
	if err != nil {
		return false
	}
	
	defer result.Free()
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		pokemon := NewPlayerPokemon()
		pokemonId := row[0].(int)
		pokemon.Base = g_PokemonManager.GetPokemon(pokemonId)
		pokemon.Nickname = row[1].(string)
		pokemon.IsBound = row[2].(int)
		pokemon.Experience = row[3].(int)
		pokemon.Stats[0] = row[4].(int)
		pokemon.Stats[1] = row[5].(int)
		pokemon.Stats[2] = row[7].(int)
		pokemon.Stats[3] = row[6].(int)
		pokemon.Stats[4] = row[8].(int)
		pokemon.Stats[5] = row[9].(int)
		pokemon.Happiness = row[10].(int)
		pokemon.Gender = row[11].(int)
		pokemon.InParty = row[12].(int)
		pokemon.Slot = row[13].(int)
		
		// Add to party if needed
		if pokemon.InParty == 1 {
			p.PokemonParty.AddSlot(pokemon, pokemon.Slot)
		}
	}
	
	return true
}

func (p *Player) GetType() int {
	return CTYPE_PLAYER
}

func (p *Player) SetConnection(_conn *Connection) {
	p.Conn = _conn
	p.Conn.Owner = p
	go _conn.HandleConnection()
}

// Called by Connection to remove itself from its owner
// when the player disconnects
func (p *Player) removeConnection() {
	if !p.Conn.IsOpen {
		g_game.OnPlayerLoseConnection(p)
	}
}

func (p *Player) SetMoney(_money int) int {
	if p.Money += _money; p.Money < 0 {
		p.Money = 0
	}
	return p.Money
}

func (p *Player) OnCreatureMove(_creature ICreature, _from *Tile, _to *Tile, _teleport bool) {
	if _creature.GetUID() == p.GetUID() {
		p.lastStep = PUSYS_TIME()
		return
	}

	canSeeFromTile := CanSeePosition(p.GetPosition(), _from.Position)
	canSeeToTile := CanSeePosition(p.GetPosition(), _to.Position)

	if canSeeFromTile && !canSeeToTile { // Leaving viewport
		p.sendCreatureMove(_creature, _from, _to)

		p.RemoveVisibleCreature(_creature)
		_creature.RemoveVisibleCreature(p)
	} else if canSeeToTile && !canSeeFromTile { // Entering viewport
		p.AddVisibleCreature(_creature)
		_creature.RemoveVisibleCreature(p)

		p.sendCreatureMove(_creature, _from, _to)
	} else { // Moving inside viewport
		p.AddVisibleCreature(_creature)
		_creature.AddVisibleCreature(p)

		p.sendCreatureMove(_creature, _from, _to)
	}
}

func (p *Player) OnCreatureTurn(_creature ICreature) {
	if _creature.GetUID() != p.GetUID() {
		p.sendCreatureTurn(_creature)
	}
}

func (p *Player) OnCreatureAppear(_creature ICreature, _isLogin bool) {
	canSeeCreature := CanSeeCreature(p, _creature)
	if !canSeeCreature {
		return
	}

	// We're checking inside the AddVisibleCreature method so no need to check here
	p.AddVisibleCreature(_creature)
	_creature.AddVisibleCreature(p)
}

func (p *Player) OnCreatureDisappear(_creature ICreature, _isLogout bool) {
	// TODO: Have to do something here with _isLogout

	p.RemoveVisibleCreature(_creature)
}

func (p *Player) AddVisibleCreature(_creature ICreature) {
	if _, found := p.VisibleCreatures[_creature.GetUID()]; !found {
		p.VisibleCreatures[_creature.GetUID()] = _creature
		p.sendCreatureAdd(_creature)
	}
}

func (p *Player) RemoveVisibleCreature(_creature ICreature) {
	// No need to check if the key actually exists because Go is awesome
	// http://golang.org/doc/effective_go.html#maps
	p.VisibleCreatures[_creature.GetUID()] = nil, false
	p.sendCreatureRemove(_creature)
}

// ------------------------------------------------------ //
func (p *Player) sendMapData(_dir int) {
	if p.Conn != nil {
		p.Conn.Send_Tiles(_dir, p.GetPosition())
	}
}

func (p *Player) sendCreatureMove(_creature ICreature, _from, _to *Tile) {
	if p.Conn != nil {
		p.Conn.Send_CreatureWalk(_creature, _from, _to)
	}
}

func (p *Player) sendCreatureTurn(_creature ICreature) {
	if p.Conn != nil {
		p.Conn.Send_CreatureTurn(_creature, p.GetDirection())
	}
}

func (p *Player) sendCreatureAdd(_creature ICreature) {
	if p.Conn != nil {
		p.Conn.Send_CreatureAdd(_creature)
	}
}

func (p *Player) sendCreatureRemove(_creature ICreature) {
	if p.Conn != nil {
		p.Conn.Send_CreatureRemove(_creature)
	}
}

func (p *Player) sendPlayerWarp() {
	if p.Conn != nil {
		p.Conn.Send_PlayerWarp(p.GetPosition())
	}
}

func (p *Player) sendCreatureSay(_creature ICreature, _speakType int, _message string, _channelId int) {
	if p.Conn != nil {
		p.Conn.Send_CreatureChat(_creature, _channelId, _speakType, _message)
	}
}
