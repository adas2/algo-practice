package practice

// BinarySearchVanilla finds target within int arrary
func BinarySearchVanilla(arr []int, target int, low int, high int) int {
	if low > high{
		return -1 //error case 
	}

	// begin iteration in while style in C
	for low <= high {
		mid := low+(high-low)/2
		guess := arr[mid]
		if guess == target{
			return mid
		} else if guess < target {
			low = mid+1
		} else { // gues > target
			high = mid-1
		}
	}
	// no match found
	return -1
}

// BinarySearchClosest find the index of last closest element <= target
func BinarySearchClosest(arr []int, target int, low int, high int) int {
	if low > high{
		return -1 //error case 
	}
	candidate := -1 //init

	// begin iteration in while style in C
	for low <= high {
		mid := low+(high-low)/2
		guess := arr[mid]

		if guess <= target {
			low = mid+1
			candidate = mid
		} else { // guess > target
			high = mid-1
		}
	}
	// no match found
	return candidate
}

// BinarySearchRepeated find target in array of repeated elements 
// with index of the last occurence
func BinarySearchLast(arr []int, target int, low int, high int) int {

	if low>high{
		return -1 //err case
	}
	candidate := -1 //init err value

	// iterative impl
	for low <= high{
		mid := low+(high-low)/2
		if arr[mid] == target {
			// potentital last element
			candidate = mid
		}
		if arr[mid] <= target{ // even if a match search right half for better candidate
			low = mid+1
		}else{ // arr[mid] > target
			high = mid-1
		}
	}
	return candidate
}

