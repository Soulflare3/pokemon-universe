package main

import (
	"database/sql"
	"fmt"
	"flag"
	"os" 
	
	"github.com/astaxie/beedb"
	_ "github.com/ziutek/mymysql/godrv"
	
	"nonamelib/config"
	"nonamelib/log"
	
	"pulogic/pokemon"
)

const (
	IS_DEBUG = false
)

var (
	g_map *Map = NewMap()
	g_npc *NpcList = NewNpcList()
	g_locations *LocationsList = NewLocationsList()
	g_orm beedb.Model
	g_server *Server
	g_config *config.ConfigFile
	g_newTileId int64
	g_newTileLayerId int64
	version string
)

func initConfig(configFile *string) bool {
	c, err := config.ReadConfigFile("data/" + *configFile)
	if err != nil {
		fmt.Printf("Could not load config file: %v\n\r", err)
		return false
	}

	g_config = c

	return true
}

func initLogger() {
	var flags int
	var err error

	toConsole, err := g_config.GetBool("log", "console")
	if err != nil || toConsole {
		flags = log.L_CONSOLE
	}
	toFile, err := g_config.GetBool("log", "file")
	if err != nil || toFile {
		flags = flags | log.L_FILE
	}
	showDebug, err := g_config.GetBool("log", "debug")
	if err != nil || showDebug {
		flags = flags | log.F_DEBUG
	}

	logFile, err := g_config.GetString("log", "filename")
	if err != nil || len(logFile) <= 0 {
		logFile = "log.txt"
	}
	logFile = "logs/" + logFile
	os.MkdirAll("logs", os.ModePerm)

	log.LogFilename = logFile
	log.Flags = flags
}

func initDatabase() bool {
	// Fetch database info from conf file
	username, _ := g_config.GetString("database", "user")
	password, _ := g_config.GetString("database", "pass")
	scheme, _ := g_config.GetString("database", "db")

	// Connect
	db, err := sql.Open("mymysql", fmt.Sprintf("%v/%v/%v", scheme, username, password))
	if err != nil {
		log.Error("main", "setupDatabase", "Error when opening sql connection: %v", err.Error())
	} else {
		g_orm = beedb.New(db)
		beedb.OnDebug, _ = g_config.GetBool("database", "debug_mode")
	}

	return (err == nil)
}

func main() {
	fmt.Println("*******************************************")
	fmt.Println("** Pokemon Universe - Mapserver v0.5.r1  **")
	fmt.Println("*******************************************")
	
	// Flags
	configFile := flag.String("config", "server.conf", "Name of the config file to load")
	flag.Parse()
	
	// Load config file
	fmt.Printf("Loading config file...")
	if initConfig(configFile) == false {
		fmt.Printf("[FAILED]\n")
		return
	}
	fmt.Printf("[Succeeded]\n")
	
	initLogger()

	// Connect to database 
	fmt.Printf("Connecting to database...")
	if initDatabase() == false {
		fmt.Printf("[FAILED]\n")
		return
	}
	fmt.Printf("[Succeeded]\n")
	
	fmt.Printf("Loading all Pokemon data...")
	pokemon.G_orm = &g_orm
	pokemonManager := pokemon.GetInstance()
	if !pokemonManager.Load() {
		return
	}
	fmt.Printf("[Succeeded]\n")

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
	
	fmt.Printf("Retrieving Locations...")
	if !g_locations.LoadLocations(){
		return
	}
	fmt.Printf("[Succeeded] (Loaded %d pokecenters, %d music and %d locations)\n", g_locations.GetNumPokecenters(), g_locations.GetNumMusic(), g_locations.GetNumLocations())
	
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
