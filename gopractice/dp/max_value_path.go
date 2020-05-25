package dp

import "fmt"

// Given the 2D matrix of coordinates with path values
// find the max value traversal path
// path from NE to SW
// accepted moves: left, down
// assumption: no -ve values

// type Point struct {
// 	x int
// 	y int
// }

func FindMaxValuePath(valueMatrix [][]int) {
	rows := len(valueMatrix)
	cols := len(valueMatrix[0])

	path := make([][]int, rows)
	for i := range path {
		path[i] = make([]int, cols)
	}
	// init first value path[]
	path[0][cols-1] = valueMatrix[0][cols-1]

	for i := 0; i < rows; i++ {
		for j := cols - 1; j >= 0; j-- {
			// if starting row
			if i == 0 {
				if j != cols-1 {
					path[i][j] = valueMatrix[i][j] + path[i][j+1]
				}
			} else if j == cols-1 {
				if i != 0 {
					path[i][j] = valueMatrix[i][j] + path[i-1][j]
				}
			} else {
				path[i][j] = findMax(path[i-1][j], path[i][j+1]) + valueMatrix[i][j]

			}
		}
	}

	fmt.Println(path[rows-1][0])
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Logic:
// 1 3 7
// 4 2 5
// 0 6 1

// expected path : 7->5->2->6->0

// start v[0][n] --> find v[m][0]
// path[0][2] = 7
// path[1][2] = 7+5, path[0][1] = 7+2
// path[1][1] = max (path[0][1], path[1][2]) ...
