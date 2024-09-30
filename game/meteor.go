package game

import (
	"math/rand"
	"meteor/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Meteor struct {
	image *ebiten.Image
	speed int64
	position Vector
}

func NewMeteor() *Meteor {
	image := assets.MeteorSprites[rand.Intn(len(assets.MeteorSprites))]
	speed := (rand.Float64() * 13)

	positon := Vector{
		x: rand.Float64() * screamWith,
		y: -100,
	}

	return &Meteor{
		image: image,
		speed: int64(speed),
		position: positon,
	}
}

func (m *Meteor) Update() {
	m.position.y += float64(m.speed)
}

func (m *Meteor) Draw(scream *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(m.position.x, m.position.y)

	scream.DrawImage(m.image, op)
}