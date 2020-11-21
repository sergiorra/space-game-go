package spacegame

import "github.com/faiface/pixel"

type Player struct {
	direction Direction
	world     *World
	sprite    *pixel.Sprite
	life      int
	pos       *pixel.Vec
	vel       float64
}

func NewPlayer(path string, life int, world *World) (*Player, error) {
	// Initialize sprite to use with the player
	pic, err := loadPicture(path)
	if err != nil {
		return nil, err
	}
	spr := pixel.NewSprite(pic, pic.Bounds())
	initialPos := pixel.V(world.Bounds().W()/2, spr.Frame().H())

	return &Player{
		life:   life,
		sprite: spr,
		world:  world,
		pos:    &initialPos,
		vel:    250.0,
	}, nil
}

func (p Player) Draw(t pixel.Target) {
	p.sprite.Draw(t, pixel.IM.Moved(*p.pos))
}
