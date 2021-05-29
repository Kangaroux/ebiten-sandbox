package main

import (
	"github.com/Kangaroux/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Title")

	g := game.Game{}
	g.Init(&game.TestScene{})

	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}
}
