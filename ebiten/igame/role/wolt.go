package role

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"log"
	"pixel_AStart/ebiten/common"
	role "pixel_AStart/ebiten/igame/sprite"
)

func Init_wolt() *role.Sprite {
	img, _, err := ebitenutil.NewImageFromFile(common.RealPath + "/resource/02/Wolt.png")
	if err != nil {
		log.Fatal(err)
	}

	sprite := role.Sprite{}

	sprite.Name = "wolt"
	sprite.Image = img

	// 初始位置
	sprite.X = 2 * 16
	sprite.Y = 16 * 16

	sprite.Count = 0
	sprite.Scale = common.Scale
	sprite.Status = 1
	sprite.MoveSpeed = 0.1

	// 16 像素
	sprite.ImgStatus1 = make([]*ebiten.Image, 3)
	sprite.ImgStatus1[0] = img.SubImage(image.Rect(8, 255, 24, 271)).(*ebiten.Image)
	sprite.ImgStatus1[1] = img.SubImage(image.Rect(32, 255, 48, 271)).(*ebiten.Image)
	sprite.ImgStatus1[2] = img.SubImage(image.Rect(60, 255, 76, 271)).(*ebiten.Image)

	// 最大宽度 24
	// 最大高度 24
	sprite.ImgStatus2 = make([]*ebiten.Image, 3)
	sprite.ImgStatus2[0] = img.SubImage(image.Rect(98, 250, 122, 274)).(*ebiten.Image)
	sprite.ImgStatus2[1] = img.SubImage(image.Rect(123, 250, 147, 274)).(*ebiten.Image)
	sprite.ImgStatus2[2] = img.SubImage(image.Rect(149, 251, 173, 275)).(*ebiten.Image)

	// 宽 20
	// 高 20
	sprite.ImgStatus3 = make([]*ebiten.Image, 4)
	sprite.ImgStatus3[0] = img.SubImage(image.Rect(7, 295, 38, 59)).(*ebiten.Image)
	sprite.ImgStatus3[1] = img.SubImage(image.Rect(41, 295, 61, 59)).(*ebiten.Image)
	sprite.ImgStatus3[2] = img.SubImage(image.Rect(66, 295, 86, 59)).(*ebiten.Image)
	sprite.ImgStatus3[3] = img.SubImage(image.Rect(91, 295, 111, 59)).(*ebiten.Image)

	// 移动上
	sprite.ImgStatus4 = make([]*ebiten.Image, 4)
	sprite.ImgStatus4[0] = img.SubImage(image.Rect(126, 39, 146, 59)).(*ebiten.Image)
	sprite.ImgStatus4[1] = img.SubImage(image.Rect(149, 39, 169, 59)).(*ebiten.Image)
	sprite.ImgStatus4[2] = img.SubImage(image.Rect(173, 39, 193, 59)).(*ebiten.Image)
	sprite.ImgStatus4[3] = img.SubImage(image.Rect(196, 39, 216, 59)).(*ebiten.Image)
	// 移动左 缺少素材临时用右代替
	sprite.ImgStatus5 = make([]*ebiten.Image, 4)
	sprite.ImgStatus5[0] = img.SubImage(image.Rect(229, 39, 249, 59)).(*ebiten.Image)
	sprite.ImgStatus5[1] = img.SubImage(image.Rect(256, 39, 276, 59)).(*ebiten.Image)
	sprite.ImgStatus5[2] = img.SubImage(image.Rect(286, 39, 306, 59)).(*ebiten.Image)
	sprite.ImgStatus5[3] = img.SubImage(image.Rect(314, 39, 334, 59)).(*ebiten.Image)
	// 移动右
	sprite.ImgStatus6 = make([]*ebiten.Image, 4)
	sprite.ImgStatus6[0] = img.SubImage(image.Rect(229, 39, 249, 59)).(*ebiten.Image)
	sprite.ImgStatus6[1] = img.SubImage(image.Rect(256, 39, 276, 59)).(*ebiten.Image)
	sprite.ImgStatus6[2] = img.SubImage(image.Rect(286, 39, 306, 59)).(*ebiten.Image)
	sprite.ImgStatus6[3] = img.SubImage(image.Rect(314, 39, 334, 59)).(*ebiten.Image)

	return &sprite
}
