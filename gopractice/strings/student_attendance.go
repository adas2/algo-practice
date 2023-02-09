package strs

import (
	"fmt"
)

// LC #552

func checkRecord(n int) int {
	res := 0
	findAttendancePermutations("", 0, n, &res)
	fmt.Println(res)

	res = countStateTransitions(n)
	fmt.Printf("state transitions: %d\n", res)
	return res % 1000000007
}

// using recursive backtracking
func findAttendancePermutations(candidate string, countA, n int, total *int) {
	cLen := len(candidate)
	// check the case where A > 2 or "LLL" present
	if countA > 1 {
		return
	}
	if cLen >= 3 && string(candidate[cLen-3:]) == "LLL" {
		return
	}

	// all cases that can be rewarded ends here
	if cLen == n {
		fmt.Printf("%s ", candidate)
		*total++
		return
	}

	rec := []byte{'A', 'L', 'P'}
	// form candidate
	for _, r := range rec {
		candidate += string(r)
		if r == 'A' {
			countA++
		}
		findAttendancePermutations(candidate, countA, n, total)
		// backtrack
		if r == 'A' {
			countA--
		}
		candidate = string(candidate[:len(candidate)-1])

	}
}

// this is recursive solution with complexity O(3^n)
// dp version faster
// state transition count is most optimized in space and time

//  	 L	      L        L
// A0L0 --> A0L1 --> A0L2 --> Invalid

// A1L0 --> A1L1 --> A1L2 --> Invalid
func countStateTransitions(n int) int {
	// good state
	m := 1000000007
	A0L0 := 1
	A0L1, A0L2, A1L0, A1L1, A1L2 := 0, 0, 0, 0, 0

	// count transitions for each letter
	for i := 0; i < n; i++ {
		new_A0L0 := A0L0 + A0L1 + A0L2
		new_A0L1 := A0L0
		new_A0L2 := A0L1
		new_A1L0 := A0L0 + A0L1 + A0L2 + A1L1 + A1L2 + A1L0
		new_A1L1 := A1L0
		new_A1L2 := A1L1

		A0L0 = new_A0L0
		A0L1 = new_A0L1
		A0L2 = new_A0L2
		A1L0 = new_A1L0
		A1L1 = new_A1L1
		A1L2 = new_A1L2
	}
	res := A0L0 + A0L1 + A0L2 + A1L0 + A1L1 + A1L2

	return res % m
}
