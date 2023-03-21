package practice

// BinarySearchVanilla finds target within sorted arrary
func BinarySearchVanilla(arr []int, target int, low int, high int) int {
	if low > high {
		return -1 //error case
	}

	// begin iteration in while style in C
	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]
		if guess == target {
			return mid
		} else if guess < target {
			low = mid + 1
		} else { // guess > target
			high = mid - 1
		}
	}
	// no match found
	return -1
}

// BinarySearchClosest: index of closest element <= target
func BinarySearchClosest(arr []int, target int, low int, high int) int {
	if low > high {
		return -1 //error case
	}
	candidate := -1 //init

	// begin iteration in while style in C
	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]

		if guess <= target {
			low = mid + 1
			candidate = mid
		} else { // guess > target
			high = mid - 1
		}
	}
	// return best candidate
	return candidate
}

// BinarySearchLast: index of target in array of repeated elements
// or index of the next larger element >= target
func BinarySearchLast(arr []int, target int, low int, high int) int {
	if low > high {
		return -1 //err case
	}
	candidate := -1 //init err value

	// iterative impl
	for low <= high {
		mid := (low + high) / 2

		if arr[mid] >= target { // even if a match search right half for better candidate
			high = mid - 1
			candidate = mid
		} else { // arr[mid] < target
			low = mid + 1
		}
	}
	return candidate //returns -1 if target not present
}
