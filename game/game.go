package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{
	player *Player
	lasers []*Laser
	meteors []*Meteor
	meteorSpawnTimer *Timer
}

func NewGame() *Game {
	g := &Game{
		meteorSpawnTimer: NewTimer(24),
	}
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

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()
		m := NewMeteor()
		g.meteors  = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}

	return nil
}

// desenha os objetos na tela
func (g *Game) Draw(scream *ebiten.Image) {
	g.player.Draw(scream)

	for _, l := range g.lasers {
		l.Draw(scream)
	}

	for _, m := range g.meteors {
		m.Draw(scream)
	}

	for _, m := range g.meteors {
		// verifica colisao da nave com meteoro
		if (m.Collider().Intersects(g.player.Collider())){
			fmt.Println("voce perdeu")
			g.Reset() // reseta tudo caso perca
		}
	}
  
	for i, m := range g.meteors {
		for j, l := range g.lasers {
			if m.Collider().Intersects(l.Collider()) {
				g.meteors = append(g.meteors[:i], g.meteors[i + 1:]...)
				g.lasers = append(g.lasers[:j], g.lasers[j + 1:]...)
			}
		}
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

func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.meteors = nil
	g.lasers = nil
	g.meteorSpawnTimer.Reset()
}