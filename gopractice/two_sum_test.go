package practice

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	arr := []int{2, 4, -6, 3, 5, 1}
	var out_index []int
	var tg int = 7
	out_index = TwoSum(arr, tg)
	fmt.Println(arr)
	fmt.Println(out_index)

}
