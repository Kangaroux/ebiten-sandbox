package main

import "github.com/hajimehoshi/ebiten/v2"

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Title")

	game := Game{}
	game.ChangeScene(&TestScene{})

	if err := ebiten.RunGame(&game); err != nil {
		panic(err)
	}
}
