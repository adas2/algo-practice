package practice

// incorrect impl: using recursive backtracking
func canJumpRecursive(nums []int) bool {
	// edge case:
	if len(nums) == 1 {
		return true
	}

	// termination case
	if nums[0] >= len(nums)-1 {
		return true
	}

	if nums[0] == 0 {
		return false
	}

	// at position i it can choose one of 1 to nums[i]
	for i := 1; i <= nums[0]; i++ {
		if canJumpRecursive(nums[i:]) {
			return true
		}
	}
	return false
}

func canJumpDp(nums []int) bool {

	// create dp array to keep track of canJump
	jump_dp := make([]bool, len(nums))

	// edge case:
	jump_dp[len(nums)-1] = true

	for i := len(nums) - 2; i >= 0; i-- {
		// it can jump as long any of the reachable index can jump
		furthest_jump := min(i+nums[i], len(nums)-1)
		for j := i + 1; j <= furthest_jump; j++ {
			// jump_dp[i] = jump_dp[i] || jump_dp[j]
			if jump_dp[j] {
				jump_dp[i] = true
			}
		}
	}

	return jump_dp[0]
}

// logic:
// from each index find index that can lead to successful jump to last index
// jump_len = max(nums[i], highest_val(nums[i+1..nums[i]]))
// if there is no fwd progress, return false
// if jump_len is >= remaining lenght return true
// DP solution,
// T=O(N^2) S=O(N)

// Memory optimized O(N) solution possible using greedy approach
