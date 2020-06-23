package practice

import (
	"fmt"
	"strconv"
)

// NumIslands : Given 2-D grid of 0/1 find the number of islands of 1
func NumIslands(grid [][]byte) int {
	// convert from byte to int format
	intGrid := convertGridFormat(grid)

	// initialize visited array
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	return numIslandsHelper(intGrid, visited)
}

// helper function
func numIslandsHelper(grid [][]int, visited [][]bool) int {

	var numIslands int = 0
	// traverse each element in grid, mark as visited after processing numIslands
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 && !visited[i][j] {
				visitNbr(grid, visited, i, j)
				numIslands++
				fmt.Printf("Found island: %d\n", numIslands)
			}
			// else 0's or visited neighbors
		}
	}

	return numIslands
}

// are the neighbors all unvisited? if searched mark as visited
func visitNbr(grid [][]int, visited [][]bool, i, j int) {
	row := len(grid)
	col := len(grid[0])
	// visit node
	if !visited[i][j] {
		visited[i][j] = true
	}

	// recursively traverse all 1 value nodes
	if i > 0 && grid[i-1][j] == 1 && !visited[i-1][j] {
		visitNbr(grid, visited, i-1, j)
	}
	if i < row-1 && grid[i+1][j] == 1 && !visited[i+1][j] {
		visitNbr(grid, visited, i+1, j)

	}
	if j > 0 && grid[i][j-1] == 1 && !visited[i][j-1] {
		visitNbr(grid, visited, i, j-1)
	}
	if j < col-1 && grid[i][j+1] == 1 && !visited[i][j+1] {
		visitNbr(grid, visited, i, j+1)
	}
	// processed all nbrs
	return

}

// convert input grid from [][]byte to [][]int
func convertGridFormat(grid [][]byte) [][]int {
	// intialize output
	out := make([][]int, len(grid))
	for i := range out {
		out[i] = make([]int, len(grid[0]))
	}

	for i := range grid {
		for j := range grid[i] {
			// todo: error check
			intVal, _ := strconv.Atoi(string(grid[i][j]))
			out[i][j] = intVal
		}
	}

	return out
}

// simpler func (unused): does the grid point have neighbors?
func hasNbr(grid [][]int, i, j int) bool {
	row := len(grid)
	col := len(grid[0])

	if (i > 0 && grid[i-1][j] == 1) ||
		(i < row-1 && grid[i+1][j] == 1) ||
		(j > 0 && grid[i][j-1] == 1) ||
		(j < col-1 && grid[i][j+1] == 1) {
		return true
	}

	return false
}
