package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SceneChanger interface {
	ChangeScene(s Scene)
}

type Drawer interface {
	Draw(canvas *ebiten.Image)
}

type Updater interface {
	Update(deltaTime float64)
}

type Scene interface {
	Drawer
	Updater
	Init(game *Game)
}

type TestScene struct {
	game       *Game
	lastUpdate float64
	updateFreq float64
}

var _ Scene = (*TestScene)(nil)

func (s *TestScene) Draw(canvas *ebiten.Image) {
	canvas.Fill(color.RGBA{50, 50, 50, 255})
	ebitenutil.DebugPrint(canvas, fmt.Sprintf("elapsed: %.3f\ntps: %.2f", s.lastUpdate, ebiten.CurrentTPS()))
}

func (s *TestScene) Init(game *Game) {
	s.game = game
	s.lastUpdate = -1
	s.updateFreq = 10
}

func (s *TestScene) Update(deltaTime float64) {
	if s.game.Elapsed()-s.lastUpdate >= (1 / s.updateFreq) {
		s.lastUpdate = s.game.Elapsed()
	}
}
