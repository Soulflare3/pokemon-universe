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
	"log"
	"flag"
	"runtime"
	"os"
	"time"

	"conf"
	"mysql"

	"logger"       // PU.Logger package
	pos "position" // PU.Position package	
)

const (
	IS_DEBUG = true
	PO_DEBUG = true
)

var (
	configFile *string

	g_config 	*conf.ConfigFile
	g_logger 	*log.Logger
	g_db     	*mysql.Client

	g_game				*Game
	g_server			*Server
	g_map    			*Map
	g_PokemonManager 	*PokemonManager

	// Client viewport variables. The Z position doesn't matter in this case
	CLIENT_VIEWPORT        pos.Position = pos.Position{28, 22, 0}
	CLIENT_VIEWPORT_CENTER pos.Position = pos.Position{14, 11, 0}
)

func initConfig() bool {
	c, err := conf.ReadConfigFile("data/" + *configFile)
	if err != nil {
		fmt.Printf("Could not load config file: %v\n\r", err)
		return false
	}

	g_config = c

	return true
}

func initLogger() bool {
	var flags int

	toConsole, err := g_config.GetBool("log", "console")
	if err != nil || toConsole {
		flags = logger.L_CONSOLE
	}
	toFile, err := g_config.GetBool("log", "file")
	if err != nil || toFile {
		flags = flags | logger.L_FILE
	}

	logFile, err := g_config.GetString("log", "filename")
	if err != nil || len(logFile) <= 0 {
		logFile = "log.txt"
	}
	myLog, err := logger.NewLogger(logFile, flags)
	if err != nil || myLog == nil {
		fmt.Printf("[Error] Could not initialize logger: %v\n\r", err)
		return false
	}
	g_logger = log.New(myLog, "", log.Ltime)
	if toFile {
		fmt.Printf(" - Start logging to file: %v\n\r", logFile)
	}

	return true
}

func initDatabase() bool {
	// Fetch database info from conf file
	SQLHost, _ := g_config.GetString("database", "host")
	SQLUser, _ := g_config.GetString("database", "user")
	SQLPass, _ := g_config.GetString("database", "pass")
	SQLDB, _ := g_config.GetString("database", "db")

	// Enable/Disable intern mysql logging system
	// g_db.Logging, _ = g_config.GetBool("database", "show_log")

	// Connect to database
	var err os.Error
	g_db, err = mysql.DialTCP(SQLHost, SQLUser, SQLPass, SQLDB)
	if err != nil {
		g_logger.Printf("[Error] Could not connect to database: %v\n\r", err)
		return false
	}

	g_db.Reconnect = true

	return true
}

func main() {
	// Use all cpu cores
	runtime.GOMAXPROCS(2)

	// Flags
	configFile = flag.String("config", "server.conf", "Name of the config file to load")
	flag.Parse()

	fmt.Println("***********************************************")
	fmt.Println("**          Pokemon Universe Server          **")
	fmt.Println("**                                           **")
	fmt.Println("** http://code.google.com/p/pokemon-universe **")
	fmt.Println("**       GNU General Public License V2       **")
	fmt.Println("***********************************************")

	// Load config file
	fmt.Println(" - Loading config file")
	if initConfig() == false {
		return
	}

	// Setup logger
	fmt.Println(" - Setting up logging system")
	if initLogger() == false {
		return
	}

	// Connect to database
	g_logger.Println("Connecting to databeast")
	if initDatabase() == false {
		return
	}
	
	g_logger.Println("Loading pokemon data")
	g_PokemonManager = NewPokemonManager()
	if !g_PokemonManager.Load() {
		g_logger.Println("Failed to load pokemon data...")
		return
	}

	if !PO_DEBUG {
	// Load data
	g_logger.Println("Loading game data...")
	g_game = NewGame()
	if !g_game.Load() {
		g_logger.Println("Failed to load game data...")
		return
	}

	// Start server
	g_game.State = GAME_STATE_NORMAL
	g_logger.Println("--- SERVER STARTING ---")
	g_server = NewServer()
	g_server.Start()
	} else {
		POTestClientDoIt()
		
		for {
			time.Sleep(1e9)
		}
	}
}
