package trees

// Description:
// 1.define binary indexed trees data structure
// 2. Update(arr, x, val) implement with O(logN) time
// 3. Sum(arr, x, y) implement in O(logN) time

// BIT global struct
// type BIT []int

// init BIT initializes a binary indexed tree from a given array
func initBIT(arr []int) []int {
	// deaclare the struct: extra element to incoporate the dummt node at index 0
	biTree := make([]int, len(arr)+1)
	// biTree := &BIT{}

	// call update N times to include each element of arr
	// time = O(NlogN)
	for index, value := range arr {
		// fmt.Println("Updating index", index, value)
		Update(biTree, index, value)
	}

	return biTree
}

// Update (arr, x, val): updates value at arr[x] to += val
// time = O(logN)
// Note actual array is not updated
func Update(biTree []int, x, val int) {
	// adjust index for biTree
	x = x + 1
	for x < len(biTree) {
		biTree[x] += val
		x = x + (x & -x)
	}
}

// GetSum computes the prefix sum from index 0 to x, i.e. first x+1 elements
// time = O(logN)
func GetSum(biTree []int, x int) int {
	sum := 0
	// adjust index for biTree
	x = x + 1
	for x > 0 {
		sum += biTree[x]
		x = x - (x & -x)
	}
	return sum
}

// RangeSum calculates sum of array from indices x to y i.e Sum(arr[x...y]) both inclusive
// i.e. y-x+1 elements
func RangeSum(biTree []int, x, y int) int {
	// range query for arr[x..y] both inclusive
	return (GetSum(biTree, y) - GetSum(biTree, x-1))
}
