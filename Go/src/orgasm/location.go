package main

import (	
	"pulogic/models"
	"nonamelib/log"
)

type LocationsList struct {
	locations map[int]models.Location
	musics map[int]models.Music
	pokecenters map[int]models.Pokecenter	
}

func NewLocationsList() *LocationsList {
	return &LocationsList { 	locations: make(map[int]models.Location),
					musics: make(map[int]models.Music),
					pokecenters: make(map[int]models.Pokecenter) }
}

func (l *LocationsList) LoadLocations() bool {
	if !l.LoadPokecenters() || !l.LoadMusic() {
		return false
	}

	var locats []models.Location
	if err := g_orm.FindAll(&locats); err != nil {
		log.Error("location", "LoadLocations", "Error while loading locations: %v", err.Error())
		return false
	}
	
	for _, locationEntity := range locats {
		l.locations[locationEntity.Idlocation] = locationEntity
	}
	return true
}

func (l *LocationsList) LoadPokecenters() bool {
	var pokec []models.Pokecenter
	if err := g_orm.FindAll(&pokec); err != nil {
		log.Error("location", "LoadPokecenters", "Error while loading pokecenters: %v", err.Error())
		return false
	}
	
	for _, pokecenterEntity := range pokec {
		l.pokecenters[pokecenterEntity.Idpokecenter] = pokecenterEntity
	}
	return true
}

func (l *LocationsList) LoadMusic() bool {
	var mus []models.Music
	if err := g_orm.FindAll(&mus); err != nil {
		log.Error("location", "LoadPokecenters", "Error while loading music: %v", err.Error())
		return false
	}
	
	for _, musicEntity := range mus {
		l.musics[musicEntity.Idmusic] = musicEntity
	}
	return true
}

func (l *LocationsList) GetNumMusic() int {
	return len(l.musics)
}

func (l *LocationsList) GetNumPokecenters() int {
	return len(l.pokecenters)
}

func (l *LocationsList) GetNumLocations() int {
	return len(l.locations)
}