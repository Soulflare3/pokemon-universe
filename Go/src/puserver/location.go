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
	pos "putools/pos"
	puh "puhelper"
)

type Location struct {
	ID         int
	Name       string
	Music      int
	PokeCenter pos.Position
}

func (l *Location) GetId() int {
	return l.ID
}

func (l *Location) GetName() string {
	return l.Name
}

func (l *Location) GetMusicId() int {
	return l.Music
}

func (l *Location) GetPokecenter() pos.Position {
	return l.PokeCenter
}

type LocationMap map[int]*Location
type LocationStore struct {
	Locations LocationMap
}

func NewLocationStore() *LocationStore {
	return &LocationStore{Locations: make(LocationMap)}
}

func (store *LocationStore) Load() error {
	var query string = "SELECT t.idlocation, t.name, t.idmusic, p.position FROM location t LEFT JOIN pokecenter p ON p.idpokecenter = t.idpokecenter"
	result, err := puh.DBQuerySelect(query)
	if err != nil {
		return err
	}

	defer result.Free()
	for {
		row := result.FetchMap()
		if row == nil {
			break
		}

		idlocation := puh.DBGetInt(row["idlocation"])
		name := puh.DBGetString(row["name"])
		music := puh.DBGetInt(row["idmusic"])
		pokecenter := puh.DBGetInt64(row["position"]) // Hash
		pcposition := pos.NewPositionFromHash(pokecenter)

		location := &Location{ID: idlocation,
			Name:       name,
			Music:      music,
			PokeCenter: pcposition}
		store.addLocation(location)
	}

	return nil
}

func (store *LocationStore) addLocation(_location *Location) {
	_, found := store.Locations[_location.ID]
	if found == false {
		store.Locations[_location.ID] = _location
	}
}

func (store *LocationStore) GetLocation(_idx int) (location *Location, found bool) {
	location, found = store.Locations[_idx]
	return
}
