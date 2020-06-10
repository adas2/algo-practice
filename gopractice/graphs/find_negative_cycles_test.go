package graphs

import (
	"fmt"
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
	g := initEdgeGraph(5)
	// 0 1 -5 1 2 -6 2 0 -9
	g.addEdge(0, 1, -5)
	g.addEdge(1, 2, -6)
	g.addEdge(2, 0, -9)

	if isNegativeCycle(g) {
		fmt.Printf("CYCLE EXISTS\n")
	} else {
		fmt.Printf("NO CYCLE EXISTS")
	}

}
