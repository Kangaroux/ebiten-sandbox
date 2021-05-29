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
	game *Game

	fpsSamples    *Sampler
	updateSamples *Sampler
	lastDraw      float64
	lastUpdate    float64
	nextRefresh   float64
	refreshRate   int
	text          string
}

var _ Scene = (*TestScene)(nil)

func (s *TestScene) Draw(canvas *ebiten.Image) {
	now := s.game.ElapsedReal()

	s.fpsSamples.Add(1 / (now - s.lastDraw))
	canvas.Fill(color.RGBA{50, 50, 50, 255})

	if now >= s.nextRefresh {
		s.nextRefresh = now + (1 / float64(s.refreshRate))
		s.text = fmt.Sprintf("elapsed: %.3f\ntps: %.2f\nfps: %.2f",
			s.game.ElapsedGame(),
			s.updateSamples.Average(),
			s.fpsSamples.Average(),
		)
	}

	ebitenutil.DebugPrint(canvas, s.text)

	s.lastDraw = now
}

func (s *TestScene) Init(game *Game) {
	s.game = game
	s.fpsSamples = NewSampler(10)
	s.updateSamples = NewSampler(10)
	s.refreshRate = 10
}

func (s *TestScene) Update(deltaTime float64) {
	now := s.game.ElapsedReal()
	s.updateSamples.Add(1 / (now - s.lastUpdate))
	s.lastUpdate = now
}
