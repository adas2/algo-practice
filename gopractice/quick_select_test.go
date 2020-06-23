package practice

import (
	"fmt"
	"testing"
)

func TestQuickSelect(t *testing.T) {
	v := []int{3, 7, 19, -2, 5, 10, 343}

	fmt.Println("array:", v)
	res, val := quickSelect(&v, 0, len(v)-1)
	fmt.Printf("Chosen pivot: %d with val: %v\n", res, val)

	fmt.Println("final array:", v)
}

func TestFindkthVal(t *testing.T) {
	v := []int{3, 7, 19, -2, 5, 10, 343}
	k := 2
	fmt.Printf("K: %d value: %d\n", k, findkthValue(v, 0, len(v)-1, k))
}
