package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	_ "image/png"
)

var sq *ebiten.Image

func init() {
	var err error

	sq, _, err = ebitenutil.NewImageFromFile("./img/sq3.png")

	if err != nil {
		panic(err)
	}
}

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

func (g *Grid) Draw(canvas *ebiten.Image) {
	canvas.DrawImage(g.isometric, nil)
}

func NewGrid(w, h int) *Grid {
	g := &Grid{W: w, H: h}
	tileSize := sq.Bounds().Dx()
	g.img = ebiten.NewImage(g.W*tileSize, g.H*tileSize)

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			geo := ebiten.GeoM{}
			geo.Translate(float64(i*tileSize), float64(j*tileSize))
			g.img.DrawImage(sq, &ebiten.DrawImageOptions{
				GeoM: geo,
			})
		}
	}

	g.isometric = ImageToIsometric(g.img)

	return g
}

type GridScene struct {
	Grid *Grid
}

var _ Scene = (*GridScene)(nil)

func (s *GridScene) Draw(canvas *ebiten.Image) {
	s.Grid.Draw(canvas)

	curX, curY := ebiten.CursorPosition()
	fCur := FPoint{float64(curX), float64(curY)}
	fCur.X -= 91
	iso := OrthoToIso(fCur)
	isoPt := iso.ImagePoint()

	// canvas.Set(isoPt.X, isoPt.Y, color.RGBA{255, 0, 0, 255})

	ebitenutil.DebugPrint(canvas,
		fmt.Sprintf("screen: %d, %d\nworld: %d, %d", curX, curY, isoPt.X, isoPt.Y),
	)
}

func (s *GridScene) Update(deltaTime float64) {
}

func (s *GridScene) Init(game *Game) {
	s.Grid = NewGrid(8, 8)
}
