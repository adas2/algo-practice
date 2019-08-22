package practice

// Given an array of integers, find the indices of the elements which add up to a given value
// e.g. {2, 4, -6, 3, 5} Target = 7; O/p: {1,3}

// TwoSum problem solution
func TwoSum(arr []int, target int) []int {
	// return invalid indices for nil array
	if len(arr) == 0 {
		return []int{-1, -1}
	}

	// store in a hash-table
	htable := make(map[int]int)

	// traverse array and check if dff = (target-element) is present in memory hash
	for index, elem := range arr {
		// if element present in the hash return the indices
		// else make an entry in the hash
		if val, isPresent := htable[elem]; isPresent {
			return []int{val, index}
		} else {
			htable[(target - elem)] = index
		}
	}
	// if not found
	return []int{-1, -1}
}
