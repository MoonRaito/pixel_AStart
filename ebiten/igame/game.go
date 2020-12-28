package igame

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"pixel_AStart/ebiten/camera"
	"pixel_AStart/ebiten/common"
	_map "pixel_AStart/ebiten/igame/map"
	"time"
)

func init() {
	// 常量
	common.Init()
	// 地图
	_map.Init()

}

// Game implements ebiten.Game interface.
type Game struct {
	last time.Time
	dt   float64

	// 摄像头
	camera camera.Camera
}

// NewGame generates a new Game object.
func NewGame() (*Game, error) {
	g := &Game{}
	return g, nil
}

// Update proceeds the igame state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {

	return nil
}

// Draw draws the igame screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {

	_map.Draw(screen)

	// tps: 每秒调用多少次 更新update
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS()),
	)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return common.ScreenWidth, common.ScreenHeight
}
