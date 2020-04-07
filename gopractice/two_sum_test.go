package practice

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	arr := []int{2, 4, -6, 7, 3, 1}
	// arr := []int{7, 7, 7, 7, 7}
	var out_index []int
	var tg int = 7
	out_index = TwoSum(arr, tg)
	fmt.Println(arr)
	fmt.Println(out_index)
}

func TestTwoSumOpt(t *testing.T) {
	arr := []int{2, 4, -6, 7, 3, 1}
	// arr := []int{7, 7, 7, 7, 7}
	var out_index []int
	var tg int = 11
	out_index = TwoSumOpt(arr, tg)
	fmt.Println(arr)
	fmt.Println(out_index)
}
