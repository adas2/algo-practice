package practice

import (
	"fmt"
	"testing"
)

// "1000000009"
// 9
func TestAddOperators(t *testing.T) {
	num := "1000000009"
	out := addOperators(num, 9)
	fmt.Println(out)
}

func TestEvalExp(t *testing.T) {
	s := "2+3*7"
	fmt.Println(evalExp(s))
}

func TestGenExp(t *testing.T) {
	num := "123"
	tgt := 6
	out := []string{}
	genExp2(num, 0, "", tgt, &out)

	fmt.Println(out)
}
