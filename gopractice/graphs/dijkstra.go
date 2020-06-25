package graphs

import (
	"errors"
	"fmt"
	"math"
)

// aGraph is adjacency matrix represeantion of graph
type aGraph struct {
	V     int     // number of vertices/nodes
	edges [][]int // adjacency matrix
}

func initAGraph(n int) (*aGraph, error) {
	if n == 0 {
		return nil, errors.New("num nodes is zero")
	}
	// declare empty matrix
	edges := make([][]int, n)
	for i := range edges {
		edges[i] = make([]int, n)
	}
	return &aGraph{
		V:     n,
		edges: edges,
	}, nil
}

func (g *aGraph) insertEdge(u, v, wt int) error {
	if u == v {
		return errors.New("cannot have edges from same node")
	}
	g.edges[u][v] = wt
	return nil
}

// util to add all edges
func (g *aGraph) addEdges(edges [][]int) error {
	g.edges = edges
	return nil
}

// print graph
func (g *aGraph) printGraph() {
	fmt.Println(g.edges)
}

// time complexity O(V^2) can be improved to O(ElogV) with minHeap and Adj list impl
// return the shortest path weights from a given src
func (g *aGraph) dijkstraUtil(cov []bool, wt []int, src int) []int {
	// start with src node's wt = 0
	wt[src] = 0

	// for all vertices repeat (1 node covered after each iteration )
	for i := 1; i < g.V; i++ {

		// pick the uncovered vertex min wt
		u := minDist(cov, wt)

		// for all adajacent edges update the shorted path
		for v := 0; v < g.V; v++ {
			//  if edge exists, update the wt
			if g.edges[u][v] != 0 && (wt[u]+g.edges[u][v]) < wt[v] {
				wt[v] = wt[u] + g.edges[u][v]
			}
		}

		// mark the vertex u as covered
		cov[u] = true

	}
	// final shortest path array
	return wt
}

// find the index with min wt that is not in covered array
func minDist(cov []bool, wt []int) int {
	min := math.MaxInt32
	idx := -1
	for i, v := range wt {
		if v < min && cov[i] == false {
			min = v
			idx = i
		}
	}
	return idx
}

// Dijkstra shortest path
func findShortestPath(g *aGraph, src int) []int {
	// define a covered array with default 'false'
	cov := make([]bool, g.V)
	// define a weight array with shortest path weights
	wt := make([]int, g.V)
	// initialize with max_int
	for i := range wt {
		wt[i] = math.MaxInt32
	}

	return g.dijkstraUtil(cov, wt, src)

}
