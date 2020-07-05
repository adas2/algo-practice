package practice

import (
	"fmt"
	"testing"
)

func Test3sumClosest(t *testing.T) {
	nums := []int{0, 2, 1, -3}
	target := 1

	fmt.Println("Result:", threeSumClosest(nums, target))
}
