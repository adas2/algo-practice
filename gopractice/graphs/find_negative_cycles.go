package graphs

import (
	"fmt"
	"math"
)

// Bellman Ford shortest path algorithm that finds:
// i. shortest path of a directed graph with negative edge weights
// ii. detects if negative cycle is present in the graph

// Algorithm:
// declare dist array, and init to inf
// dist[u] = 0
// iterate for num_vertices - 1
// for each edge e (u->v)
// relax edge: if dist[v] > d[u]+weight(u,v) { dist[v] = d[u]+weight(u,v) }
// output dist array after V.E iteration ends

// const noEdge = 0
const inf = math.MaxInt32

// edge representation of graph
type eGraph struct {
	V     int     //num virtex
	edges [][]int //edge list
}

func initEdgeGraph(num int) *eGraph {
	return &eGraph{
		V:     num,
		edges: make([][]int, 0),
	}
}

func (g *eGraph) addEdge(u, v, w int) {
	g.edges = append(g.edges, []int{u, v, w})
}

func (g *eGraph) printEdges() {
	for _, e := range g.edges {
		fmt.Printf("%d --> %d : %d\n", e[0], e[1], e[2])
	}
}

// using adjacency matrix representaion
func findShortedPathFromSrc(g *eGraph, src int) ([]int, []int) {
	numV := g.V
	// dist array intit to inf
	dist := make([]int, numV)
	for i := range dist {
		dist[i] = inf
	}
	// parent array
	parent := make([]int, numV)
	for i := range dist {
		parent[i] = -1
	}

	// start from src
	dist[src] = 0

	// for path len with max of |V|-1 edges
	for i := 1; i < numV; i++ {
		// for all edges
		for _, e := range g.edges {
			// relax edge
			u, v, w := e[0], e[1], e[2]
			if dist[u] != inf && w+dist[u] < dist[v] {
				dist[v] = w + dist[u]
				// update parent
				parent[v] = u
			}
		}
	}

	return dist, parent
}

func printPath(dist, parent []int) {
	// print shortest path from src with parent info
	for i := 0; i < len(dist); i++ {
		fmt.Printf("Node: %d, Parent: %d, MinDist: %d\n", i, parent[i], dist[i])
	}
}

// Algo:
func isNegativeCycle(g *eGraph) bool {

	// find shortest path
	dist, _ := findShortedPathFromSrc(g, 0)

	// iterate an additional time over each edge
	// if the dist/weight of any node changes return true else false
	for _, e := range g.edges {
		u, v, w := e[0], e[1], e[2]
		if dist[u] != inf && w+dist[u] < dist[v] {
			return true
		}
	}

	return false
}
