package practice

import "testing"

func TestCreateNum(t *testing.T) {
	var n1, n2 int = 999, 2
	num1 := CreateNum(n1)
	num2 := CreateNum(n2)
	PrintNum(num1)
	PrintNum(num2)
	PrintNum(NumAdd(num1, num2))
}
