package tiled

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	_ "image/png"
	"log"
)

type Cursor struct {
	image  *ebiten.Image
	images []*ebiten.Image
	Count  int8
}

func (c *Cursor) Init(url string) {
	img, _, err := ebitenutil.NewImageFromFile(url)
	if err != nil {
		log.Fatal(err)
	}

	c.image = img

	c.images[0] = img.SubImage(image.Rect(115, 666, 131, 682)).(*ebiten.Image)
	c.images[1] = img.SubImage(image.Rect(135, 666, 151, 682)).(*ebiten.Image)
	c.images[2] = img.SubImage(image.Rect(156, 666, 172, 682)).(*ebiten.Image)
	c.images[3] = img.SubImage(image.Rect(135, 666, 151, 682)).(*ebiten.Image)
}
