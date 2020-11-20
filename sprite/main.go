package main

import (
	"encoding/csv"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/pkg/errors"
	"golang.org/x/image/colornames"
	"image"
	_ "image/png"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func loadAnimationSheet(sheetPath, descPath string, frameWidth float64) (sheet pixel.Picture, anims map[string][]pixel.Rect, err error) {
	// total hack, nicely format the error at the end, so I don't have to type it every time
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "error loading animation sheet")
		}
	}()

	// open and load the spritesheet
	sheetFile, err := os.Open(sheetPath)
	if err != nil {
		return nil, nil, err
	}
	defer sheetFile.Close()
	sheetImg, _, err := image.Decode(sheetFile)
	if err != nil {
		return nil, nil, err
	}
	sheet = pixel.PictureDataFromImage(sheetImg)

	// create a slice of frames inside the spritesheet
	var frames []pixel.Rect
	for x := 0.0; x+frameWidth <= sheet.Bounds().Max.X; x += frameWidth {
		frames = append(frames, pixel.R(
			x,
			0,
			x+frameWidth,
			sheet.Bounds().H(),
		))
	}

	descFile, err := os.Open(descPath)
	if err != nil {
		return nil, nil, err
	}
	defer descFile.Close()

	anims = make(map[string][]pixel.Rect)

	// load the animation information, name and interval inside the spritesheet
	desc := csv.NewReader(descFile)
	for {
		anim, err := desc.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, err
		}

		name := anim[0]
		start, _ := strconv.Atoi(anim[1])
		end, _ := strconv.Atoi(anim[2])

		anims[name] = frames[start : end+1]
	}

	return sheet, anims, nil
}

func run() {
	//rand.Int
	//获取随机数，不加随机种子，每次遍历获取都是重复的一些随机数据
	//rand.Seed(time.Now().UnixNano())
	//设置随机数种子，加上这行代码，可以保证每次随机都是随机的
	rand.Seed(time.Now().UnixNano())

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Mygame Moon!",
		Bounds: pixel.R(0, 0, 1000, 1000),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// FPS
	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	/******main code start*******/

	// 清晰设置
	win.SetSmooth(true)

	dir, _ := os.Getwd()
	fmt.Println("当前路径：", dir)
	fmt.Println("当前路径：", dir+"/sheet.png")

	sheet, anims, err := loadAnimationSheet("sheet.png", "sheet.csv", 12)
	if err != nil {
		panic(err)
	}
	imd := imdraw.New(sheet)
	imd.Precision = 32

	// 计数器
	counter := 0.0
	// 比率
	rate := 1.0 / 10
	// 精灵
	sprite := pixel.NewSprite(nil, pixel.Rect{})
	/******main code end*******/

	win.Clear(colornames.Burlywood)

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		counter += dt

		i := int(math.Floor(counter / rate))
		fmt.Println("i:" + strconv.Itoa(i))
		fmt.Println(counter)
		frame := anims["Run"][i%len(anims["Run"])]

		fmt.Println("i%len(anims[\"Run\"]:" + strconv.Itoa(len(anims["Run"])))
		fmt.Println("i%len(anims[\"Run\"]:" + strconv.Itoa(i%len(anims["Run"])))
		sprite.Set(sheet, frame)

		mat := pixel.IM
		mat = mat.ScaledXY(pixel.ZV, pixel.V(5, 5))
		mat = mat.Moved(win.Bounds().Center())
		sprite.Draw(win, mat)
		//sprite.Draw(imd,)

		//imd.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

		if i == 26 {
			sprite.Draw(win, pixel.IM.Moved(pixel.V(100, 100)))
			counter = 0
		}

		win.Update()

		// FPS
		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d | dt: %d", cfg.Title, frames, dt))
			frames = 0
		default:
		}

	}
}

func main() {
	pixelgl.Run(run)
}
