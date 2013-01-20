package main

import (
	"bytes"
//	"database/sql"
	"fmt"
	"net"
	"os"
	"container/list"
	"sync"
	
//	"github.com/astaxie/beedb"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
	
//	"nonamelib/log"
)

type Server struct {
	port int
	clients map[int]*Client
	tileChangeChan chan *Packet
	tileLock sync.Mutex
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
	// Fetch database info from conf file
	username, _ := g_config.GetString("database", "user")
	password, _ := g_config.GetString("database", "pass")
	scheme, _ := g_config.GetString("database", "db")

//	var tileOrm beedb.Model
//	
//	// Connect
//	db, err := sql.Open("mymysql", fmt.Sprintf("%v/%v/%v", scheme, username, password))
//	if err != nil {
//		log.Error("main", "setupDatabase", "Error when opening sql connection: %v", err.Error())
//		return
//	} else {
//		tileOrm = beedb.New(db)
//	}

	db := mysql.New("tcp", "", "127.0.0.1:3306", username, password, scheme)
	err := db.Connect()
	if err != nil {
		panic(err)
	}

	for {		
		packet := <-s.tileChangeChan
		
		if packet == nil {
			break
		}

		numTiles := int(packet.ReadUint16())
		if numTiles <= 0 { // Zero tile selected bug
			return
		}

		s.tileLock.Lock()
		//defer s.tileLock.Unlock()

		var query bytes.Buffer
		query.WriteString("SET foreign_key_checks = 0;\n")
		
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
					tile.DbId = g_newTileId
					fmt.Printf("Current TileID: %d\n", g_newTileId)
					g_newTileId++
				} else if IS_DEBUG {
					fmt.Printf("Update Tile - X: %d - Y: %d - Z: %d - DbId: %d\n", x, y, z, tile.DbId)
				}

				// Set/update blocking
				tile.SetBlocking(blocking)
				
				// Save tile to database
				buffer := tile.Save()
				query.Write(buffer.Bytes())

				for j := 0; j < numLayers; j++ {
					layerId := int(packet.ReadUint16())
					sprite := int(packet.ReadUint16())

					tileLayer := tile.GetLayer(layerId)
					if tileLayer == nil {
					
						// Add and save new tile layer
						tileLayer = tile.AddLayer(layerId, sprite)
						
						//Save the tile layer
						buffer := tileLayer.Save()
						query.Write(buffer.Bytes())
						
						if IS_DEBUG {
							fmt.Printf("Add Layer - Tile Id: %d - Layer: %d - DbId: %d\n", tile.DbId, layerId, tileLayer.DbId)
						}
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
						
						//Save the tile layer
						buffer := tileLayer.Save()
						query.Write(buffer.Bytes())
					}
				}
			} else {
				if exists {
					if IS_DEBUG {
						fmt.Printf("Delete Tile - Tile Id: %d\n", tile.DbId)
					}

					// Remove tile from database
					buffer := tile.Delete()
					query.Write(buffer.Bytes())
					g_map.RemoveTile(tile)
				}
			}

			updatedTiles.PushBack(tile)
		}
		
		query.WriteString("SET foreign_key_checks = 1;")
		// Execute
		if IS_DEBUG {
			fmt.Println(query.String())
		}
		//_, err := tileOrm.Exec(query.String())
		res, err := db.Start(query.String())
		if err != nil {
			fmt.Println(err.Error())
		} else {
			for ; res != nil; res, _ = res.NextResult() {
				fmt.Println("Getting result..")
			}
		}
		
		fmt.Println("Done... Waiting for next.")
		
		s.tileLock.Unlock()
		
		//Send the updated tiles to all clients
		//s.SendTileUpdateToClients(updatedTiles, 0)
	}
	
	fmt.Println("Out of process loop")
}

func (s *Server) SendTileUpdateToClients(_tiles *list.List, _sender int) {
//	for e := _tiles.Front(); e != nil; e = e.Next() {
//		tile := e.Value.(*Tile)

		// Send to connected clients, except sender
//		for id, client := range(s.clients) {
//			if id != _sender {
//				// Send to client
//			}
//		}
//	}
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