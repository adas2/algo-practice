package practice

import (
	"fmt"
	"testing"
)

func TestWiggleSort(t *testing.T) {
	nums := []int{1, 3, 2, 2, 1, 3}

	wiggleSort(nums)
	fmt.Println(nums)
}
