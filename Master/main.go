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
	"runtime"
	"fmt"
	"flag"
	"log"
	
	"mysql"
	"conf"
	"logger" // PU.Logger package	
)

var (
	configFile = flag.String("config", "server.conf", "Name of the config file to load")
	
	g_config *conf.ConfigFile
	g_logger *log.Logger
	g_db     *mysql.MySQL
	g_store	 *ServerStore
)

func main() {
	// Use all cpu cores
	runtime.GOMAXPROCS(4)
	
	fmt.Println("***********************************************")
	fmt.Println("**          Pokemon Universe Master          **")
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
	
	server := NewServer()
	server.Run()
}

func initConfig() bool {
	c, err := conf.ReadConfigFile(*configFile)
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
		flags = flags|logger.L_FILE
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
	// Create new instance
	g_db = mysql.New()

	// Enable logging
	g_db.Logging = true
	
	// Fetch database info from conf file
	dbHost, _ := g_config.GetString("database", "host")
	dbUser, _ := g_config.GetString("database", "user")
	dbPass, _ := g_config.GetString("database", "pass")
	dbData, _ := g_config.GetString("database", "db")

	// Enable/Disable intern mysql logging system
	g_db.Logging, _ = g_config.GetBool("database", "show_log")

	// Connect to database
	err := g_db.Connect(dbHost, dbUser, dbPass, dbData)
	if err != nil {
		g_logger.Printf("[Error] Could not connect to database: %v\n\r", err)
		return false
	}

	return true
}
