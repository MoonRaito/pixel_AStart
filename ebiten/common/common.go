package common

import "os"

var (
	// 偏移值
	OffsetX = 0
	OffsetY = 0

	// 真实路径
	RealPath = ""
)

const (
	// 屏幕大小
	ScreenWidth  = 16 * 30
	ScreenHeight = 16 * 20
	// 缩放倍数
	Scale = 2
)

func Init() {
	dir, _ := os.Getwd()
	RealPath = dir
}
