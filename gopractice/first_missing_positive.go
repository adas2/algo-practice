package practice

import "math"

// LC # 41 (hard)
// given an unsorted array find the first missing positive integer
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
*/

// Better space optimization; in-place array manipulation
// intead of a separate array same array is used with some filtering logic
// in-place entry is used by changing the sign of the elements.

// Trick is to do the following (i) ignore all the negative and zero values,
// 	(ii) within the positive values in the array, the answer can be either <= (n=len(arr))
// 	or n+1 if all (1-n) values are present.
// 	(iii) for the case when ans < n; we have to keep track of the positive values using indices
// we will use the 0th index for value "n" number, for other value we will use the abs(value) as index

// Slightly cleaner impl based on LC editorial
func firstMissingPositive2(nums []int) int {
	// preprocess the nums to ignore 1 and -ve values
	onePresent := false
	for i, v := range nums {
		if v <= 0 {
			nums[i] = 1
		}

		if v == 1 {
			onePresent = true
		}
	}

	// good case
	if !onePresent {
		return 1
	}

	// use the abs(value) as hash keys/index to keep track of the value
	for _, val := range nums {

		var index int
		// can use math.Abs too but convert to float
		if val < 0 {
			index = -val
		} else {
			index = val
		}

		// if the val at index is -ve val that means it has been already tracked
		// else change to -ve
		if index < len(nums) && nums[index] > 0 {
			nums[index] = -nums[index]
		}
		// use the 0th index for value n
		if index == len(nums) {
			nums[0] = -nums[0]
		}
	}

	// now we will check in order (1..n-1,0) first index which is +ve is out answer
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			return i
		}
	}

	if nums[0] > 0 {
		return len(nums)
	}

	// in case all 1...n values are in num, the next +ve is n+1
	return len(nums) + 1

}
