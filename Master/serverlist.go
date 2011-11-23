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
	"sync"
	"time"

	"mysql"
)

type ServerInfo struct {
	Name      string
	Ip        string
	Online    uint8
	LastCheck int64
}

type ServerStore struct {
	servers        *ServerList
	DataChangeTime int64
}

func NewServerStore() *ServerStore {
	s := &ServerStore{servers: NewServerList(), DataChangeTime: time.Seconds()}

	if err := s.load(); err != nil {
		g_logger.Printf("ServerStore: %v\n\r", err)
	}

	// Disabled for now. Just restart the server to reload gameservers
	//go s.reloadLoop()
	return s
}

func (store *ServerStore) load() (err error) {
	var res *mysql.MySQLResult
	var row map[string]interface{}

	var queryString string = "SELECT idserver, name, address FROM servers"
	if res, err = g_db.Query(queryString); err != nil {
		return err
	}

	for {
		if row = res.FetchMap(); row == nil {
			break
		}

		idserver := row["idserver"].(int)
		name := row["name"].(string)
		address := row["address"].(string)

		info := ServerInfo{Name: name, Ip: address, Online: 0, LastCheck: 0}
		store.servers.Set(idserver, info)
	}

	return nil
}

func (store *ServerStore) reloadLoop() {
	for {

		// Check if all the servers are online
		store.servers.CheckStatus()

		// Sleep for 5 min
		time.Sleep(30e10)

		// Reload server list from database
		// ---
	}
}

type ServerList struct {
	mu      sync.RWMutex
	servers map[int]ServerInfo
}

func NewServerList() *ServerList {
	return &ServerList{servers: make(map[int]ServerInfo)}
}

func (list *ServerList) CheckStatus() {
	/*for _, value := range list.servers {
		// Ping each server
		// Need to find a good/cheap way to do this...
		println(value.Name)
	}*/
}

func (list *ServerList) Get(key int) (info ServerInfo, ok bool) {
	list.mu.RLock()
	info, ok = list.servers[key]
	list.mu.RUnlock()

	return
}

func (list *ServerList) Set(key int, server ServerInfo) {
	if _, ok := list.Get(key); !ok {
		list.mu.Lock()
		list.servers[key] = server
		list.mu.Unlock()
	}

	return
}
