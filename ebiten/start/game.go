package start

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"pixel_AStart/ebiten/camera"
	"pixel_AStart/ebiten/common"
	"pixel_AStart/ebiten/role/roy"

	"pixel_AStart/ebiten/tiled"
	"time"
)

// 光标
var cursor = &tiled.Cursor{
	Count: 0,
	Scale: common.Scale,
}

// 主角 罗伊
var sroy = &roy.Roy{
	Count:  0,
	Scale:  common.Scale,
	Status: 1,
}

func init() {

	common.Init()

	// 地图初始化
	tiled.Init()
	// 光标
	cursor.Init(common.RealPath + "/resource/images/cursor.png")
	// 罗伊
	sroy.Init(common.RealPath + "/resource/02/Map_Lord_Roy.png")

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

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	g.dt = time.Since(g.last).Seconds()
	g.last = time.Now()

	//fmt.Println(g.dt)

	// 鼠标选中
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		fmt.Println(tiled.GetKey(int(float64(x)/(16*common.Scale)), int(float64(y)/(16*common.Scale))))
		tile := tiled.Tiles[tiled.GetKey(int(float64(x)/(16*common.Scale)), int(float64(y)/(16*common.Scale)))]
		if tile != nil {
			fmt.Println("tile name type:" + tile.Name + "**" + tile.Type)
		} else {
			fmt.Println("tile is nil")
		}
	}

	// 光标
	cursor.Update(g.dt)
	// 罗伊
	sroy.Update(g.dt)

	// 光标 是否选中 精灵  后期可改为 循环多个角色
	if cursor.X == sroy.X && cursor.Y == sroy.Y {
		cursor.IsSelected = true
		sroy.IsSelected = true
	} else {
		cursor.IsSelected = false
		sroy.IsSelected = false
	}

	//if ebiten.IsKeyPressed(ebiten.KeyQ) {
	//	if g.camera.ZoomFactor > -2400 {
	//		g.camera.ZoomFactor -= 1
	//	}
	//}
	//if ebiten.IsKeyPressed(ebiten.KeyE) {
	//	if g.camera.ZoomFactor < 2400 {
	//		g.camera.ZoomFactor += 1
	//	}
	//}

	// Write your game's logical update.
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.

	// 加载地图
	tiled.Draw(screen)

	// 光标
	cursor.Draw(screen)
	// 罗伊
	sroy.Draw(screen)

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
