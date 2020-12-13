package path

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"pixel_AStart/ebiten/common"
	"pixel_AStart/ebiten/queue"
	"pixel_AStart/ebiten/tiled"
	"time"
)

type Path struct {
	image *ebiten.Image
	// 当前
	X, Y int

	// 父节点
	PX, PY int

	// 移动力
	movePower int
}

// close
var paths map[string]*Path

// open 先进 先出
var patch = queue.NewQueue()

func NewPath() (*Path, error) {
	//img, _, err := ebitenutil.NewImageFromFile(common.RealPath + "/resource/02/Restore.png")
	//if err != nil {
	//	log.Fatal(err)
	//}

	imgPath := ebiten.NewImage(15, 15)
	//imgPath.Fill(color.White)

	//color.RGBA{R, G, B , A}
	//R、G、B三个参数，正整数值的取值范围为：0 - 255。百分数值的取值范围为：0.0% - 100.0%。
	//超出范围的数值将被截至其最接近的取值极限。这里是16进制
	//A为透明度参数，取值在0~1之间，不可为负值
	//imgPath.Fill(color.RGBA{0x00, 0x80, 0x00, 0x80})

	// 7AC5CD -> 122,197,205    	#4682B4
	imgPath.Fill(color.RGBA{0x46, 0x82, 0xB4, 0xC8})
	//imgPath.Fill(color.RGBA{0x66, 0xcc, 0xff, 0xff})
	//imgPath.Fill(color.RGBA{0x66, 0xcc, 0xff, 0xff})

	pa := &Path{
		image: imgPath,
	}

	return pa, nil
}

// 所有路径
func (p *Path) Find(x, y int) {

	// close
	paths = make(map[string]*Path)

	// 记录已查找过的
	open = make(map[string]*Path)

	pa := &Path{
		X: x,
		Y: y,

		PX: 0,
		PY: 0,

		movePower: 6,
	}

	// 加入队列
	push(pa)

	// 查询路径  直到  open 列表为 空
	fmt.Println(time.Now())
	for {
		if !findPath() {
			break
		}
	}
	fmt.Println(time.Now())
	//go func() {
	//	for {
	//
	//		// 500 毫秒
	//		time.Sleep(time.Duration(100)*time.Millisecond)
	//		fmt.Println("")
	//		fmt.Println("************************")
	//		fmt.Println(time.Now())
	//
	//		fmt.Println("open:"+strconv.Itoa(patch.Size()))
	//		fmt.Println("close:"+strconv.Itoa(len(paths)))
	//		if !findPath() {
	//			break
	//		}
	//		fmt.Println("open:"+strconv.Itoa(patch.Size()))
	//		fmt.Println("close:"+strconv.Itoa(len(paths)))
	//
	//		// 一秒
	//		//time.Sleep(time.Duration(1)*time.Second)
	//		//time.Sleep(2e9)
	//		fmt.Println(time.Now())
	//		fmt.Println("************************")
	//		fmt.Println("")
	//	}
	//}()
}

func findPath() bool {

	// 取一个值
	p := pop()
	if p == nil {
		return false
	}

	paths[tiled.GetKey(p.X, p.Y)] = p

	// 上下左右
	u := NewPathByPare(p, 1)
	d := NewPathByPare(p, 2)
	l := NewPathByPare(p, 3)
	r := NewPathByPare(p, 4)

	// 不在close
	if _, ok := paths[tiled.GetKey(u.X, u.Y)]; !ok {
		// 不在open 中 加入 队列
		if !containsKey(tiled.GetKey(u.X, u.Y)) {
			// 检查 当前地图块是否可用
			tile := tiled.Tiles[tiled.GetKey(u.X, u.Y)]
			// 当前地图块 必须是移动 并且 行动力 允许
			if tile.Property.Mp != 0 && p.movePower-tile.Property.Mp > 0 {
				u.movePower = p.movePower - tile.Property.Mp
				push(u)
			}

		}
	}
	if _, ok := paths[tiled.GetKey(d.X, d.Y)]; !ok {
		if !containsKey(tiled.GetKey(d.X, d.Y)) {
			// 检查 当前地图块是否可用
			tile := tiled.Tiles[tiled.GetKey(d.X, d.Y)]
			// 当前地图块 必须是移动 并且 行动力 允许
			if tile.Property.Mp != 0 && p.movePower-tile.Property.Mp > 0 {
				d.movePower = p.movePower - tile.Property.Mp
				push(d)
			}
		}
	}
	if _, ok := paths[tiled.GetKey(l.X, l.Y)]; !ok {
		if !containsKey(tiled.GetKey(l.X, l.Y)) {
			// 检查 当前地图块是否可用
			tile := tiled.Tiles[tiled.GetKey(l.X, l.Y)]
			// 当前地图块 必须是移动 并且 行动力 允许
			if tile.Property.Mp != 0 && p.movePower-tile.Property.Mp > 0 {
				l.movePower = p.movePower - tile.Property.Mp
				push(l)
			}
		}
	}
	if _, ok := paths[tiled.GetKey(r.X, r.Y)]; !ok {
		if !containsKey(tiled.GetKey(r.X, r.Y)) {
			// 检查 当前地图块是否可用
			tile := tiled.Tiles[tiled.GetKey(r.X, r.Y)]
			// 当前地图块 必须是移动 并且 行动力 允许
			if tile.Property.Mp != 0 && p.movePower-tile.Property.Mp > 0 {
				r.movePower = p.movePower - tile.Property.Mp
				push(r)
			}
		}
	}

	return true
}

// 根据 父 创建
func NewPathByPare(pare *Path, t int) *Path {
	x := pare.X
	y := pare.Y

	// 边界必须大于0
	if t == 1 && pare.Y-1 >= 0 {
		y = pare.Y - 1
	}
	if t == 2 && pare.Y+1 < 21 {
		y = pare.Y + 1
	}
	// 边界必须大于0
	if t == 3 && pare.X-1 >= 0 {
		x = pare.X - 1
	}
	if t == 4 && pare.X+1 < 15 {
		x = pare.X + 1
	}

	return &Path{
		X: x,
		Y: y,

		PX: pare.X,
		PY: pare.Y,
	}
}

// 画路径
func (p *Path) Draw(screen *ebiten.Image) {
	// close
	for _, v := range paths {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(v.X*16)+1, float64(v.Y*16)+1+float64(common.OffsetY))
		op.GeoM.Scale(common.Scale, common.Scale)
		//screen.DrawImage(c.img_status1[(c.Count/5)%4], op)
		//screen.DrawImage(p.image.SubImage(image.Rect(19, 10, 35, 26)).(*ebiten.Image), op)

		//screen.DrawImage(p.image.SubImage(image.Rect(0, 0, 10, 10)).(*ebiten.Image), op)
		screen.DrawImage(p.image, op)

		//v, i := rect(float32(v.X*16), float32(v.Y*16)+float32(common.OffsetY), 16, 16, color.RGBA{0x00, 0x80, 0x00, 0x80})
		//screen.DrawTriangles(v, i, p.image, nil)
	}
}

func (p *Path) Clear() {
	// close
	paths = make(map[string]*Path)
	// 临时记录
	open = make(map[string]*Path)

	// 先进先出队列
	patch = queue.NewQueue()
}

var open = make(map[string]*Path)

// 是否存在
func containsKey(key string) bool {
	_, ok := open[key]
	return ok
}

// 出
func pop() *Path {
	deQueue := patch.DeQueue()
	if deQueue == nil {
		return nil
	}
	// 强制类型转换
	return deQueue.(*Path)
}

// 入
func push(p *Path) {
	patch.EnQueue(p)
	open[tiled.GetKey(p.X, p.Y)] = p
}

func rect(x, y, w, h float32, clr color.RGBA) ([]ebiten.Vertex, []uint16) {
	r := float32(clr.R) / 0xff
	g := float32(clr.G) / 0xff
	b := float32(clr.B) / 0xff
	a := float32(clr.A) / 0xff
	x0 := x
	y0 := y
	x1 := x + w
	y1 := y + h

	return []ebiten.Vertex{
		{
			DstX:   x0,
			DstY:   y0,
			SrcX:   common.Scale,
			SrcY:   common.Scale,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX:   x1,
			DstY:   y0,
			SrcX:   common.Scale,
			SrcY:   common.Scale,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX:   x0,
			DstY:   y1,
			SrcX:   common.Scale,
			SrcY:   common.Scale,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX:   x1,
			DstY:   y1,
			SrcX:   common.Scale,
			SrcY:   common.Scale,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
	}, []uint16{0, 1, 2, 1, 2, 3}
}
