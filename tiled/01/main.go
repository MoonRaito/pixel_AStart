package main

import (
	"fmt"
	"image/color"
	"os"
	"strconv"
	"time"

	// We must use blank imports for any image formats in the tileset image sources.
	// You will get an error if a blank import is not made; TilePix does not import
	// specific image formats, that is the responsibility of the calling code.
	_ "image/png"

	"github.com/bcvery1/tilepix"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var (
	winBounds    = pixel.R(0, 0, 800, 600)
	camPos       = pixel.ZV
	camSpeed     = 64.0 * 10
	camZoom      = 4.0
	camZoomSpeed = 1.2
)

func run() {

	// FPS
	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	cfg := pixelgl.WindowConfig{
		Title:  "TilePix",
		Bounds: pixel.R(0, 0, 240*3, 336*3),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	dir, _ := os.Getwd()
	fmt.Println("当前路径：", dir)

	// Load and initialise the map.
	m, err := tilepix.ReadFile(dir + "/resource/01/base01.tmx")
	if err != nil {
		panic(err)
	}

	// m 属性 研究
	fmt.Println("m.Height:" + strconv.Itoa(m.Height))
	fmt.Println(m.Height)

	index := 1

	last := time.Now()
	for !win.Closed() {
		win.Clear(color.White)

		dt := time.Since(last).Seconds()
		last = time.Now()

		//mat := pixel.IM
		//mat = mat.ScaledXY(pixel.ZV, pixel.V(2, 2))
		//mat = mat.Moved(win.Bounds().Center().Add(pixel.V(-120,-168)))

		//fmt.Println(m.Bounds())

		// Draw all layers to the window.
		if err := m.DrawAll(win, color.Black, pixel.IM); err != nil {
			panic(err)
		}

		//canvas := pixelgl.NewCanvas(m.Bounds())
		//canvas.Clear(color.Black)
		//canvas.Draw(win,mat.Moved(win.Bounds().Center()))

		//cam := pixel.IM.Scaled(win.Bounds().Center(), scaled).Moved(win.Bounds().Center())

		if index == 1 {

			fmt.Println(win.Bounds())
			fmt.Println(win.Bounds().Center())
			fmt.Println(m.Bounds().Center())

		}
		//cam := pixel.IM.Scaled(win.Bounds().Center(), scaled).Moved(m.Bounds().Center().ScaledXY(pixel.V(scaled,scaled)))
		//cam := pixel.IM.Moved(m.Bounds().Center().ScaledXY(pixel.V(scaled,scaled)))
		//cam := pixel.IM.Moved(pixel.ZV)

		// 标准大小
		//cam := pixel.IM.Scaled(pixel.ZV, scaled).Moved(pixel.ZV)

		// 添加移动
		cam := pixel.IM.Scaled(pixel.ZV, camZoom).Moved(pixel.ZV.Sub(camPos))
		win.SetMatrix(cam)

		if win.Pressed(pixelgl.KeyA) || win.Pressed(pixelgl.KeyLeft) {
			camPos.X -= camSpeed * dt
			if camPos.X < 0 {
				camPos.X = 0
			}
		}
		if win.Pressed(pixelgl.KeyD) || win.Pressed(pixelgl.KeyRight) {
			camPos.X += camSpeed * dt
			if camPos.X > 240 {
				camPos.X = 240
			}
		}
		if win.Pressed(pixelgl.KeyS) || win.Pressed(pixelgl.KeyDown) {
			camPos.Y -= camSpeed * dt
			if camPos.Y < 0 {
				camPos.Y = 0
			}
		}
		if win.Pressed(pixelgl.KeyW) || win.Pressed(pixelgl.KeyUp) {
			camPos.Y += camSpeed * dt
			if camPos.Y > 336 {
				camPos.Y = 336
			}
		}

		if index == 1 {

			index++
		}

		win.Update()

		// FPS
		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		default:
		}
	}
}

func main() {
	pixelgl.Run(run)
}
