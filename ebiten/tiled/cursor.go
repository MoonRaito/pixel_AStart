package tiled

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image"
	_ "image/png"
	"log"
	"pixel_AStart/ebiten/common"
)

type Cursor struct {
	image  *ebiten.Image
	images []*ebiten.Image
	// X,Y 游戏世界地图位置
	Count, X, Y int
	dt          float64
	// 缩放倍数
	Scale float64

	// 是否按下
	isPressed bool
	// 是否选中精灵
	IsSelected bool

	// 场景 x 也就是 屏幕的 x
	screenX int
	// 场景 y 也就是 屏幕的 y
	screenY int
}

func (c *Cursor) Init(url string) {
	img, _, err := ebitenutil.NewImageFromFile(url)
	if err != nil {
		log.Fatal(err)
	}

	c.image = img
	c.X = 16 * 5
	c.Y = 16 * 5
	c.Count = 0
	c.screenX = 16 * 5
	c.screenY = 16 * 5

	// 精灵动画
	c.images = make([]*ebiten.Image, 2)
	c.images[0] = img.SubImage(image.Rect(19, 10, 35, 26)).(*ebiten.Image)
	c.images[1] = img.SubImage(image.Rect(43, 12, 59, 28)).(*ebiten.Image)
}

func (c *Cursor) Update(dt float64) {
	c.Count++
	if c.Count > 120 {
		c.Count = 0
	}

	// 1秒 60帧
	c.dt += dt
	if c.dt > 1.0 {
		c.dt = 0
	}

	// 按键移动光标
	if inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		if c.X-16 >= 0 {
			c.X -= 16
			c.screenX -= 16
			// 设置偏移量
			//if c.X - 16 < 32 {
			//	common.OffsetX += 16
			//}

		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		newX := c.X + 16
		if newX <= 240 {
			c.X += 16

			// 设置偏移量
			//if c.X + 16 > 320-(16*2) {
			//	common.OffsetX -= 16
			//}
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		if c.Y-16 >= 0 {
			c.Y -= 16

			//if c.Y - 16 < 32 {
			//	common.OffsetY += 16
			//}
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		// 下一个坐标点
		next := c.screenY + 16

		// 光标高度必须小于 场景的高度
		if c.screenY <= common.ScreenHeight/common.Scale {
			// 是否是在屏幕的下两格的位置
			if next >= common.ScreenHeight/common.Scale-(16*2) {
				// 是否是整个地图的下两格
				if next >= MapHeight-(16*2) {
					// 移动光标
					c.screenY = next
				}
				// 地图的下两格 记录偏移量
				common.OffsetY -= 16
			} else {
				// 非 屏幕下两格 移动光标
				c.screenY = next
			}
		}
	}
}

func (c *Cursor) Draw(screen *ebiten.Image) {

	i := 0
	if c.dt <= 0.50 {
		i = 0
	}

	if 0.50 < c.dt && c.dt <= 1 {
		i = 1
	}

	// 选中时 只使用最大图片
	if c.IsSelected {
		i = 0
	}

	//fmt.Println((c.Count/5)%4)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.screenX), float64(c.screenY))
	op.GeoM.Scale(c.Scale, c.Scale)
	//screen.DrawImage(c.images[(c.Count/5)%4], op)
	screen.DrawImage(c.images[i], op)
}
