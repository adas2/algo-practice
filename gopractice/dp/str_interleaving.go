package dp

// LC #97
// given 3 strs s1,s2,s3 find if s3 is an interleaving of s1 and s2
func isInterleave(s1 string, s2 string, s3 string) bool {

	m, n := len(s1), len(s2)

	// error case
	if m+n != len(s3) {
		return false
	}
	// define dp array: default is false
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// iterate over m and n
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 && j == 0 {
				// base case: nil strs match
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] && s2[j-1] == s3[i+j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1]
			} else {
				//  non-zero i,j
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1] || dp[i][j-1] && s2[j-1] == s3[i+j-1]
			}
		}
	}

	return dp[m][n]
}

// classic DP problem with one trick:
// natural intution is to include dp[i][j][k] but note k = i+j+1 the index
// so the dp array doesn't need to be 3d but can just be 2d, hence time/space = O(mn)
// in fact LC has a memory optimized version with space O(m) where m is the smaller string
// 2 cases:
//  i.
// dp[i][j] = dp[i-1][j] && s1[i]==s3[k] (when ith char matches with k)
// 				i.e., second term becomes s1[i]==s3[i+j+1]
// ii.
// dp [i][j = dp[i][j-1] && s2[j]==s3[i+j+1] (when jth char matches with k)
// if no match then dp[i][j] = false
