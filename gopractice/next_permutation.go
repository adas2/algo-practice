package practice

import (
	"fmt"
)

// LC 31

func nextPermutation(nums []int) {

	// find the first low value
	index := -1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			index = i - 1
			break
		}
	}

	if index >= 0 {
		// find the next highest value index
		j := len(nums) - 1
		for nums[j] <= nums[index] {
			j--
		}
		// swap
		nums[index], nums[j] = nums[j], nums[index]
	}

	fmt.Println("After swap", nums)

	// now reverse the order index+1..n
	i, j := index+1, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

}

// logic: tricky
// find the first a[i] from the end such that a[i-1] < a[i]
// from index i..n-1, find next largest index j such that a[j] > a[i-1], i.e. a[i-1] > a[j+1], but a[i-1] < a[j-1]
// swap(a[j], a[i-1])--> this fixes the index i-1
// now sort the indices i..n; but since i...n were already in descending order and swapping doesn't break that order
// so e simply reverse the indices i...n to get a sorted order
