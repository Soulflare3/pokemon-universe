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
	"os"
	"mysql"
	pos "position"
)

type Location struct {
	ID				int32
	Name			string
	Music			int32
	PokeCenter		pos.Position
}

type LocationMap map[int32]*Location
type LocationStore struct {
	Locations	LocationMap
}

func NewLocationStore() *LocationStore {
	return &LocationStore{ Locations: make(LocationMap) }
}

func (store *LocationStore) Load() (err os.Error) {
	var res *mysql.MySQLResult
	var row map[string]interface{}
	
	var query string = "SELECT t.idlocation, t.name, t.idmusic, p.position FROM location t LEFT JOIN pokecenter p ON p.idpokecenter = t.idpokecenter"
	if res, err = g_db.Query(query); err != nil {
		return
	}
	
	for {
		if row = res.FetchMap(); row == nil {
			break
		}
		
		idlocation	:= row["idlocation"].(int32)
		name		:= row["name"].(string)
		music		:= row["idmusic"].(int32)
		pokecenter	:= row["position"].(int64) // Hash
		pcposition 	:= pos.NewPositionFromHash(pokecenter)
		
		location := &Location { ID: idlocation,
								Name: name,
								Music: music,
								PokeCenter: pcposition }
		store.addLocation(location)
	}
	
	return
}

func (store *LocationStore) addLocation(_location *Location) {
	_, found := store.Locations[_location.ID]
	if found == false {
		store.Locations[_location.ID] = _location
	}
}

func (store *LocationStore) GetLocation(_idx int32) (location *Location, found bool) {
	location, found = store.Locations[_idx]
	return
}
