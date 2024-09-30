package game

import "github.com/hajimehoshi/ebiten/v2"

type Game struct{
	player *Player
}

func NewGame() *Game {
	player := NewPlayer()
	return &Game{
		player: player,
	}
}

// metodos que a lib espera para funcionar --------------------------------------
// atualiza a logica do jogo
func (g *Game) Update() error {
	return nil
}

// desenha os objetos na tela
func (g *Game) Draw(scream *ebiten.Image) {
	g.player.Draw(scream)
}

// tamanaho da tela
func (g *Game) Layout(outsideWith, outsideHeith int) (screamWith, screamHeith int) {
	return outsideWith, outsideHeith
}
//---------------------------------------------------------------------------------