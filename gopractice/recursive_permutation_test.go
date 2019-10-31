package practice

import (
	"fmt"
	"testing"
)

// func TestRecursivePermutations(t *testing.T) {

// 	s := "abcd"

// 	RecursivePermutations(s, 0, 3)

// }

func TestFindPermutaionsMain(t *testing.T) {

	s := "abc"
	output := FindPermutaionsMain(s)

	fmt.Println(output, len(output))
}
