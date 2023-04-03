package strs

import (
	"strings"
	"unicode"
)

// LC # 394 (medium --> hard IMO)

// Using recursive implementaion with memory/stack S=O(n)
func decodeString(s string) string {

	// stack to monitor open paran '[' with the count associated
	open := make([]int, 0)
	// shared index passed by reference
	index := 0

	return decodeUtil(s, &open, &index)
}

// recursive util func
func decodeUtil(s string, open *[]int, index *int) string {

	r := []rune(s)
	res := ""

	for *index < len(r) {
		k := 0
		for *index < len(s) && unicode.IsDigit(r[*index]) == true {
			k = (k * 10) + int(r[*index]-'0')
			*index += 1
		}

		if r[*index] == '[' {
			// start of a new repeated seq, recursively call to get output
			// s[i-1] is the cnt push to stack
			*open = append(*open, k)
			// fmt.Printf("start [ slen %d res %s\n", len(*open), res)
			*index += 1
			res += decodeUtil(s, open, index)
		} else if r[*index] == ']' {
			// pop from stack
			slen := len(*open)
			cnt := (*open)[slen-1]
			*open = (*open)[:slen-1]
			*index += 1
			// repeat str cnt times
			res = strings.Repeat(res, cnt)
			// fmt.Printf("end ] slen %d cnt: %d, res: %s\n", slen, cnt, res)
			return res
		} else if unicode.IsDigit(r[*index]) == false {
			// non-digit char
			res += string(r[*index])
			*index += 1
		}
	}

	return res
}

// logic:
// use a recursive solution, when we see a '[' we push the k val onto stack,
// when we see ']' we pop that corresponding k val and decode the substring seen so far
// we terminate when all the str is read
// Note: we pass around the index by reference otherwise the same index will be covered multiple times (each time for a new recusion )
// Same for the stack, ptr to the stack is passed since we want to keep track of the porcessed '[' across all recursions
// Optional optimization: get rid of the stack, but using the k value before the recursive call itself
