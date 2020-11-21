package spacegame

import "github.com/faiface/pixel"

type World struct {
	width  float64
	height float64
	bg     pixel.Picture
}

// NewWorld initializes a new world with its properties
func NewWorld(w, h float64) *World {
	return &World{
		width:  w,
		height: h,
	}
}

// AddBackground adds background image to the world
func (w *World) AddBackground(path string) error {
	bg, err := loadPicture(path)
	if err != nil {
		return err
	}

	w.bg = bg
	return nil
}

// Bounds returns world's bounds
func (w World) Bounds() pixel.Rect {
	return pixel.R(0, 0, w.width, w.height)
}

// Draw draws the world
func (w World) Draw(t pixel.Target) {
	spriteBg := pixel.NewSprite(w.bg, w.bg.Bounds())
	bgBatch := pixel.NewBatch(&pixel.TrianglesData{}, w.bg)

	for x := 0.0; x <= w.width; x += spriteBg.Frame().W() {
		for y := 0.0; y <= w.height; y += spriteBg.Frame().H() {
			pos := pixel.V(x, y)
			spriteBg.Draw(bgBatch, pixel.IM.Moved(pos))
		}
	}

	bgBatch.Draw(t)
}