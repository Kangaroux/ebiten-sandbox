package game

import (
	"image"
	"math"
)

type FPoint struct {
	X, Y float64
}

func (fp FPoint) ImagePoint() image.Point {
	return image.Pt(int(math.Ceil(fp.X)), int(math.Ceil(fp.Y)))
}
