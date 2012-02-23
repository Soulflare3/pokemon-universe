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
	pnet "network"
)

type SendPokemonData struct {
	Pokemon *PokemonParty
}

func (m *SendPokemonData) GetHeader() uint8 {
	return pnet.HEADER_POKEMONPARTY
}

// WritePacket write the needed object data to a Packet and returns it
func (m *SendPokemonData) WritePacket() pnet.IPacket {
	packet := pnet.NewPacketExt(m.GetHeader())
	
	for i := 0; i < 6; i++ {
		packet.AddUint8(uint8(i)) // Slot
		
		// Retreive pokemon from party
		pokemon := m.Pokemon.GetFromSlot(i)
		if pokemon == nil {
			packet.AddUint32(0) // Zero if slot is empty
		} else {
			pokemonLevel := pokemon.GetLevel()
		
			packet.AddUint32(uint32(pokemon.IdDb)) // Database ID
			packet.AddUint16(uint16(pokemon.Base.Species.SpeciesId)) // Real ID
			packet.AddString(pokemon.GetNickname()) // Name/Nickname
			packet.AddUint16(uint16(pokemonLevel)) // Level
			packet.AddUint32(uint32(pokemon.Experience)) // Current EXP 
			packet.AddUint32(uint32(ExperienceForLevel(pokemonLevel + 1))) // Exp for next level
			
			packet.AddUint16(uint16(pokemon.Base.Types[0]))
			packet.AddUint16(uint16(pokemon.Base.Types[1]))
			packet.AddUint16(uint16(pokemon.Nature))
			packet.AddUint8(uint8(pokemon.Gender))
			
			packet.AddUint8(uint8(pokemon.Stats[0]))
			packet.AddUint8(uint8(pokemon.Stats[1]))
			packet.AddUint8(uint8(pokemon.Stats[2]))
			packet.AddUint8(uint8(pokemon.Stats[3]))
			packet.AddUint8(uint8(pokemon.Stats[4]))
			packet.AddUint8(uint8(pokemon.Stats[5]))
			
			// Pokemon moves
			for j := 0; j < 4; j++ {
				packet.AddUint8(uint8(j))
				
				// Get pokemon move
				pokemonMove := pokemon.Moves[j]
				if pokemonMove == nil {
					packet.AddUint32(0)
				} else {
					packet.AddUint32(uint32(pokemonMove.Move.MoveId))
					packet.AddString(pokemonMove.Move.Identifier)
					packet.AddString(pokemonMove.Move.FlavorText)
					
					moveTypeStr := GetTypeValueById(pokemonMove.Move.TypeId)
					packet.AddString(moveTypeStr)
					
					packet.AddUint8(uint8(pokemonMove.CurrentPP))
					packet.AddUint8(uint8(pokemonMove.Move.PP))
					packet.AddUint8(uint8(pokemonMove.Move.Power))
					packet.AddUint8(uint8(pokemonMove.Move.Accuracy))
					packet.AddUint8(uint8(pokemonMove.Move.TargetId))
				}
			}
		}
	}

	return packet
}