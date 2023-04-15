package graphs

func possibleBipartition(n int, dislikes [][]int) bool {
	// create adjacency matrix
	g := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		g[i] = make([]int, n+1)
	}

	// create graph: tricky make sure add undirected edges both u->v, v->u
	for _, d := range dislikes {
		g[d[0]][d[1]] = 1
		g[d[1]][d[0]] = 1
	}

	// color slice
	color := make([]int, n+1)
	for i := range color {
		color[i] = -1
	}

	// check for all unvisited nodes
	for v := 1; v <= n; v++ {
		if color[v] == -1 {
			if bipartitionUtil(g, n, v, &color) == false {
				return false
			}
		}
	}
	// all nodes could be colored
	return true

}

func bipartitionUtil(g [][]int, n, src int, color *[]int) bool {
	// queue for BFS
	queue := []int{}

	queue = append(queue, src)
	(*color)[src] = 0

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		// self edge cannot be bipartite
		if g[u][u] == 1 {
			return false
		}

		for v := 1; v <= n; v++ {
			// edge
			if g[u][v] == 1 && (*color)[v] == -1 {
				(*color)[v] = 1 - (*color)[u]

				// push
				queue = append(queue, v)
			} else if g[u][v] == 1 && (*color)[v] == (*color)[u] {
				return false
			}
		}
		// all v processed
	}
	return true
}

// time complexity := O(N^2)
// TODO: to make it O(N+E) use adjacency list repr
