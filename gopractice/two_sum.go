package practice

import (
	"sort"
)

// Given an array of integers, find the indices of the elements which add up to a given value
// e.g. {2, 4, -6, 3, 5} Target = 7; O/p: {1,3}

// TwoSum problem solution
// memory unoptmized: T=O(N) S=O(N)
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

// memory optimized version
// outputs the value not index
// S=O(1), T=O(NlogN)
func TwoSumOpt(arr []int, target int) []int {
	// return invalid indices for nil array
	if len(arr) == 0 {
		return []int{-1, -1}
	}

	// sort the input arr
	sort.Ints(arr)
	// fmt.Println(arr)

	result := []int{-1, -1}
	// user two pointers front and back and adjust them to match the target
	for front, back := 0, len(arr)-1; front < back; {
		if arr[front]+arr[back] == target {
			result[0] = arr[front]
			result[1] = arr[back]
			break
		} else if arr[front]+arr[back] > target {
			back--
		} else { // front+back < target
			front++
		}
	}
	return result
}
