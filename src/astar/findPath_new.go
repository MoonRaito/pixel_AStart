package astar

//
//import (
//	"container/heap"
//	"fmt"
//	"github.com/faiface/pixel"
//	"image/color"
//	"strconv"
//	"strings"
//)
//
//
//
//var open_new = &priorityQueue{}
//var close_new = make(map[string]*Iblock)
//var first = 1
//// 获取一个块
//func FindPathOneOpen_new(walls map[string]*Iblock) (bool, *Iblock) {
//
//	// 第一次时
//	if first==1 {
//		heap.Init(open_new)
//		heap.Push(open_new, start)
//	}
//
//	openOneMin := heap.Pop(open_new).(*Iblock)
//
//	// 放入close，从open中删除
//	key := GetKey(openOneMin.X, openOneMin.Y)
//	if _, ok := close[key]; !ok {
//		close[key] = start
//	}
//
//	// 从open中删除
//	//if _, ok := open[key]; ok {
//	//	delete(open, key)
//	//}
//
//	// 上 必须在边界内
//	if openOneMin.Y+1 < 8 {
//		b, iblock := checkBlock(walls, openOneMin, openOneMin.X, openOneMin.Y+1)
//		if b {
//			return b, iblock
//		}
//	}
//
//	// 下  必须在边界内
//	if openOneMin.Y-1 >= 0 {
//		b, iblock := checkBlock(walls, openOneMin, openOneMin.X, openOneMin.Y-1)
//		if b {
//			return b, iblock
//		}
//	}
//
//	// 左  必须在边界内
//	if openOneMin.X-1 >= 0 {
//		b, iblock := checkBlock(walls, openOneMin, openOneMin.X-1, openOneMin.Y)
//		if b {
//			return b, iblock
//		}
//	}
//	// 右
//	if openOneMin.X+1 < 10 {
//		b, iblock := checkBlock(walls, openOneMin, openOneMin.X+1, openOneMin.Y)
//		if b {
//			return b, iblock
//		}
//	}
//
//	first++
//
//	return false, nil
//
//}
