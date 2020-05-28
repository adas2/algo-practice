package trees

import (
	"fmt"
	"testing"
)

func TestBit(t *testing.T) {
	nums := []int{3, 5, 0, 6, 7, 8, 1, -3}

	bit := initBIT(nums)

	fmt.Println("BItree:", bit)

	fmt.Println("Prefix Sum:", GetSum(bit, 3))
	Update(bit, 2, 3)
	fmt.Println("New Sum:", GetSum(bit, 3))
	fmt.Println("RangeSum:", RangeSum(bit, 1, 5))
	Update(bit, 7, 10)
	fmt.Println("Updated RangeSum:", RangeSum(bit, 6, 7))
}
