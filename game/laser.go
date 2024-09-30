package game

import (
	"meteor/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Laser struct {
	image *ebiten.Image
	position Vector
}

func NewLaser(position Vector) *Laser {
	image := assets.LaserSprite

	bounds := image.Bounds()

	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	position.X -= halfW
	position.Y -= halfH

	return &Laser{
		image: image,
		position: position,
	}
}

func (l *Laser) Update() {
	speed := 7.0

	l.position.Y += -speed
}

func (l *Laser) Draw(scream *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(l.position.X, l.position.Y)

	scream.DrawImage(l.image, op)
}

func (l *Laser) Collider() Rect {
	bounds := l.image.Bounds()

	return NewRect(
		l.position.X, 
		l.position.Y, 
		float64(bounds.Dx()), 
		float64(bounds.Dy()))
} 