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

	//Some test code
	img := g_engine.LoadImage("test.png")
	font := g_engine.LoadFont("MyriadPro-Regular.ttf",20)
	font.SetStyle(true,false,false)

	//Handle events 
	for g_running {
		event, present := sdl.PollEvent()
		if present {
			EventHandler(event)
		}
		sdl.RenderClear()

		//Some more test code 
		img.Draw(0, 0)
		
		font.SetColor(255,255,255)
		font.SetAlpha(100)
		font.DrawText("Hello world!", 50,180)

		//Even more test code
		sdl.SetRenderDrawColor(255, 0, 0, 100)
		sdl.SetRenderDrawBlendMode(sdl.SDL_BLENDMODE_BLEND)
		sdl.RenderFillRect(sdl.Rect{10, 10, 100, 100})
		sdl.SetRenderDrawBlendMode(sdl.SDL_BLENDMODE_NONE)
		sdl.SetRenderDrawColor(0, 0, 0, 255)

		sdl.RenderPresent()
		time.Sleep(10)
	}
	sdl.Quit()
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
