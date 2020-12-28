package role

import (
	"github.com/hajimehoshi/ebiten/v2"
	"pixel_AStart/ebiten/path"
	_ "pixel_AStart/ebiten/path"
)

type Sprite struct {
	Image *ebiten.Image

	Name string

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
	ImgStatus1 []*ebiten.Image
	// 指向
	ImgStatus2 []*ebiten.Image
	// 选中
	ImgStatus3 []*ebiten.Image

	// 移动上
	ImgStatus4 []*ebiten.Image
	// 移动左
	ImgStatus5 []*ebiten.Image
	// 移动右
	ImgStatus6 []*ebiten.Image

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

func (t *Sprite) FindPath(x, y int) {

}
