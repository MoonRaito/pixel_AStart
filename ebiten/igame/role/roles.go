package role

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"pixel_AStart/ebiten/common"
	"pixel_AStart/ebiten/igame/cursor"
	"pixel_AStart/ebiten/igame/path"
	role "pixel_AStart/ebiten/igame/sprite"
	"pixel_AStart/ebiten/tiled"
)

var roles map[string]*role.Sprite
var cursorSelect *role.Sprite

func Init() {
	roles = make(map[string]*role.Sprite)
	roles["roy"] = Init_roy()
	roles["wolt"] = Init_wolt()
}

func Update(dt float64) {
	cursor.Icursor.IsSelected = false
	for key := range roles {
		roles[key].Update(dt)

		// 记录每个角色的 坐标
		path.Roles_XY[tiled.GetKey(roles[key].X/16, roles[key].Y/16)] = roles[key].Name

		//// 空格是否被选中 光标选中 时 寻路
		//if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		//
		//	if roles[key].Status == 3 {
		//		// 光标选中角色时  显示物品  等操作
		//		if cursor.Icursor.X == roles[key].X && cursor.Icursor.Y == roles[key].Y {
		//
		//		} else {
		//			// 光标 是否 是在路径范围内
		//			if path.In(cursor.Icursor.X/common.TileSize, cursor.Icursor.Y/common.TileSize) {
		//
		//				// 角色被指向时
		//				if roles[key].Status == 2 {
		//				}
		//
		//				// 如果 已被选中 ，那么开始移动角色
		//				if roles[key].Status == 3 {
		//					roles[key].MoveX = float64(roles[key].X)
		//					roles[key].MoveY = float64(roles[key].Y)
		//					roles[key].Moving()
		//					roles[key].MoveTo()
		//					roles[key].Status = 4
		//				}
		//			} else {
		//				// 声音特效
		//			}
		//		}
		//
		//		// 待机
		//		//sroy.Status = 8
		//	}
		//
		//	// 状态 2 寻路
		//	if roles[key].Status == 2 {
		//		path.IPath.Find(roles[key].X/common.TileSize, roles[key].Y/common.TileSize,roles[key].AttackRange)
		//		// 选中
		//		roles[key].Status = 3
		//	}
		//}

	}

	// 光标 是否选中 精灵 时 更改光标状态
	if name, ok := path.Roles_XY[tiled.GetKey(cursor.Icursor.X/common.TileSize, cursor.Icursor.Y/common.TileSize)]; ok {
		cursor.Icursor.IsSelected = true

		// 光标指向精灵后  精灵动画为2
		sprite := roles[name]

		// 被选中时按下空格 寻路
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			// 状态 2 寻路
			if sprite.Status == 2 {
				path.IPath.Find(sprite.X/common.TileSize, sprite.Y/common.TileSize, sprite.AttackRange)
				cursorSelect = sprite
				// 选中
				sprite.Status = 3
			}
		}

	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if cursorSelect != nil {

			if cursorSelect.Status == 3 {
				// 光标选中角色时  显示物品  等操作
				if cursor.Icursor.X == cursorSelect.X && cursor.Icursor.Y == cursorSelect.Y {

				} else {
					// 光标 是否 是在路径范围内
					if path.In(cursor.Icursor.X/common.TileSize, cursor.Icursor.Y/common.TileSize) {

						// 角色被指向时
						if cursorSelect.Status == 2 {
						}

						// 如果 已被选中 ，那么开始移动角色
						if cursorSelect.Status == 3 {
							cursorSelect.MoveStartX = cursorSelect.X
							cursorSelect.MoveStartY = cursorSelect.Y
							cursorSelect.MoveX = float64(cursorSelect.X)
							cursorSelect.MoveY = float64(cursorSelect.Y)
							cursorSelect.Moving()
							cursorSelect.MoveTo()
							// 移动过程
							cursorSelect.Status = 10

							// 移动角色时隐藏路径
							path.IsShow = false
						}
					} else {
						// 声音特效
					}
				}

				// 待机
				//sroy.Status = 8
			}
		}

	}

	// 模拟按钮B
	if inpututil.IsKeyJustReleased(ebiten.KeyB) {
		if cursorSelect != nil {
			// 移动角色时隐藏路径
			path.IsShow = true

			// back 时 修改精灵状态
			if cursorSelect.Status != 1 {
				cursorSelect.Status -= 1
				fmt.Println("test")
			}
			cursorSelect.X = cursorSelect.MoveStartX
			cursorSelect.Y = cursorSelect.MoveStartY
		}
	}

	// 选中角色后的移动光标的路径
	path.MovePath(cursor.Icursor.X/common.TileSize, cursor.Icursor.Y/common.TileSize)

	// 临时 用c 清图
	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		// 初始化状态
		for key := range roles {
			roles[key].Status = 1
		}
		// 初始化路径
		path.IPath.Clear()
	}

}
func Draw(screen *ebiten.Image) {
	for key := range roles {
		roles[key].Draw(screen)
	}
}
