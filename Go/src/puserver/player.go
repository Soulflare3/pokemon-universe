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
	"container/list"
	
	"putools/log"
	pul "pulogic"
	pkmn "pulogic/pokemon"
	puh "puhelper"
)

type PlayerList map[uint64]*Player

type Player struct {
	Creature     		// Inherit generic creature data
	dbid     		int // database id

	Conn 			*Connection

	Pokemon			pkmn.PlayerPokemonList
	PokemonParty	*pkmn.PokemonParty

	Location       	*Location
	LastPokeCenter 	*Tile

	Money          	int
	TimeoutCounter	int
	GroupFlags		int64
}

func NewPlayer(_name string) *Player {
	p := &Player{}
	p.uid = puh.GenerateUniqueID()
	p.Conn = nil
	p.Outfit = NewOutfit()
	p.name = _name

	p.Pokemon = make(pkmn.PlayerPokemonList)
	p.PokemonParty = pkmn.NewPokemonParty()
	p.lastStep = PUSYS_TIME()
	p.moveSpeed = 250
	p.VisibleCreatures = make(pul.CreatureList)
	p.ConditionList = list.New()
	p.TimeoutCounter = 0
	
	// Add self to visible creatures
	p.VisibleCreatures[p.GetUID()] = p

	return p
}

func (p *Player) LoadData() bool {
	// Load player info
	if !p.loadPlayerInfo() {
		return false
	}

	// Load all pokemon player has
	if !p.loadPokemon() {
		return false
	}

	return true
}

func (p *Player) loadPlayerInfo() bool {
	var query string = "SELECT p.idplayer, p.name, p.position, p.movement, p.idpokecenter, p.money, p.idlocation," +
		" g.group_idgroup, o.head, o.nek, o.upper, o.lower, o.feet FROM player `p`" +
		" INNER JOIN player_outfit `o` ON o.idplayer = p.idplayer" +
		" INNER JOIN player_group `g` ON g.player_idplayer = p.idplayer" +
		" WHERE p.idplayer=%d"
	result, err := puh.DBQuerySelect(fmt.Sprintf(query, p.dbid))
	if err != nil {
		return false
	}

	defer result.Free()
	row := result.FetchRow()
	if row == nil {
		logger.Printf("[Error] No pokemon data for player %s (DB ID: %d)\n", p.name, p.dbid)
		return false
	}

	p.dbid = puh.DBGetInt(row[0])
	p.name = puh.DBGetString(row[1])
	tile, ok := g_map.GetTile(puh.DBGetInt64(row[2]))
	if !ok {
		logger.Printf("[Warning] Could not load position info for player %s (%d)\n", p.name, p.dbid)
		tile, _ = g_map.GetTileFrom(-510, -236, 0)
		if tile == nil {
			logger.Println("[Error] Could not load default position")
			return false
		}
	}
	p.Position = tile
	p.SetDirection(DIR_SOUTH)
	p.Movement = puh.DBGetInt(row[3])
	// TODO: p.LastPokeCenter = row[4].(int)
	p.Money = puh.DBGetInt(row[5])
	location, ok := g_game.Locations.GetLocation(puh.DBGetInt(row[6]))
	if !ok {
		logger.Printf("[Error] Could not load location info for player %s (%d)\n", p.name, p.dbid)
		return false
	}
	p.Location = location 

	// Group/Right stuff : row[7].(int)

	p.SetOutfitKey(pul.OUTFIT_HEAD, puh.DBGetInt(row[8]))
	p.SetOutfitKey(pul.OUTFIT_NEK, puh.DBGetInt(row[9]))
	p.SetOutfitKey(pul.OUTFIT_UPPER, puh.DBGetInt(row[10]))
	p.SetOutfitKey(pul.OUTFIT_LOWER, puh.DBGetInt(row[11]))
	p.SetOutfitKey(pul.OUTFIT_FEET, puh.DBGetInt(row[12]))

	return true
}

func (p *Player) loadPokemon() bool {
	var query string = "SELECT idpokemon, nickname, bound, experience, iv_hp, iv_attack, iv_attack_spec, iv_defence, iv_defence_spec," +
		" iv_speed, happiness, gender, in_party, party_slot, idplayer_pokemon, shiny, abilityId, damagedHp FROM player_pokemon WHERE idplayer='%d'"
	result, err := puh.DBQuerySelect(fmt.Sprintf(query, p.dbid))
	if err != nil {
		return false
	}

	logger.Println("Loading player pokemon..")

	for {
		row := result.FetchRow()
		if row == nil {
			break
		}

		pokemon := pkmn.NewPlayerPokemon(p.dbid)
		pokemon.IdDb = puh.DBGetInt(row[14])
		pokemonId := puh.DBGetInt(row[0])
		pokemon.Base = pkmn.GetInstance().GetPokemon(pokemonId)
		pokemon.Nickname = puh.DBGetString(row[1])
		pokemon.IsBound = puh.DBGetInt(row[2])
		pokemon.Experience = puh.DBGetFloat64(row[3])
		pokemon.Stats[0] = puh.DBGetInt(row[4])
		pokemon.Stats[1] = puh.DBGetInt(row[5])
		pokemon.Stats[2] = puh.DBGetInt(row[7])
		pokemon.Stats[3] = puh.DBGetInt(row[6])
		pokemon.Stats[4] = puh.DBGetInt(row[8])
		pokemon.Stats[5] = puh.DBGetInt(row[9])
		pokemon.Happiness = puh.DBGetInt(row[10])
		pokemon.Gender = puh.DBGetInt(row[11])
		pokemon.InParty = puh.DBGetInt(row[12])
		pokemon.Slot = puh.DBGetInt(row[13])
		pokemon.IsShiny = puh.DBGetInt(row[15])
		abilityId := puh.DBGetInt(row[16])
		pokemon.DamagedHp = puh.DBGetInt(row[17])
		
		pokemon.Ability = pkmn.GetInstance().GetAbilityById(abilityId)
		if pokemon.Ability == nil {
			logger.Printf("[Warning] Pokemon (%d) has an invalid abilityId (%d)\n", pokemon.IdDb, abilityId)
			pokemon.Ability = pkmn.GetInstance().GetAbilityById(96)
		}

		// Add to party if needed
		if pokemon.InParty == 1 {
			p.PokemonParty.AddSlot(pokemon, pokemon.Slot)
		}
	}
	result.Free()

	// Load moves for each pokemon
	for index, pokemon := range p.PokemonParty.Party {
		if pokemon != nil {
			fmt.Printf("Load pokemon moves for: %d - %s\n\r", index, pokemon.GetNickname())
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
}

// Called by Connection to remove itself from its owner
// when the player disconnects
func (p *Player) removeConnection() {
	if p.Conn == nil || !p.Conn.IsOpen {
		g_game.OnPlayerLoseConnection(p)
	}
}

func (p *Player) SetMoney(_money int) int {
	if p.Money += _money; p.Money < 0 {
		p.Money = 0
	}
	return p.Money
}

func (p *Player) GetPokemonParty() *pkmn.PokemonParty {
	return p.PokemonParty
}

func (p *Player) OnCreatureMove(_creature pul.ICreature, _from pul.ITile, _to pul.ITile, _teleport bool) {
	if _creature.GetUID() == p.GetUID() {
		p.lastStep = PUSYS_TIME()
		return
	}
	
	from := _from.(*Tile)
	to := _to.(*Tile)

	canSeeFromTile := CanSeePosition(p.GetPosition(), from.Position)
	canSeeToTile := CanSeePosition(p.GetPosition(), to.Position)

	if canSeeFromTile && !canSeeToTile { // Leaving viewport
		p.sendCreatureMove(_creature, from, to)

		p.RemoveVisibleCreature(_creature)
		_creature.RemoveVisibleCreature(p)
	} else if canSeeToTile && !canSeeFromTile { // Entering viewport
		p.AddVisibleCreature(_creature)
		_creature.RemoveVisibleCreature(p)

		p.sendCreatureMove(_creature, from, to)
	} else { // Moving inside viewport
		p.AddVisibleCreature(_creature)
		_creature.AddVisibleCreature(p)

		p.sendCreatureMove(_creature, from, to)
	}
}

func (p *Player) OnCreatureTurn(_creature pul.ICreature) {
	if _creature.GetUID() != p.GetUID() {
		p.sendCreatureTurn(_creature)
	}
}

func (p *Player) OnCreatureAppear(_creature pul.ICreature, _isLogin bool) {
	if _creature.GetUID() == p.GetUID() {
		return
	}
	
	canSeeCreature := CanSeeCreature(p, _creature)
	if !canSeeCreature {
		return
	}

	// We're checking inside the AddVisibleCreature method so no need to check here
	p.AddVisibleCreature(_creature)
	_creature.AddVisibleCreature(p)
}

func (p *Player) OnCreatureDisappear(_creature pul.ICreature, _isLogout bool) {
	if _creature.GetUID() == p.GetUID() {
		return
	}
	
	// TODO: Have to do something here with _isLogout

	p.RemoveVisibleCreature(_creature)
}

func (p *Player) AddVisibleCreature(_creature pul.ICreature) {
	if _, found := p.VisibleCreatures[_creature.GetUID()]; !found {
		p.VisibleCreatures[_creature.GetUID()] = _creature
		p.sendCreatureAdd(_creature)
	}
}

func (p *Player) RemoveVisibleCreature(_creature pul.ICreature) {
	// No need to check if the key actually exists because Go is awesome
	// http://golang.org/doc/effective_go.html#maps
	delete(p.VisibleCreatures, _creature.GetUID())
	p.sendCreatureRemove(_creature)
}


// ------------------------------------------------------ //
func (p *Player) sendCancelMessage(_message int) {
	switch _message {		
		case RET_YOUAREEXHAUSTED:
			p.Conn.SendCancel("You are exhausted.")
		case RET_NOTPOSSIBLE:
			fallthrough
		default:
			p.Conn.SendCancel("Sorry, not possible.")
	}
}

func (p *Player) sendTextMessage(_mclass int, _message string) {
	if p.Conn != nil {
		// p.Conn.sendTextMessage(_mclass, _message)
	}
}

func (p *Player) sendMapData(_dir int) {
	if p.Conn != nil {
		p.Conn.SendMapData(_dir, p.GetPosition())
	}
}

func (p *Player) sendCreatureMove(_creature pul.ICreature, _from, _to pul.ITile) {
	if p.Conn != nil {
		p.Conn.SendCreatureMove(_creature, _from.(*Tile), _to.(*Tile))
	}
}

func (p *Player) sendCreatureTurn(_creature pul.ICreature) {
	if p.Conn != nil {
		p.Conn.SendCreatureTurn(_creature, p.GetDirection())
	}
}

func (p *Player) sendCreatureAdd(_creature pul.ICreature) {
	if p.Conn != nil {
		p.Conn.SendCreatureAdd(_creature)
	}
}

func (p *Player) sendCreatureRemove(_creature pul.ICreature) {
	if p.Conn != nil {
		p.Conn.SendCreatureRemove(_creature)
	}
}

func (p *Player) sendPlayerWarp() {
	if p.Conn != nil {
		p.Conn.SendPlayerWarp(p.GetPosition())
	}
}

func (p *Player) sendCreatureSay(_creature pul.ICreature, _speakType int, _message string) {
	if p.Conn != nil {
		//p.Conn.SendCreatureSay(_creature, _speakType, _message)
	}
}

func (p *Player) sendCreatureChangeVisibility(_creature pul.ICreature, _visible bool) {
	if _creature.GetUID() != p.GetUID() {
		if _visible {
			p.AddVisibleCreature(_creature)
		} else if !p.hasFlag(PlayerFlag_CanSenseInvisibility) {
			p.RemoveVisibleCreature(_creature)
		}
	}
}

// --------------------- CHAT ----------------------------//
func (p *Player) sendClosePrivateChat(_channelId int) {

}

func (p *Player) sendToChannel(_fromPlayer pul.ICreature, _type int, _text string, _channelId int, _time int) {
	p.Conn.SendCreatureSay(_fromPlayer, _type, _text, _channelId, _time)
}

// ------------------------------------------------------ //
func (p *Player) HealParty() {
	p.PokemonParty.HealParty()
	
	// TODO: Send update to client
}

func (p *Player) setFlags(_flags int64) {
	p.GroupFlags = _flags
}

func (p *Player) hasFlag(_value uint64) bool {
	return (0 != (p.GroupFlags & (1 << _value)))
}