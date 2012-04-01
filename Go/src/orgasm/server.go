package main

import (
	"fmt"
	"net"
	"os"
	"container/list"
)

type Server struct {
	port int
	
	clients map[int]*Client
}

func NewServer(_port int) *Server {
	return &Server{port: _port, clients: make(map[int]*Client) }
}

func (s *Server) RunServer() {
	sock, err := net.Listen("tcp", ":" + fmt.Sprintf("%d", s.port))
	if err != nil {
		fmt.Printf("Server error: %v", err)
		os.Exit(1)
	}
	fmt.Printf("[Succeeded]\nWaiting for clients\n")

	for {
		clientsock, err := sock.Accept()
		if err != nil {
			fmt.Printf("Server error: %v", err)
			break
		}
		
		client := NewClient(clientsock)
		fmt.Printf("Client connected: %d\n", client.id)
		
		s.clients[client.id] = client

		go client.HandleClient()
	}
	sock.Close()
}

func (s *Server) SendTileUpdateToClients(_tiles *list.List, _sender int) {
	for e := _tiles.Front(); e != nil; e = e.Next() {
		tile := e.Value.(*Tile)
		
		// Send to connected clients, except sender
//		for id, client := range(s.clients) {
//			if id != _sender {
//				// Send to client
//			}
//		}
		
		// If tile is set to remove, do it now
		if tile.IsRemoved {
			g_map.RemoveTile(tile)
		}
	}
}

func (s *Server) SendMapListUpdateToClients() {
	for _, client := range(s.clients) {
		client.SendMapList()
	}
}

func (s *Server) SendNpcToClients(_id int) {
	for _, client := range(s.clients) {
		client.SendNpc(_id)
	}
}

func (s *Server) SendDeleteNpcToClients(_id int) {
	for _, client := range(s.clients) {
		client.SendDeleteNpc(_id)
	}
}