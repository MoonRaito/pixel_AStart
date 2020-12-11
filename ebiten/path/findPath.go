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

var paths map[string]*Path
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

	pa := &Path{
		X: x,
		Y: y,

		PX: 0,
		PY: 0,
	}

	patch.EnQueue(pa)

	// 查询路径  直到  open 列表为 空
	go func() {
		for {
			//if !findPath() {
			//	break
			//}

			fmt.Println(time.Now())
			fmt.Println(len(paths))
			time.Sleep(10000)
			fmt.Println(time.Now())
		}
	}()
}

func findPath() bool {

	// 取一个值
	deQueue := patch.DeQueue()
	if deQueue == nil {
		return false
	}

	// 强制类型转换
	p := deQueue.(*Path)
	paths[tiled.GetKey(p.X, p.Y)] = p

	// 上下左右
	u := NewPathByPare(p, 1)
	d := NewPathByPare(p, 2)
	l := NewPathByPare(p, 3)
	r := NewPathByPare(p, 4)

	if _, ok := paths[tiled.GetKey(u.X, u.Y)]; !ok {
		patch.EnQueue(u)
	}
	if _, ok := paths[tiled.GetKey(d.X, d.Y)]; !ok {
		patch.EnQueue(d)
	}
	if _, ok := paths[tiled.GetKey(l.X, l.Y)]; !ok {
		patch.EnQueue(l)
	}
	if _, ok := paths[tiled.GetKey(r.X, r.Y)]; !ok {
		patch.EnQueue(r)
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
	for _, v := range paths {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(v.X*16), float64(v.Y*16)+float64(common.OffsetY))
		op.GeoM.Scale(common.Scale, common.Scale)
		//screen.DrawImage(c.img_status1[(c.Count/5)%4], op)
		screen.DrawImage(p.image.SubImage(image.Rect(19, 10, 35, 26)).(*ebiten.Image), op)
	}
}
