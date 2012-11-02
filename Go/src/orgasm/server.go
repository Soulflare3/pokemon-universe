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
	tileChangeChan chan *Packet
}

func NewServer(_port int) *Server {
	return &Server{port: _port,
		clients:        make(map[int]*Client),
		tileChangeChan: make(chan *Packet)}
}

func (s *Server) RunServer() {
	sock, err := net.Listen("tcp", ":" + fmt.Sprintf("%d", s.port))
	if err != nil {
		fmt.Printf("Server error: %v", err)
		os.Exit(1)
	}
	
	go s.HandleTileChange()
	fmt.Printf("[Succeeded]\nWaiting for clients\n")

	for {
		clientsock, err := sock.Accept()
		if err != nil {
			fmt.Printf("Server error: %v", err)
			break
		}
		
		client := NewClient(clientsock, s.tileChangeChan)
		fmt.Printf("Client connected: %d\n", client.id)
		
		s.clients[client.id] = client

		go client.HandleClient()
	}
	sock.Close()
}

func (s *Server) HandleTileChange() {
	for {
		packet := <-s.tileChangeChan
		if packet == nil {
			break
		}

		numTiles := int(packet.ReadUint16())
		if numTiles <= 0 { // Zero tile selected bug
			return
		}

		updatedTiles := list.New()

		for i := 0; i < numTiles; i++ {
			x := int(packet.ReadInt16())
			y := int(packet.ReadInt16())
			z := int(packet.ReadUint16())
			blocking := int(packet.ReadUint16())
			numLayers := int(packet.ReadUint16())

			// Check if tile already exists
			tile, exists := g_map.GetTileFromCoordinates(x, y, z)

			if IS_DEBUG {
				fmt.Printf("Tile Exists - %v - Layers: %d\n", exists, numLayers)
			}

			if numLayers > 0 {
				if !exists { // Tile does not exists, create it		
					if IS_DEBUG {
						fmt.Printf("New Tile - X: %d - Y: %d - Z: %d\n", x, y, z)
					}

					tile = NewTileExt(x, y, z)
				} else if IS_DEBUG {
					fmt.Printf("Update Tile - X: %d - Y: %d - Z: %d - DbId: %d\n", x, y, z, tile.DbId)
				}

				// Set/update blocking
				tile.SetBlocking(blocking)

				for j := 0; j < numLayers; j++ {
					layerId := int(packet.ReadUint16())
					sprite := int(packet.ReadUint16())

					tileLayer := tile.GetLayer(layerId)
					if tileLayer == nil {
						if IS_DEBUG {
							fmt.Printf("Add Layer - Tile Id: %d - Layer: %d - DbId: %d\n", tile.DbId, layerId, tileLayer.DbId)
						}

						// Add and save new tile layer
						tileLayer = tile.AddLayer(layerId, sprite)
					} else {
						if sprite == 0 {
							if IS_DEBUG {
								fmt.Printf("Delete Layer - Tile Id: %d - DbId: %d\n", tile.DbId, tileLayer.DbId)
							}

							// Remove layer, this will also remove the layer from database
							tile.RemoveLayer(tileLayer)
						} else {
							if IS_DEBUG {
								fmt.Printf("Update Layer - Tile Id: %d - DbId: %d\n", tile.DbId, tileLayer.DbId)
							}

							// Update tile layer with new sprite id
							tileLayer.SetSpriteId(sprite)
						}
					}
				}

				// Save tile to database
				tile.Save()
			} else {
				if exists {
					if IS_DEBUG {
						fmt.Printf("Delete Tile - Tile Id: %d\n", tile.DbId)
					}

					// Remove tile from database
					tile.Delete()
				}
			}

			updatedTiles.PushBack(tile)
		}

		s.SendTileUpdateToClients(updatedTiles, 0)
	}
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

func (s *Server) SendNpcToClients(_id int64) {
	for _, client := range(s.clients) {
		client.SendNpc(_id)
	}
}

func (s *Server) SendDeleteNpcToClients(_id int64) {
	for _, client := range(s.clients) {
		client.SendDeleteNpc(_id)
	}
}