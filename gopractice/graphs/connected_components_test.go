package graphs

import "testing"

func TestFindConnectedComponents(t *testing.T) {
	g := initGraph(7)

	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(3, 4)
	g.addEdge(2, 5)
	g.addEdge(5, 6)

	FindConnectedComponents(g)
}
