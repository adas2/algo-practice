package dp

import (
	"fmt"
	"math"
)

func maxDotProduct(nums1 []int, nums2 []int) int {
	// error case
	if len(nums1) == 0 || len(nums2) == 0 {
		return -1
	}
	// initialize dp array
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	// init incl. base cases, when no element is chosen
	for i := 0; i <= len(nums1); i++ {
		for j := 0; j <= len(nums2); j++ {
			dp[i][j] = math.MinInt32
		}
	}

	// iterate over the arrays
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			fmt.Printf("i=%d, j=%d: ", i, j)
			fmt.Println(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]+(nums1[i-1]*nums2[j-1]))
			dp[i][j] = FourMax(
				dp[i-1][j],
				dp[i][j-1],
				dp[i-1][j-1],
				Max(dp[i-1][j-1], 0)+(nums1[i-1]*nums2[j-1]),
			)
		}
	}
	// return for the full length
	return dp[len(nums1)][len(nums2)]
}

// util func
func FourMax(a, b, c, d int) int {
	return Max(Max(Max(a, b), c), d)
}

// logic:
// Base case
// When i == 0 or j == 0, we return INT_MIN. (Because this is an empty case, which is intolerable)
// State Transition, for any i > 0 and j > 0, there are 4 possibilities
// nums1[i - 1] is not selected, dp[i][j] = dp[i - 1][j]
// nums2[j - 1] is not selected, dp[i][j] = dp[i][j - 1]
// Neither nums1[i - 1] or nums2[j - 1] is selected, dp[i][j] = dp[i - 1][j - 1]
// Both nums1[i - 1] and nums2[j - 1] are selected, dp[i][j] = max(dp[i - 1][j - 1], 0) + nums1[i - 1] * nums2[j - 1]
// Since we already selected one pair (nums1[i - 1], nums2[j - 1]), we can assume the minimal proceeding value is 0
