package path

import "pixel_AStart/ebiten/tiled"

type IPath interface {
	// PathNeighbors returns the direct neighboring nodes of this node which
	// can be pathed to.
	PathNeighbors() []*tiled.Tile
	// PathNeighborCost calculates the exact movement cost to neighbor nodes.
	PathNeighborCost(to tiled.Tile) int
	// PathEstimatedCost is a heuristic method for estimating movement costs
	// between non-adjacent nodes.
	PathEstimatedCost(to tiled.Tile) float64
}
