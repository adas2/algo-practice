package strs

import (
	"math"
)

func myAtoi(s string) int {
	// fmt.Printf("IntMax %d, IntMin %d \n", math.MaxInt32, math.MinInt32)
	r := []rune(s)
	num := 0
	negative := false
	numStart := false
	var overflow bool

	for i := range r {
		// leading whitespace
		if !numStart && r[i] == ' ' {
			continue
		}

		// mark start of num if sign is seen
		if !numStart && (r[i] == '+' || r[i] == '-') {
			numStart = true
			if r[i] == '-' {
				negative = true
			}
			continue
		}

		if !isDigit(r[i]) {
			// end of digit
			break
		}
		if !numStart {
			numStart = true
		}
		digit := int(r[i] - '0')
		// check overflow: IntMax 2147483647, IntMin -2147483648
		if num > 214748364 || (!negative && num == 214748364 && digit > 7) ||
			(negative && num == 214748364 && digit > 8) {
			overflow = true
			break
		}
		num = (num * 10) + digit
	}

	if negative {
		if overflow {
			return math.MinInt32
		}
		return -1 * num
	}

	if overflow {
		return math.MaxInt32
	}
	return num
}

func isDigit(r rune) bool {
	if (r-'0' >= 0) && (r-'0' <= 9) {
		return true
	}
	return false
}
