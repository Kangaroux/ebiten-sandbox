package game

import (
	"image"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) ImagePoint() image.Point {
	return image.Pt(int(math.Ceil(p.X)), int(math.Ceil(p.Y)))
}
