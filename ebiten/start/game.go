package start

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"pixel_AStart/ebiten/camera"
	"pixel_AStart/ebiten/common"
	"pixel_AStart/ebiten/path"
	"pixel_AStart/ebiten/role/roy"
	"pixel_AStart/ebiten/tiled"
	"strconv"
	"time"
)

// 光标
var cursor = &tiled.Cursor{
	Count: 0,
	Scale: common.Scale,
}

// 主角 罗伊
var sroy = &roy.Roy{
	Count:     0,
	Scale:     common.Scale,
	Status:    1,
	MoveSpeed: 0.1,
}

// path 可行进 路径
var paths = &path.Path{}

func init() {
	// 设置偏移量
	common.OffsetY = -176

	common.Init()

	// 地图初始化
	tiled.Init()
	// 光标
	cursor.Init(common.RealPath + "/resource/images/cursor.png")
	// 罗伊
	sroy.Init(common.RealPath + "/resource/02/Map_Lord_Roy.png")

	paths, _ = path.NewPath()

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
		//x, y := ebiten.CursorPosition()
		//fmt.Println("x:" + strconv.Itoa(x) + "    y:" + strconv.Itoa(y))
		//fmt.Println(tiled.GetKey(int(float64(x)/(16*common.Scale)), int(float64(y)/(16*common.Scale))))
		//tile := tiled.Tiles[tiled.GetKey(int(float64(x)/(16*common.Scale)), int(float64(y)/(16*common.Scale)))]
		//if tile != nil {
		//	fmt.Println("tile name type:" + tile.Name + "**" + tile.Type)
		//} else {
		//	fmt.Println("tile is nil")
		//}
		//
		//sroy.X = x/common.Scale - ((x / common.Scale) % 16)
		//sroy.Y = y/common.Scale - ((y / common.Scale) % 16) + common.OffsetY
		fmt.Println("pianyiliang:" + strconv.Itoa(common.OffsetY))
	}

	// 光标
	cursor.Update(g.dt)
	// 罗伊
	sroy.Update(g.dt)

	// 选中角色后的路径
	path.MovePath(cursor.X/common.TileSize, cursor.Y/common.TileSize)

	// 光标 是否选中 精灵  后期可改为 循环多个角色
	if cursor.X == sroy.X && cursor.Y == sroy.Y {
		cursor.IsSelected = true

		// 当前角色 活跃状态 改为 指向
		if sroy.Status == 1 {
			sroy.Status = 2
		}
	} else {
		cursor.IsSelected = false
		// 当前角色 状态 改为 指活跃
		if sroy.Status == 2 {
			sroy.Status = 1
		}
	}

	// 选中并按下空格
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if sroy.Status == 3 {
			// 光标选中角色时  显示物品  等操作
			if cursor.X == sroy.X && cursor.Y == sroy.Y {

			} else {
				// 光标 是否 是在路径范围内
				if path.In(cursor.X/common.TileSize, cursor.Y/common.TileSize) {
					//sroy.X = cursor.X
					//sroy.Y = cursor.Y

					// 角色被指向时
					if sroy.Status == 2 {
					}

					// 如果 已被选中 ，那么开始移动角色
					if sroy.Status == 3 {
						sroy.MoveX = float64(sroy.X)
						sroy.MoveY = float64(sroy.Y)
						sroy.Moving()
						sroy.MoveTo()
						sroy.Status = 4
					}
				} else {
					// 声音特效
				}
			}

			// 待机
			//sroy.Status = 8
		}

		// 状态 1 寻路
		if sroy.Status == 2 {
			paths.Find(sroy.X/16, sroy.Y/16)
			// 选中
			sroy.Status = 3
		}

	}
	// 临时 用c 清图
	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		paths.Clear()

		sroy.Status = 1
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

	// 路径
	paths.Draw(screen)

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
