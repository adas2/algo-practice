package graphs

// LC # 1293. Shortest Path in a Grid with Obstacles Elimination
// Hard/tricky

// struct to keep track of visited grid
type state struct {
	r int
	c int
	k int
}

func shortestPath(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])

	// define queue for BFS: entry={row, col, k, steps}
	queue := [][]int{}

	// define set for visited; note this could be a 3d array, but saves mem
	visited := map[state]bool{}

	// keep the neighboring index offsets
	rowOff := []int{-1, 0, 1, 0}
	colOff := []int{0, -1, 0, 1}

	// add the src cell in grid with dist 0
	queue = append(queue, []int{0, 0, k, 0})
	visited[state{0, 0, k}] = true

	// start BFS
	for len(queue) > 0 {
		// pop the leftmost
		curr := queue[0]
		queue = queue[1:]
		obs := curr[2]
		steps := curr[3]
		// terminating condition
		if curr[0] == m-1 && curr[1] == n-1 {
			return steps
		}
		// else check all neighbors
		for i := range rowOff {
			r, c := curr[0]+rowOff[i], curr[1]+colOff[i]
			if isSafeNeighbor(r, c, k, grid) {
				// check if obstacle
				if grid[r][c] == 1 && obs > 0 && !isVisited(r, c, obs-1, visited) {
					queue = append(queue, []int{r, c, obs - 1, steps + 1})
					// mark new cell,k as visited
					visited[state{r, c, obs - 1}] = true
				}
				if grid[r][c] == 0 && !isVisited(r, c, obs, visited) {
					// no obstacles
					queue = append(queue, []int{r, c, obs, steps + 1})
					visited[state{r, c, obs}] = true
				}
			}
		}
	}

	// if no such path exists it will come here
	return -1
}

func isSafeNeighbor(i, j, k int, grid [][]int) bool {
	// valid bounds
	if i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0]) {
		// unvisited state
		return true
	}
	return false
}

func isVisited(i, j, k int, visited map[state]bool) bool {
	// exists in set?
	if _, exists := visited[state{i, j, k}]; exists {
		return true
	}
	return false
}

// Logic: This is classic BFS/Lee's algorithm for flood-fill but with extra state i.e. k
// Hence for each valid cell, k options has to be considered at max
// this makes the time complexity O(m.n.k) and Space = O(m.n.k)
// Algo: use the queue to keep the (cell, k, steps) and use BFS
// add to queue only when (next_cell,k) is unvisited and k >= 0, where if cell_val == 1, k=k-1
// When the curr elem in the queue == dest return the steps
// To keep track of the already visited cell, maintain set visited(cell, k)
