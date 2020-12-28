package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"strconv"
	"time"
)

const (
	screenWidth  = 1080
	screenHeight = 2280
)

var (
	ebitenImage *ebiten.Image
	luobo       *ebiten.Image
	bu1         *ebiten.Image
	bu2         *ebiten.Image
	bu3         *ebiten.Image
	draw        *ebiten.Image

	card      [24]int
	card_Rand [4]int

	card_ int

	last   time.Time
	dt     float64
	dtRand float64

	isTouch bool
	isRand  int
)

func init() {
	img, _, err := ebitenutil.NewImageFromFile("./tuzi.png")
	if err != nil {
		log.Fatal(err)
	}
	ebitenImage = img

	// w:110    h:166
	luobo = ebitenImage.SubImage(image.Rect(15, 410, 125, 566)).(*ebiten.Image)
	bu1 = ebitenImage.SubImage(image.Rect(420, 410, 530, 566)).(*ebiten.Image)
	bu2 = ebitenImage.SubImage(image.Rect(285, 410, 395, 566)).(*ebiten.Image)
	bu3 = ebitenImage.SubImage(image.Rect(155, 410, 265, 566)).(*ebiten.Image)

	// 一步
	for i := 0; i < 24; i++ {
		if i < 12 {
			card[i] = 1
		}
		if i >= 12 && i < 16 {
			card[i] = 2
		}
		if i >= 16 && i < 18 {
			card[i] = 3
		}
		if i >= 18 {
			card[i] = 4
		}
	}

	card_Rand[0] = 1
	card_Rand[1] = 2
	card_Rand[2] = 3
	card_Rand[3] = 4

	// 初始画面
	draw = luobo
}

type Game struct{}

func (g *Game) Update() error {
	dt = time.Since(last).Seconds()
	last = time.Now()

	if inpututil.JustPressedTouchIDs() != nil {
		card_ = card[rand.Intn(24)]

		fmt.Println("I'm touch:" + strconv.Itoa(card_))
		if card_ == 1 {
			draw = bu1
		}
		if card_ == 2 {
			draw = bu2
		}
		if card_ == 3 {
			draw = bu3
		}
		if card_ == 4 {
			draw = luobo
		}

		isTouch = true
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		card_ = card[rand.Intn(24)]

		fmt.Println("I'm touch:" + strconv.Itoa(card_))
		if card_ == 1 {
			draw = bu1
		}
		if card_ == 2 {
			draw = bu2
		}
		if card_ == 3 {
			draw = bu3
		}
		if card_ == 4 {
			draw = luobo
		}

		isTouch = true
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.NRGBA{0xff, 0xff, 0xff, 0xff})

	if isTouch && dtRand < 1 {
		dtRand += dt

		// Draw the image with 'Source Alpha' composite mode (default).
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, 30)
		op.GeoM.Scale(10, 10)
		i := card_Rand[rand.Intn(4)]
		if i == 1 {
			screen.DrawImage(bu1, op)
		}
		if i == 2 {
			screen.DrawImage(bu2, op)
		}
		if i == 3 {
			screen.DrawImage(bu3, op)
		}
		if i == 4 {
			screen.DrawImage(luobo, op)
		}

		return
	} else {
		dtRand = 0
		isTouch = false
	}

	// Draw the image with 'Source Alpha' composite mode (default).
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 30)
	op.GeoM.Scale(10, 10)
	screen.DrawImage(draw, op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {

	ebiten.SetWindowSize(screenWidth/2, screenHeight)
	ebiten.SetWindowTitle("happy igame")
	// Call ebiten.RunGame to start your igame loop.
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
