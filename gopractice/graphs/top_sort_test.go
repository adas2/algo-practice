package graphs

import (
	"fmt"
	"testing"
)

func TestTopSort(t *testing.T) {
	// create a graph
	g := initGraph(6)
	g.dir = true
	g.addEdge(5, 2)
	g.addEdge(5, 0)
	g.addEdge(4, 0)
	g.addEdge(4, 1)
	g.addEdge(2, 3)
	g.addEdge(3, 1)

	g.printGraph()

	// in dgree for nodes 0...5
	inD := []int{2, 2, 1, 1, 0, 0}

	tSort := topSortGraph(g, inD)

	fmt.Println(tSort)

}
