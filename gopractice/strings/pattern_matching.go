package strs

import (
	"fmt"
	"math"
)

// given a test T and a pattern P, find all the occurences of P in T.
// E.g. T = "ABABABACBABA", P = "BABA" Output: {1,3,8}. Overlaps are allowed
// Expected best case complexity O(m+n), worst case is O(mn)
// If no match return empty array

const prime = 101

func findPattern(t, p string) []int {

	if len(t) == 0 || len(p) == 0 {
		return []int{}
	}
	result := make([]int, 0, len(t)) //optional: defining max capacity

	n := len(t)
	m := len(p)
	// precompute hashes: O(m)
	pH := computeHash(p, 0, m-1)
	tH := computeHash(t, 0, m-1)

	// iterate over the string: O(n)
	for i := 0; i < n-m; i++ {
		// fmt.Printf("char: %d\n", i)
		// fmt.Printf("text hash: %v\n", tH)
		// fmt.Printf("pattern hash: %v\n", pH)
		if pH == tH {
			// potential match
			fmt.Printf("Potential match for substr: %s at index: %d\n", t[i:i+m], i)
			// verify full match
			if verifyMatch(t, p, i, i+m-1) {
				result = append(result, i)
			}
			// record index
		}
		// shift window by 1 char and update hash
		tH = updateHash(tH, t, i+1, i+m)
	}

	return result
}

// util: given a string calculate hash
/**
H(0…k-1) = (c0*a^k-1 + c1*a^k-2 + c2*ak-3 …+… ck-1*a^0) mod b
for next char:
H(1…k)   = (c1*a^k-1 + c2*a^k-2 …+… ck-1*a^1 + ck*a^0) mod b
i.e.,
H(1…k)   = (a * H(0…k-1) - c0*a^k + ck) mod b
precompute m=a^k
H(1…k)   = (a * H(0…k-1) - m * c0 + ck) mod b
**/
func computeHash(str string, start, end int) int {
	hash := 0
	// len
	k := end - start + 1
	// num chars in alphabet
	d := 256

	for i := 1; i <= k; i++ {
		// modular arithmetic
		hash += (int(str[i-1]) * int(math.Pow(float64(d), float64(k-i)))) % prime
	}

	return hash % prime
}

// rolling hash, update existing hash with new character
// start, end new indices for hashing
func updateHash(oldHash int, s string, start, end int) int {
	// precompute
	// fmt.Printf("old hash: %v\n", oldHash)
	k := end - start + 1
	d := 256
	v := int(math.Pow(float64(d), float64(k))) % prime
	newHash := (oldHash * d) - (v * int(s[start-1])) + int(s[end])

	return newHash % prime
}

// verify hash match
func verifyMatch(t, p string, start, end int) bool {
	j := 0
	for i := start; i <= end; i++ {
		if p[j] != t[i] {
			return false
		}
		j++
	}
	return true
}

// Algo:
// compute hash over pattern
// compute a rolling hash over a sliding window of m = len(pattern)
// check if the values match, match every char, record location
// shift by 1 and repeat till end
