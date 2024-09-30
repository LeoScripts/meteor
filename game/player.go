package game

import (
	"meteor/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image *ebiten.Image
	position Vector
	game *Game
}

func NewPlayer(game *Game) *Player {
	image := assets.PlayerSprite

	bounds := image.Bounds() // pega largura e altura da imagem
	halfw := float64(bounds.Dx()) /2

	position := Vector{
		x: (screamWith /2) - halfw,
		y: 500,
	}

	return &Player{
		image: image,
		game: game,
		position: position,
	}
}

func (p *Player) Update() {
	speed := 6.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.x -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.x += speed
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.position.y -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.position.y += speed
	}

	// logica do laser
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		bounds := p.image.Bounds()
		halfW := float64(bounds.Dx()) /2
		halfH := float64(bounds.Dx()) /2

		spawnPos := Vector{
			p.position.x + halfW,
			p.position.y + halfH,
		}

		laser := NewLaser(spawnPos)
		p.game.AddLasers(laser)
	}
}

func (p *Player) Draw(scream *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(p.position.x, p.position.y)

	scream.DrawImage(p.image, op)
}