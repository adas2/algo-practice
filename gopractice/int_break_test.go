package practice

import (
	"fmt"
	"testing"
)

func TestIntBreak(t *testing.T) {
	for i := 2; i <= 10; i++ {

		fmt.Println("Num::", i)
		res := integerBreak2(i)
		fmt.Println(res)
	}
}
