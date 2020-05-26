package practice

import "math"

// given an array find the first missing positive integer
func firstMissingPositive(nums []int) int {
	// edge cases
	if len(nums) == 0 {
		return 1 // error case
	}
	min := math.MaxInt32
	// case when min element is > 1
	for i, v := range nums {
		// record the min value
		if v < min {
			min = v
		}
		// filter out the 0's and -ve elements by replacing MAX_INT
		if v <= 0 {
			nums[i] = math.MaxInt32
		}
	}
	// if min is postive and not 1
	if min > 1 {
		return 1
	}

	// all other cases
	// 0. add an extra element to index 0 to incorporate all indices
	nums = append([]int{math.MaxInt32}, nums...)
	// 1. if value < len(nums) check the sign for nums[value]
	// if not -ve flip the sign
	for _, v := range nums {
		// note: to avoid skipping values whose signs were changed
		index := Abs(v)
		if index < len(nums) && nums[index] > 0 {
			nums[index] = -1 * nums[index]
		}
	}
	// 2. first element which is positve
	// skip index zero
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			return i
		}
	}
	// case when missing element exceeds size if array
	return len(nums)
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*
// space inefficient approach:
find the range of integers in the array
create a bit array/hash_map for the entire range where key is integer value
map[key=int value] = true is present, false else
traverse map to find out the first int missing
space = O(K) and Time = O(N+K) where K is the range of ints in the array
inefficient when K >> N; e.g. N=10 K=100000000

// in place array manipulation
intead of a separate array same array is used with some filtering logic
in-place entry is used by changing the sign of the elements
*/
// TODO: can you make it less than 3 passes?
