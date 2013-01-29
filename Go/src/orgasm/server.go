package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"container/list"
	"sync"
	"time"
	
	"github.com/ziutek/mymysql/mysql"
    _ "github.com/ziutek/mymysql/native"
	
	"nonamelib/log"
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
	
	db := mysql.New("tcp", "", "127.0.0.1:3306", username, password, scheme)
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	
	rows, _, err := db.Query("SELECT MAX(idtile_layer) AS max_id FROM tile_layer LIMIT 50")
    if err != nil {
        panic(err)
    }
    
    for _, row := range rows {
        // You can get converted value
        g_newTileLayerId = row.Int64(0) + 1      // Zero value
    }

	log.Verbose("Server", "HandleTileChange", "Determined next tilelayer ID: %d", g_newTileLayerId)
	
	for {		
		packet := <-s.tileChangeChan
		
		if packet == nil {
			log.Error("Server", "HandleTileChange", "Error! Packet empty!")
			break
		}
		
		s.tileLock.Lock()
		
		start := time.Now().UnixNano()

		//Generate all tiles to update / insert for database
		updatedTiles := s.CreateUpdatedTilesList(packet);
		
		//Prepare batch
		var query bytes.Buffer
		query.WriteString("SET foreign_key_checks = 0;")
		
		for e := updatedTiles.Front(); e != nil; e = e.Next() {
			tile := e.Value.(*Tile)
			
			if(tile.IsNew) {
				//create insert batch for all new tiles
	        	query.WriteString(fmt.Sprintf(QUERY_INSERT_TILE, tile.DbId, tile.Position.X, tile.Position.Y, tile.Position.Z, tile.Blocking, 0))	        	
        		tile.IsNew = false
        	} else {
        		//create update batch for all o tiles
        		query.WriteString(fmt.Sprintf(QUERY_UPDATE_TILE, tile.Blocking, 0, tile.DbId))
        	}
        	
        	for _, tileLayer := range tile.Layers {
        		if(tileLayer.IsNew) {
        			query.WriteString(fmt.Sprintf(QUERY_INSERT_TILELAYER, tileLayer.TileId, tileLayer.Layer, tileLayer.SpriteId))
        			tileLayer.IsNew = false
        		} else {
        			query.WriteString(fmt.Sprintf(QUERY_UPDATE_TILELAYER, tileLayer.SpriteId, tileLayer.DbId))
        		}
        	}
    	}
		
		//Finish batch
		query.WriteString("SET foreign_key_checks = 1;")
		
		startQuery := time.Now().UnixNano()
		
		// Execute
        res, err := db.Start(query.String())
        if err != nil {
            fmt.Println(err.Error())
        } else if res != nil {
			for ; res != nil; {
                res2, err2 := res.NextResult()
                if err2 != nil {
                	fmt.Printf("Error getting result: %s\n", err2.Error())
				} else if res2 == nil {
                	break
                }
                       
            	res = res2
			}
        }
        
        end := time.Now().UnixNano()
        total := float64((end - start)) * 0.000001
        totalQuery := float64((startQuery - start)) * 0.000001
		
		log.Verbose("Server", "HandleTileChange", "Done adding tiles, waiting for next. Total: %dms | Query: %dms | Tiles: %d", int64(total), int64(totalQuery), updatedTiles.Len())
		
		s.tileLock.Unlock()
	}
	
	log.Error("Server", "HandleTileChange", "Error! Out of process!")
}

func (s *Server) CreateUpdatedTilesList(_packet *Packet) *list.List {
	updatedTiles := list.New()

	numTiles := int(_packet.ReadUint16())
	if numTiles <= 0 { // Zero tile selected bug
		log.Error("Server", "CreateUpdatedTilesList", "Error zero tile selected bug")
	}

	for i := 0; i < numTiles; i++ {
		x := int(_packet.ReadInt16())
		y := int(_packet.ReadInt16())
		z := int(_packet.ReadUint16())
		blocking := int(_packet.ReadUint16())
		numLayers := int(_packet.ReadUint16())

		// Check if tile already exists
		tile := g_map.getOrAddTile(x, y, z)
		
		if(tile.IsNew) {
			if IS_DEBUG {
				log.Verbose("Server", "CreateUpdatedTilesList", "Adding new tile with id %d", g_newTileId)
			}
			tile.DbId = g_newTileId
			g_newTileId++
		}

		if numLayers > 0 {

			// Set/update blocking
			tile.SetBlocking(blocking)

			for j := 0; j < numLayers; j++ {
				layerId := int(_packet.ReadUint16())
				sprite := int(_packet.ReadUint16())

				tileLayer := tile.GetLayer(layerId)
				if tileLayer == nil {
				
					// Add and save new tile layer
					tileLayer = tile.AddLayer(layerId, sprite)
					
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
				}
			}
		} else {
			if !tile.IsNew {
				if IS_DEBUG {
					fmt.Printf("Delete Tile - Tile Id: %d\n", tile.DbId)
				}
				//////////////////////////////////
				//TODO Remove tile from database//
				//////////////////////////////////

				g_map.RemoveTile(tile)
			}
		}

		updatedTiles.PushBack(tile)
	}
	
	return updatedTiles;
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