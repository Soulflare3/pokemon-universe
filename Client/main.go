package main
 
import (
    "fmt"
    "os"
    "sdl"
    "time"
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
    for { 
        event, present := sdl.PollEvent()
		if present {
			EventHandler(event)		
		}  
        sdl.RenderClear()
       
        //draw
   
        sdl.RenderPresent() 
        time.Sleep(10)
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
