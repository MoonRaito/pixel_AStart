package role

import (
	"github.com/hajimehoshi/ebiten/v2"
	"pixel_AStart/ebiten/path"
	_ "pixel_AStart/ebiten/path"
	"pixel_AStart/ebiten/queue"
	"pixel_AStart/ebiten/tiled"
)

type Sprite struct {
	image *ebiten.Image

	name string

	Count, X, Y int
	dt          float64

	// 状态 1：活跃 2:指向 3：选中 4：移动上 5：移动下 6:移动左 7：移动 8：攻击  9：待机
	// 状态 1：活跃 2:指向 3：选中 4：移动
	Status int
	// 是否选中
	IsSelected bool
	// 缩放
	Scale float64

	// 待行动
	imgStatus1 []*ebiten.Image
	// 指向
	imgStatus2 []*ebiten.Image
	// 选中
	imgStatus3 []*ebiten.Image

	// 移动上
	imgStatus4 []*ebiten.Image
	// 移动左
	imgStatus5 []*ebiten.Image
	// 移动右
	imgStatus6 []*ebiten.Image

	// 需要移动到
	MoveX, MoveY           float64
	MoveNumber             float64
	MoveStartX, MoveStartY float64
	MoveEndX, MoveEndY     float64

	// 移动速度
	MoveSpeed float64

	movePower int

	// 路径  close
	Paths map[string]*path.Path

	// 攻击范围 open
	AttackRange map[string]*path.Path
}

// open 先进 先出
var nq = queue.NewQueue()

func (t *Sprite) FindPath(x, y int) {
	t.AttackRange = make(map[string]*path.Path)
	t.Paths = make(map[string]*path.Path)

	pa := &path.Path{
		X: x,
		Y: y,

		PX: 0,
		PY: 0,

		MovePower: t.movePower,
	}

	// 加入队列
	nq.EnQueue(pa)

	for {
		if nq.Size() == 0 {
			// There's no path, return found false.
			return
		}

		p := nq.DeQueue().(*path.Path)

		// open中删除
		delete(t.AttackRange, tiled.GetKey(p.X, p.Y))
		// 加入close
		t.Paths[tiled.GetKey(p.X, p.Y)] = p

		//for _, neighbor := range t.PathNeighbors() {
		//
		//}
	}
}
