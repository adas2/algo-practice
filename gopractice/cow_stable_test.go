package practice

import (
	"fmt"
	"testing"
)

func TestFindMaxSeparation(t *testing.T){
	k := 3
	arr := []int{32, 46, 128,200,300}

	s := FindMaxSeparation(arr, k)

	fmt.Println("max sep", s)
}