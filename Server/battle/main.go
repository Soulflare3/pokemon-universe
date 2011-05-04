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
	"time"
)

const (
	BATTLESERVER_IP = "127.0.0.1:5080"
)

func createTempData() {

}

func main() {
	// Use all cpu cores
	runtime.GOMAXPROCS(2)

	fmt.Println("***********************************************")
	fmt.Println("**      Pokemon Universe Battle Client       **")
	fmt.Println("**                                           **")
	fmt.Println("** http://code.google.com/p/pokemon-universe **")
	fmt.Println("**      GNU General Public License V2.1      **")
	fmt.Println("***********************************************")
	
	client, err := NewPOClient()
	if err != nil {
		println("Error creating client: " + err.String())
		return
	}
	println("Created client and connected to battle server")
	client.SendLoginInfo()
	
	for ; ; {
		time.Sleep(5e9)
	}
}