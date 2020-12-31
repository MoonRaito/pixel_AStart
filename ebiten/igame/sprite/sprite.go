package role

import (
	"github.com/hajimehoshi/ebiten/v2"
	"pixel_AStart/ebiten/common"
	"pixel_AStart/ebiten/igame/cursor"
	"pixel_AStart/ebiten/igame/path"
	"pixel_AStart/ebiten/queue"
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
	MoveStartX, MoveStartY int
	MoveEndX, MoveEndY     float64

	// 移动速度
	MoveSpeed float64

	movePower int

	// 攻击范围 1为1格 2为两格
	AttackRange int
}

func (s *Sprite) FindPath(x, y int) {

}

func (s *Sprite) Update(dt float64) {
	s.Count++
	if s.Count > 120 {
		s.Count = 0
	}

	// 1秒 60帧
	s.dt += dt
	// 每个状态的帧数不一致
	if s.Status == 1 {
		if s.dt > 1.0 {
			s.dt = 0
		}
	}
	if s.Status == 2 {
		if s.dt > 1.0 {
			s.dt = 0
		}
	}

	// 光标选中和移开
	s.cursorSelect()

}

// 光标选中和移开时
func (s *Sprite) cursorSelect() {
	// 是否被光标 指向
	// 光标 是否选中 精灵
	if cursor.Icursor.X == s.X && cursor.Icursor.Y == s.Y {
		// 当前角色 活跃状态 改为 指向
		if s.Status == 1 {
			s.Status = 2
		}
	} else {
		// 当前角色 状态 改为 指活跃
		if s.Status == 2 {
			s.Status = 1
		}
	}
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	// 未选中时 并且 状态为 活跃
	if s.Status == 1 {
		s.status1(screen)
	}
	// 光标 指向时
	if s.Status == 2 {
		s.status2(screen)
	}
	// 光标 选中时
	if s.Status == 3 {
		s.status3(screen)
	}

	// 开始移动角色
	if s.Status == 10 {
		s.UpdateMove(screen)
	}
	// 移动完成
	if s.Status == 4 {
		s.MoveFinish(screen)
	}
}

func (c *Sprite) status1(screen *ebiten.Image) {
	i := 0
	if c.dt < 0.45 {
		i = 0
	} else if 0.45 <= c.dt && c.dt <= 0.50 {
		i = 1
	} else if 0.50 < c.dt && c.dt < 0.95 {
		i = 2
	} else if 0.95 <= c.dt {
		i = 1
	}

	//fmt.Println((c.Count/5)%4)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.X), float64(c.Y)+float64(common.OffsetY))
	op.GeoM.Scale(common.Scale, common.Scale)
	//screen.DrawImage(c.img_status1[(c.Count/5)%4], op)
	screen.DrawImage(c.ImgStatus1[i], op)
}

func (c *Sprite) status2(screen *ebiten.Image) {
	i := 0
	if c.dt < 0.45 {
		i = 0
	} else if 0.45 <= c.dt && c.dt <= 0.50 {
		i = 1
	} else if 0.50 < c.dt && c.dt < 0.95 {
		i = 2
	} else if 0.95 <= c.dt {
		i = 1
	}

	//fmt.Println((c.Count/5)%4)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.X)-3, float64(c.Y)-5+float64(common.OffsetY))
	op.GeoM.Scale(c.Scale, c.Scale)
	//screen.DrawImage(c.imgStatus2[(c.Count/5)%3], op)
	screen.DrawImage(c.ImgStatus2[i], op)
}

func (c *Sprite) status3(screen *ebiten.Image) {

	//fmt.Println((c.Count/5)%4)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.X)-1, float64(c.Y)-5+float64(common.OffsetY))
	op.GeoM.Scale(c.Scale, c.Scale)
	screen.DrawImage(c.ImgStatus3[(c.Count/10)%4], op)
}

// 移动
var Moving = queue.NewQueue()

func (c *Sprite) Moving() {
	d := 0
	for i := path.MovepathList.Front(); i != nil; i = i.Next() {
		if d == 0 {
			d++
			continue
		}
		Moving.EnQueue(i.Value.(*path.Path))
	}
}
func (c *Sprite) MoveTo() bool {
	// 取第一个
	p := Moving.DeQueue()

	if p == nil {
		// 设置 角色位置 为终点
		c.X = int(c.MoveEndX)
		c.Y = int(c.MoveEndY)
		return false
	}

	X := p.(*path.Path).X * 16
	Y := p.(*path.Path).Y * 16

	c.MoveEndX = float64(X)
	c.MoveEndY = float64(Y)

	return true
}

var moveLast int

func (c *Sprite) UpdateMove(screen *ebiten.Image) {

	// 到达指定节点后 开始下一个
	if int(c.MoveX) == int(c.MoveEndX) && int(c.MoveY) == int(c.MoveEndY) {
		// 到达终点
		if !c.MoveTo() {
			c.Status = 4
			// 隐藏路径
			//未实现

			return
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(c.MoveX-1, c.MoveY-5+float64(common.OffsetY))
	op.GeoM.Scale(c.Scale, c.Scale)
	// 下
	if int(c.MoveX) == int(c.MoveEndX) && int(c.MoveY) < int(c.MoveEndY) {
		c.MoveY = c.MoveY + (c.MoveSpeed * c.dt)
		screen.DrawImage(c.ImgStatus3[(c.Count/10)%4], op)
		moveLast = 1
	}

	// 上
	if int(c.MoveX) == int(c.MoveEndX) && int(c.MoveY) > int(c.MoveEndY) {
		c.MoveY = c.MoveY - (c.MoveSpeed * c.dt)
		screen.DrawImage(c.ImgStatus4[(c.Count/10)%4], op)
		moveLast = 2
	}

	// 左
	if int(c.MoveX) > int(c.MoveEndX) && int(c.MoveY) == int(c.MoveEndY) {
		c.MoveX = c.MoveX - (c.MoveSpeed * c.dt)
		screen.DrawImage(c.ImgStatus5[(c.Count/10)%4], op)
		moveLast = 3
	}

	// 右
	if int(c.MoveX) < int(c.MoveEndX) && int(c.MoveY) == int(c.MoveEndY) {
		c.MoveX = c.MoveX + (c.MoveSpeed * c.dt)
		screen.DrawImage(c.ImgStatus6[(c.Count/10)%4], op)
		moveLast = 4
	}

}
func (c *Sprite) MoveFinish(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(c.MoveX-1, c.MoveY-5+float64(common.OffsetY))
	op.GeoM.Scale(c.Scale, c.Scale)
	// 下
	if moveLast == 1 {
		screen.DrawImage(c.ImgStatus3[(c.Count/10)%4], op)
	}

	// 上
	if moveLast == 2 {
		screen.DrawImage(c.ImgStatus4[(c.Count/10)%4], op)
	}

	// 左
	if moveLast == 3 {
		screen.DrawImage(c.ImgStatus5[(c.Count/10)%4], op)
	}

	// 右
	if moveLast == 4 {
		screen.DrawImage(c.ImgStatus6[(c.Count/10)%4], op)
	}

}
