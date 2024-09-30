package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct{}

// metodos que a lib espera para funcionar --------------------------------------
// atualiza a logica do jogo
func (g *Game) Update() error {
	return nil
}

// desenha os objetos na tela
func (g *Game) Draw(scream *ebiten.Image) {
	
}

// tamanaho da tela
func (g *Game) Layout(outsideWith, outsideHeith int) (screamWith, screamHeith int) {
	return outsideWith, outsideHeith
}
//---------------------------------------------------------------------------------

func main () {
	g := &Game{}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}