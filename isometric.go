package game

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	isoScale = 1 / math.Sqrt2
)

func IsoMatrix() ebiten.GeoM {
	geo := ebiten.GeoM{}

	geo.Rotate(45 * math.Pi / 180)
	geo.Scale(1, isoScale)

	return geo
}

// ImageToIsometric returns a new copy of src using an isometric projection.
func ImageToIsometric(src *ebiten.Image) *ebiten.Image {
	srcDim := image.Pt(src.Bounds().Dx(), src.Bounds().Dy())
	isoDim := image.Pt(
		int(math.Ceil(math.Sqrt(float64(srcDim.X*srcDim.X+srcDim.Y*srcDim.Y)))),
		srcDim.Y,
	)

	out := ebiten.NewImage(isoDim.X, isoDim.Y)
	geo := IsoMatrix()
	geo.Translate(float64(isoDim.X)/2, 0)

	out.DrawImage(src, &ebiten.DrawImageOptions{
		GeoM:   geo,
		Filter: ebiten.FilterLinear,
	})

	return out
}

func OrthoToIso(p1 Point) (p2 Point) {
	m := IsoMatrix()
	x, y := m.Apply(p1.X, p1.Y)
	return Point{x, y}
}
