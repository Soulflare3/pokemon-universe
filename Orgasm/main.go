package main

import (
	"fmt"
	"mysql"
	"sync"
	"flag"
	
	"conf"
)

var g_map *Map = NewMap()
var g_db *mysql.Client
var g_dblock sync.Mutex
var g_server *Server
var g_config *conf.ConfigFile

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
	g_db, err = mysql.DialTCP(SQLHost, SQLUser, SQLPass, SQLDB)
	if err != nil {
		fmt.Printf("[Error] Could not connect to database:\n %v\n", err)
		return false
	}

	g_db.Reconnect = true

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
	fmt.Println(" - Loading config file")
	if initConfig(configFile) == false {
		return
	}

	// Connect to database 
	fmt.Printf("Connecting to database...")
	if initDatabase() == false {
		return
	}
	fmt.Printf("[Succeeded]\n")

	// Load images
	fmt.Printf("Loading tile images...")
	LoadImages()
	fmt.Printf("[Succeeded] (%d images loaded)\n", len(ImagesMap))

	// Get maps 
	fmt.Printf("Retrieving map names..")
	g_map.LoadMapList()
	fmt.Printf("[Succeeded] (%d maps loaded)\n", g_map.GetNumMaps())
	
	// Retrieve all tiles
	fmt.Printf("Retrieving tiles...")
	g_map.LoadTiles()
	fmt.Printf("[Succeeded] (%d tiles loaded)\n", g_map.GetNumTiles())

	// Setup http handler
	// http.HandleFunc("/add/", addFormHandler)
	// http.HandleFunc("/adduser/", addHandler)
	// http.HandleFunc("/users/", userHandler)
	// http.HandleFunc("/list/", listHandler)
	// http.HandleFunc("/inc/", SourceHandler)
	// http.HandleFunc("/", handleIndex)
	// http.ListenAndServe(":8080", nil)
	
	// Set up server
	fmt.Printf("Running server...")
	serverPort, ok := g_config.GetInt("default", "port")
	if ok == nil || serverPort <= 0 {
		serverPort = 6171
	}
	g_server = NewServer(serverPort)
	g_server.RunServer()
}
