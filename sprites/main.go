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
 
    delta := float32(0)
    ticker := time.NewTicker(time.Second / 60)
    renderWindow := sf.NewRenderWindow(sf.VideoMode{800, 600, 32}, "Sprites", sf.StyleDefault, sf.DefaultContextSettings())
    renderWindow.SetVSyncEnabled(true)
    texture, _ := sf.NewTextureFromFile("megaman.png", nil)
    sprite, _ := sf.NewSprite(texture)
    sprite.SetPosition(sf.Vector2f{10, 50})
 
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
            if sprite.GetPosition().X == 10 {
                delta = 10
            } 
            if sprite.GetPosition().X == 700 {
                delta = -10
            }
            sprite.Move(sf.Vector2f{delta, 0})
        }
 
        renderWindow.Clear(sf.ColorWhite())
        renderWindow.Draw(sprite, sf.DefaultRenderStates())
        renderWindow.Display()
    }
}