package main
 
import (
    "fmt"
    "os"
    "sdl"
    "time"
)
 
func main() {
    //Initialize SDL
    err := sdl.Init()
    if err != "" { 
        fmt.Printf("Error in Init: %v", err)
        return
    }
 
	//Initialize the engine
 	InitEngine()

	//Some test code
    s := sdl.LoadImage("test.png")
    texture := s.CreateTexture() 
 
    //Handle events
    for {  
        event, present := sdl.PollEvent()
		if present {
			EventHandler(event)		
		}  
        sdl.RenderClear()
       
		//Some more test code
        texture.RenderCopy(sdl.Rect{0,0,int32(s.W),int32(s.H)}, sdl.Rect{0,0,int32(s.W),int32(s.H)})
   
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
