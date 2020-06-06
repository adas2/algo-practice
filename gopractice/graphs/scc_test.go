package graphs

import (
	"fmt"
	"testing"
)

func TestDFSTraversal(t *testing.T) {
	g := initGraph(4)
	// make directed graph
	g.dir = true

	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 2)
	g.addEdge(2, 0)
	g.addEdge(2, 3)
	g.addEdge(3, 3)

	findDFSTraversal(g, 2)

}

func TestTranspose(t *testing.T) {

	g := initGraph(5)
	g.dir = true
	g.addEdge(0, 1)
	g.addEdge(0, 4)
	g.addEdge(0, 3)
	g.addEdge(2, 0)
	g.addEdge(3, 2)
	g.addEdge(4, 1)
	g.addEdge(4, 3)

	g.printGraph()
	gT := g.transposeGraph()
	fmt.Println("Transpose")
	gT.printGraph()
}

func TestSCC(t *testing.T) {

	g := initGraph(9)
	g.dir = true

	// example for geeksforgeeks
	g.addEdge(0, 1)
	g.addEdge(1, 2)
	g.addEdge(2, 3)
	g.addEdge(3, 0)
	g.addEdge(2, 4)
	g.addEdge(4, 5)
	g.addEdge(5, 6)
	g.addEdge(6, 4)
	g.addEdge(7, 6)
	g.addEdge(7, 8)

	g.printGraph()

	findSCC(g)

}
