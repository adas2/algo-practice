package dp

import (
	"fmt"
	"math"
)

// LC #322

// memory inefficient version:
func coinChange(coins []int, amount int) int {
	// define dp[i][j] matrix for storing the optimal result
	// i-th index refers to amount, j-th index the coin value
	dp := make([][]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = make([]int, len(coins)+1)
	}

	// initialize the dp array with base cases
	// no coins will not yield any value
	for i := range dp {
		dp[i][0] = -1
	}
	// zero amount requires zero coins
	for j := range dp[0] {
		dp[0][j] = 0
	}

	// iterate over each amount and each coin value
	for a := 1; a <= amount; a++ {
		for c := 1; c <= len(coins); c++ {
			//  two cases:
			var first, second int
			// 1. 1 or more coin[c-1] is used
			if a-coins[c-1] >= 0 && dp[a-coins[c-1]][c] != -1 {
				first = dp[a-coins[c-1]][c] + 1
			} else {
				first = -1
			}
			// 2. coin[c-1] is not used at all
			second = dp[a][c-1]
			dp[a][c] = getMin(first, second)
		}
	}
	fmt.Println(dp)
	return dp[amount][len(coins)]
}

func getMin(a, b int) int {
	if a == -1 {
		return b
	}
	if b == -1 {
		return a
	}
	if a < b {
		return a
	}
	return b
}

// Memory efficient version:
// USe 1-D array for DP instead of 2D, since only count matters, order if coins don't
func coinChange2(coins []int, amount int) int {
	// define dp array
	dp := make([]int, amount+1)
	// case when amount in 0
	dp[0] = 0

	// bottom up using value
	for v := 1; v <= amount; v++ {
		// init to max
		dp[v] = math.MaxInt32
		for _, c := range coins {
			// case 1: coin c can be used
			if (v-c) >= 0 && dp[v-c] != math.MaxInt32 && dp[v-c]+1 < dp[v] {
				dp[v] = dp[v-c] + 1
			}
			// else: optimize for next iter
		}
	}

	if dp[amount] != math.MaxInt32 {
		return dp[amount]
	}

	return -1
}
