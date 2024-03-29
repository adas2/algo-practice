package practice

import (
	"math/rand"
	"time"
)

// find the kth smallest value in an unsorted array using randomized quick select

// return the index and value
func quickSelect(arr *[]int, l, r int) (int, int) {

	if l > r {
		return -1, -1
	}

	// randomly select an pivot between [0, len(arr)]
	rand.Seed(time.Now().UnixNano())
	pivot := rand.Intn(r - l + 1)
	// place the pivot to end of arr (swap)
	swap(arr, pivot, r)

	// variables
	i := l          // first index larger than pivot, init to 0th index
	pv := (*arr)[r] //value of pivot

	// now iterate to place pivot in the right location
	for j := l; j < r; j++ {
		// move element less than pivot to the left of the first large index
		if (*arr)[j] <= pv {
			swap(arr, j, i)
			i++ // next larger elem index incremented
		}
	}
	// once all elements covered swap pivot with next larger elem
	swap(arr, r, i)

	// return the new pivot location
	return i, (*arr)[i]
}

// util
func swap(arr *[]int, i, j int) {
	if i == j {
		return
	}
	temp := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = temp
}

// find kth smallest value
// l,r: left and right index of array
// k: index to find
func findkthValue(arr []int, l, r, k int) int {
	// quick select on random partition
	// if pivot == k, return answer
	// if pivot < k, try (k-pivot) in arr[pivot+1, n]
	// else pivot > k, try k in arr[0, pivot-1]

	// Todo: fix base case
	// if r-l+1 <= 0 || k < 0 || k > r-l+1 {
	// 	fmt.Println("Error", l, r, k)
	// 	return math.MaxInt32
	// }

	// randomized quick select partition
	pivot, val := quickSelect(&arr, l, r)

	if pivot == k {
		return val
	} else if pivot > k {
		return findkthValue(arr, l, pivot-1, k)
	}
	// else, adjust the k value
	return findkthValue(arr, pivot+1, r, k-pivot-1)

}

// Time complexity : O(N) in the average case, O(N^2)in the worst case.
// Space complexity : O(1)
