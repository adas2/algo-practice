package strs

import "fmt"

// LC #336

// hash table impl
// total complexity: O(n.l^2)
func palindromePairs(words []string) [][]int {
	//  store all the words in a hash table for fast lookup, keep index
	wMap := map[string]int{}
	for i, w := range words {
		wMap[w] = i
	}
	result := [][]int{}

	// now for each word find the all suffixes and prefixes which are palindromes
	// O(n.l^2) where l = avg_len(word) n is total # words
	// this can be improved with Manacher's algo?
	for idx, w := range words {
		// Case 1: when reverse of word is in map
		// O(l)
		if val, exists := wMap[reverseWord(w)]; exists {
			// note in this case single letter words can match with itself
			// (not allowed since given words are unique)
			if val != idx {
				fmt.Printf("found case 1: %s, %s in list\n", w, reverseWord(w))
				result = append(result, []int{idx, val})
			}
		}
		// Case 2: prefix is a palindrome (ignore full word prefix captured by previous case)
		// O(l^2)
		for i := 0; i <= len(w)-1; i++ {
			if checkPalindrome(w, 0, i) == true {
				// search for remaining words reverse
				rem := string(w[i+1:])
				if val, exists := wMap[reverseWord(rem)]; exists {
					fmt.Printf("found case 2: %s, %s in list\n", w, reverseWord(rem))
					result = append(result, []int{val, idx})
				}
			}
		}
		// Case 3: suffix is a palindrome
		// O(l^2)
		for j := len(w) - 1; j >= 0; j-- {
			if checkPalindrome(w, j, len(w)-1) == true {
				// search for remaining words reverse
				rem := string(w[:j])
				if val, exists := wMap[reverseWord(rem)]; exists {
					fmt.Printf("found case 3: %s, %s in list\n", w, reverseWord(rem))
					result = append(result, []int{idx, val})
				}
			}
		}

	}

	return result
}

// util check if a string is a palindrome
func checkPalindrome(word string, start, end int) bool {
	// 2 pointers from two ends
	i := start
	j := end
	for i < j {
		if word[i] != word[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func reverseWord(word string) string {
	r := []rune(word)
	for i, j := 0, len(r)-1; i < j; {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
	return string(r)
}

// Note this is better than bruteforce only is n >> l i.e. avg len of word is small compared to num words
// brute force is O(n^2.l)
