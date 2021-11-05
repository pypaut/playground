package main

import (
    "fmt"
    "github.com/veandco/go-sdl2/sdl"
)

const (
    screenWidth = 1000
    screenHeight = 800
    playerSpeed = 0.3
    ballSpeed = 0.15
)

func main() {
    if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
        fmt.Println("SDL initialization:", err)
        return
    }

    window, err := sdl.CreateWindow(
        "GoPong",
        sdl.WINDOWPOS_UNDEFINED,
        sdl.WINDOWPOS_UNDEFINED,
        screenWidth,
        screenHeight,
        sdl.WINDOW_OPENGL,
    )
    if err != nil {
        fmt.Println("Window initialization:", err)
        return
    }
    defer func(window *sdl.Window) {
        err := window.Destroy()
        if err != nil {
            return
        }
    }(window)

    renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
    if err != nil {
        fmt.Println("Renderer initialization:", err)
        return
    }
    defer func(renderer *sdl.Renderer) {
        err := renderer.Destroy()
        if err != nil {
            return
        }
    }(renderer)

    playerY := float64(screenHeight / 2 - 50)
    p1PosY := playerY
    p1Rect := sdl.Rect{X: 100, Y: int32(p1PosY), W: 10, H: 100}
    p2PosY := playerY
    p2Rect := sdl.Rect{X: screenWidth - 100, Y: int32(playerY), W: 10, H: 100}
    ballPosX := float64(screenWidth / 2)
    ballPosY := float64(screenHeight / 2)
    ballRect := sdl.Rect{X: int32(ballPosX), Y: int32(ballPosY), W: 10, H: 10}

    ballDirX := 0.0
    ballDirY := 0.0

    isRunning := false

    for {
        for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            switch event.(type) {
            case *sdl.QuitEvent:
                return
            }
        }

        keys := sdl.GetKeyboardState()

        // Player 1
        if keys[sdl.SCANCODE_W] == 1 && p1Rect.Y > 0 {
            p1PosY -= playerSpeed
        }
        if keys[sdl.SCANCODE_S] == 1 && p1Rect.Y + 100 < screenHeight {
            p1PosY += playerSpeed
        }
        p1Rect.Y = int32(p1PosY)

        // Player 2
        if keys[sdl.SCANCODE_UP] == 1 && p2Rect.Y > 0 {
            p2PosY -= playerSpeed
        }
        if keys[sdl.SCANCODE_DOWN] == 1  && p2Rect.Y + 100 < screenHeight {
            p2PosY += playerSpeed
        }
        p2Rect.Y = int32(p2PosY)

        // Ball
        if keys[sdl.SCANCODE_SPACE] == 1 && !isRunning {
            isRunning = true
            ballDirX = ballSpeed
        }

        /*
        * COLLISIONS
        */
        // Player collision
        p1Collision := ballRect.HasIntersection(&p1Rect)
        p2Collision := ballRect.HasIntersection(&p2Rect)
        if p1Collision || p2Collision {
            ballDirX *= -1
            ballDirY *= -1

            yBallMiddle := ballPosY + float64(ballRect.H / 2)
            yPlayerMiddle := 0.0

            if p1Collision {
                yPlayerMiddle = p1PosY + float64(p1Rect.H / 2)
            }

            if p2Collision {
                yPlayerMiddle = p2PosY + float64(p2Rect.H / 2)
            }

            yPlayerCollision := yBallMiddle - yPlayerMiddle
            ballDirY += 0.002 * yPlayerCollision
        }

        // Wall collision
        if ballPosY < 0 || ballPosY+ float64(ballRect.H) > screenHeight {
            ballDirY *= -1
        }

        if ballPosX < 0 || ballPosX+ float64(ballRect.W) > screenWidth {
            return
        }

        ballPosX += ballDirX
        ballPosY += ballDirY

        ballRect.X = int32(ballPosX)
        ballRect.Y = int32(ballPosY)

        /*
        * DRAW
        */

        // Background
        err := renderer.SetDrawColor(0, 0, 0, 255)
        if err != nil {
            return
        }
        err = renderer.Clear()
        if err != nil {
            return
        }

        // Player 1
        err = renderer.SetDrawColor(150, 0, 150, 255)
        if err != nil {
            return
        }
        err = renderer.FillRect(&p1Rect)
        if err != nil {
            return
        }

        // Player 2
        err = renderer.SetDrawColor(150, 0, 150, 255)
        if err != nil {
            return
        }
        err = renderer.FillRect(&p2Rect)
        if err != nil {
            return
        }

        // Ball
        err = renderer.SetDrawColor(150, 0, 150, 255)
        if err != nil {
            return
        }
        err = renderer.FillRect(&ballRect)
        if err != nil {
            return
        }

        renderer.Present()
    }
}
