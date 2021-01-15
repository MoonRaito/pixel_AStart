package util_image

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

func flip(src *ebiten.Image) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, 16, 16))
	for x := 0; x < dst.Bounds().Dx(); x++ {
		for y := 0; y < dst.Bounds().Dy(); y++ {
			dst.Set(x, y, src.At(src.Bounds().Dx()-x, y))
		}
	}
	return dst
}

func Flip(src *ebiten.Image) *ebiten.Image {
	return ebiten.NewImageFromImage(flip(src))
}
