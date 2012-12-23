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
	"nonamelib/log"
	pos "nonamelib/pos"
	"pulogic/models"
)

type Warp struct {
	dbid int64
	destination pos.Position
	
	IsNew bool
	IsModified bool
	IsRemoved bool
}

func NewWarp(_destination pos.Position) *Warp {
	return &Warp{dbid: 0, destination: _destination, IsNew: true}
}

func NewWarpFromPacket(_packet *Packet) *Warp {
	toX := int(_packet.ReadInt16())
	toY := int(_packet.ReadInt16())
	toZ := int(_packet.ReadInt16())
	tp_pos := pos.NewPositionFrom(toX, toY, toZ)
	
	return NewWarp(tp_pos)
}

func (e *Warp) GetDbId() int64 {
	return e.dbid
}

func (e *Warp) GetEventType() int {
	return EVENTTYPES_TELEPORT
}

func (e *Warp) SetDestination(_pos pos.Position) {
	e.destination = _pos
	e.IsModified = true
}

func (e *Warp) ToPacket(_packet *Packet) {
	_packet.AddUint8(uint8(e.GetEventType()))
	_packet.AddUint16(uint16(e.destination.X))
	_packet.AddUint16(uint16(e.destination.Y))
	_packet.AddUint16(uint16(e.destination.Z))
}

func (e *Warp) UpdateFromPacket(_packet *Packet) {
	e.destination.X = int(_packet.ReadInt16())
	e.destination.Y = int(_packet.ReadInt16())
	e.destination.Z = int(_packet.ReadInt16())
	
	e.IsModified = true
}

func (e *Warp) Save() bool {
//	var query string
//	if e.IsNew { // Update
//		query = fmt.Sprintf(QUERY_INSERT_EVENT, e.GetEventType(), e.destination.X, e.destination.Y, e.destination.Z, "", "", "", "", "")
//	} else if e.IsModified { // Insert
//		query = fmt.Sprintf(QUERY_UPDATE_EVENT, e.GetEventType(), e.destination.X, e.destination.Y, e.destination.Z, "", "", "", "", "", e.dbid)
//	}
//	
//	if len(query) > 0 {
//		if err := puh.DBQuery(query); err != nil {
//			return false
//		}
//		
//		if e.IsNew {
//			e.dbid = int64(puh.DBGetLastInsertId())
//		}		
//	}

	entity := models.TileEvents { IdtileEvents: int(e.dbid),
								  Eventtype: e.GetEventType(),
								  Param1: fmt.Sprintf("%d", e.destination.X),
								  Param2: fmt.Sprintf("%d", e.destination.Y),
								  Param3: fmt.Sprintf("%d", e.destination.Z) }
	if err := g_orm.Save(&entity); err != nil {
		log.Error("Warp", "Save", "Error saving: %s", err.Error())
		return false
	}
	
	if e.IsNew {
		e.dbid = int64(entity.IdtileEvents)
	}
	
	e.IsNew = false
	e.IsModified = false

	return true
}

func (e *Warp) Delete() bool {
//	query := fmt.Sprintf(QUERY_DELETE_EVENT, e.dbid)
//	if err := puh.DBQuery(query); err != nil {
//		return false
//	}

	entity := models.TileEvents { IdtileEvents: int(e.dbid) }
	if _, err := g_orm.Delete(&entity); err != nil {
		log.Error("Warp", "Save", "Error deleting: %s", err.Error())
		return false
	}
	
	e.IsRemoved = true
	return true
}