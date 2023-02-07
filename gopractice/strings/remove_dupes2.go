package strs

import (
	"bytes"
)

// LC #1209 Medium

// stack element
type nCnt struct {
	alpha byte
	cntr  int
}

func removeDuplicates(s string, k int) string {

	// stack to keep track of the <letter, count> seen in the order
	stk := []nCnt{}

	// traverse the str
	for i := range s {
		// letter seen before and top of stack, incr cnt only
		if len(stk) > 0 && stk[len(stk)-1].alpha == s[i] {

			stk[len(stk)-1].cntr++
		} else {
			// if seen before but not on stack top, add again init cnt = 1
			stk = append(stk, nCnt{s[i], 1})
		}

		// if the stk_top cnt == k pop from the stack
		if len(stk) > 0 && stk[len(stk)-1].cntr == k {
			stk = stk[:len(stk)-1]
		}

	}

	// construct result str from the remaining letters in stack and cnt values
	res := []byte{}
	for _, val := range stk {
		res = append(res, bytes.Repeat([]byte{val.alpha}, val.cntr)...)
	}

	return string(res)
}

// Logic
// Input: s = "deeedbbcccbdaa", k = 3
// Output: "aa"
// keep a map of char to total seen count
// keep a stack with the last letter seen on the top
// do as following:
// 		when the count == 3 for stack-top pop()
// 		else add to the top of stack and increase count for the letter
// resultant stack will have final string
