package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Draw(canvas *ebiten.Image)
}

type TestScene struct{}

func (s *TestScene) Draw(canvas *ebiten.Image) {
	canvas.Fill(color.RGBA{50, 50, 50, 255})
}
