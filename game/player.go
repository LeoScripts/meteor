package game

import (
	"meteor/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image *ebiten.Image
}

func NewPlayer() *Player {
	image := assets.PlayerSprite
	return &Player{
		image: image,
	}
}

func (p *Player) Update() {

}

func (p *Player) Draw(scream *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(100, 100)

	scream.DrawImage(p.image, op)
}