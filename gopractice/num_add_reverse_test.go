package practice

import "testing"

func TestCreateNum(t *testing.T) {
	var n1, n2 int = 999, 2
	num1 := createNum(n1)
	num2 := createNum(n2)
	printNum(num1)
	printNum(num2)
	printNum(numAdd(num1, num2))
}
