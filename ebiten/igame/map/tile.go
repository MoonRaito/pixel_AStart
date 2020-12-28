package _map

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	"log"
	"pixel_AStart/ebiten/common"
	"pixel_AStart/resource/images"
	"strconv"
)

/*
地图块
*/
var (
	MapWidth, MapHeight int
)

type Tile struct {
	Eimage *ebiten.Image

	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
	X    int    `xml:"x,attr"`
	Y    int    `xml:"y,attr"`

	Property *property
}

type property struct {
	Name string
	// 防御
	Def int
	// 闪避
	Avo int
	// 移动力 move power  到达此处的移动力
	Mp int
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
	m, err := ReadFile("./resource/01/base01.tmx")
	if err != nil {
		log.Fatal(err)
	}

	MapWidth = m.Width * m.TileWidth
	MapHeight = m.Height * m.TileHeight

	// 基础图片
	//img, _, err := ebitenutil.NewImageFromFile(common.RealPath + "/resource/01/base01.png")
	img1, _, err := image.Decode(bytes.NewReader(images.Base_png))
	if err != nil {
		log.Fatal(err)
	}
	img := ebiten.NewImageFromImage(img1)

	// 加载地图中的 object 属性
	Worlds = World{}
	for _, og := range m.ObjectGroups {
		for _, o := range og.Objects {
			//image := img.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image)

			x := int(o.X)
			y := int(o.Y) - 16
			eimage := img.SubImage(image.Rect(x, y, int(o.X)+16, int(o.Y))).(*ebiten.Image)
			tile := &Tile{
				Eimage:   eimage,
				Name:     o.Name,
				Type:     o.Type,
				X:        x,
				Y:        y,
				Property: initProperty(o.Type),
			}

			Worlds.SetTile(tile, x/16, y/16)
		}
	}
}

func GetKey(x int, y int) string {
	return strconv.Itoa(x) + "_" + strconv.Itoa(y)
}

func Draw(screen *ebiten.Image) {
	for _, tiles := range Worlds {
		for _, tile := range tiles {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.X), float64(tile.Y)+float64(common.OffsetY))
			op.GeoM.Scale(common.Scale, common.Scale)
			screen.DrawImage(tile.Eimage, op)
		}
	}
}

// 砖块属性
/*
type0:不可抵达    Def:0   Avo:0
type1:城门    Def:3   Avo:20
type2:废墟    Def:0   Avo:0
type3:民家    Def:0   Avo:10
type4:村    Def:0   Avo:10


type11:平地    Def:0   Avo:0
type12:山     Def:1  Avo:30
type13:高山    Def:2   Avo:40
type14:森    Def:1   Avo:20
type15:沟壑    Def:0   Avo:40
type16:海    Def:0   Avo:10
*/
func initProperty(t string) *property {
	switch t {
	case "0":
		return &property{
			Name: "不可抵达",
			Def:  0,
			Avo:  0,
			Mp:   0,
		}
	case "1":
		return &property{
			Name: "城门",
			Def:  3,
			Avo:  20,
			Mp:   1,
		}
	case "2":
		return &property{
			Name: "废墟",
			Def:  0,
			Avo:  0,
			Mp:   1,
		}
	case "3":
		return &property{
			Name: "民家",
			Def:  0,
			Avo:  10,
			Mp:   1,
		}
	case "4":
		return &property{
			Name: "村",
			Def:  0,
			Avo:  10,
			Mp:   1,
		}

	case "11":
		return &property{
			Name: "平地",
			Def:  0,
			Avo:  0,
			Mp:   1,
		}
	case "12":
		return &property{
			Name: "山",
			Def:  1,
			Avo:  30,
			Mp:   2,
		}
	case "13":
		return &property{
			Name: "高山",
			Def:  2,
			Avo:  40,
			Mp:   3,
		}
	case "14":
		return &property{
			Name: "森",
			Def:  1,
			Avo:  20,
			Mp:   2,
		}
	case "15":
		return &property{
			Name: "沟壑",
			Def:  0,
			Avo:  0,
			Mp:   0,
		}
	case "16":
		return &property{
			Name: "海",
			Def:  0,
			Avo:  10,
			Mp:   0,
		}

	default:
		return nil
	}
}

// 世界地图
type World map[int]map[int]*Tile

var Worlds World

// SetTile sets a tile at the given coordinates in the world.
func (w World) SetTile(t *Tile, x, y int) {
	if w[x] == nil {
		w[x] = map[int]*Tile{}
	}
	w[x][y] = t
}

// Tile gets the tile at the given coordinates in the world.
func (w World) Tile(x, y int) *Tile {
	if w[x] == nil {
		return nil
	}
	return w[x][y]
}
