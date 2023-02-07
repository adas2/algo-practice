package graphs

import (
	"fmt"

	custom "adas2.io/practice/custom_impl"
)

// max size for stack
const maxSize = 100

// strongly connected components with Kosaraju's Algo
// Input: is a directed graph
// Output: list of strongly connected components

// DFSTraversal implements DFS traversal for graph
// use adjacency list format
func (g *Graph) DFSTraversal(src int, visited []bool) {
	// visited[src] =
	u := src
	// mark u as visited
	visited[u] = true
	fmt.Printf("Node visited: %d\n", u)
	// for adjacent verticies
	for _, v := range g.adjV[u] {
		// recursive call DFS
		if visited[v] == false {
			g.DFSTraversal(v, visited)
		}
	}
	return
}

// DFSUtil is custom DFS util for SCC algo
// returns the updated stack
// queue is similar to Top sort (?)
func (g *Graph) DFSUtil(src int, visited []bool, s *custom.Sstack) {
	u := src
	// mark u as visited
	visited[u] = true
	fmt.Printf("Node visited: %d\n", u)
	// for adjacent verticies
	for _, v := range g.adjV[u] {
		// recursive call DFS
		if visited[v] == false {
			g.DFSUtil(v, visited, s)
		}
	}
	s.Push(u)
}

// get transpose of a graph
func (g *Graph) transposeGraph() *Graph {
	// only valid for a directed graph
	if g.dir == false {
		return nil
	}

	gT := initGraph(g.V)
	gT.dir = true
	// for all edges u->v delete edge and add edge v->u
	for u := 0; u < g.V; u++ {
		for _, v := range g.adjV[u] {
			gT.addEdge(v, u)
		}
	}

	return gT
}

func (g *Graph) printGraph() {
	for u := 0; u < g.V; u++ {
		fmt.Printf("%d --> ", u)
		for _, v := range g.adjV[u] {
			fmt.Printf("%d, ", v)
		}
		fmt.Println()
	}
}

func findDFSTraversal(g *Graph, src int) {
	// define visited array
	visited := make([]bool, g.V)

	g.DFSTraversal(src, visited)
}

func findSCC(g *Graph) {
	// define visited array
	visited := make([]bool, g.V)
	// traverse time stack
	st := custom.InitStack(maxSize)

	// call DFS with traverse time stack untill all nodes are covered
	for i := 0; i < g.V; i++ {
		if visited[i] == false {
			g.DFSUtil(i, visited, st)
		}
	}
	fmt.Printf("Traverse time stack: %v\n", *st)

	// create transpose graph
	gT := g.transposeGraph()

	// array to keep already visited nodes in transpose graph
	discovered := make([]bool, gT.V)

	// while all nodes are visited
	for i := 0; i < gT.V; i++ {
		// pop stack and
		src, _ := st.Pop()

		if discovered[src] == false {
			fmt.Println("New SCC:")
			// traverse DFS with popped src on transpose graph
			gT.DFSTraversal(src, discovered)
		}
	}

}
