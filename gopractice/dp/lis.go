package dp

// find the longest increasing subsequence in a given array (int)

// method 1: usng dp with O(N^2) complexity
// output: length of the LIS
func lengthOfLIS(nums []int) int {
	// error case
	if len(nums) == 0 {
		return 0
	}
	// adjust size to include whole in array
	dp := make([]int, len(nums))
	// base case: LIS len for only 1 element
	dp[0] = 1
	var maxLIS int
	result := 1

	for i := 1; i < len(nums); i++ {
		maxLIS = 0 //init
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				// new sequence foound
				maxLIS = Max(maxLIS, dp[j])
			}
		}
		// in case no sequence, LIS is 1
		dp[i] = Max(maxLIS+1, 1)
		// fmt.Printf("maxLIS: %d, dp[%d] = %d\n", maxLIS, i, dp[i])
		// store the max len seen so far
		result = Max(result, dp[i])
		// maxLIS = Max(maxLIS+1, dp[i-1])
	}

	return result
}

// explanation:
// dp[i] ==> LIS len for arr[0..i]
// dp[i] ==> Max(DP[j])+1, for all 0<=j<i and arr[i]>arr[j])
// else if no larger element dp[i]=dp[i-1]

// method 2: using patience sorting

// method 3: using segment trees
