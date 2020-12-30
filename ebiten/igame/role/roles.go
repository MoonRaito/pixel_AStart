package role

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"pixel_AStart/ebiten/common"
	"pixel_AStart/ebiten/igame/cursor"
	"pixel_AStart/ebiten/igame/path"
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

		// 空格是否被选中 光标选中 时 寻路
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			// 状态 2 寻路
			if roles[key].Status == 2 {
				path.IPath.Find(roles[key].X/common.TileSize, roles[key].Y/common.TileSize)
				// 选中
				roles[key].Status = 3
			}
		}

		// 临时 用c 清图
		if inpututil.IsKeyJustPressed(ebiten.KeyC) {
			roles[key].Status = 1
		}
	}
	// 选中角色后的路径
	path.MovePath(cursor.Icursor.X/common.TileSize, cursor.Icursor.Y/common.TileSize)

	// 临时 用c 清图
	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		path.IPath.Clear()
	}

}
func Draw(screen *ebiten.Image) {
	for key := range roles {
		roles[key].Draw(screen)
	}
}
