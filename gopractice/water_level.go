// leetcode #42 (hard)
package practice

func Trap(height []int) int {
	// first pass find the left max for each index i
	max_left := make([]int, len(height))
	max_right := make([]int, len(height))
	curr_left_max, curr_right_max := -1, -1
	res := 0
	for i, j := 0, len(height)-1; i < len(height); {
		if curr_left_max < height[i] {
			curr_left_max = height[i]

		}
		max_left[i] = max(height[i], curr_left_max)

		if curr_right_max < height[j] {
			curr_right_max = height[j]

		}
		max_right[j] = max(height[j], curr_right_max)
		i++
		j--
	}

	// fmt.Println("Max left height:", max_left)
	// fmt.Println("Max right height:", max_right)

	// traverse height array to calculate water trapped
	for index, val := range height {
		min_level := min(max_left[index], max_right[index])
		if val < min_level {
			res += min_level - val
		}
	}
	return res
}

// logic: Multi-pass T=O(N), S=O(N) solution (optmize memory if possible)
// keep the left max height and right max height for each index
// traverse the array, calculate following:
// if val[i] < min(left_max, right_max) trap_vol = +=min(left_max, right_max)-val[i]

// O(1) space solution hint: use moving pointers to keep track of left_max
// and right_max without storin in array.
