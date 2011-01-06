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
    "os"
    "sdl"
)

const (
    WINDOW_WIDTH = 964
    WINDOW_HEIGHT = 720
)

func main() {
    //Initialize SDL
    err := sdl.Init()
    if err != "" {
        fmt.Printf("Error in Init: %v", err)
		return
    }

    //Create the window
    err = sdl.CreateWindow("Pokemon Universe", WINDOW_WIDTH, WINDOW_HEIGHT)
    if err != "" {
        fmt.Printf("Error in CreateWindow: %v", err)
		return
    }

	//Handle events
    events := make(chan *sdl.SDLEvent)
    go sdl.StartEventLoop(events)
    for {
        event := <- events
		EventHandler(event)
    }
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
			os.Exit(0)
	}
}