package dp

// EPI: DP pp: 274
// Given the possible play scores in football and a final score
// Calculate the total number of combinations of the outcome of the plays
// E.g. Plays : {2,3,7} Final score = 12, Output: 4 ways
//  1. 2X6 all dafety
//  2. 2x3 + 3x2 2 safety, 2 FGs
//  3. 7+3+2 TD,FG, safety
//  4. 3x4 4 FGs

func scoreCombinations(final int, scores []int) int {
	// Using 2D array
	dp := make([][]int, final+1)
	// init the array to
	for i := range dp {
		// default inits to 0
		dp[i] = make([]int, len(scores)+1)
	}
	// final score of 0 in one way only
	for j := range scores {
		dp[0][j] = 1
	}

	// bottom up DP[i][j] reprs the tot score combinations
	// 	for final score i using score[1..j] plays
	for i := 1; i <= final; i++ {
		for j := 1; j <= len(scores); j++ {
			// case 1: winning scores[j-1] is used
			if i >= scores[j-1] && dp[i-scores[j-1]][j] != 0 {
				dp[i][j] += dp[i-scores[j-1]][j]
			}
			// case 2: score[j-1] is not used
			if dp[i][j-1] != 0 {
				dp[i][j] += dp[i][j-1]
			}
			// total
			// dp[i][j] = with + without scores[j-1]
		}
	}

	// fmt.Println(dp)

	return dp[final][len(scores)]
}
