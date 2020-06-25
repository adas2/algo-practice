package practice

import (
	"fmt"
	"math"
)

// Memory scalability CCI pp: 350

// bitSet is a bit vector for encoding large array of ints
// each int in vector represents 32 bits in the vector
type bitSet struct {
	vector []int32
}

// creates a new bit set based othe number required
func initBitSet(numSize int) *bitSet {
	// creates a vector with vSize
	f := float64(numSize) / float64(32)
	vSize := math.Ceil(f)
	return &bitSet{
		vector: make([]int32, int(vSize)),
	}
}

// add a new number
func (b *bitSet) add(num int) {
	// calcuate the index of the vector array
	idx := uint32(num / 32)
	// calculate offset
	off := uint32(num % 32)
	// update bit in the location
	b.vector[idx] |= (1 << off)
	fmt.Printf("Updating [%d] %d, val: %016b\n", idx, off, b.vector[idx])
}

// check if number is present from bit vector
func (b *bitSet) check(num int) bool {
	// calcuate indx and offset
	idx := uint32(num / 32)
	off := uint32(num % 32)
	if b.vector[idx]&(1<<off) > 0 {
		return true
	}
	return false
}

func findDuplicateNums(arr []int) {
	// create bitSet
	b := initBitSet(len(arr))

	// update bSet
	for _, v := range arr {
		if b.check(v) {
			fmt.Println("duplicate found:", v)
		} else {
			b.add(v)
		}
	}
	// check for
}
