package sprite

import "github.com/faiface/pixel"

type msprite struct {
	pixel.Sprite
	X, Y int8
}

func (ms *msprite) IsMe(v pixel.Vec) bool {
	return false
}
