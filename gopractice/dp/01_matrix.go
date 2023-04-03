package dp

// lc # 542 (grid problem)
// O(n^2) 2 pass
func updateMatrix(mat [][]int) [][]int {
	// declare dp array (init to 0 by default)
	dp := make([][]int, len(mat))
	for i := range mat {
		dp[i] = make([]int, len(mat[0]))
	}

	// Imp: init to Max val and then generalize the min func
	// note: by prob def max value of a path is 10^4
	// Using Inf or MaxInt could cause overflow since +1 op is done below
	for i := range dp {
		for j := range dp[0] {
			dp[i][j] = 10000
		}
	}

	// first pass: top left
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				dp[i][j] = 0
			} else {
				// top and left neighbors
				if i >= 1 {
					dp[i][j] = Min(dp[i][j], dp[i-1][j]+1)
				}

				if j >= 1 {
					dp[i][j] = Min(dp[i][j], dp[i][j-1]+1)
				}
			}
		}
	}

	// second pass
	for i := len(mat) - 1; i >= 0; i-- {
		for j := len(mat[0]) - 1; j >= 0; j-- {
			if mat[i][j] == 0 {
				dp[i][j] = 0
			} else {
				// bottom and right neighbors
				if i < len(mat)-1 {
					dp[i][j] = Min(dp[i][j], dp[i+1][j]+1)
				}

				if j < len(mat[0])-1 {
					dp[i][j] = Min(dp[i][j], dp[i][j+1]+1)
				}
			}
		}
	}

	return dp
}

// logic:
// 2 travsersal dp with optimization;one from top right
// dp[i,j] = min(dp[i-1,j], dp[i, j-1])
// dp[i,j] = 0 if mat[i,j] --> 0
// dp[i,j] = 1 if mat[i-1,j] == 0 or mat[i, j-1] == 0
// second pass one from bottom left
// dp[i,j] = min(dp[i+1,j], dp[i, j+1], dp[i,j])
// other 2 conditions same as first pass
