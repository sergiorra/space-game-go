package main

import (
	"github.com/faiface/pixel"
	"github.com/sergiorra/space-game-go/internal"
	"log"
	"time"

	"github.com/faiface/pixel/pixelgl"
)

const (
	windowWidth  = 1024
	windowHeight = 768
)

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Space Game",
		Bounds: pixel.R(0, 0, windowWidth, windowHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatal(err)
	}

	world := spacegame.NewWorld(windowWidth, windowHeight)
	if err := world.AddBackground("resources/background.png"); err != nil {
		log.Fatal(err)
	}

	player, err := spacegame.NewPlayer("resources/player.png", 5, world)
	if err != nil {
		log.Fatal(err)
	}

	direction := spacegame.Idle
	world.Draw(win)

	last := time.Now()

	// infinite loop
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		if win.Pressed(pixelgl.KeyLeft) {
			direction = spacegame.LeftDirection
		}

		if win.Pressed(pixelgl.KeyRight) {
			direction = spacegame.RightDirection
		}

		player.Update(direction, dt)

		world.Draw(win)
		player.Draw(win)
		direction = spacegame.Idle

		win.Update()
	}
}
