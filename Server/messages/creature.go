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

func (c *Connection) Send_CreatureAdd(_creature ICreature) {
	if _creature.GetUID() == c.Owner.GetUID() {
		return
	}

	creatureData := pnet.NewData_AddCreature()
	creatureData.AddCreature.UID = _creature.GetUID()
	creatureData.AddCreature.Name = _creature.GetName()
	creatureData.AddCreature.X = _creature.GetPosition().X
	creatureData.AddCreature.Y = _creature.GetPosition().Y
	creatureData.AddCreature.Direction = _creature.GetDirection()

	if player, is_player := _creature.(*Player); is_player {
		for i := 0; i < 6; i++ {
			outfit := pnet.NewBodyPart(player.GetOutfitStyle(OutfitPart(i)), uint32(player.GetOutfitColour(OutfitPart(i))))
			creatureData.AddCreature.Outfit[i] = outfit
		}
	}

	c.SendMessage(creatureData)
}

func (c *Connection) Send_CreatureRemove(_creature ICreature) {
	if _creature.GetUID() == c.Owner.GetUID() {
		return
	}

	msg := pnet.NewData_RemoveCreature()
	msg.RemoveCreature.UID = _creature.GetUID()
	c.SendMessage(msg)
}
