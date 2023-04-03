package dp

// LC # 198

func rob(nums []int) int {

	N := len(nums)
	dp := make([]int, N+1)

	// base case
	dp[N-1] = nums[N-1]
	dp[N] = 0

	// fill dp array backwards from end to start
	for i := N - 2; i >= 0; i-- {
		dp[i] = Max(nums[i]+dp[i+2], dp[i+1])
	}

	return dp[0]

}

// logic:
// dp[i] --> max amt robbed for nums[i...N-1] where len(nums) = N
// at each step two options: rob the house at i OR skip the house ;
// choose whichever maximizes the total amt
//  dp[i] = max(nums[i] + dp[i+2], dp[i+1])
// base case: dp[N-1] = nums[N-1] since only option is rob 1 house
// dp[N] = 0 since no houses left

// Note optimized version just needs 2 vars
func rob2(nums []int) int {
	N := len(nums)

	if N == 0 {
		return 0
	}

	// base case
	robNext := nums[N-1]
	robNextPlusOne := 0

	// fill dp array backwards from end to start
	for i := N - 2; i >= 0; i-- {
		curr := Max(nums[i]+robNextPlusOne, robNext)

		robNextPlusOne = robNext
		robNext = curr
	}

	return robNext
}
