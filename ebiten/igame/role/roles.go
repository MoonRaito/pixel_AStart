package role

import (
	"github.com/hajimehoshi/ebiten/v2"
	"pixel_AStart/ebiten/igame/cursor"
	role "pixel_AStart/ebiten/igame/sprite"
)

var roles map[string]*role.Sprite

func Init() {
	roles = make(map[string]*role.Sprite)
	roles["roy"] = Init_roy()
	roles["wolt"] = Init_wolt()
}

func Update(dt float64) {
	cursor.Icursor.IsSelected = false
	for key := range roles {
		roles[key].Update(dt)

		// 光标 是否选中 精灵 时 更改光标状态
		if cursor.Icursor.X == roles[key].X && cursor.Icursor.Y == roles[key].Y {
			cursor.Icursor.IsSelected = true
		}
	}
}
func Draw(screen *ebiten.Image) {
	for key := range roles {
		roles[key].Draw(screen)
	}
}
