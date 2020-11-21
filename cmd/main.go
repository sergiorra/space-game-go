package main

import (
	"github.com/faiface/pixel"
	"github.com/sergiorra/space-game-go/internal"
	"log"

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

	world.Draw(win)

	// infinite loop
	for !win.Closed() {
		win.Update()
	}
}
