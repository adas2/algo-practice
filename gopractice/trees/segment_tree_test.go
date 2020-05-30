package trees

import (
	"fmt"
	"testing"
)

func TestSegmentTree(t *testing.T) {
	// array
	arr := []int{3, 6, -2, 73, 231, 13, 7}
	sTr := initSegmentTree(arr)
	fmt.Println("seg tree", sTr)

	fmt.Println("query", sTr.query(arr, 1, 3))
	fmt.Println("update", sTr.update(arr, 3, 101))
	fmt.Println("query", sTr.query(arr, 1, 3))
	fmt.Println("update", sTr.update(arr, 6, 50))
	fmt.Println("query", sTr.query(arr, 5, 6))
	// corner cases
	fmt.Println("query", sTr.query(arr, 0, 6))
	fmt.Println("query", sTr.query(arr, 6, 6))
}
