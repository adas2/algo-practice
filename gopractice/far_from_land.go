package practice

// lc # 1162

// BFS solution:
func maxDistance(grid [][]int) int {
	// create a visited array (note/; all cells inits to zero)
	visited := make([][]int, len(grid))
	for i := 0; i < len(grid[0]); i++ {
		visited[i] = make([]int, len(grid[0]))
	}

	// define bfs queue/slice {x,y} each element of queue
	queue := make([][]int, 0)

	// add the 1's to queue and mark them as visited
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				queue = append(queue, []int{i, j})
				visited[i][j] = 1
			}
		}
	}

	// tricky: distance per iteration, set to -1 since it counts 1 for the inital 1 cells
	dist := -1

	// offsets
	roff := []int{-1, 0, 0, 1}
	coff := []int{0, -1, 1, 0}

	// while queue is not empty
	for len(queue) > 0 {
		// note the len such we keep track of each iteration in the queue
		qlen := len(queue)
		// 1 itera to process the neighbors of the cells
		for qlen > 0 {
			// pop cell
			x, y := queue[0][0], queue[0][1]
			queue = queue[1:]
			qlen -= 1

			// traverse all neighbors of x,y, if not 1 visit and add to queue
			for i := 0; i < 4; i++ {
				if x+roff[i] >= 0 && y+coff[i] >= 0 && x+roff[i] < len(grid) && y+coff[i] < len(grid[0]) &&
					visited[x+roff[i]][y+coff[i]] == 0 {
					visited[x+roff[i]][y+coff[i]] = 1
					queue = append(queue, []int{x + roff[i], y + coff[i]})
				}
			}
		}
		// ietration ends here, we have visited +1 dist
		dist += 1
	}

	// 0 dist means no zeros in grid or all zeros
	if dist == 0 {
		return -1
	}

	return dist
}

// Logic:
// multi source bfs traversal from every 1, at each tsep we traverse the nearest zero
// if already visited skip, else add to queue for next iterations.
// every iteration will this visit the cells which are +1 distance from the neartest 1
// after n iterations the +n dist, this way the largest iteration is the max dist
// since we cover all 1's we are guranteed to get the min dist from 1 for every 0's
