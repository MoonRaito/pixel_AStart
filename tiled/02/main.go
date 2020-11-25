package main

import (
	"fmt"
	"image/color"
	"math"
	"os"
	"time"

	// We must use blank imports for any image formats in the tileset image sources.
	// You will get an error if a blank import is not made; TilePix does not import
	// specific image formats, that is the responsibility of the calling code.
	_ "image/png"

	"github.com/bcvery1/tilepix"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	"pixel_AStart/utils"
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

	// 添加精灵
	sheet, _, err := utils.LoadAnimationSheet(dir+"/resource/02/Map_Lord_Roy.png", dir+"/resource/02/sheet.csv", 16)
	if err != nil {
		panic(err)
	}
	//imd := imdraw.New(sheet)
	//imd.Precision = 32

	// 计数器
	counter := 0.0
	// 比率
	rate := 0.5 / 10

	// 精灵
	sprite := pixel.NewSprite(nil, pixel.Rect{})

	//Map_Roy,h,16,minx 695-13-16
	// 精灵 帧
	var s_frames []pixel.Rect
	for i := 0; i < 12; i++ {
		s_frames = append(s_frames, pixel.R(
			115,
			666,
			131,
			682,
		))
	}

	s_frames = append(s_frames, pixel.R(
		135,
		666,
		151,
		682,
	))

	for i := 0; i < 12; i++ {
		s_frames = append(s_frames, pixel.R(
			156,
			666,
			172,
			682,
		))
	}

	s_frames = append(s_frames, pixel.R(
		135,
		666,
		151,
		682,
	))

	last := time.Now()
	for !win.Closed() {
		win.Clear(color.White)

		dt := time.Since(last).Seconds()
		last = time.Now()

		// Draw all layers to the window.
		if err := m.DrawAll(win, color.Black, pixel.IM); err != nil {
			panic(err)
		}

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

		// region 添加精灵 start
		counter += dt

		i := int(math.Floor(counter / rate))
		frame := s_frames[i%len(s_frames)]
		sprite.Set(sheet, frame)

		mat := pixel.IM
		//mat = mat.ScaledXY(pixel.ZV, pixel.V(2, 2))

		//mat = mat.Moved(pixel.V(8,8))
		mat = mat.Moved(pixel.V(30, 30))
		sprite.Draw(win, mat)

		// endregion 添加精灵end

		// 点击事件 选中精灵
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			fmt.Println(win.MousePosition())

			// sprite frame 在图片的截取的图片大小
			fmt.Println(sprite.Frame())
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
