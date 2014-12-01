package main

import (
    sf "bitbucket.org/krepa098/gosfml2"
    "runtime"
    "time"
    "sergio/animated_sprites/animatedSprite"
)

func init() {
    runtime.LockOSThread()
}
 
func main() {
    screenDimensions := sf.Vector2i{800, 600}
    ticker := time.NewTicker(time.Second / 60)
    renderWindow := sf.NewRenderWindow(sf.VideoMode{screenDimensions.x, screenDimensions.y, 32}, "Animated sprites", sf.StyleDefault, sf.DefaultContextSettings())
    renderWindow.SetVSyncEnabled(true)
    texture, _ := sf.NewTextureFromFile("animated_sprite_sheet.png", nil)
    sprite, _ := sf.NewSprite(texture)
    sprite.SetPosition(sf.Vector2f{10, 50})
 
    var walkingAnimationDown Animation
    walkingAnimationDown.setSpriteSheet(texture)
    walkingAnimationDown.addFrame(sf.IntRect(32, 0, 32, 32))
    walkingAnimationDown.addFrame(sf.IntRect(64, 0, 32, 32))
    walkingAnimationDown.addFrame(sf.IntRect(32, 0, 32, 32))
    walkingAnimationDown.addFrame(sf.IntRect( 0, 0, 32, 32))

    var walkingAnimationLeft Animation
    walkingAnimationLeft.setSpriteSheet(texture)
    walkingAnimationLeft.addFrame(sf.IntRect(32, 32, 32, 32))
    walkingAnimationLeft.addFrame(sf.IntRect(64, 32, 32, 32))
    walkingAnimationLeft.addFrame(sf.IntRect(32, 32, 32, 32))
    walkingAnimationLeft.addFrame(sf.IntRect( 0, 32, 32, 32))

    var walkingAnimationRight Animation
    walkingAnimationRight.setSpriteSheet(texture)
    walkingAnimationRight.addFrame(sf.IntRect(32, 64, 32, 32))
    walkingAnimationRight.addFrame(sf.IntRect(64, 64, 32, 32))
    walkingAnimationRight.addFrame(sf.IntRect(32, 64, 32, 32))
    walkingAnimationRight.addFrame(sf.IntRect( 0, 64, 32, 32))

    var walkingAnimationUp Animation
    walkingAnimationUp.setSpriteSheet(texture)
    walkingAnimationUp.addFrame(sf.IntRect(32, 96, 32, 32))
    walkingAnimationUp.addFrame(sf.IntRect(64, 96, 32, 32))
    walkingAnimationUp.addFrame(sf.IntRect(32, 96, 32, 32))
    walkingAnimationUp.addFrame(sf.IntRect( 0, 96, 32, 32))

    currentAnimation *Animation = &walkingAnimationDown

    // set up AnimatedSprite
    animatedSprite := animatedSprite.NewAnimatedSprite(2000, true, false)
    animatedSprite.setPosition(sf.Vector2f(screenDimensions / 2))

    frameClock := time.NewTimer()

    speed := 80.0
    noKeyWasPressed := true

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
        frameTime := frameClock.Restart()

        renderWindow.Clear(sf.ColorWhite())
        renderWindow.Draw(sprite, sf.DefaultRenderStates())
        renderWindow.Display()
    }
}