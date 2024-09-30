package game

import "github.com/hajimehoshi/ebiten/v2"

type Game struct{
	player *Player
	lasers []*Laser
}

func NewGame() *Game {
	g := &Game{}
	player := NewPlayer(g)
	g.player = player
	return g
	
}

// metodos que a lib espera para funcionar --------------------------------------
// atualiza a logica do jogo
func (g *Game) Update() error {
	g.player.Update()

	for _, l := range g.lasers {
		l.Update()
	}

	return nil
}

// desenha os objetos na tela
func (g *Game) Draw(scream *ebiten.Image) {
	g.player.Draw(scream)

	for _, l := range g.lasers {
		l.Draw(scream)
	}
}

// tamanaho da tela
func (g *Game) Layout(outsideWith, outsideHeith int) (int,  int) {
	return screamWith, screamHeith
}
//---------------------------------------------------------------------------------

func (g *Game) AddLasers(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}