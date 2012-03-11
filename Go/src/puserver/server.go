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
	"crypto/sha1"
	"fmt"
	"hash"
	"strings"
	"net/http"
	"net/websocket"
	"time"
	
	pnet "network"
	pnetmsg "pulogic/netmsg"
	puh "puhelper"
	"putools/log"
)

type Server struct {
	Port          string
	ClientVersion int
	TimeoutCount  int
}

func NewServer() *Server {
	server := Server{}

	port, err := g_config.GetString("default", "port")
	if err != nil || len(port) <= 0 {
		port = "1337"
	}
	server.Port = port

	version, err := g_config.GetInt("default", "clientversion")
	if err != nil {
		version = 0
	}
	server.ClientVersion = version
	server.TimeoutCount = 0

	return &server
}

func (s *Server) Start() {
	// Start timeout loop here
	logger.Println("[Message] Idle player checker goroutine started")
	go s.TimeoutLoop()
	go g_game.CheckCreatures()

	// Open new socket listener
	logger.Println("Opening websocket server on :" + s.Port + "/puserver")

	logger.Println("Server ready to accept new connections")
	http.Handle("/puserver", websocket.Handler(ClientConnection));
	err := http.ListenAndServe(":" + s.Port, nil);
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func ClientConnection(clientsock *websocket.Conn) {
	packet := pnet.NewPacket()
	buffer := make([]uint8, pnet.PACKET_MAXSIZE)
	recv, err := clientsock.Read(buffer)
	if err == nil {
		copy(packet.Buffer[0:recv], buffer[0:recv])
		packet.GetHeader()
		g_server.ParseFirstMessage(clientsock, packet)
	} else {
		if(err.Error() != "EOF") {
			logger.Println("Client connection error: " + err.Error())
		}
	}
}

// This function checks players without a connection every 5 seconds
// and idle players every 10 seconds
func (s *Server) TimeoutLoop() {
	// Check connectionless players
//	g_game.mutexDisconnectList.Lock()
//	defer g_game.mutexDisconnectList.Unlock()
//	for guid, value := range g_game.PlayersDiscon {
//		value.TimeoutCounter++
//		if value.TimeoutCounter >= 30 {
//			delete(g_game.PlayersDiscon, guid)
//			go g_game.RemoveCreature(guid)
//		}
//	}

	if s.TimeoutCount == 10 {
		s.TimeoutCount = 0

		// Check idle players
		g_game.mutexPlayerList.RLock()
		defer g_game.mutexPlayerList.RUnlock()
		for _, player := range g_game.Players {
			if player.Conn != nil {
				// TODO: send ping
				if player.GetTimeSinceLastMove() > 9e5 { // (900000sec / 15 min)
					go g_game.OnPlayerLoseConnection(player)
				}
			}
		}
	}
	
	s.TimeoutCount++
	
	// Run again after 1 second
	go func () {
		time.Sleep(1e9)
		g_server.TimeoutLoop()
	}()
}

// Check login credentials before creating a player object
func (s *Server) ParseFirstMessage(conn *websocket.Conn, packet *pnet.Packet) {
	// Read packet header
	header := packet.ReadUint8()
	if header != pnet.HEADER_LOGIN {
		return
	}
	
	// Make new Connection object to handle net.Conn
	connection := NewConnection(conn)
	// Parse packet
	// We can use the same packet for sending the return status
	firstMessage := &pnetmsg.LoginMessage{}
	firstMessage.ReadPacket(packet)

	if g_game.State == GAME_STATE_CLOSING || g_game.State == GAME_STATE_CLOSED {
		firstMessage.Status = pnetmsg.LOGINSTATUS_SERVERCLOSED
	} else if firstMessage.ClientVersion < s.ClientVersion {
		firstMessage.Status = pnetmsg.LOGINSTATUS_WRONGVERSION
	} else {
		// Load account info
		ret, playerId := s.CheckAccountInfo(firstMessage.Username, firstMessage.Password)
		
		if !ret {
			firstMessage.Status = pnetmsg.LOGINSTATUS_WRONGACCOUNT
			logger.Printf("[LOGIN] %s - Wrong Account - %s", firstMessage.Username, firstMessage.Password)
		} else {
			// Account exists and password is correct
			logger.Println("LoadPlayerProfile: 0")
			ret, player := s.LoadPlayerProfile(playerId)

			if !ret || player == nil {
				firstMessage.Status = pnetmsg.LOGINSTATUS_FAILPROFILELOAD
				logger.Printf("[LOGIN] %s - Failed to load profile", firstMessage.Username)
			} else if player.Conn != nil {
				firstMessage.Status = pnetmsg.LOGINSTATUS_ALREADYLOGGEDIN
				logger.Println("[LOGIN] %s - Already logged in", firstMessage.Username)
			} else {
				firstMessage.Status = pnetmsg.LOGINSTATUS_READY
				logger.Printf("[LOGIN] %d - %v logged in", player.GetUID(), player.GetName())
				
				// Assign Connection to Player object
				player.SetConnection(connection)

				// AddCreature sends few messages to the player so,
				// quickly sending the status message before adding the player.
				connection.SendMessage(firstMessage)
				g_game.AddCreature(player)
				
				player.Conn.HandleConnection()
				return
			}
		}
	}

	connection.SendMessage(firstMessage)
}

func (s *Server) CheckAccountInfo(_username string, _password string) (bool, int64) {
	_username = puh.Escape(_username)
	_password = puh.Escape(_password)

	var queryString string = "SELECT idplayer, password, password_salt FROM player WHERE name='" + _username + "'"
	result, err := puh.DBQuerySelect(queryString);
	if err != nil {
		return false, 0
	}

	row := result.FetchMap()
	defer result.Free()
	if row == nil {
		return false, 0
	}

	idplayer := puh.DBGetInt64(row["idplayer"])
	password := puh.DBGetString(row["password"])
	salt := puh.DBGetString(row["password_salt"])
	_password = _password + salt

	passCheck := s.PasswordTest(_password, password)
	return passCheck, idplayer
}

func (s *Server) PasswordTest(_plain string, _hash string) bool {
	var h hash.Hash = sha1.New()
	h.Write([]byte(_plain))

	var sha1Hash string = strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))
	var original string = strings.ToUpper(_hash)

	return (sha1Hash == original)
}

func (s *Server) LoadPlayerProfile(_playerId int64) (ret bool, p *Player) {
	p = nil
	ret = false

	var queryString string = fmt.Sprintf("SELECT idplayer, name FROM player WHERE idplayer=%d", _playerId)
	result, err := puh.DBQuerySelect(queryString);
	if err != nil {
		return
	}

	row := result.FetchMap()
	if row == nil {
		result.Free()
		return
	}
	idPlayer := puh.DBGetInt(row["idplayer"])
	name := puh.DBGetString(row["name"])
	result.Free()

	value, found := g_game.GetPlayerByName(name)
	if found {
		p = value
		ret = true
	} else {
		p = NewPlayer(name)
		p.dbid = idPlayer
		ret = p.LoadData()
	}

	return
}
