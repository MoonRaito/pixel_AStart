package tiled

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image"
	_ "image/png"
	"log"
)

type Cursor struct {
	image       *ebiten.Image
	images      []*ebiten.Image
	Count, X, Y int
	dt          float64
	// 缩放倍数
	Scale float64

	// 是否按下
	isPressed bool
	// 是否选中精灵
	IsSelected bool
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
		c.X -= 16
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		c.X += 16
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		c.Y -= 16
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) || inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		c.Y += 16
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
	op.GeoM.Translate(float64(c.X), float64(c.Y))
	op.GeoM.Scale(c.Scale, c.Scale)
	//screen.DrawImage(c.images[(c.Count/5)%4], op)
	screen.DrawImage(c.images[i], op)
}
