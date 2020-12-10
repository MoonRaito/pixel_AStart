package roy

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"log"
	"pixel_AStart/ebiten/common"
)

type Roy struct {
	image       *ebiten.Image
	Count, X, Y int
	dt          float64

	// 状态 1:活跃 2：待机 3：移动上 4：移动下 5:移动左 6：移动6 7：攻击
	Status int
	// 是否选中
	IsSelected bool
	// 缩放
	Scale float64

	// 待行动
	imgStatus1 []*ebiten.Image
	// 选中
	imgStatus2 []*ebiten.Image
}

func (c *Roy) Init(url string) {
	img, _, err := ebitenutil.NewImageFromFile(url)
	if err != nil {
		log.Fatal(err)
	}

	c.image = img
	c.X = 2 * 16
	c.Y = 10 * 16
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
}

func (c *Roy) Update(dt float64) {
	c.Count++
	if c.Count > 120 {
		c.Count = 0
	}

	// 1秒 60帧
	c.dt += dt
	if c.dt > 1.0 {
		c.dt = 0
	}

}

func (c *Roy) Draw(screen *ebiten.Image) {

	// 未选中时 并且 状态为 活跃
	if !c.IsSelected && c.Status == 1 {
		c.status1(screen)
	}
	// 选中时 并且 状态为 活跃
	if c.IsSelected && c.Status == 1 {
		c.status2(screen)
	}

}

func (c *Roy) status1(screen *ebiten.Image) {
	i := 0
	if c.dt < 0.45 {
		i = 0
	}

	if 0.45 <= c.dt && c.dt <= 0.50 {
		i = 1
	}

	if 0.50 < c.dt && c.dt < 0.95 {
		i = 2
	}

	if 0.95 <= c.dt {
		i = 1
	}

	//fmt.Println((c.Count/5)%4)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.X), float64(c.Y)+float64(common.OffsetY))
	op.GeoM.Scale(c.Scale, c.Scale)
	//screen.DrawImage(c.img_status1[(c.Count/5)%4], op)
	screen.DrawImage(c.imgStatus1[i], op)
}

func (c *Roy) status2(screen *ebiten.Image) {
	i := 0
	if c.dt < 0.45 {
		i = 0
	}

	if 0.45 <= c.dt && c.dt <= 0.50 {
		i = 1
	}

	if 0.50 < c.dt && c.dt < 0.95 {
		i = 2
	}

	if 0.95 <= c.dt {
		i = 1
	}

	//fmt.Println((c.Count/5)%4)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.X)-3, float64(c.Y)-5+float64(common.OffsetY))
	op.GeoM.Scale(c.Scale, c.Scale)
	//screen.DrawImage(c.imgStatus2[(c.Count/5)%3], op)
	screen.DrawImage(c.imgStatus2[i], op)
}
