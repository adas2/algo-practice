package graphs

import (
	"fmt"
	"testing"
)

func TestDijkstraShortedPath(t *testing.T) {

	g, err := initAGraph(9)
	if err != nil {
		t.Errorf("graph init failed\n")
	}

	g.addEdges(
		[][]int{
			{0, 4, 0, 0, 0, 0, 0, 8, 0},
			{4, 0, 8, 0, 0, 0, 0, 11, 0},
			{0, 8, 0, 7, 0, 4, 0, 0, 2},
			{0, 0, 7, 0, 9, 14, 0, 0, 0},
			{0, 0, 0, 9, 0, 10, 0, 0, 0},
			{0, 0, 4, 14, 10, 0, 2, 0, 0},
			{0, 0, 0, 0, 0, 2, 0, 1, 6},
			{8, 11, 0, 0, 0, 0, 1, 0, 7},
			{0, 0, 2, 0, 0, 0, 6, 7, 0},
		},
	)

	g.printGraph()

	result := findShortestPath(g, 0)

	for i, v := range result {
		fmt.Printf("Vertex: %d, Wt: %d\n", i, v)
	}
}
