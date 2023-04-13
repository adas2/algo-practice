package custom

import (
	"hash"
	"hash/fnv"
	"math"
)

type BloomFilter struct {
	bitset []bool      // m : bit in the array
	hashes uint        // k : num hashes used
	size   uint        // n : num inserted elements
	f      hash.Hash32 // h : hash val computed
}

func NewBloomFilter(size uint, hashes uint) *BloomFilter {
	return &BloomFilter{
		bitset: make([]bool, size),
		hashes: hashes,
		size:   size,
		f:      fnv.New32(),
	}
}

func (bf *BloomFilter) Add(item string) {
	for i := uint(0); i < bf.hashes; i++ {
		bf.f.Write([]byte(item))
		index := bf.f.Sum32() % uint32(bf.size)
		bf.bitset[index] = true
		bf.f.Reset()
	}
}

func (bf *BloomFilter) Contains(item string) bool {
	for i := uint(0); i < bf.hashes; i++ {
		bf.f.Write([]byte(item))
		index := bf.f.Sum32() % uint32(bf.size)
		if !bf.bitset[index] {
			return false
		}
		bf.f.Reset()
	}
	return true
}

// for detailed analysis check Wikipedia link:
// https://en.wikipedia.org/wiki/Bloom_filter#Probability_of_false_positives
// False positive rate: (1-e^(-kn/m))^k
func (bf *BloomFilter) FalsePositiveProbability() float64 {
	return math.Pow(1-math.Exp(-float64(bf.hashes)*float64(len(bf.bitset))/float64(bf.size)), float64(bf.hashes))
}

// Note: Above code is ChatGPT generated (do not use for production)
