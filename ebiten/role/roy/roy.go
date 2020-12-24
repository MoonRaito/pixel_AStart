package roy

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"log"
	"pixel_AStart/ebiten/common"
	"pixel_AStart/ebiten/path"
	"pixel_AStart/ebiten/queue"
)

type Roy struct {
	image       *ebiten.Image
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
}

func (c *Roy) PathNeighborCost(to path.IPath) float64 {
	//toT := to.(*Roy)
	//return KindCosts[toT.Kind]
	return 0
}

func (c *Roy) Init(url string) {
	img, _, err := ebitenutil.NewImageFromFile(url)
	if err != nil {
		log.Fatal(err)
	}

	c.image = img
	c.X = 2 * 16
	c.Y = 18 * 16
	//c.X = 5 * 16
	//c.Y = 5 * 16
	c.Count = 0

	c.imgStatus1 = make([]*ebiten.Image, 3)
	c.imgStatus1[0] = img.SubImage(image.Rect(115, 13, 131, 29)).(*ebiten.Image)
	c.imgStatus1[1] = img.SubImage(image.Rect(135, 13, 151, 29)).(*ebiten.Image)
	c.imgStatus1[2] = img.SubImage(image.Rect(156, 13, 172, 29)).(*ebiten.Image)

	// 最大宽度 26
	// 最大高度 21
	// 中心分割 右15 左11
	c.imgStatus2 = make([]*ebiten.Image, 3)
	c.imgStatus2[0] = img.SubImage(image.Rect(22, 8, 48, 29)).(*ebiten.Image)
	c.imgStatus2[1] = img.SubImage(image.Rect(52, 8, 78, 29)).(*ebiten.Image)
	c.imgStatus2[2] = img.SubImage(image.Rect(79, 8, 105, 29)).(*ebiten.Image)

	// 宽 20
	// 高 20
	c.imgStatus3 = make([]*ebiten.Image, 4)
	c.imgStatus3[0] = img.SubImage(image.Rect(18, 39, 38, 59)).(*ebiten.Image)
	c.imgStatus3[1] = img.SubImage(image.Rect(41, 39, 61, 59)).(*ebiten.Image)
	c.imgStatus3[2] = img.SubImage(image.Rect(66, 39, 86, 59)).(*ebiten.Image)
	c.imgStatus3[3] = img.SubImage(image.Rect(91, 39, 111, 59)).(*ebiten.Image)

	// 移动上
	c.imgStatus4 = make([]*ebiten.Image, 4)
	c.imgStatus4[0] = img.SubImage(image.Rect(126, 39, 146, 59)).(*ebiten.Image)
	c.imgStatus4[1] = img.SubImage(image.Rect(149, 39, 169, 59)).(*ebiten.Image)
	c.imgStatus4[2] = img.SubImage(image.Rect(173, 39, 193, 59)).(*ebiten.Image)
	c.imgStatus4[3] = img.SubImage(image.Rect(196, 39, 216, 59)).(*ebiten.Image)
	// 移动左 缺少素材临时用右代替
	c.imgStatus5 = make([]*ebiten.Image, 4)
	c.imgStatus5[0] = img.SubImage(image.Rect(229, 39, 249, 59)).(*ebiten.Image)
	c.imgStatus5[1] = img.SubImage(image.Rect(256, 39, 276, 59)).(*ebiten.Image)
	c.imgStatus5[2] = img.SubImage(image.Rect(286, 39, 306, 59)).(*ebiten.Image)
	c.imgStatus5[3] = img.SubImage(image.Rect(314, 39, 334, 59)).(*ebiten.Image)
	// 移动右
	c.imgStatus6 = make([]*ebiten.Image, 4)
	c.imgStatus6[0] = img.SubImage(image.Rect(229, 39, 249, 59)).(*ebiten.Image)
	c.imgStatus6[1] = img.SubImage(image.Rect(256, 39, 276, 59)).(*ebiten.Image)
	c.imgStatus6[2] = img.SubImage(image.Rect(286, 39, 306, 59)).(*ebiten.Image)
	c.imgStatus6[3] = img.SubImage(image.Rect(314, 39, 334, 59)).(*ebiten.Image)
}

func (c *Roy) Update(dt float64) {
	c.Count++
	if c.Count > 120 {
		c.Count = 0
	}

	// 1秒 60帧
	c.dt += dt
	// 每个状态的帧数不一致
	if c.Status == 1 {
		if c.dt > 1.0 {
			c.dt = 0
		}
	}
	if c.Status == 2 {
		if c.dt > 1.0 {
			c.dt = 0
		}
	}

}

func (c *Roy) Draw(screen *ebiten.Image) {

	// 未选中时 并且 状态为 活跃
	if c.Status == 1 {
		c.status1(screen)
	}
	// 光标 指向时
	if c.Status == 2 {
		c.status2(screen)
	}
	// 光标 选中时
	if c.Status == 3 {
		c.status3(screen)
	}

	// 开始移动角色
	if c.Status == 4 {
		c.UpdateMove(screen)
	}

}

func (c *Roy) status1(screen *ebiten.Image) {
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
	screen.DrawImage(c.imgStatus1[i], op)
}

func (c *Roy) status2(screen *ebiten.Image) {
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
	screen.DrawImage(c.imgStatus2[i], op)
}

func (c *Roy) status3(screen *ebiten.Image) {
	//i := 0
	//s := 0.25
	//sp := 0.25
	//if c.dt <= s {
	//	i = 0
	//} else if s < c.dt && c.dt <= s+sp {
	//	i = 1
	//} else if s+sp < c.dt && c.dt <= s+(sp*2) {
	//	i = 2
	//} else if s+(sp*3) <= c.dt {
	//	i = 3
	//}

	//fmt.Println((c.Count/5)%4)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.X)-1, float64(c.Y)-5+float64(common.OffsetY))
	op.GeoM.Scale(c.Scale, c.Scale)
	screen.DrawImage(c.imgStatus3[(c.Count/10)%4], op)
	//screen.DrawImage(c.imgStatus3[i], op)
}

// 移动
var Moving = queue.NewQueue()

func (c *Roy) Moving() {
	d := 0
	for i := path.MovepathList.Front(); i != nil; i = i.Next() {
		if d == 0 {
			d++
			continue
		}
		Moving.EnQueue(i.Value.(*path.Path))
	}
}
func (c *Roy) MoveTo() bool {
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

func (c *Roy) UpdateMove(screen *ebiten.Image) {

	// 到达指定节点后 开始下一个
	if int(c.MoveX) == int(c.MoveEndX) && int(c.MoveY) == int(c.MoveEndY) {
		// 到达终点
		if !c.MoveTo() {
			c.Status = 3

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
		screen.DrawImage(c.imgStatus3[(c.Count/10)%4], op)
	}

	// 上
	if int(c.MoveX) == int(c.MoveEndX) && int(c.MoveY) > int(c.MoveEndY) {
		c.MoveY = c.MoveY - (c.MoveSpeed * c.dt)
		screen.DrawImage(c.imgStatus4[(c.Count/10)%4], op)
	}

	// 左
	if int(c.MoveX) > int(c.MoveEndX) && int(c.MoveY) == int(c.MoveEndY) {
		c.MoveX = c.MoveX - (c.MoveSpeed * c.dt)
		screen.DrawImage(c.imgStatus5[(c.Count/10)%4], op)
	}

	// 右
	if int(c.MoveX) < int(c.MoveEndX) && int(c.MoveY) == int(c.MoveEndY) {
		c.MoveX = c.MoveX + (c.MoveSpeed * c.dt)
		screen.DrawImage(c.imgStatus6[(c.Count/10)%4], op)
	}

}

//func (c *Roy) MoveTo(screen *ebiten.Image) {
//	oldX := c.X
//	oldY := c.Y
//
//	//movingX := float64(c.X) + (c.MoveSpeed * c.dt)
//	//movingY := float64(c.Y) + (c.MoveSpeed * c.dt)
//	for i := path.MovepathList.Front(); i != nil; i = i.Next() {
//		X := i.Value.(*path.Path).X
//		Y := i.Value.(*path.Path).Y
//
//		// 上
//		if X == oldX && Y<oldY {
//
//		}
//
//		// 下
//		if X == oldX && Y>oldY {
//
//		}
//
//		// 左
//		if X > oldX && Y==oldY {
//
//		}
//
//		// 右
//		if X < oldX && Y==oldY {
//
//		}
//
//
//	}
//
//
//	//op := &ebiten.DrawImageOptions{}
//	//op.GeoM.Translate(movingX-1, movingY-5+float64(common.OffsetY))
//	//op.GeoM.Scale(c.Scale, c.Scale)
//	//screen.DrawImage(c.imgStatus3[(c.Count/10)%4], op)
//}
