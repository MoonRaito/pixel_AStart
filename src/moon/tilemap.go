package moon

type tilemap struct {
	X, Y int

	//F = G + H 中的每项值：
	//F（方块的和值）：左上角
	//G（从A点到方块的移动量）：左下角
	//H（从方块到B点的估算移动量): 右下角
	F, G, H int

	// 父坐标
	px, py int
}
