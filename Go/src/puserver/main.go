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
	"flag"
	"fmt"
	"runtime"

	"goconf"
	"gomysql"
	puh "puhelper"
	"pulogic/pokemon"
	"putools/log"
)

var (
	configFile *string = flag.String("config", "server.conf", "Name of the configuration file to load")

	g_config *conf.ConfigFile

	g_game   *Game
	g_server *Server
	g_map    *Map
	g_npc	 *NpcManager
)

func main() {
	flag.Parse()

	// Always use the maximum available CPU/Cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("***********************************************")
	fmt.Println("**          Pokemon Universe Server          **")
	fmt.Println("**                                           **")
	fmt.Println("** http://code.google.com/p/pokemon-universe **")
	fmt.Println("**       GNU General Public License V2       **")
	fmt.Println("***********************************************")

	if !initConfig() {
		return
	}

	if !initLogger() {
		return
	}

	logger.Println("Initial setup complete.")
	logger.Println("Connecting to database.")
	if !initDatabase() {
		return
	}

	logger.Println("Loadig Pokemon data")
	pokemonManager := pokemon.GetInstance()
	if !pokemonManager.Load() {
		return
	}

	logger.Println("Loading game data")
	g_game = NewGame()
	if !g_game.Load() {
		logger.Println("Failed to load game data...")
		return
	}

	logger.Println("Starting server")
	g_server = NewServer()
	g_server.Start()
}

// Initialize configuration file
func initConfig() bool {
	c, err := conf.ReadConfigFile("data/" + *configFile)
	if err != nil {
		fmt.Printf("Could not load configuration file: %v\n", err)

		// TODO: Create default configuration file

		return false
	}

	g_config = c

	return true
}

// Initialize log to file
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

	logger.LogFilename = logFile
	logger.Flags = flags

	result := logger.Init()
	if !result {
		fmt.Println("Failed to initialize logger")
	}

	return result
}

// Initialize and connect to database
func initDatabase() bool {
	// Fetch database info from conf file
	SQLHost, _ := g_config.GetString("database", "host")
	SQLUser, _ := g_config.GetString("database", "user")
	SQLPass, _ := g_config.GetString("database", "pass")
	SQLDB, _ := g_config.GetString("database", "db")

	// Enable/Disable intern mysql logging system
	// g_db.Logging, _ = g_config.GetBool("database", "show_log")

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
