package practice

import "fmt"

// Given an array of similar objects (small, medium, large) i.e. S, M, L, sort the objects in order
// E.g. Input := {L,M,S,L,L,S,M,M} Ouptut := {S,S,M,M,M,L,L,L}
// Hint: try to do this in linear time
// Bonus: Is single pass possible?

// Solution logic: (2 pass algo)
// First pass: move S elements in their proper place, i.e. left side of array
// Second pass: move the M elements in their propoer place, i.e. middle of the array
// Large elements get automatically sorted

// In first pass: two pointers from two ends, left: first occurence of non-S object
// right: first occurence of S-object
// swap left and right, left++, right--

// SortSimilarItems sorts
func SortSimilarItems(items []string) []string {

	var left, right int
	left, right = 0, len(items)-1
	// 1 pass: scan the array for S items
	for left <= right {
		// move the first index which is non-S
		for items[left] == "S" {
			left++
		}
		// move to the last index which is S
		for items[right] != "S" {
			right--
		}
		// now swap if indices are valid
		if left < len(items) && right >= 0 && left < right {
			// swapItems(items, left, right)
			temp := items[left]
			items[left] = items[right]
			items[right] = temp

			// adjust indices
			left++
			right--
			continue
		}

	}
	fmt.Println(items)

	// 2 pass: scan for M ietms
	left, right = 0, len(items)-1
	for left <= right {
		// move the first index which is non-S
		for items[left] != "L" {
			left++
		}
		// move to the last index which is S
		for items[right] != "M" {
			right--
		}
		// now swap if indices are valid
		if left < len(items) && right >= 0 && left < right {
			// swapItems(items, left, right)
			temp := items[left]
			items[left] = items[right]
			items[right] = temp

			// adjust indices
			left++
			right--
			continue
		}

	}
	fmt.Println(items)

	return items
}

// Single pass algorithm:
// sort only the S objects to the left and L objects to the right
// M will be automatically sorted
