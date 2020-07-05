package practice

// LC #85

// convert to histogram area problem
func maximalRectangle(matrix [][]byte) int {

	heights := convertByteToIntMatrix(matrix)

	// sum up the columns for each row
	for i := 1; i < len(heights); i++ {
		for j := range heights[0] {
			heights[i][j] += heights[i-1][j]
		}
	}

	// now for each row find largest histogram area
	maxRect := 0
	for i := range heights {
		rect := largestRectangleArea(heights[i])
		maxRect = max(maxRect, rect)
	}
	return maxRect
}

func convertByteToIntMatrix(matrix [][]byte) [][]int {
	// convert to int matrix (needed to resue the histogram problem)
	heights := make([][]int, len(matrix))
	for i := range matrix {
		heights[i] = make([]int, len(matrix[0]))
	}
	for i := range matrix {
		for j := range matrix[0] {
			heights[i][j] = int(matrix[i][j]) - '0'
		}
	}

	return heights
}
