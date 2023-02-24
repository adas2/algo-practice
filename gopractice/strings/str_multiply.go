package strs

import (
	"fmt"
)

// LC #43

func multiply(num1 string, num2 string) string {
	// corner case
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	r1 := []rune(num1)
	r2 := []rune(num2)
	// create digit list
	dig1 := make([]int8, len(r1))
	dig2 := make([]int8, len(r2))

	j := 0
	for i := len(r1) - 1; i >= 0; i-- {
		dig1[j] = int8(r1[i] - '0')
		j++
	}
	j = 0
	for i := len(r2) - 1; i >= 0; i-- {
		dig2[j] = int8(r2[i] - '0')
		j++
	}

	shift := 0
	var partial []int8
	// partial := make([][]int8, len(dig2))
	result := []int8{0}
	for i := range dig2 {
		partial = digMultiply(dig1, dig2[i])
		zeros := make([]int8, shift)
		partial = append(zeros, partial...)
		fmt.Printf("Part%d: %v\n", shift, partial)
		result = digAdd(result, partial)
		shift++
	}

	// reconstruct rune from []int8
	resStr := []rune{}
	// k := 0
	for j := len(result) - 1; j >= 0; j-- {
		resStr = append(resStr, rune(result[j]+'0'))
		// k++
	}

	return string(resStr)
}

// in reverse
func digMultiply(digits []int8, num int8) []int8 {
	// product for each digit can only be max 2 digits len (max 9x9=81)
	res := make([]int8, len(digits))
	// start from right to left
	var carry int8 = 0
	for i := range digits {
		prod := (digits[i] * num)
		res[i] = (prod + carry) % 10
		carry = (prod + carry) / 10
	}
	if carry > 0 {
		res = append(res, carry)
	}
	return res
}

func digAdd(digit1 []int8, digit2 []int8) []int8 {
	var carry int8 = 0
	res := []int8{}
	if len(digit1) > len(digit2) {
		digit1, digit2 = digit2, digit1
	}
	var i int
	for i = range digit1 {
		res = append(res, (digit1[i]+digit2[i]+carry)%10)
		carry = (digit1[i] + digit2[i] + carry) / 10
	}
	// remaining digits
	for j := i + 1; j < len(digit2); j++ {
		res = append(res, (digit2[j]+carry)%10)
		carry = (digit2[j] + carry) / 10
	}
	if carry > 0 {
		res = append(res, carry)
	}

	return res
}

// elegant soln from LC go impl
func multiply2(num1 string, num2 string) string {

	n1 := len(num1)
	n2 := len(num2)

	r := make([]byte, n1+n2)

	for i := 0; i < len(r); i++ {
		r[i] = '0'
	}

	for i := n1 - 1; i >= 0; i-- {
		for j := n2 - 1; j >= 0; j-- {
			prod := (num1[i] - '0') * (num2[j] - '0')
			sum := r[i+j+1] - '0' + prod
			r[i+j+1] = sum%10 + '0'
			r[i+j] = sum/10 + r[i+j]
		}
	}

	for len(r) > 1 && r[0] == '0' {
		r = r[1:]
	}
	return string(r)

}

// Logic:
// str mult num1="4" num2="6" result = "24"
// naive: convert str to int, multiply and reconvert back to str
// however limited by IntMax size 19 digits, whereas 0 < num1,num2 <= 200 digits
// so large number multiplication is needed
//   		  123
// 			x 456
// 			------
//            738
// 			 615.
//			492..
// 			------
// 			56088
// play with rune
