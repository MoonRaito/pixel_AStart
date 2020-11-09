package moon

import (
	"fmt"
	"github.com/faiface/pixel"
	"image/color"
	"strconv"
	"strings"
)

func GetKey(x int, y int) string {
	return strconv.Itoa(x) + "_" + strconv.Itoa(y)
}

// 初始化 start end
var start = &Iblock{}
var end = &Iblock{}

func InitStart(s *Iblock) {
	start = s
}
func InitEnd(e *Iblock) {
	end = e
}

//
var open = make(map[string]*Iblock)
var close = make(map[string]*Iblock)
var blockSize = 0

func FindPath(walls map[string]*Iblock, maps [][]string) map[string]*Iblock {

	// 初始化块大小
	blockSize = 100

	// 起始位置
	start := &Iblock{}
	//end := &Iblock{}
	for _, b := range walls {
		if b.Btype == -1 {
			start = b
		}
		if b.Btype == -2 {
			end = b
		}
	}
	close[GetKey(start.X, start.Y)] = start

	// 测试
	left := getBlock(start, start.X-1, start.Y)
	left.UpdateIblock()
	walls[strconv.Itoa(left.X)+"_"+strconv.Itoa(left.Y)] = left

	fmt.Println(len(walls))

	return nil
}

func getBlock(p *Iblock, x int, y int) *Iblock {

	g1 := abs(x - start.X)
	g2 := abs(y - start.Y)

	h1 := abs(x - end.X)
	h2 := abs(y - end.Y)

	g := g1 + g2
	h := h1 + h2
	f := g + h

	blockX := float64(x) * 100
	blockY := float64(y) * 100

	block := &Iblock{
		X:     x,
		Y:     y,
		Btype: 3,
		Color: color.Black,
		Rect:  pixel.R(blockX+1, blockY+1, blockX+100-1, blockY+100-1),
		G:     g,
		F:     f,
		H:     h,
		PX:    p.X,
		PY:    p.Y,
	}
	return block
}

func abs(num int) int {
	if num < 0 {
		num = -num
	}
	return num
}

// 获取一个块
func FindPathOneOpen(walls map[string]*Iblock, maps [][]string) (bool, *Iblock) {

	openOneMin := getOpenOne()

	// 放入close，从open中删除
	key := GetKey(openOneMin.X, openOneMin.Y)
	if _, ok := close[key]; !ok {
		close[key] = start
	}

	// 从open中删除
	if _, ok := open[key]; ok {
		delete(open, key)
	}

	// 上
	b, iblock := checkBlock(walls, openOneMin, openOneMin.X, openOneMin.Y+1)
	if b {
		return b, iblock
	}

	// 下  必须在边界内
	if openOneMin.Y-1 >= 0 {
		b, iblock = checkBlock(walls, openOneMin, openOneMin.X, openOneMin.Y-1)
		if b {
			return b, iblock
		}
	}

	// 左  必须在边界内
	if openOneMin.X-1 >= 0 {
		b, iblock = checkBlock(walls, openOneMin, openOneMin.X-1, openOneMin.Y)
		if b {
			return b, iblock
		}
	}
	// 右
	b, iblock = checkBlock(walls, openOneMin, openOneMin.X+1, openOneMin.Y)
	if b {
		return b, iblock
	}

	return false, nil

}

// 获取一个块
func FindPathAll(walls map[string]*Iblock, maps [][]string) (bool, *Iblock) {
	b, iblock := FindPathOneOpen(walls, maps)
	if b {
		return b, iblock
	}
	return FindPathAll(walls, maps)
}

// 获取open排序后的第一个块
func getOpenOne() *Iblock {
	var block *Iblock
	if len(open) > 0 {
		for _, v := range open {
			if block == nil {
				block = v
			}
			if block.F > v.F {
				block = v
			}
		}
		return block
	} else {
		return start
	}

}

func checkBlock(walls map[string]*Iblock, b *Iblock, x int, y int) (bool, *Iblock) {
	key := GetKey(x, y)
	// 是否已放入close
	if _, ok := close[key]; !ok {
		if strings.EqualFold(key, GetKey(end.X, end.Y)) {
			end.PX = b.X
			end.PY = b.Y
			return true, end
		}

		// 是否是墙
		if wall, ok := walls[key]; ok {
			if wall.Btype == 1 {
				return false, nil
			}
		}

		rigth := getBlock(b, x, y)
		rigth.UpdateIblock()
		walls[key] = rigth

		open[key] = rigth
	}

	return false, nil
}

func InitOpenClose() {
	open = make(map[string]*Iblock)
	close = make(map[string]*Iblock)
}

func DrawPath(walls map[string]*Iblock, iblock *Iblock) {
	if iblock.Btype == -1 {
		return
	}
	if road, ok := walls[GetKey(iblock.PX, iblock.PY)]; ok {
		road.Color = pixel.RGB(0.5, 0.2, 0.1)
		road.UpdateIblock()
		DrawPath(walls, road)
	}
	return
}
