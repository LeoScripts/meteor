package game

import (
	"meteor/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image *ebiten.Image
	position Vector
	game *Game
	laserLoadingTimer *Timer
}

func NewPlayer(game *Game) *Player {
	image := assets.PlayerSprite

	bounds := image.Bounds() // pega largura e altura da imagem
	halfw := float64(bounds.Dx()) /2

	position := Vector{
		X: (screamWith /2) - halfw,
		Y: 500,
	}

	return &Player{
		image: image,
		game: game,
		position: position,
		laserLoadingTimer: NewTimer(12),
	}
}

func (p *Player) Update() {
	speed := 6.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.X -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.X += speed
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.position.Y -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.position.Y += speed
	}

	// atualizando o tempo de execução do laser
	p.laserLoadingTimer.Update()
	// logica do laser
	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.laserLoadingTimer.IsReady() {
		p.laserLoadingTimer.Reset()
		
		bounds := p.image.Bounds()
		halfW := float64(bounds.Dx()) /2
		halfH := float64(bounds.Dx()) /2

		spawnPos := Vector{
			p.position.X + halfW,
			p.position.Y + halfH,
		}

		laser := NewLaser(spawnPos)
		p.game.AddLasers(laser)
	}
}

func (p *Player) Draw(scream *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(p.position.X, p.position.Y)

	scream.DrawImage(p.image, op)
}

func (p *Player) Collider() Rect {
	bounds := p.image.Bounds()

	return NewRect(
		p.position.X, 
		p.position.Y, 
		float64(bounds.Dx()), 
		float64(bounds.Dy()))
} 