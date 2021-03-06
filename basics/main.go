package main
 
import (
    sf "bitbucket.org/krepa098/gosfml2"
    "runtime"
    "time"
)
 
func init() {
    runtime.LockOSThread()
}
 
func main() {
 
    ticker := time.NewTicker(time.Second / 30)
 
    renderWindow := sf.NewRenderWindow(sf.VideoMode{800, 600, 32}, "Base Go SFML framework", sf.StyleDefault, sf.DefaultContextSettings())
 
    for renderWindow.IsOpen() {
        select {
        case <-ticker.C:
            //poll events
            for event := renderWindow.PollEvent(); event != nil; event = renderWindow.PollEvent() {
                switch ev := event.(type) {
                case sf.EventKeyPressed:
 
                    //exit on ESC
                    if ev.Code == sf.KeyEscape {
                        renderWindow.Close()
                    }
                case sf.EventClosed:
                    renderWindow.Close()
 
                }
            }
        }
 
        renderWindow.Clear(sf.ColorWhite())
        renderWindow.Display()
    }
}