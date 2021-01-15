package role

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"log"
	"pixel_AStart/ebiten/common"
	role "pixel_AStart/ebiten/igame/sprite"
)

func Init_roy() *role.Sprite {
	img, _, err := ebitenutil.NewImageFromFile(common.RealPath + "/resource/02/Map_Lord_Roy.png")
	if err != nil {
		log.Fatal(err)
	}

	sprite := role.Sprite{}

	sprite.Name = "roy"
	sprite.Image = img

	// 初始位置
	sprite.X = 2 * 16
	sprite.Y = 18 * 16

	sprite.Count = 0
	sprite.Scale = common.Scale
	sprite.Status = 1
	sprite.MoveSpeed = 0.1
	sprite.AttackRange = 1

	// 测试图片翻转
	//sprite.ImgStatus1 = make([]*ebiten.Image, 3)
	//a := util_image.Flip(img.SubImage(image.Rect(115, 13, 131, 29)).(*ebiten.Image))

	sprite.ImgStatus1 = make([]*ebiten.Image, 3)
	sprite.ImgStatus1[0] = img.SubImage(image.Rect(115, 13, 131, 29)).(*ebiten.Image)
	sprite.ImgStatus1[1] = img.SubImage(image.Rect(135, 13, 151, 29)).(*ebiten.Image)
	sprite.ImgStatus1[2] = img.SubImage(image.Rect(156, 13, 172, 29)).(*ebiten.Image)

	// 最大宽度 26
	// 最大高度 21
	// 中心分割 右15 左11
	sprite.ImgStatus2 = make([]*ebiten.Image, 3)
	sprite.ImgStatus2[0] = img.SubImage(image.Rect(22, 8, 48, 29)).(*ebiten.Image)
	sprite.ImgStatus2[1] = img.SubImage(image.Rect(52, 8, 78, 29)).(*ebiten.Image)
	sprite.ImgStatus2[2] = img.SubImage(image.Rect(79, 8, 105, 29)).(*ebiten.Image)

	// 宽 20
	// 高 20
	sprite.ImgStatus3 = make([]*ebiten.Image, 4)
	sprite.ImgStatus3[0] = img.SubImage(image.Rect(18, 39, 38, 59)).(*ebiten.Image)
	sprite.ImgStatus3[1] = img.SubImage(image.Rect(41, 39, 61, 59)).(*ebiten.Image)
	sprite.ImgStatus3[2] = img.SubImage(image.Rect(66, 39, 86, 59)).(*ebiten.Image)
	sprite.ImgStatus3[3] = img.SubImage(image.Rect(91, 39, 111, 59)).(*ebiten.Image)

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
