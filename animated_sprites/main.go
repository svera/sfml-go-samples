package main

import (
    sf "bitbucket.org/krepa098/gosfml2"
    "runtime"
    "time"
    "sergio/animated_sprites/animatedSprite"
    "sergio/animated_sprites/animation"
)

func init() {
    runtime.LockOSThread()
}
 
func main() {
    screenDimensions := sf.Vector2i{800, 600}
    ticker := time.NewTicker(time.Second / 60)
    renderWindow := sf.NewRenderWindow(sf.VideoMode{uint(screenDimensions.X), uint(screenDimensions.Y), 32}, "Animated sprites", sf.StyleDefault, sf.DefaultContextSettings())
    renderWindow.SetVSyncEnabled(true)
    texture, _ := sf.NewTextureFromFile("animated_sprite_sheet.png", nil)
    sprite, _ := sf.NewSprite(texture)
    sprite.SetPosition(sf.Vector2f{10, 50})
 
    var walkingAnimationDown animation.Animation
    walkingAnimationDown.SetSpriteSheet(texture)
    walkingAnimationDown.AddFrame(sf.IntRect{32, 0, 32, 32})
    walkingAnimationDown.AddFrame(sf.IntRect{64, 0, 32, 32})
    walkingAnimationDown.AddFrame(sf.IntRect{32, 0, 32, 32})
    walkingAnimationDown.AddFrame(sf.IntRect{0, 0, 32, 32})

    var walkingAnimationLeft animation.Animation
    walkingAnimationLeft.SetSpriteSheet(texture)
    walkingAnimationLeft.AddFrame(sf.IntRect{32, 32, 32, 32})
    walkingAnimationLeft.AddFrame(sf.IntRect{64, 32, 32, 32})
    walkingAnimationLeft.AddFrame(sf.IntRect{32, 32, 32, 32})
    walkingAnimationLeft.AddFrame(sf.IntRect{0, 32, 32, 32})

    var walkingAnimationRight animation.Animation
    walkingAnimationRight.SetSpriteSheet(texture)
    walkingAnimationRight.AddFrame(sf.IntRect{32, 64, 32, 32})
    walkingAnimationRight.AddFrame(sf.IntRect{64, 64, 32, 32})
    walkingAnimationRight.AddFrame(sf.IntRect{32, 64, 32, 32})
    walkingAnimationRight.AddFrame(sf.IntRect{0, 64, 32, 32})

    var walkingAnimationUp animation.Animation
    walkingAnimationUp.SetSpriteSheet(texture)
    walkingAnimationUp.AddFrame(sf.IntRect{32, 96, 32, 32})
    walkingAnimationUp.AddFrame(sf.IntRect{64, 96, 32, 32})
    walkingAnimationUp.AddFrame(sf.IntRect{32, 96, 32, 32})
    walkingAnimationUp.AddFrame(sf.IntRect{0, 96, 32, 32})

    currentAnimation := &walkingAnimationDown

    // set up AnimatedSprite
    as := animatedSprite.NewAnimatedSprite(2000, true, false)
    as.SetPosition(sf.Vector2f{float32(screenDimensions.X / 2), float32(screenDimensions.Y / 2)})

    speed := float32(80.0)
    noKeyWasPressed := true
    tickNumber := 0

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
                    movement := sf.Vector2f{0.0, 0.0}
                    if ev.Code == sf.KeyUp {
                        currentAnimation = &walkingAnimationUp
                        movement.Y -= speed
                        noKeyWasPressed = false
                    }
                    if ev.Code == sf.KeyDown {
                        currentAnimation = &walkingAnimationDown
                        movement.Y += speed
                        noKeyWasPressed = false
                    }
                    if ev.Code == sf.KeyLeft {
                        currentAnimation = &walkingAnimationLeft
                        movement.X -= speed
                        noKeyWasPressed = false
                    }
                    if ev.Code == sf.KeyRight {
                        currentAnimation = &walkingAnimationRight
                        movement.X += speed;
                        noKeyWasPressed = false;
                    }
                    as.Play(currentAnimation)
                    as.Move(movement)

                    // if no key was pressed stop the animation
                    if noKeyWasPressed {
                        as.Stop()
                    }
                    noKeyWasPressed = true

                case sf.EventClosed:
                    renderWindow.Close()
 
                }
            }
            tickNumber += 1
            if tickNumber == 60 {
                tickNumber = 0
            }
            if (tickNumber % 2 == 0) {
                as.Update()
            }
        }

        renderWindow.Clear(sf.ColorWhite())
        renderWindow.Draw(as, sf.DefaultRenderStates())
        renderWindow.Display()
    }
}