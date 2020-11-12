package moon

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	"image/color"
)

type Iblock struct {
	X, Y                   int
	PX, PY                 int
	TxtF, TxtG, TxtH, TxtC *text.Text
	F, G, H                int

	Rect  pixel.Rect
	Vel   pixel.Vec
	Color color.Color

	// 块类型  -2结束位 -1:起始位  0：正常  1墙 2：障碍 3:路线
	Btype int
	// 权重等级
	level int

	Block *imdraw.IMDraw

	// 箭头 显示父节点xy坐标
	TxtX, TxtY, PXY *text.Text

	// 权重
	rank   float64
	index  int
	Open   bool
	Closed bool
	Cost   float64
	Parent *Iblock
}

func (ib *Iblock) UpdateIblock() {

	ib.Block = NewRectangle(ib.Color, ib.Rect)

	ib.TxtC = text.New(pixel.V(float64(ib.X)*100, float64(ib.Y)*100+80), text.NewAtlas(basicfont.Face7x13, text.ASCII))
	ib.TxtF = text.New(pixel.V(float64(ib.X)*100, float64(ib.Y)*100+50), text.NewAtlas(basicfont.Face7x13, text.ASCII))
	ib.TxtG = text.New(pixel.V(float64(ib.X)*100, float64(ib.Y)*100), text.NewAtlas(basicfont.Face7x13, text.ASCII))
	ib.TxtH = text.New(pixel.V(float64(ib.X)*100+50, float64(ib.Y)*100), text.NewAtlas(basicfont.Face7x13, text.ASCII))

	ib.PXY = text.New(pixel.V(float64(ib.X)*100+50, float64(ib.Y)*100+50), text.NewAtlas(basicfont.Face7x13, text.ASCII))

	ib.TxtX = text.New(pixel.V(float64(ib.X)*100, float64(ib.Y)*100+80), text.NewAtlas(basicfont.Face7x13, text.ASCII))
	ib.TxtY = text.New(pixel.V(float64(ib.X)*100+70, float64(ib.Y)*100+80), text.NewAtlas(basicfont.Face7x13, text.ASCII))

}

//首字母小写不能被其他包调用
func New(Color color.Color) *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = Color
	imd.Push(pixel.V(100, 100))
	imd.Color = Color
	imd.Push(pixel.V(200, 100))
	imd.Color = Color
	imd.Push(pixel.V(200, 200))
	imd.Color = Color
	imd.Push(pixel.V(100, 200))
	imd.Polygon(0)

	return imd
}

//首字母小写不能被其他包调用
func Update(imd *imdraw.IMDraw, Color color.Color) *imdraw.IMDraw {
	imd.Color = Color
	imd.Push(pixel.V(100, 100))
	imd.Color = Color
	imd.Push(pixel.V(200, 100))
	imd.Color = Color
	imd.Push(pixel.V(200, 200))
	imd.Color = Color
	imd.Push(pixel.V(100, 200))
	imd.Polygon(0)

	return imd
}
func NewSquare(Color color.Color, point pixel.Vec, size float64) *imdraw.IMDraw {
	return UpdateSquare(imdraw.New(nil), Color, point, size)
}

//首字母小写不能被其他包调用
func UpdateSquare(imd *imdraw.IMDraw, Color color.Color, point pixel.Vec, size float64) *imdraw.IMDraw {
	imd.Color = Color
	imd.Push(point)
	imd.Color = Color
	imd.Push(pixel.V(point.X+size, point.Y))
	imd.Color = Color
	imd.Push(pixel.V(point.X+size, point.Y+size))
	imd.Color = Color
	imd.Push(pixel.V(point.X, point.Y+size))
	imd.Polygon(0)

	return imd
}

// 长方形
func NewRectangle(Color color.Color, r pixel.Rect) *imdraw.IMDraw {
	return UpdateRectangle(imdraw.New(nil), Color, r)
}
func UpdateRectangle(imd *imdraw.IMDraw, Color color.Color, r pixel.Rect) *imdraw.IMDraw {
	imd.Color = Color
	imd.Push(r.Min, r.Max)
	imd.Rectangle(0)
	return imd
}
func NewRectangle2(Color color.Color, r pixel.Rect, b float64) *imdraw.IMDraw {
	return UpdateRectangle2(imdraw.New(nil), Color, r, b)
}
func UpdateRectangle2(imd *imdraw.IMDraw, Color color.Color, r pixel.Rect, b float64) *imdraw.IMDraw {
	imd.Color = Color
	imd.Push(r.Min, r.Max)
	imd.Color = colornames.Skyblue
	imd.Rectangle(b)
	return imd
}
