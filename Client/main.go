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
	"os"
	"exec"
	"path"
	"runtime"
)

const (
	CLIENT_VERSION = 4
)

var g_running bool = true
var g_engine *PU_Engine = NewEngine()
var g_game *PU_Game = NewGame()
var g_gui *PU_Gui = NewGui()
var g_conn *PU_Connection = NewConnection()

func main() {
	//Make sure that resources get released
	defer g_engine.Exit()
	
	//Permission to run on 2 CPU cores
	runtime.GOMAXPROCS(2)

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
	
	//Gui test code
	g_loginControls.Show()

	//Handle events 
	for g_running {
		event, present := sdl.PollEvent()
		if present {
			EventHandler(event)
		}
		
		//Render everything on screen
		Draw()
		
		//Give the CPU some time to do other stuff
		sdl.Delay(10)
		
		//Handle a network packet
		g_conn.HandlePacket()
	}
	sdl.Quit()
}

func Draw() {
	sdl.RenderClear()
	g_game.Draw()
	sdl.RenderPresent()
}

func GetPath() string {
	file, _ := exec.LookPath(os.Args[0])
	dir, _ := path.Split(file)
	os.Chdir(dir)
	path, _ := os.Getwd()
	return path+"/"
}

func EventHandler(_event *sdl.SDLEvent) {
	switch _event.Evtype {
	case sdl.SDL_WINDOWEVENT:
		HandleWindowEvent(_event.Window())
		
	case sdl.SDL_KEYDOWN, sdl.SDL_TEXTINPUT:
		HandleKeyboardEvent(_event.Keyboard())
		
	case sdl.SDL_MOUSEBUTTONUP, sdl.SDL_MOUSEBUTTONDOWN:
		HandleMouseButtonEvent(_event.MouseButton())
		
	case sdl.SDL_MOUSEMOTION:
		HandleMouseMotionEvent(_event.MouseMotion())
	}
}

func HandleWindowEvent(_event *sdl.WindowEvent) {
	switch _event.Event {
	case sdl.SDL_WINDOWEVENT_CLOSE:
		g_running = false
	}
}

func HandleKeyboardEvent(_event *sdl.KeyboardEvent) {
	switch _event.Evtype {
		case sdl.SDL_KEYDOWN:
			g_gui.KeyDown(0, int(_event.Keysym().Sym))
			
		case sdl.SDL_TEXTINPUT:
			g_gui.KeyDown(int(_event.State), int(_event.Keysym().Scancode));
	}
}

func HandleMouseButtonEvent(_event *sdl.MouseButtonEvent) {
	switch _event.Evtype {
		case sdl.SDL_MOUSEBUTTONUP:
			g_gui.MouseUp(int(_event.X), int(_event.Y))
			
		case sdl.SDL_MOUSEBUTTONDOWN:
			g_gui.MouseDown(int(_event.X), int(_event.Y))
	}
}

func HandleMouseMotionEvent(_event *sdl.MouseMotionEvent) {
	g_gui.MouseMove(int(_event.X), int(_event.Y))
}
