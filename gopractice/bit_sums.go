package practice

import (
	"fmt"
	"math/bits"
)

// utility funcs for Bitwise addtions

type BitNum struct {
	// slice with bits from LSB to MSB (reverse order)
	bitSlice []uint
}

func NewBitNum() BitNum {
	return BitNum{bitSlice: []uint{}}
}

// Fix this
func AddBits(a, b BitNum, carry uint) (BitNum, uint) {
	var maxLen int
	// way of finding max
	if maxLen = len(b.bitSlice); len(a.bitSlice) > len(b.bitSlice) {
		maxLen = len(a.bitSlice)
	}

	// out := make([]uint, maxLen+1) //space for carry
	out := NewBitNum()
	var sum, bit uint
	for i := 0; i < maxLen; i++ {
		if i > len(a.bitSlice)-1 {
			sum, carry = bits.Add(0, b.bitSlice[i], carry)
		} else if i > len(b.bitSlice)-1 {
			sum, carry = bits.Add(a.bitSlice[i], 0, carry)
		} else {
			sum, carry = bits.Add(a.bitSlice[i], b.bitSlice[i], carry)
		}
		bit = sum % 2
		carry = sum / 2

		// out = AppenBit(out, bit)
		fmt.Println(bit)
	}
	return out, carry
}
