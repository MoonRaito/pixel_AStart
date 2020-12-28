package path

import (
	"container/list"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"pixel_AStart/ebiten/common"
	"pixel_AStart/ebiten/queue"
	"pixel_AStart/ebiten/tiled"
	"strconv"
	"time"
)

type Path struct {
	image     *ebiten.Image
	imgAttack *ebiten.Image
	imgMove   *ebiten.Image
	// 当前
	X, Y int

	// 父节点
	PX, PY int

	// 移动力
	movePower int
	MovePower int
}

// close
var paths map[string]*Path

// 攻击范围
var attackRange map[string]*Path

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

	//DC143C
	imgAttack := ebiten.NewImage(15, 15)
	imgAttack.Fill(color.RGBA{0xDC, 0x14, 0x3C, 0xC8})

	//66CD00
	imgMove := ebiten.NewImage(15, 15)
	imgMove.Fill(color.RGBA{0x66, 0xCD, 0x00, 0xC8})
	pa := &Path{
		image:     imgPath,
		imgAttack: imgAttack,
		imgMove:   imgMove,
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

	findAttack()

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

func findAttack() bool {
	attackRange = make(map[string]*Path)
	// 循环路径
	for _, v := range paths {

		// 简易计算 只有一格攻击力

		// 上下左右 是否在 路径中
		u := NewPathByPare(v, 1)
		d := NewPathByPare(v, 2)
		l := NewPathByPare(v, 3)
		r := NewPathByPare(v, 4)
		if _, ok := paths[tiled.GetKey(u.X, u.Y)]; !ok {
			attackRange[tiled.GetKey(u.X, u.Y)] = u
		}
		if _, ok := paths[tiled.GetKey(d.X, d.Y)]; !ok {
			attackRange[tiled.GetKey(d.X, d.Y)] = d
		}
		if _, ok := paths[tiled.GetKey(l.X, l.Y)]; !ok {
			attackRange[tiled.GetKey(l.X, l.Y)] = l
		}
		if _, ok := paths[tiled.GetKey(r.X, r.Y)]; !ok {
			attackRange[tiled.GetKey(r.X, r.Y)] = r
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

	// 攻击范围
	for _, v := range attackRange {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(v.X*16)+1, float64(v.Y*16)+1+float64(common.OffsetY))
		op.GeoM.Scale(common.Scale, common.Scale)
		screen.DrawImage(p.imgAttack, op)
	}

	// 移动路径
	//for _, v := range smovepath {
	//	op := &ebiten.DrawImageOptions{}
	//	op.GeoM.Translate(float64(v.X*16)+1, float64(v.Y*16)+1+float64(common.OffsetY))
	//	op.GeoM.Scale(common.Scale, common.Scale)
	//	screen.DrawImage(p.imgMove, op)
	//}

	for i := MovepathList.Front(); i != nil; i = i.Next() {

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(i.Value.(*Path).X*16)+1, float64(i.Value.(*Path).Y*16)+1+float64(common.OffsetY))
		op.GeoM.Scale(common.Scale, common.Scale)
		screen.DrawImage(p.imgMove, op)
	}
}

func (p *Path) Clear() {
	// close
	paths = make(map[string]*Path)
	// 临时记录
	open = make(map[string]*Path)
	// 攻击范围
	attackRange = make(map[string]*Path)

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

func In(x, y int) bool {
	_, ok := paths[tiled.GetKey(x, y)]
	return ok
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

// 选择移动的路径 select move path
//var smovepath map[string]*Path
//// 移动路径
//func MovePath(x,y int) {
//	key := tiled.GetKey(x, y)
//	if p, ok := paths[key]; ok {
//		smovepath = make(map[string]*Path)
//		smovepath[key] = p
//		for {
//			key = tiled.GetKey(p.PX, p.PY)
//			p, ok = paths[key]
//			if ok {
//				smovepath[key] = p
//			}else {
//				break
//			}
//		}
//	}
//}
var MovepathList list.List

// 移动路径
func MovePath(x, y int) {
	key := tiled.GetKey(x, y)
	if p, ok := paths[key]; ok {
		MovepathList.Init()
		MovepathList.PushFront(p)
		for {
			key = tiled.GetKey(p.PX, p.PY)
			p, ok = paths[key]
			if ok {
				MovepathList.PushFront(p)
			} else {
				break
			}
		}
	}
}

/**** 新架构 **********/

// 实现接口
func (p *Path) PathNeighbors() []IPath {
	fmt.Println(strconv.Itoa(p.X) + "  " + strconv.Itoa(p.Y))
	var neighbors []IPath
	//neighbors := []IPath{}
	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		// 排查掉在地图内 并且 不可抵达
		if n := tiled.Worlds.Tile(p.X+offset[0], p.Y+offset[1]); n != nil &&
			n.Property.Mp != 0 {

			pp := &Path{
				X: 1,
				Y: 2,

				PX: 1,
				PY: 2,
			}

			neighbors = append(neighbors, pp)
		}
	}
	return neighbors
}

// PathNeighborCost returns the movement cost of the directly neighboring tile.
func (p *Path) PathNeighborCost(to IPath) int {
	return 0
}

func (p *Path) PathEstimatedCost(to IPath) int {

	return 0
}
