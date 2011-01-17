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
	"sdl"
	"time"
)

var g_running bool = true
var g_engine *PU_Engine = NewEngine()
var g_game *PU_Game = NewGame()
var g_gui *PU_Gui = NewGui()

func main() {
	//Make sure that resources get released
	defer g_engine.Exit()

	//Initialize SDL
	err := sdl.Init()
	if err != "" {
		fmt.Printf("Error in Init: %v", err)
		return
	} 

	//Initialize the engine
	g_engine.Init()
	
	//Load data
	g_game.LoadFonts()
	Draw() //Draw the "please wait" text after loading the fonts
	g_game.LoadGuiImages()
	g_game.LoadTileImages()
	g_game.SetState(GAMESTATE_LOGIN)

	//Handle events 
	for g_running {
		event, present := sdl.PollEvent()
		if present {
			EventHandler(event)
		}
		
		Draw()
		
		time.Sleep(10)
	}
	sdl.Quit()
}

func Draw() {
	sdl.RenderClear()
	g_game.Draw()
	sdl.RenderPresent()
}

func EventHandler(_event *sdl.SDLEvent) {
	switch _event.Evtype {
	case sdl.SDL_WINDOWEVENT:
		HandleWindowEvent(_event.Window())
	}
}

func HandleWindowEvent(_event *sdl.WindowEvent) {
	switch _event.Event {
	case sdl.SDL_WINDOWEVENT_CLOSE:
		g_running = false
	}
}
