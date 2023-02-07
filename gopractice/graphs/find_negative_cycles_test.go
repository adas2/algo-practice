package graphs

import (
	"fmt"
	"math"
	"testing"
)

func TestFindShortestPathFromSource(t *testing.T) {
	g := initEdgeGraph(5)
	g.addEdge(0, 1, -1)
	g.addEdge(0, 2, 4)
	g.addEdge(1, 2, 3)
	g.addEdge(1, 3, 2)
	g.addEdge(1, 4, 2)
	g.addEdge(3, 2, 5)
	g.addEdge(3, 1, 1)
	g.addEdge(4, 3, -3)

	dist, parent := findShortedPathFromSrc(g, 0)
	printPath(dist, parent)

}

func TestIfNegativeCycle(t *testing.T) {
	g := initEdgeGraph(3)
	// 0 1 -5 1 2 -6 2 0 -9
	g.addEdge(0, 1, -5)
	g.addEdge(1, 2, -6)
	g.addEdge(0, 2, -9)
	g.addEdge(2, 2, -1)

	if isNegativeCycle(g) {
		fmt.Printf("NEGATIVE CYCLE EXISTS\n")
	} else {
		fmt.Printf("NO NEGATIVE CYCLES\n")
	}

}

func TestDetectNegativeCycles(t *testing.T) {
	g := initEdgeGraph(3)
	// 0 1 -5 1 2 -6 2 0 -9
	g.addEdge(0, 1, 5)
	g.addEdge(1, 2, 6)
	g.addEdge(2, 0, 9)
	g.addEdge(2, 2, -1)

	d := detectNegativeCycles(g)

	fmt.Printf("Nodes part of negative cycle:\n")

	// show nodes that form negative cycle
	for i := range d {
		if d[i] == int(math.Inf(-1)) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println(d)
}
