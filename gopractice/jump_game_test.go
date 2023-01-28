package practice

import (
	"fmt"
	"testing"
)

func TestCanJump(t *testing.T) {
	nums := []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}
	fmt.Println("Nums:", nums)
	fmt.Println("Can Jump?: ", canJumpRecursive(nums))
	fmt.Println("Can Jump?: ", canJumpDp(nums))

}
