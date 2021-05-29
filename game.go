package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(canvas *ebiten.Image) {

}

func (g *Game) Layout(w, h int) (int, int) {
	return 320, 240
}
