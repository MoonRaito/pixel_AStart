package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image"
	_ "image/png"
	"log"
	"os"
	"pixel_AStart/ebiten/tiled"
	"strconv"
	"time"
)

var img *ebiten.Image

var tiles map[string]*tiled.Tile

func getKey(x int, y int) string {
	return strconv.Itoa(x) + "_" + strconv.Itoa(y)
}

// 光标
var cursor = &tiled.Cursor{
	Count: 0,
}

func init() {

	dir, _ := os.Getwd()
	fmt.Println("当前路径：", dir)
	//img_, _, err := image.Decode(bytes.NewReader(images.Ebiten_png))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//img = ebiten.NewImageFromImage(img_)

	// 加载地图
	m, err := tiled.ReadFile(dir + "/resource/01/base01.tmx")
	if err != nil {
		log.Fatal(err)
	}

	var err_ error
	img, _, err_ = ebitenutil.NewImageFromFile(dir + "/resource/01/base.png")
	if err_ != nil {
		log.Fatal(err_)
	}

	//screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)

	// 加载地图中的 object 属性
	tiles = make(map[string]*tiled.Tile)
	for _, og := range m.ObjectGroups {
		for _, o := range og.Objects {
			//image := img.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image)

			x := int(o.X)
			y := int(o.Y) - 16
			eimage := img.SubImage(image.Rect(x, y, int(o.X)+16, int(o.Y))).(*ebiten.Image)
			tile := tiled.Tile{
				Eimage: eimage,
				Name:   o.Name,
				Type:   o.Type,
				X:      x,
				Y:      y,
			}
			tiles[getKey(x/16, y/16)] = &tile
		}
	}

	// 恢复图片 暂当光标使用
	img, _, err_ = ebitenutil.NewImageFromFile(dir + "/resource/02/Restore.png")
	if err_ != nil {
		log.Fatal(err_)
	}

	cursor.Init(dir + "/resource/02/Map_Lord_Roy.png")

}

// Game implements ebiten.Game interface.
type Game struct {
	last time.Time
	dt   float64
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
		fmt.Println(getKey(x/16, y/16))
		tile := tiles[getKey(x/16, y/16)]
		if tile != nil {
			fmt.Println(tile.Name + "**" + tile.Type)
		}
	}

	cursor.Update(g.dt)
	// Write your game's logical update.
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.

	//op := &ebiten.DrawImageOptions{}
	//op.GeoM.Translate(50, 50)
	//op.GeoM.Scale(1.5, 1)
	//screen.DrawImage(img, op)

	for _, tile := range tiles {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tile.X), float64(tile.Y))
		op.GeoM.Scale(2, 2)
		screen.DrawImage(tile.Eimage, op)
	}

	// 光标
	cursor.Draw(screen)

}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	game := &Game{}
	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Your game's title")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
