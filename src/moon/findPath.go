package moon

import (
	"container/heap"
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

func GetStart() *Iblock {
	return start
}
func GetEnd() *Iblock {
	return end
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
		PXY:   p.PXY,
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

	// 上 必须在边界内
	if openOneMin.Y+1 < 8 {
		b, iblock := checkBlock(walls, openOneMin, openOneMin.X, openOneMin.Y+1)
		if b {
			return b, iblock
		}
	}

	// 下  必须在边界内
	if openOneMin.Y-1 >= 0 {
		b, iblock := checkBlock(walls, openOneMin, openOneMin.X, openOneMin.Y-1)
		if b {
			return b, iblock
		}
	}

	// 左  必须在边界内
	if openOneMin.X-1 >= 0 {
		b, iblock := checkBlock(walls, openOneMin, openOneMin.X-1, openOneMin.Y)
		if b {
			return b, iblock
		}
	}
	// 右
	if openOneMin.X+1 < 10 {
		b, iblock := checkBlock(walls, openOneMin, openOneMin.X+1, openOneMin.Y)
		if b {
			return b, iblock
		}
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

// 检查块 如果是 终点 返回true 其他返回false
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

var open_new = &priorityQueue{}
var close_new = make(map[string]*Iblock)
var first = 1

// 获取一个块
func FindPathOneOpen_new(walls map[string]*Iblock) (bool, *Iblock) {

	// 第一次时
	if true {
		fmt.Println(first)
		heap.Init(open_new)
		heap.Push(open_new, start)
	}

	if open_new.Len() == 0 {
		// There's no path, return found false.
		return false, nil
	}

	current := heap.Pop(open_new).(*Iblock)
	current.Open = false
	current.Closed = true

	neighbors := pathNeighbors(walls, current)

	for _, neighbor := range neighbors {
		//cost := current.cost + current.pather.PathNeighborCost(neighbor)
		// 默认所有 代价为1
		cost := current.cost + 1
		neighborNode := walls[GetKey(neighbor.X, neighbor.Y)]
		if cost < neighborNode.cost {
			if neighborNode.Open {
				heap.Remove(open_new, neighborNode.index)
			}
			neighborNode.Open = false
			neighborNode.Closed = false
		}
		if !neighborNode.Open && !neighborNode.Closed {
			neighborNode.cost = cost
			neighborNode.Open = true
			neighborNode.rank = cost + neighbor.PathEstimatedCost(end)
			neighborNode.parent = current
			heap.Push(open_new, neighborNode)
		}
	}

	//first++

	return false, nil

}

// 获取邻居节点
func pathNeighbors(walls map[string]*Iblock, current *Iblock) []*Iblock {
	neighbors := []*Iblock{}
	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		if n := walls[GetKey(current.X+offset[0], current.Y+offset[1])]; n != nil &&
			n.Btype != 1 {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

func (t *Iblock) PathEstimatedCost(to *Iblock) float64 {
	absX := to.X - t.X
	if absX < 0 {
		absX = -absX
	}
	absY := to.Y - t.Y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}
