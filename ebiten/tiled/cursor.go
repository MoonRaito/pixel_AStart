package tiled

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	_ "image/png"
	"log"
)

type Cursor struct {
	image       *ebiten.Image
	images      []*ebiten.Image
	Count, X, Y int
	dt          float64
}

func (c *Cursor) Init(url string) {
	img, _, err := ebitenutil.NewImageFromFile(url)
	if err != nil {
		log.Fatal(err)
	}

	c.image = img
	c.X = 0
	c.Y = 0
	c.Count = 0

	c.images = make([]*ebiten.Image, 4)
	c.images[0] = img.SubImage(image.Rect(115, 13, 131, 29)).(*ebiten.Image)
	c.images[1] = img.SubImage(image.Rect(135, 13, 151, 29)).(*ebiten.Image)
	c.images[2] = img.SubImage(image.Rect(156, 13, 172, 29)).(*ebiten.Image)
	c.images[3] = img.SubImage(image.Rect(135, 13, 151, 29)).(*ebiten.Image)
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
}

func (c *Cursor) Draw(screen *ebiten.Image) {

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
		i = 3
	}

	//fmt.Println((c.Count/5)%4)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.X), float64(c.Y))
	op.GeoM.Scale(2, 2)
	//screen.DrawImage(c.images[(c.Count/5)%4], op)
	screen.DrawImage(c.images[i], op)
}
