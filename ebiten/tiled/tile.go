package tiled

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"image/color"
	"log"
	"pixel_AStart/ebiten/common"
	"strconv"
)

/*
地图块
*/
var (
	MapWidth, MapHeight int
	Tiles               map[string]*Tile
)

type Tile struct {
	Eimage *ebiten.Image

	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
	X    int    `xml:"x,attr"`
	Y    int    `xml:"y,attr"`
}

// In returns true if (x, y) is in the sprite, and false otherwise.
func (s *Tile) In(x, y int) bool {
	// Check the actual color (alpha) value at the specified position
	// so that the result of In becomes natural to users.
	//
	// Note that this is not a good manner to use At for logic
	// since color from At might include some errors on some machines.
	// As this is not so important logic, it's ok to use it so far.
	return s.Eimage.At(x-s.X, y-s.Y).(color.RGBA).A > 0
}

// 初始化地图块
func Init() {
	// 加载地图
	m, err := ReadFileRealPath("/resource/01/base01.tmx")
	if err != nil {
		log.Fatal(err)
	}

	MapWidth = m.Width * m.TileWidth
	MapHeight = m.Height * m.TileHeight

	// 基础图片
	img, _, err := ebitenutil.NewImageFromFile(common.RealPath + "/resource/01/base.png")
	if err != nil {
		log.Fatal(err)
	}

	// 加载地图中的 object 属性
	Tiles = make(map[string]*Tile)
	for _, og := range m.ObjectGroups {
		for _, o := range og.Objects {
			//image := img.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image)

			x := int(o.X)
			y := int(o.Y) - 16
			eimage := img.SubImage(image.Rect(x, y, int(o.X)+16, int(o.Y))).(*ebiten.Image)
			tile := Tile{
				Eimage: eimage,
				Name:   o.Name,
				Type:   o.Type,
				X:      x,
				Y:      y,
			}
			Tiles[GetKey(x/16, y/16)] = &tile
		}
	}
}

func GetKey(x int, y int) string {
	return strconv.Itoa(x) + "_" + strconv.Itoa(y)
}

func Draw(screen *ebiten.Image) {
	for _, tile := range Tiles {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tile.X), float64(tile.Y)+float64(common.OffsetY))
		op.GeoM.Scale(common.Scale, common.Scale)
		//op.GeoM.Scale(
		//	math.Pow(1.01, float64(g.camera.ZoomFactor)),
		//	math.Pow(1.01, float64(g.camera.ZoomFactor)),
		//)
		screen.DrawImage(tile.Eimage, op)

		//g.camera.Render(tile.Eimage, screen)
	}
}
