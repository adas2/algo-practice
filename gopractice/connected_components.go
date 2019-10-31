package practice

import "fmt"

// type GNode struct {
// }

// Queue struct used in BST
type Queue []int

func (q *Queue) Push(v int) {

	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	n := len(*q)
	old := *q
	*q = old[:n-1]

	return old[n-1]
}

func (q *Queue) IsEmpty() bool {
	return (len(*q) == 0)
}

func (q *Queue) PrintQueue() {
	for i := range *q {
		fmt.Printf("%d ", (*q)[i])
	}
	fmt.Println()
}

// Graph struct
type gList []int

type Graph struct {
	V    int     // num vertices
	adjV []gList // adjacency lists for each vertex
}

func initGraph(N int) *Graph {
	g := &Graph{
		V:    N,
		adjV: nil,
	}

	// initialize empty lists
	g.adjV = make([]gList, N)
	// probably don't need
	// for i := range g.adjV {
	// 	g.adjV[i] = nil
	// }

	// fmt.Println(g.adjV)
	return g
}

func (g *Graph) addEdge(u, v int) {
	// undirected graph add two edges
	g.adjV[u] = append(g.adjV[u], v)
	g.adjV[v] = append(g.adjV[v], u)
}

// BST traversal
func (g *Graph) TraverseBST(source int, discovered []bool) {
	fmt.Println("BST travresal:")

	// var u int
	// init Queue
	q := Queue([]int{})

	// start with src node
	discovered[source] = true
	q.Push(source)

	// fmt.Println(discovered)
	// fmt.Println(g.adjV)
	// q.PrintQueue()

	// while queue has nodes
	for q.IsEmpty() == false {
		// find adjacent nodes
		u := q.Pop()
		// pre-process u
		fmt.Printf("%v ", u)
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

func FindConnectedComponents(g *Graph) {

	// create discovered array
	numNodes := g.V
	discovered := make([]bool, numNodes)

	for i := range discovered {
		discovered[i] = false
	}
	var c int = 1

	// find components by BST traversal
	for i := 0; i < numNodes; i++ {
		if discovered[i] == false {
			// new component
			fmt.Println("Component:", c)
			g.TraverseBST(i, discovered)
			c++
		}
	}

	return
}
