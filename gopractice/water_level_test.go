package practice

import (
	"fmt"
	"testing"
)

func TestTrap(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	fmt.Println("Water trapped:", Trap(height))
}
