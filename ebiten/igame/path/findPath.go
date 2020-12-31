package path

import (
	"container/list"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"pixel_AStart/ebiten/common"
	_map "pixel_AStart/ebiten/igame/map"
	"pixel_AStart/ebiten/queue"
	"pixel_AStart/ebiten/tiled"
	"time"
)

var IPath = &Path{}

type Path struct {
	image     *ebiten.Image
	imgAttack *ebiten.Image
	imgMove   *ebiten.Image
	// 当前
	X, Y int

	// 父节点
	PX, PY int

	// 移动力
	MovePower int
}

// 显示true或隐藏false
var IsShow bool

// close
var paths map[string]*Path

// 攻击范围
var attackRange map[string]*Path

// open 先进 先出
var patch = queue.NewQueue()

func Init() {

	imgPath := ebiten.NewImage(15, 15)

	// 7AC5CD -> 122,197,205    	#4682B4
	imgPath.Fill(color.RGBA{0x46, 0x82, 0xB4, 0xC8})

	//DC143C
	imgAttack := ebiten.NewImage(15, 15)
	imgAttack.Fill(color.RGBA{0xDC, 0x14, 0x3C, 0xC8})

	//66CD00
	imgMove := ebiten.NewImage(15, 15)
	imgMove.Fill(color.RGBA{0x66, 0xCD, 0x00, 0xC8})
	IPath = &Path{
		image:     imgPath,
		imgAttack: imgAttack,
		imgMove:   imgMove,
	}

	Roles_XY = make(map[string]string)

	IsShow = true
}

// 所有路径
func (p *Path) Find(x, y, attackRange int) {

	// close
	paths = make(map[string]*Path)

	// 记录已查找过的
	open = make(map[string]*Path)

	pa := &Path{
		X: x,
		Y: y,

		PX: 0,
		PY: 0,

		MovePower: 6,
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

	findAttack(attackRange)

	fmt.Println(time.Now())

}

var Roles_XY map[string]string

func findPath() bool {

	// 取一个值
	p := pop()
	if p == nil {
		return false
	}

	paths[tiled.GetKey(p.X, p.Y)] = p

	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		if n := _map.Worlds.Tile(p.X+offset[0], p.Y+offset[1]); n != nil {
			x := n.X / common.TileSize
			y := n.Y / common.TileSize

			if _, ok := Roles_XY[tiled.GetKey(x, y)]; ok {
				continue
			}

			// 不在open 和 close 中
			if _, ok := paths[tiled.GetKey(x, y)]; !ok {
				if !containsKey(tiled.GetKey(x, y)) {
					// 当前地图块 必须是移动 并且 行动力 允许
					if n.Property.Mp != 0 && p.MovePower-n.Property.Mp > 0 {

						np := &Path{
							X: x,
							Y: y,

							PX: p.X,
							PY: p.Y,

							MovePower: p.MovePower - n.Property.Mp,
						}

						push(np)
					}
				}
			}
		}
	}

	return true
}

func findAttack(ar int) bool {
	attackRange = make(map[string]*Path)

	// 循环路径
	for _, v := range paths {

		// 简易计算 只有一格或者两格
		for _, offset := range [][]int{
			{-ar, 0},
			{ar, 0},
			{0, -ar},
			{0, ar},
		} {
			if n := _map.Worlds.Tile(v.X+offset[0], v.Y+offset[1]); n != nil {
				x := n.X / common.TileSize
				y := n.Y / common.TileSize

				if _, ok := Roles_XY[tiled.GetKey(x, y)]; ok {
					continue
				}

				if _, ok := paths[tiled.GetKey(x, y)]; !ok {
					p := &Path{
						X: x,
						Y: y,
					}
					attackRange[tiled.GetKey(x, y)] = p
				}

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
	if IsShow {
		// close
		for _, v := range paths {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(v.X*16)+1, float64(v.Y*16)+1+float64(common.OffsetY))
			op.GeoM.Scale(common.Scale, common.Scale)
			screen.DrawImage(p.image, op)
		}

		// 攻击范围
		for _, v := range attackRange {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(v.X*16)+1, float64(v.Y*16)+1+float64(common.OffsetY))
			op.GeoM.Scale(common.Scale, common.Scale)
			screen.DrawImage(p.imgAttack, op)
		}

		// 移动路径
		for i := MovepathList.Front(); i != nil; i = i.Next() {

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i.Value.(*Path).X*16)+1, float64(i.Value.(*Path).Y*16)+1+float64(common.OffsetY))
			op.GeoM.Scale(common.Scale, common.Scale)
			screen.DrawImage(p.imgMove, op)
		}
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
