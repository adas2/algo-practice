package practice

import "testing"

func TestFindConnectedComponents(t *testing.T) {
	g := initGraph(5)

	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(3, 4)

	FindConnectedComponents(g)
}
