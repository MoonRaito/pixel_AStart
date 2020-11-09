package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	"image/color"
	"moon"
	"strconv"
	"time"
)

func addWall(walls map[string]*moon.Iblock, v pixel.Vec) {
	addBlockMap(walls, v, pixel.RGB(0.4, 0.4, 0.4), 1)
}

func addBlockMap(walls map[string]*moon.Iblock, v pixel.Vec, color color.Color, btype int) *moon.Iblock {
	bsize := 100

	/* 画图 */
	x := int(v.X) / bsize
	y := int(v.Y) / bsize

	key := strconv.Itoa(x) + "_" + strconv.Itoa(y)
	fmt.Println(key)

	wall, ok := walls[key]

	// 不存在添加墙
	if !ok {

		wallX := float64(x) * 100
		wallY := float64(y) * 100

		wall = &moon.Iblock{
			X:     x,
			Y:     y,
			Btype: btype,
			Color: color,
			Rect:  pixel.R(wallX+1, wallY+1, wallX+100-1, wallY+100-1),
			G:     1,
			F:     0,
			H:     0,
		}

		wall.UpdateIblock()

		// 添加墙
		walls[key] = wall
		//fmt.Println(wall)

	}

	return wall
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Mygame Moon!",
		Bounds: pixel.R(0, 0, 1005, 805),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	//imd.Color = pixel.RGB(1, 0, 0)
	//imd.Push(pixel.V(200, 100))
	//imd.Color = pixel.RGB(0, 1, 0)
	//imd.Push(pixel.V(800, 100))
	//imd.Color = pixel.RGB(0, 0, 1)
	//imd.Push(pixel.V(500, 700))
	//imd.Polygon(0)

	imd.Color = pixel.RGB(1, 0, 0)
	imd.Push(pixel.V(100, 100))
	imd.Color = pixel.RGB(0, 1, 0)
	imd.Push(pixel.V(200, 100))
	imd.Color = pixel.RGB(0, 0, 1)
	imd.Push(pixel.V(200, 200))
	imd.Color = pixel.RGB(0, 0, 1)
	imd.Push(pixel.V(100, 200))
	imd.Polygon(0)

	imd2 := imdraw.New(nil)
	imd2.Push(pixel.V(200, 100))
	imd2.Push(pixel.V(300, 100))
	imd2.Push(pixel.V(300, 200))
	imd2.Color = pixel.RGB(0, 1, 0)
	imd2.Push(pixel.V(200, 200))
	imd2.Polygon(0)

	imd2.Color = pixel.RGB(0, 1, 0)

	imd.Color = colornames.Blueviolet
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(100, 100), pixel.V(700, 100))
	imd.EndShape = imdraw.SharpEndShape
	imd.Color = pixel.RGB(0, 1, 0)
	imd.Push(pixel.V(100, 500), pixel.V(700, 500))
	imd.Line(30)

	imd.Draw(win)
	imd2.Draw(win)

	s_index := pixel.V(500, 500)
	square := moon.NewSquare(pixel.RGB(1, 0, 0), s_index, 100)
	square.Draw(win)

	//basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	//basicTxt := text.New(pixel.V(100, 500), basicAtlas)

	txt := text.New(s_index, text.NewAtlas(basicfont.Face7x13, text.ASCII))

	fmt.Fprintln(txt, "Hello, text!")
	//fmt.Fprintln(txt, "I support multiple lines!")
	//fmt.Fprintf(txt, "And I'm an %s, yay!", "io.Writer")
	//
	//txt.WriteString(win.Typed())

	txt.Draw(win, pixel.IM)
	//txt.Draw(square,pixel.IM)

	//win.Clear(colornames.Skyblue)

	// 测试 Rect 矩阵？长方形？
	rectangle := imdraw.New(nil)
	rect := pixel.R(0, 0, 100, 100)
	rectangle.Color = pixel.RGB(0, 0, 1)
	rectangle.Push(rect.Min, rect.Max)
	rectangle.Rectangle(2)
	//rectangle.Draw(win)

	// 画布 场景？   画布相对于屏幕的起始位置 默认00 最好为00方便调整
	canvas := pixelgl.NewCanvas(pixel.R(0, 0, 1005, 805))
	canvas.Clear(colornames.Royalblue)

	//txt.Draw(canvas,pixel.IM)

	//canvas.Draw(win, pixel.IM)
	// 设置画布的偏移位置
	matrices := pixel.IM.Moved(canvas.Bounds().Center())
	fmt.Println(matrices)
	//canvas.Draw(win, pixel.IM.Moved(pixel.V(350,400)))
	//canvas.Draw(win, matrices)

	// 画 砖块
	maps := moon.InitMaps()
	for h := range maps {
		for l := range maps[h] {
			x, y := float64(h*100), float64(l*100)
			imDraw := moon.NewRectangle2(pixel.RGB(1, 1, 1), pixel.R(x, y, x+100, y+100), float64(1))
			imDraw.Draw(canvas)
		}
	}

	// 初始化画布
	canvas.Draw(win, matrices)

	//msg := text.New(pixel.V(100,750), text.NewAtlas(basicfont.Face7x13, text.ASCII))
	msg := text.New(pixel.V(100, 750), text.NewAtlas(basicfont.Face7x13, text.ASCII))

	imd3 := imdraw.New(nil)
	times := 0

	// 批量
	//var walls map [string]*moon.Iblock
	//walls = make(map[string]*moon.Iblock)
	// 和上面等价
	walls := make(map[string]*moon.Iblock)

	//testwallone := &moon.Iblock{
	//	X: 1,
	//	Y:1,
	//	Btype:1,
	//	Color : pixel.RGB(1, 0, 0),
	//	Rect:pixel.R(600,600,700,700),
	//}
	//
	//testwallone.UpdateIblock()
	//
	//walls["1_1"] = testwallone
	// 画墙 test end

	///*查看元素在集合中是否存在 */
	//capital, ok := walls [ "American" ] /*如果确定是真实的,则存在,否则不存在 */
	///*fmt.Println(capital) */
	///*fmt.Println(ok) */
	//if (ok) {
	//	fmt.Println("American 的首都是", capital)
	//} else {
	//	fmt.Println("American 的首都不存在")
	//}

	// FPS
	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	var (
		b      = false
		iblock = &moon.Iblock{}
	)

	for !win.Closed() {

		//if win.JustPressed(pixelgl.MouseButtonLeft) {
		if win.Pressed(pixelgl.MouseButtonLeft) {
			if times == 0 {
				imd3 = moon.Update(imd3, pixel.RGB(1, 0, 0))
				times = 1
			} else {
				imd3 = moon.Update(imd3, pixel.RGB(1, 0, 1))
				times = 0
			}

			/* 添加墙 */
			addWall(walls, win.MousePosition())

		}

		if win.Pressed(pixelgl.MouseButtonRight) {
			//win.Clear(colornames.Skyblue)

			// 删除墙
			x := int(win.MousePosition().X) / 100
			y := int(win.MousePosition().Y) / 100
			key := strconv.Itoa(x) + "_" + strconv.Itoa(y)
			delete(walls, key)

		}

		// 设置 开始位置
		if win.JustPressed(pixelgl.KeyS) {
			blockMap := addBlockMap(walls, win.MousePosition(), pixel.RGB(0.2, 0.6, 0.2), -1)
			moon.InitStart(blockMap)
		}

		// 设置 结束位置
		if win.JustPressed(pixelgl.KeyE) {
			blockMap := addBlockMap(walls, win.MousePosition(), pixel.RGB(0.8, 0.8, 0.8), -2)
			moon.InitEnd(blockMap)
		}

		// 清空map
		if win.JustPressed(pixelgl.KeyC) {
			walls = make(map[string]*moon.Iblock)
			moon.InitOpenClose()
		}

		win.Clear(colornames.Skyblue)
		canvas.Draw(win, matrices)

		// 获取寻路
		if win.JustPressed(pixelgl.KeySpace) {
			//moon.FindPath(walls,nil)
			b, iblock = moon.FindPathOneOpen(walls, nil)
			//b, iblock = moon.FindPathAll(walls, nil)
		}

		//imd3.Draw(win)

		for _, wall := range walls {
			wall.Block.Draw(win)

			wall.TxtF.Clear()
			fmt.Fprintln(wall.TxtF, wall.F)
			wall.TxtF.Draw(win, pixel.IM)

			wall.TxtG.Clear()
			fmt.Fprintln(wall.TxtG, wall.G)
			wall.TxtG.Draw(win, pixel.IM)

			wall.TxtH.Clear()
			fmt.Fprintln(wall.TxtH, wall.H)
			wall.TxtH.Draw(win, pixel.IM)

		}

		// 完成
		if b {
			// 文字
			msg.Clear()
			fmt.Fprintln(msg, "寻路到达终点：x:"+strconv.Itoa(iblock.X)+"   y:"+strconv.Itoa(iblock.Y))
			msg.Draw(win, pixel.IM)

			//road := walls[moon.GetKey(iblock.PX, iblock.PY)]
			//road.Color = pixel.RGB(0.5, 0.2, 0.1)
			//road.UpdateIblock()

			path := walls
			moon.DrawPath(path, iblock)
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
