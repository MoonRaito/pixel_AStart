package cursor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image"
	"log"
	"pixel_AStart/ebiten/common"
	_map "pixel_AStart/ebiten/igame/map"
)

/*
地图中的光标
*/
// 光标
var Icursor = &Cursor{}

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

func Cursor_Init() {
	Icursor = &Cursor{
		Count: 0,
		Scale: common.Scale,
	}

	img, _, err := ebitenutil.NewImageFromFile(common.RealPath + "/resource/images/cursor.png")
	if err != nil {
		log.Fatal(err)
	}

	Icursor.image = img
	Icursor.Count = 0

	Icursor.X = 2 * 16
	Icursor.Y = 18 * 16
	Icursor.screenX = 16 * 2
	Icursor.screenY = 16 * 7

	// 精灵动画
	Icursor.images = make([]*ebiten.Image, 2)
	Icursor.images[0] = img.SubImage(image.Rect(19, 10, 35, 26)).(*ebiten.Image)
	Icursor.images[1] = img.SubImage(image.Rect(43, 12, 59, 28)).(*ebiten.Image)
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
		if c.screenX-16 >= 0 {
			c.X -= 16
			c.screenX -= 16
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		if c.screenX+16 < common.ScreenWidth/common.Scale {
			c.X += 16
			c.screenX += 16
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		// 下一个坐标点 屏幕位置
		next := c.screenY - 16

		// 下一个坐标点 地图配置
		nextMap := c.Y - 16

		// 光标高度必须大于 场景的高度
		if next >= 0 {
			// 是否是在屏幕的上两格的位置
			if next < 16*2 {
				// 是否是整个地图的上两格
				if nextMap < 16*2 {
					// 移动光标
					c.screenY = next
				} else {
					// 地图的下两格 记录偏移量
					common.OffsetY += 16
				}
			} else {
				// 非 屏幕上两格 移动光标
				c.screenY = next
			}

			// 记录地图位置
			c.Y = nextMap
		}

	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		// 下一个坐标点 屏幕位置
		next := c.screenY + 16

		// 下一个坐标点 地图配置
		nextMap := c.Y + 16

		// 光标高度必须小于 场景的高度
		if next < common.ScreenHeight/common.Scale {
			// 是否是在屏幕的下两格的位置
			if next >= common.ScreenHeight/common.Scale-(16*2) {

				// 整个地图的下两格 不再偏移 只移动光标

				// 是否是整个地图的下两格
				if nextMap >= _map.MapHeight-(16*2) {
					// 移动光标
					c.screenY = next
				} else {
					// 地图的下两格 记录偏移量
					common.OffsetY -= 16
				}
			} else {
				// 非 屏幕下两格 移动光标
				c.screenY = next
			}

			// 记录地图位置
			c.Y = nextMap
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
	screen.DrawImage(c.images[i], op)
}
