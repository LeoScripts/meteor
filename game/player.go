package game

import (
	"meteor/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image *ebiten.Image
	position Vector
}

func NewPlayer() *Player {
	image := assets.PlayerSprite

	bounds := image.Bounds() // pega largura e altura da imagem
	halfw := float64(bounds.Dx()) /2

	position := Vector{
		x: (screamWith /2) - halfw,
		y: 500,
	}

	return &Player{
		image: image,
		position: position,
	}
}

func (p *Player) Update() {

}

func (p *Player) Draw(scream *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(p.position.x, p.position.y)

	scream.DrawImage(p.image, op)
}