package practice

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	arr := []int{2, 4, -6, 7, 3, 1}
	// arr := []int{7, 7, 7, 7, 7}
	var outIdx []int
	var tg int = 7
	outIdx = TwoSum(arr, tg)
	fmt.Println(arr)
	fmt.Println(outIdx)
}

func TestTwoSumOpt(t *testing.T) {
	arr := []int{2, 4, -6, 7, 3, 1}
	// arr := []int{7, 7, 7, 7, 7}
	var outIdx []int
	var tg int = 11
	outIdx = TwoSumOpt(arr, tg)
	fmt.Println(arr)
	fmt.Println(outIdx)
}
