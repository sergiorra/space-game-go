package spacegame

import (
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/pixel"
)

type Laser struct {
	pic       pixel.Picture
	pos       *pixel.Vec
	vel       float64
	isVisible bool
	world     *World
	sprite    *pixel.Sprite
	sfxPath   string
}

// NewBaseLaser initializes a new base laser with all common properties
func NewBaseLaser(path, sfxPath string, vel float64, world *World) (*Laser, error) {
	pic, err := loadPicture(path)
	if err != nil {
		return nil, err
	}

	return &Laser{
		pic:     pic,
		vel:     vel,
		world:   world,
		sfxPath: sfxPath,
	}, nil
}

// NewLaser initializes a new laser with common properties and all specific properties
func (l *Laser) NewLaser(pos pixel.Vec) *Laser {
	spr := pixel.NewSprite(l.pic, l.pic.Bounds())
	return &Laser{
		pos:       &pos,
		vel:       l.vel,
		sprite:    spr,
		isVisible: true,
		world:     l.world,
		sfxPath:   l.sfxPath,
	}
}

// Draw draws laser if it is visible
func (l Laser) Draw(t pixel.Target) {
	if l.isVisible == true {
		l.sprite.Draw(t, pixel.IM.Moved(*l.pos))
	}
}

// Update updates laser position and visibility
func (l *Laser) Update() {
	l.pos.Y += l.vel
	if l.pos.Y > l.world.height {
		l.isVisible = false
	}
}

// Shoot loads shoot sound and plays it
func (l Laser) Shoot() {
	sfx, err := loadSound(l.sfxPath)
	if err != nil {
		log.Fatal(err)
	}

	speaker.Init(sfx.format.SampleRate, sfx.format.SampleRate.N(time.Second/10))
	defer sfx.streamer.Close()

	done := make(chan bool)
	speaker.Play(beep.Seq(sfx.streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}