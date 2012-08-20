package main

import (
	"fmt"
	"sync"
	"flag"
	
	"gomysql"
	"goconf"
	puh "puhelper"
	"putools/log"
)

const (
	IS_DEBUG = false
)

var (
	g_map *Map = NewMap()
	g_npc *NpcList = NewNpcList()
	g_dblock sync.Mutex
	g_server *Server
	g_config *conf.ConfigFile
	version string
)

func initConfig(configFile *string) bool {
	c, err := conf.ReadConfigFile("data/" + *configFile)
	if err != nil {
		fmt.Printf("Could not load config file: %v\n\r", err)
		return false
	}

	g_config = c

	return true
}

func initDatabase() bool {
	// Fetch database info from conf file
	SQLHost, _ := g_config.GetString("database", "host")
	SQLUser, _ := g_config.GetString("database", "user")
	SQLPass, _ := g_config.GetString("database", "pass")
	SQLDB, _ := g_config.GetString("database", "db")

	// Connect to database
	var err error
	puh.DBCon, err = mysql.DialTCP(SQLHost, SQLUser, SQLPass, SQLDB)
	if err != nil {
		logger.Printf("[Error] Could not connect to database: %v\n\r", err)
		return false
	} else {
		logger.Println("Connected to SQL server:")
		logger.Printf(" - Host: %s\n", SQLHost)
		logger.Printf(" - Database: %s\n", SQLDB)
	}

	puh.DBCon.Reconnect = true

	return true
}

func main() {
	fmt.Println("***************************************")
	fmt.Println("** Pokemon Universe - Mapserver v0.3 **")
	fmt.Println("***************************************")
	
	// Flags
	configFile := flag.String("config", "server.conf", "Name of the config file to load")
	flag.Parse()
	
	// Load config file
	fmt.Printf("Loading config file...")
	if initConfig(configFile) == false {
		return
	}
	fmt.Printf("[Succeeded]\n")

	// Connect to database 
	fmt.Printf("Connecting to database...")
	if initDatabase() == false {
		return
	}
	fmt.Printf("[Succeeded]\n")

	// Load images
	//fmt.Printf("Loading tile images...")
	//LoadImages()
	//fmt.Printf("[Succeeded] (%d images loaded)\n", len(ImagesMap))

	// Get maps 
	fmt.Printf("Retrieving map names...")
	g_map.LoadMapList()
	fmt.Printf("[Succeeded] (%d maps loaded)\n", g_map.GetNumMaps())
	
	// Retrieve all tiles
	fmt.Printf("Retrieving tiles...")
	g_map.LoadTiles()
	fmt.Printf("[Succeeded] (%d tiles loaded)\n", g_map.GetNumTiles())
	
	// Retreive all NPCs
	fmt.Printf("Retrieving NPCs...")
	g_npc.LoadNpcList()
	fmt.Printf("[Succeeded] (%d NPCs loaded)\n", g_npc.GetNumNpcs())
	
//	fmt.Printf("Retrieving NPC Pokemon...")
//	g_npc.LoadNpcPokemon()
//	fmt.Printf("[Succeeded] (%d Pokemons loaded)\n", g_npc.GetNumPokemons())
	
	fmt.Println("Initialisation completed!\n")
	
	// Set up server
	version, _ = g_config.GetString("default", "version")
	fmt.Println("Current server is suited for client version " + version)
	fmt.Printf("Running server...")
	serverPort, ok := g_config.GetInt("default", "port")
	
	if ok == nil || serverPort <= 0 {
		serverPort = 6171
	}
	g_server = NewServer(serverPort)
	g_server.RunServer()
}
