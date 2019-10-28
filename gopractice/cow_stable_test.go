package practice

import (
	"fmt"
	"testing"
)

func TestFindMaxSeparation(t *testing.T) {
	k := 4
	// arr := []int{32, 46, 128, 200, 300}
	arr := []int{0, 5, 7, 10, 12}
	fmt.Println("Input coordinates:", arr, "k =", k)

	s := FindMaxSeparation(arr, k)

	fmt.Println("max sep", s)
}
