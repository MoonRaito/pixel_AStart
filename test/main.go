package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"pixel_AStart/ebiten/common"
)

func main() {
	common.Init()
	//dir := common.RealPath + "/resource/01/base.png"
	dir := "C:\\Users\\39495\\Desktop\\work\\go\\idea_go_test_git\\resource\\01\\base.png"
	fmt.Println(dir)
	img, _, err := ebitenutil.NewImageFromFile(dir)
	//img1, _, err := image.Decode(bytes.NewReader(images.Base_png))
	if err != nil {
		log.Fatal(err)
	}
	//img := ebiten.NewImageFromImage(img1)

	fmt.Println(img)
}
