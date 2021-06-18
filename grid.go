package game

import (
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
	size      int           // The size of the grid, in tiles
	tileSize  int           // The size of each square tile, in pixels
	img       *ebiten.Image // An image of the grid
	isometric *ebiten.Image // An image of the grid using isometric projection
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

func NewGrid(size int) *Grid {
	g := &Grid{size: size, tileSize: sq.Bounds().Dx()}
	g.img = ebiten.NewImage(g.size*g.tileSize, g.size*g.tileSize)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			geo := ebiten.GeoM{}
			geo.Translate(float64(i*g.tileSize), float64(j*g.tileSize))
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

	// curX, curY := ebiten.CursorPosition()
	// fCur := Point{float64(curX), float64(curY)}
	// fCur.X -= 91
	// iso := OrthoToIso(fCur)
	// isoPt := iso.ImagePoint()

	// ebitenutil.DebugPrint(canvas,
	// 	fmt.Sprintf("screen: %d, %d\nworld: %d, %d", curX, curY, isoPt.X, isoPt.Y),
	// )
}

func (s *GridScene) Update(deltaTime float64) {
}

func (s *GridScene) Init(game *Game) {
	s.Grid = NewGrid(2)
}
