package tiled

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

/*
地图块
*/

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
