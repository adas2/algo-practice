package practice

import "fmt"

// Find median of two sorted arrays (of uequal sizes)
// for equal size array median, simpler algo exits (hint: merge sort merge op)
// complexity: O(log(min(n, m)))

func findSortedMedian(arr1, arr2 []int) float32 {

	var even bool
	// case 1: odd length
	if (len(arr1)+len(arr2))%2 != 0 {
		even = false
		fmt.Printf("Median is the %d index element\n", (len(arr1)+len(arr2))/2)
	} else {
		even = true
		fmt.Printf("Median is the %d and %d index element\n", (len(arr1)+len(arr2))/2-1, (len(arr1)+len(arr2))/2)
	}

	var minIdx, maxIdx, i, j int
	var median float32
	// choose the smaller array
	if len(arr1) > len(arr2) {
		temp := arr1
		arr1 = arr2
		arr2 = temp
	}

	fmt.Println(arr1)
	fmt.Println(arr2)

	// start from the start
	minIdx = 0
	maxIdx = len(arr1)
	// divide arr1 and arr2 into two halves with roughly same num elements
	// s.t. all elements in left half < elements in right half
	// i.e. elems in arr1[:i], arr2[:j] are less than arr1[i:], arr2[j:]
	// terminate when condition is met
	for minIdx <= maxIdx {
		// choose the mid index in arr1
		i = (minIdx + maxIdx) / 2
		// adjust index in arr2 according
		j = (len(arr1)+len(arr2)+1)/2 - i
		fmt.Println(minIdx, maxIdx, i, j)

		if j > 0 && i < len(arr1) && arr2[j-1] > arr1[i] {
			// move right in arr1,search space reduced
			minIdx = i + 1
		} else if i > 0 && j < len(arr2) && arr1[i-1] > arr2[j] {
			// move left in arr1, reduced search space
			maxIdx = i - 1
		} else {
			// desired cond is reached,
			// i.e. arr1[i-1] < arr2[j] && arr1[j-1] > arr2[i]

			// two cases even and odd
			// within each cosider when:
			// left half of arr1 is nil,
			// left half of arr2 is nil
			if i == 0 {
				median = float32(arr2[j-1])
			} else if j == 0 {
				median = float32(arr1[i-1])
			} else {
				median = float32(max(arr1[i-1], arr2[j-1]))
			}
			break
		}
	}

	// evaluate other half in necessary (even case)
	// i.e. min of the elements in right half
	if even {
		// consider
		// right half of arr1 is nil
		// right half of arr2 is nil
		if i == len(arr1) {
			median = (median + float32(arr2[j])) / 2
		} else if j == len(arr2) {
			median = (median + float32(arr1[i])) / 2
		} else {
			// normal case
			median = (median + float32(min(arr1[i], arr2[j]))) / 2
		}
	}

	return median
}

// util funcs
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
