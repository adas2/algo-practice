package practice

import (
	"fmt"
	"testing"
)

func TestMaxRect(t *testing.T) {
	m := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}

	fmt.Printf("Max rectangle: %d\n", maximalRectangle(m))
}
