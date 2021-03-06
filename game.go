package game

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	activeScene    Scene
	lastUpdateTime float64
	startTime      float64
}

var _ ebiten.Game = (*Game)(nil)
var _ SceneChanger = (*Game)(nil)

func timeNow() float64 {
	return float64(time.Now().UnixNano()) / 1_000_000_000
}

func (g *Game) ChangeScene(s Scene) {
	g.activeScene = s
	g.activeScene.Init(g)
}

func (g *Game) Init(s Scene) {
	g.ChangeScene(s)
	g.startTime = timeNow()
}

func (g *Game) ElapsedGame() float64 {
	return g.lastUpdateTime - g.startTime
}

func (g *Game) ElapsedReal() float64 {
	return timeNow() - g.startTime
}

func (g *Game) Update() error {
	now := timeNow()
	delta := now - g.lastUpdateTime
	g.lastUpdateTime = now

	g.activeScene.Update(delta)

	return nil
}

func (g *Game) Draw(canvas *ebiten.Image) {
	canvas.Fill(color.White)
	g.activeScene.Draw(canvas)
}

func (g *Game) Layout(w, h int) (int, int) {
	scale := 6
	return w / scale, h / scale
}
