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
	pnet "network" // PU Network package
)

type Server struct {
	Port 			string
	
	ClientVersion 	uint16
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
	server.ClientVersion = uint16(version)
	
	return &server
}

func (s *Server) Start() {
	// Start timeout loop here
	// go s.timeoutLoop()

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
		
		var headerbuffer [2]uint8
		recv, err := clientsock.Read(headerbuffer[0:])
		if (err != nil) || (recv == 0) {
			g_logger.Printf("[Warning] Could not read packet header: %v", err)
			continue
		}
		// Create new packet
		packet := pnet.NewPacket()
		copy(packet.Buffer[0:2], headerbuffer[0:2]) // Write header buffer to packet
		packet.GetHeader()
		
		databuffer := make([]uint8, packet.MsgSize)
		recv, err = clientsock.Read(databuffer[0:])
		if recv == 0 || err != nil {	
			g_logger.Printf("[Warning] Serer connection read error: %v", err)
			continue
		}
		copy(packet.Buffer[2:], databuffer[:]) // Write rest of the received data to packet
		
		// Read and execute the first received packet
		s.parseFirstMessage(clientsock, packet)
	}
}

// Loop which will check if players are idle for X amount of minutes
func (s *Server) timeoutLoop() {
	g_logger.Println(" - Idle player checker goroutine started")
	for ; ; time.Sleep(10e9) { // 10 Sec
		// Check if there are players who are idle for X min
	}
}

func (s *Server) parseFirstMessage(conn net.Conn, packet *pnet.Packet) {
	// Read packet header
	header := packet.ReadUint8()
	if header != pnet.HEADER_LOGIN {
		return
	}
	
	// Make new Connection object to hold net.Conn
	connection := NewConnection(conn)
	
	// Parse packet
	// We can use the same packet for sending the return status
	firstMessage := &LoginMessage{}
	firstMessage.ReadPacket(packet)
		
	if g_game.State == GAME_STATE_CLOSING || g_game.State == GAME_STATE_CLOSED {
		firstMessage.Status = LOGINSTATUS_SERVERCLOSED
	} else if firstMessage.ClientVersion < s.ClientVersion {
		firstMessage.Status = LOGINSTATUS_WRONGVERSION
	} else {
		// Load account info
		ret := CheckAccountInfo(firstMessage.Username, firstMessage.Password)
		if !ret {
			firstMessage.Status = LOGINSTATUS_WRONGACCOUNT
		} else { 
			// Account exists and password is correct
			ret, player := LoadPlayerProfile(firstMessage.Username)
			
			if !ret || player == nil {
				firstMessage.Status = LOGINSTATUS_FAILPROFILELOAD			
			} else if player.Conn != nil {
				firstMessage.Status = LOGINSTATUS_ALREADYLOGGEDIN
			} else {
				firstMessage.Status = LOGINSTATUS_READY
				
				// Assign Connection to Player object
				player.SetConnection(connection)
			}
		}
	}
	
	connection.SendMessage(firstMessage)
}
