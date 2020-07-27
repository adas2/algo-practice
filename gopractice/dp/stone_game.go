package dp

// LC # 1140
// 2 player game with pile pf stones
func stoneGameII(piles []int) int {
	// define DP
	n := len(piles)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// calculate suffix sum to store the partial sums from end to start
	for i := n - 2; i >= 0; i-- {
		piles[i] += piles[i+1]
	}

	res := optimalGameSelection(piles, dp, 0, 1)

	return res
}

// i := index of the starting point of array
// M = Max(x,M) starts with 1
func optimalGameSelection(piles []int, dp [][]int, i, M int) int {
	// base cases
	n := len(piles)
	if i >= n {
		return 0
	}
	// if this choice can take all elements that maximizes the value/points
	if 2*M >= n-i {
		// check i bounds
		return piles[i]
	}
	if dp[i][M] != 0 {
		// memoization
		return dp[i][M]
	}

	// iterate for all possible 1 <= x <= M
	for x := 1; x <= 2*M; x++ {
		// value taken with choice of x piles[i]....piles[i+x]
		taken := piles[i] - piles[i+x]
		// remaining total value
		remaining := piles[i+x] - optimalGameSelection(piles, dp, i+x, Max(x, M))
		// if Lee chooses to play optimally the value which will go to Alex after Lee's turn
		dp[i][M] = Max(taken+remaining, dp[i][M])
	}
	return dp[i][M]
}

// Similar to coin_profit game
// This type of problem is called MinMax problem
// to maximize the player 1's profit, assume player 2 also plays optimally
// thus player 1's next turn will be minimal value
// states: Current value(V), options 1<= x <= 2M, index i of remanining stone piles
// (V,x,i) = Max( for 1 <= x <= M, piles[:x] + Min( V, y, i+x+y ) for all y where 1 <= y <= 2x
// )
