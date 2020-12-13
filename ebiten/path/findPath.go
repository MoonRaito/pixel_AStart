package path

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"log"
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
}

// close
var paths map[string]*Path

// open 先进 先出
var patch = queue.NewQueue()

func NewPath() (*Path, error) {
	img, _, err := ebitenutil.NewImageFromFile(common.RealPath + "/resource/02/Restore.png")
	if err != nil {
		log.Fatal(err)
	}

	pa := &Path{
		image: img,
	}

	return pa, nil
}

// 所有路径
func (p *Path) Find(x, y int) {

	paths = make(map[string]*Path)

	// 记录已查找过的
	open = make(map[string]*Path)

	pa := &Path{
		X: x,
		Y: y,

		PX: 0,
		PY: 0,
	}

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

	if _, ok := paths[tiled.GetKey(u.X, u.Y)]; !ok {
		// 不在open 中 加入 队列
		if !containsKey(tiled.GetKey(u.X, u.Y)) {
			push(u)
		}
	}
	if _, ok := paths[tiled.GetKey(d.X, d.Y)]; !ok {
		if !containsKey(tiled.GetKey(d.X, d.Y)) {
			push(d)
		}
	}
	if _, ok := paths[tiled.GetKey(l.X, l.Y)]; !ok {
		if !containsKey(tiled.GetKey(l.X, l.Y)) {
			push(l)
		}
	}
	if _, ok := paths[tiled.GetKey(r.X, r.Y)]; !ok {
		if !containsKey(tiled.GetKey(r.X, r.Y)) {
			push(r)
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
		op.GeoM.Translate(float64(v.X*16), float64(v.Y*16)+float64(common.OffsetY))
		op.GeoM.Scale(common.Scale, common.Scale)
		//screen.DrawImage(c.img_status1[(c.Count/5)%4], op)
		screen.DrawImage(p.image.SubImage(image.Rect(19, 10, 35, 26)).(*ebiten.Image), op)
	}
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
