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
	"os"
	"net"
	"hash"
	"crypto/sha1"
	"strings"

	"mysql"
	pnet "network"
)

type Server struct {
	Port  string
	Store *ServerStore
}

func NewServer() *Server {
	s := &Server{}
	port, err := g_config.GetString("default", "port")
	if err != nil || len(port) <= 0 {
		port = "666"
	}
	s.Port = port

	s.Store = NewServerStore()

	return s
}

func (s *Server) Run() {
	// Open new socket listener
	g_logger.Println("Opening server socket on port " + s.Port)
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
		go s.ParseMessage(clientsock, packet)
	}
}

func (s *Server) ParseMessage(_socket net.Conn, _packet *pnet.Packet) {
	// Read packet header
	header := _packet.ReadUint8()
	if header != pnet.HEADER_LOGIN {
		return
	}

	firstMessage := NewLoginMessage()
	firstMessage.ReadPacket(_packet)

	ret, iduser := s.checkCredentials(firstMessage.Username, firstMessage.Password)
	if !ret {
		firstMessage.Status = LOGINSTATUS_WRONGACCOUNT
	} else {
		s.loadCharacters(firstMessage, iduser)
	}

	// Send packet back to user
	packet, _ := firstMessage.WritePacket()
	packet.SetHeader()
	_socket.Write(packet.Buffer[0:packet.MsgSize])
}

func (s *Server) checkCredentials(_username, _password string) (ret bool, iduser string) {
	iduser = ""
	_username = g_db.Escape(_username)
	_password = g_db.Escape(_password)

	var err os.Error
	var res *mysql.MySQLResult
	var rows map[string]interface{}
	var queryString string = "SELECT iduser, password, password_salt FROM users WHERE username='" + _username + "'"
	if res, err = g_db.Query(queryString); err != nil {
		ret = false
		return
	}

	if rows = res.FetchMap(); rows == nil {
		ret = false
		return
	}

	iduser = fmt.Sprintf("%d", rows["iduser"].(int))
	password, _ := rows["password"].(string)
	salt, _ := rows["password_salt"].(string)
	_password = _password + salt

	ret = PasswordTest(_password, password)
	return
}

func (s *Server) loadCharacters(_message *LoginMessage, _iduser string) {
	var err os.Error
	var res *mysql.MySQLResult
	var rows map[string]interface{}
	var queryString string = "SELECT idserver, name FROM characters WHERE idaccount='" + _iduser + "'"
	if res, err = g_db.Query(queryString); err != nil {
		_message.Status = LOGINSTATUS_FAILPROFILELOAD
		return
	}

	_message.Status = LOGINSTATUS_READY

	for {
		if rows = res.FetchMap(); rows == nil {
			break
		}

		charname := rows["name"].(string)
		idserver := rows["idserver"].(int)

		// Fetch server
		if serverinfo, ok := s.Store.servers.Get(idserver); ok {
			_message.Servers[charname] = serverinfo
		}
	}
}

func PasswordTest(_plain string, _hash string) bool {
	var h hash.Hash = sha1.New()
	h.Write([]byte(_plain))

	var sha1Hash string = fmt.Sprintf("%X", h.Sum())
	var original string = strings.ToUpper(_hash)

	return (sha1Hash == original)
}
