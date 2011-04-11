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
	"net"
	"time"
	"fmt" 
	pnet "network" // PU Network package
)

type Server struct {
	Port 			string
	ClientVersion 	int
	TimeoutCount	int
}

func NewServer() *Server {
	server := Server { }
	
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
	g_logger.Println("[Message] Idle player checker goroutine started")
	go s.timeoutLoop()

	// Open new socket listener
	g_logger.Println("Opening server socket on port "+s.Port)
	socket, err := net.Listen("tcp", ":"+s.Port)
	if err != nil {
		g_logger.Printf("[Error] Could not open socket - %v\n", err)
		return
	}
	defer socket.Close() // Defer the close function so that's get done automatically when this method breaks
	defer g_logger.Println("[Notice] Server socket closed")
	
	g_logger.Println("Server ready to accept new connections")
	for {
		clientsock, err := socket.Accept()
		if err != nil {
			g_logger.Println("[Warning] Could not accept new connection")
			continue
		}
		
		// Read and execute the first received packet
		s.parseFirstMessage(clientsock)
	}
}

// This function checks players without a connection every second
// and idle players every 10 seconds
func (s *Server) timeoutLoop() {
	s.TimeoutCount++
	
	// Check connectionless players
	g_game.mutexDisconnectList.Lock()
	defer g_game.mutexDisconnectList.Unlock()
	for guid, value := range(g_game.PlayersDiscon) {
		value.TimeoutCounter++
		if value.TimeoutCounter >= 30 {
			g_game.PlayersDiscon[guid] = nil, false
			go g_game.RemoveCreature(guid)
		}
	}
	
	if s.TimeoutCount == 10 {
		s.TimeoutCount = 0
		
		// Check idle players
		g_game.mutexPlayerList.RLock()
		defer g_game.mutexPlayerList.RUnlock()
		for _, player := range(g_game.Players) {
			if player.Conn != nil {
				player.Conn.Send_Ping()
				if player.GetTimeSinceLastMove() > 9e5 { // (900000sec / 15 min)
					go g_game.OnPlayerLoseConnection(player)
				}
			}
		}
	}
	
	time.Sleep(1e9)
	go s.timeoutLoop()
}

func (s *Server) parseFirstMessage(conn net.Conn) {
	// Make new Connection object to hold net.Conn
	connection := NewConnection(conn)
	var message *pnet.Message
	var err string
	if message, err = connection.Tranceiver.Receive(); err != "" {
		fmt.Printf("Error receiving first message: %s\n", err)
		return
	}
	
	// Read message header
	if message.Header != pnet.HEADER_LOGIN {
		return
	}
	
	//Get the login data
	loginMessage := message.Login
	
	//Create the login status message
	statusMessage := pnet.NewData_LoginStatus()
		
	if g_game.State == GAME_STATE_CLOSING || g_game.State == GAME_STATE_CLOSED {
		statusMessage.LoginStatus.Status = LOGINSTATUS_SERVERCLOSED
	} else if loginMessage.Version < s.ClientVersion {
		statusMessage.LoginStatus.Status = LOGINSTATUS_WRONGVERSION
	} else {
		// Load account info
		ret := CheckAccountInfo(loginMessage.Username, loginMessage.Password)
		if !ret {
			statusMessage.LoginStatus.Status = LOGINSTATUS_WRONGACCOUNT
		} else { 
			// Account exists and password is correct
			ret, player := LoadPlayerProfile(loginMessage.Username)
			
			if !ret || player == nil {
				statusMessage.LoginStatus.Status = LOGINSTATUS_FAILPROFILELOAD
				g_logger.Printf("[Login] Failed to load profile for %v", loginMessage.Username)
			} else if player.Conn != nil {
				statusMessage.LoginStatus.Status = LOGINSTATUS_ALREADYLOGGEDIN
				fmt.Println("Already logged in")
			} else {
				statusMessage.LoginStatus.Status = LOGINSTATUS_READY
				g_logger.Printf("[Login] %d - %v logged in", player.GetUID(), player.GetName())
				// Assign Connection to Player object
				player.SetConnection(connection)
			}
		}
	}
	
	connection.SendMessage(statusMessage)
}
