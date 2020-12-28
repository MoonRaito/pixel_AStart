package path

import (
	"container/heap"
	"fmt"
	"pixel_AStart/ebiten/queue"
	"strconv"
)

type IPath interface {
	// PathNeighbors returns the direct neighboring nodes of this node which
	// can be pathed to.
	PathNeighbors() []IPath
	// PathNeighborCost calculates the exact movement cost to neighbor nodes.
	PathNeighborCost(to IPath) int
	// PathEstimatedCost is a heuristic method for estimating movement costs
	// between non-adjacent nodes.
	PathEstimatedCost(to IPath) int
}

// node is a wrapper to store A* data for a IPath node.
type node struct {
	pather IPath
	cost   int
	rank   int
	parent *node
	open   bool
	closed bool
	index  int
}

// nodeMap is a collection of nodes keyed by IPath nodes for quick reference.
type nodeMap map[IPath]*node

// get gets the IPath object wrapped in a node, instantiating if required.
func (nm nodeMap) get(p IPath) *node {
	n, ok := nm[p]
	if !ok {
		n = &node{
			pather: p,
		}
		nm[p] = n
	}
	return n
}

// Path calculates a short path and the distance between the two IPath nodes.
//
// If no path is found, found will be false.
func GetPath(from IPath) (path []IPath, distance float64, found bool) {
	nm := nodeMap{}
	nq := &priorityQueue{}
	pq := queue.NewQueue()
	fromNode := nm.get(from)
	fromNode.open = true
	pq.EnQueue(fromNode)

	// open 先进 先出

	for {
		if nq.Len() == 0 {
			// There's no path, return found false.
			return
		}
		current := heap.Pop(nq).(*node)
		current.open = false
		current.closed = true

		for _, neighbor := range current.pather.PathNeighbors() {
			cost := current.cost + current.pather.PathNeighborCost(neighbor)
			neighborNode := nm.get(neighbor)

			fmt.Println("cost:" + strconv.Itoa(cost) + "    neighborNode:" + strconv.Itoa(neighborNode.cost))

			//if cost < neighborNode.cost {
			//	if neighborNode.open {
			//		heap.Remove(nq, neighborNode.index)
			//	}
			//	neighborNode.open = false
			//	neighborNode.closed = false
			//}
			//if !neighborNode.open && !neighborNode.closed {
			//	neighborNode.cost = cost
			//	neighborNode.open = true
			//	neighborNode.rank = cost + neighbor.PathEstimatedCost(to)
			//	neighborNode.parent = current
			//	heap.Push(nq, neighborNode)
			//}
		}

	}
}
