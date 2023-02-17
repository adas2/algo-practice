package strs

import "fmt"

// LC # 418 Fitting sentences on given screen

func wordsTyping(sentence []string, rows int, cols int) int {
	cnt := 0
	// len array
	wlen := make([]int, len(sentence))
	for i := range sentence {
		wlen[i] = len(sentence[i])
	}

	idx := 0 //index of word
	r, c := rows, cols
	for r > 0 {
		for c >= wlen[idx] {
			c -= wlen[idx]
			idx++
			c-- // space after each word
			fmt.Printf("idx: %d, c: %d\n", idx, c)
			if idx == len(sentence) { //sentence complete
				cnt++
				idx = 0 //reset index
			}
		}
		// next row
		r--
		c = cols
	}

	return cnt
}

// Logic: example
// for each word to fit l[i] = len(w[i]) + 1
// while rows > 0
// 		cols -= l[i] while cols > 0
// 			i++
// 			if i == len(sentence) cnt++; C-- (extra space)
// 		rows--
// return cnt
// E.g.:
// sentence = ["a", "bcd", "e"], rows = 3, cols = 6
// Output: 2
// Explanation:
// a-bcd-
// e-a---
// bcd-e-
// T = O(max(r*c, N)) where N is number of words
