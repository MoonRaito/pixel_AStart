package role

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"log"
	"pixel_AStart/ebiten/common"
	role "pixel_AStart/ebiten/igame/sprite"
)

func Init_roy() {
	img, _, err := ebitenutil.NewImageFromFile(common.RealPath + "/resource/02/Map_Lord_Roy.png")
	if err != nil {
		log.Fatal(err)
	}

	roy := role.Sprite{}

	roy.Name = "roy"
	roy.Image = img

	// 初始位置
	roy.X = 2 * 16
	roy.Y = 18 * 16
	//roy.X = 5 * 16
	//roy.Y = 5 * 16
	roy.Count = 0

	roy.ImgStatus1 = make([]*ebiten.Image, 3)
	roy.ImgStatus1[0] = img.SubImage(image.Rect(115, 13, 131, 29)).(*ebiten.Image)
	roy.ImgStatus1[1] = img.SubImage(image.Rect(135, 13, 151, 29)).(*ebiten.Image)
	roy.ImgStatus1[2] = img.SubImage(image.Rect(156, 13, 172, 29)).(*ebiten.Image)

	// 最大宽度 26
	// 最大高度 21
	// 中心分割 右15 左11
	roy.ImgStatus2 = make([]*ebiten.Image, 3)
	roy.ImgStatus2[0] = img.SubImage(image.Rect(22, 8, 48, 29)).(*ebiten.Image)
	roy.ImgStatus2[1] = img.SubImage(image.Rect(52, 8, 78, 29)).(*ebiten.Image)
	roy.ImgStatus2[2] = img.SubImage(image.Rect(79, 8, 105, 29)).(*ebiten.Image)

	// 宽 20
	// 高 20
	roy.ImgStatus3 = make([]*ebiten.Image, 4)
	roy.ImgStatus3[0] = img.SubImage(image.Rect(18, 39, 38, 59)).(*ebiten.Image)
	roy.ImgStatus3[1] = img.SubImage(image.Rect(41, 39, 61, 59)).(*ebiten.Image)
	roy.ImgStatus3[2] = img.SubImage(image.Rect(66, 39, 86, 59)).(*ebiten.Image)
	roy.ImgStatus3[3] = img.SubImage(image.Rect(91, 39, 111, 59)).(*ebiten.Image)

	// 移动上
	roy.ImgStatus4 = make([]*ebiten.Image, 4)
	roy.ImgStatus4[0] = img.SubImage(image.Rect(126, 39, 146, 59)).(*ebiten.Image)
	roy.ImgStatus4[1] = img.SubImage(image.Rect(149, 39, 169, 59)).(*ebiten.Image)
	roy.ImgStatus4[2] = img.SubImage(image.Rect(173, 39, 193, 59)).(*ebiten.Image)
	roy.ImgStatus4[3] = img.SubImage(image.Rect(196, 39, 216, 59)).(*ebiten.Image)
	// 移动左 缺少素材临时用右代替
	roy.ImgStatus5 = make([]*ebiten.Image, 4)
	roy.ImgStatus5[0] = img.SubImage(image.Rect(229, 39, 249, 59)).(*ebiten.Image)
	roy.ImgStatus5[1] = img.SubImage(image.Rect(256, 39, 276, 59)).(*ebiten.Image)
	roy.ImgStatus5[2] = img.SubImage(image.Rect(286, 39, 306, 59)).(*ebiten.Image)
	roy.ImgStatus5[3] = img.SubImage(image.Rect(314, 39, 334, 59)).(*ebiten.Image)
	// 移动右
	roy.ImgStatus6 = make([]*ebiten.Image, 4)
	roy.ImgStatus6[0] = img.SubImage(image.Rect(229, 39, 249, 59)).(*ebiten.Image)
	roy.ImgStatus6[1] = img.SubImage(image.Rect(256, 39, 276, 59)).(*ebiten.Image)
	roy.ImgStatus6[2] = img.SubImage(image.Rect(286, 39, 306, 59)).(*ebiten.Image)
	roy.ImgStatus6[3] = img.SubImage(image.Rect(314, 39, 334, 59)).(*ebiten.Image)
}
