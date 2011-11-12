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

func (p *Player) LoadData() bool {
	// Load player info
	if !p.loadPlayerInfo() {
		// TODO: Unload player and disconnect
		
		return false
	}
	
	// Load all pokemon player has
	if !p.loadPokemon() {
		// TODO: Unload player and disconnect
		
		return false
	}
	
	return true
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
	
	fmt.Printf("Player: %d - %s\n\r", DBGetInt(row[0]), DBGetString(row[1]))
		
	p.dbid = DBGetInt(row[0])
	p.name = DBGetString(row[1])
//	tile, ok := g_map.GetTile(row[2].(int64))
//	if !ok {
//		g_logger.Printf("[Error] Could not load position info for player %s (%d)", p.name, p.dbid)
//		return false
//	}
//	p.Position = tile
	p.Movement = DBGetInt(row[3])
	// TODO: p.LastPokeCenter = row[4].(int)
	p.Money = DBGetInt(row[5])
//	location, ok := g_game.Locations.GetLocation(DBGetInt(row[6]))
//	if !ok {
//		g_logger.Printf("[Error] Could not load location info for player %s (%d)", p.name, p.dbid)
//		return false
//	}
//	p.Location = location 
	
	// Group/Right stuff : row[7].(int)
	
	p.SetOutfitKey(OUTFIT_HEAD, DBGetInt(row[8]))
	p.SetOutfitKey(OUTFIT_NEK, DBGetInt(row[9]))
	p.SetOutfitKey(OUTFIT_UPPER, DBGetInt(row[10]))
	p.SetOutfitKey(OUTFIT_LOWER, DBGetInt(row[11]))
	p.SetOutfitKey(OUTFIT_FEET, DBGetInt(row[12]))
	
	return true
}

func (p *Player) loadPokemon() bool {
	var query string = "SELECT idpokemon, nickname, bound, experience, iv_hp, iv_attack, iv_attack_spec, iv_defence, iv_defence_spec," +
						" iv_speed, happiness, gender, in_party, party_slot, idplayer_pokemon, shiny, abilityId FROM player_pokemon WHERE idplayer='%d'"
	result, err := DBQuerySelect(fmt.Sprintf(query, p.dbid))
	if err != nil {
		fmt.Println(err)
		return false
	}
	
	fmt.Println("Loading player pokemon..")
	
	for {
		row := result.FetchRow()
		if row == nil {
			break
		}
		
		fmt.Printf("Player pokemon: %d\n\r", DBGetInt(row[14]))
		
		pokemon := NewPlayerPokemon(p.dbid)
		pokemon.IdDb = DBGetInt(row[14])
		pokemonId := DBGetInt(row[0])
		pokemon.Base = g_PokemonManager.GetPokemon(pokemonId)
		pokemon.Nickname = DBGetString(row[1])
		pokemon.IsBound = DBGetInt(row[2])
		pokemon.Experience = DBGetInt(row[3])
		pokemon.Stats[0] = DBGetInt(row[4])
		pokemon.Stats[1] = DBGetInt(row[5])
		pokemon.Stats[2] = DBGetInt(row[7])
		pokemon.Stats[3] = DBGetInt(row[6])
		pokemon.Stats[4] = DBGetInt(row[8])
		pokemon.Stats[5] = DBGetInt(row[9])
		pokemon.Happiness = DBGetInt(row[10])
		pokemon.Gender = DBGetInt(row[11])
		pokemon.InParty = DBGetInt(row[12])
		pokemon.Slot = DBGetInt(row[13])
		pokemon.IsShiny = DBGetInt(row[15])
		abilityId := DBGetInt(row[16])
		pokemon.Ability = g_PokemonManager.GetAbilityById(abilityId)
		
		if pokemon.Ability == nil {
			g_logger.Printf("[Warning] Pokemon (%d) has an invalid abilityId (%d)\n\r", pokemon.IdDb, abilityId)
			pokemon.Ability = g_PokemonManager.GetAbilityById(96)
		}
		
		// Add to party if needed
		if pokemon.InParty == 1 {
			p.PokemonParty.AddSlot(pokemon, pokemon.Slot)
		}
	}
	result.Free()
	
	// Load moves for each pokemon
	for index, pokemon := range(p.PokemonParty.Party) {
		if pokemon != nil {
			fmt.Printf("Load pokemon moves for: %d\n\r", index)
			pokemon.LoadMoves()
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
