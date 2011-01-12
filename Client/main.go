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
	font := NewFont("MyriadPro-Regular.ttf",20)

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
