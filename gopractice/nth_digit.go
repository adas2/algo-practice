package practice

import (
	"fmt"
	"math"
)

// LC #400

// find nth digit in infinite sequence: 1234567891011...
func findNthDigit(n int) int {

	// find the correct region by substracting the ranges in increasing order
	remaining := float64(n)
	// num digits in current region i
	nRange, dRange := 9.0, 9.0
	i := 1.0 // region number
	for remaining > dRange {
		fmt.Println(dRange, nRange)
		remaining -= dRange
		i++
		nRange = 9.0 * math.Pow(10, i-1)
		dRange = i * nRange
	}

	fmt.Println("digit region:", i)

	// now find index and digit offset in the  corect region with i digits
	index := (remaining - 1) / i
	offset := int(remaining-1) % int(i)
	fmt.Printf("index: %f, offset %d\n", index, offset)

	// construct the number
	num := int(math.Pow(10, i-1) + index)
	fmt.Println("num", num)

	var d, factor int
	// find the required digit
	for j := 0; j <= offset; j++ {
		i--
		factor = int(math.Pow(10, i))
		d = num / factor
		num = num % factor
	}

	return d
}

// good explanation here:
// 0-9 -> 1 digit 10 digits
// 10-99 -> 2 digit 90*2 = 180 digits
// 100-999 -> 3 digits 900*3 = 2700 digits so on..
// we find the correct range and index within the range
// let digit_r -> num digits for that range
// index-1/digit_r = the number in the range
// index-1%digit_r = the offset of the number
