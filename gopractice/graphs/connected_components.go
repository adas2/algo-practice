package graphs

import (
	"fmt"

	custom "adas2.io/practice/custom_impl"
)

// Graph adj list struct
type gList []int

// Graph struct
type Graph struct {
	V    int     // num vertices
	adjV []gList // adjacency lists for each vertex
	dir  bool    // directed graph?
}

func initGraph(N int) *Graph {
	g := &Graph{
		V:    N,
		adjV: nil,
		dir:  false,
	}

	// initialize empty lists
	g.adjV = make([]gList, N)
	// fmt.Println("adj list", g.adjV)
	return g
}

func (g *Graph) addEdge(u, v int) {
	// undirected graph add two edges
	g.adjV[u] = append(g.adjV[u], v)
	if g.dir == false {
		g.adjV[v] = append(g.adjV[v], u)
	}
}

// TraverseBFS = BFS traversal
func (g *Graph) TraverseBFS(source int, discovered []bool) {
	fmt.Println("BFS traversal:")

	// init Queue to empty
	q := custom.Queue([]int{})

	// start with src node
	discovered[source] = true
	q.Push(source)

	// while queue has nodes
	for q.IsEmpty() == false {
		// find adjacent nodes
		u := q.Pop()
		// pre-process u
		fmt.Printf("processing %v \n", u)
		for _, v := range g.adjV[u] {
			if discovered[v] == false {
				discovered[v] = true
				q.Push(v)
			}
			// else seen before
		}
		// u is processed
	}
	fmt.Println()

}

// FindConnectedComponents finds connected graphs
func FindConnectedComponents(g *Graph) {

	// create discovered array
	numNodes := g.V
	// default init to false
	discovered := make([]bool, numNodes)
	// fmt.Println("discovered", discovered)

	// init comp count
	c := 0

	// find components by BFS traversal
	for i := 0; i < numNodes; i++ {
		if discovered[i] == false {
			// new component
			c++
			fmt.Println("Component:", c)
			g.TraverseBFS(i, discovered)
		}
	}

	return
}
