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

func (c *Connection) Send_PlayerData() {
	playerData := pnet.NewData_PlayerData()
	playerData.PlayerData.UID			= c.Owner.GetUID()
	playerData.PlayerData.Name			= c.Owner.GetName()
	playerData.PlayerData.X				= c.Owner.GetPosition().X
	playerData.PlayerData.Y				= c.Owner.GetPosition().Y
	playerData.PlayerData.Direction 	= c.Owner.Direction
	playerData.PlayerData.Money			= c.Owner.Money
	
	for i := 0; i < 6; i++ {
		outfit := pnet.NewBodyPart(c.Owner.GetOutfitStyle(OutfitPart(i)), uint32(c.Owner.GetOutfitColour(OutfitPart(i))))
		playerData.PlayerData.Outfit[i] = outfit
	}
	
	c.SendMessage(playerData)
	
	//ToDo: Send PkMn
	
	//ToDo: Send items
	
	// Send map
	c.Send_Tiles(DIR_NULL, c.Owner.GetPosition())
	
	// ready
	readyMessage := pnet.NewData_LoginStatus()
	readyMessage.LoginStatus.Status = LOGINSTATUS_READY
	c.SendMessage(readyMessage)
}