package game

import (
	"fmt"
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	_ "image/png"
)

type Grid struct {
	W, H      int
	img       *ebiten.Image
	isometric *ebiten.Image
}

func (g *Grid) DisplayWidth() int {
	return g.img.Bounds().Dx()
}

func (g *Grid) DisplayHeight() int {
	return g.img.Bounds().Dy()
}

var sq *ebiten.Image

func init() {
	var err error

	sq, _, err = ebitenutil.NewImageFromFile("./img/sq2.png")

	if err != nil {
		panic(err)
	}
}

func (g *Grid) Draw(canvas *ebiten.Image) {
	canvas.DrawImage(g.isometric, nil)
}

func toIsometric(img *ebiten.Image) *ebiten.Image {
	yScale := 1 / math.Sqrt2
	srcDim := image.Pt(img.Bounds().Dx(), img.Bounds().Dy())
	isoDim := image.Pt(
		int(math.Ceil(math.Sqrt(float64(srcDim.X*srcDim.X+srcDim.Y*srcDim.Y)))),
		srcDim.Y,
	)

	fmt.Println("src: ", srcDim)
	fmt.Println("iso: ", isoDim)

	out := ebiten.NewImage(isoDim.X, isoDim.Y)
	geo := ebiten.GeoM{}

	geo.Rotate(45 * math.Pi / 180)
	geo.Scale(1, yScale)
	geo.Translate(float64(isoDim.X)/2, 0)
	out.DrawImage(img, &ebiten.DrawImageOptions{
		GeoM: geo,
	})

	return out
}

func NewGrid(w, h int) *Grid {
	g := &Grid{W: w, H: h}
	g.img = ebiten.NewImage(64, 64)

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			geo := ebiten.GeoM{}
			geo.Translate(float64(i*9), float64(j*9))
			g.img.DrawImage(sq, &ebiten.DrawImageOptions{
				GeoM: geo,
			})
		}
	}

	g.isometric = toIsometric(g.img)

	return g
}

type GridScene struct {
	Grid *Grid
}

var _ Scene = (*GridScene)(nil)

func (s *GridScene) Draw(canvas *ebiten.Image) {
	s.Grid.Draw(canvas)
}

func (s *GridScene) Update(deltaTime float64) {
}

func (s *GridScene) Init(game *Game) {
	s.Grid = NewGrid(7, 7)
}
