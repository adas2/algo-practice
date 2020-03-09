package practice

import (
	"sort"
)

// Leetcode
// 472. Concatenated Words

func findAllConcatenatedWordsInADict(words []string) []string {

	var result []string

	// set of words already seen
	wSet := make(map[string]bool)

	// sort the words according to len
	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })
	minWlen := len(words[0])
	if minWlen == 0 {
		minWlen = 1
	}

	// ietrate over words
	for _, w := range words {

		if w == "" {
			continue
		}

		// for current w find if it matches using alreadys seen words
		if wordMatchHelper(wSet, w, minWlen) {
			result = append(result, w)
		}

		// add word to set
		if _, exists := wSet[w]; !exists {
			wSet[w] = true
		}

	}

	return result
}

func wordMatchHelper(wSet map[string]bool, w string, minLen int) bool {

	// ignore empty words
	if len(w) == 0 {
		return false
	}

	dp := make([]bool, len(w))
	// fmt.Println("word", w, "set", wSet)

	for i := minLen; i < len(w); i++ {

		// if whole string matches
		str := w[:i]
		if _, match := wSet[str]; match {
			// fmt.Printf("match at index %d, for string %s\n", i-1, str)
			dp[i-1] = match
		}
		// partial string matches
		for j := i - minLen; j >= 0; j-- {

			str := w[j+1 : i+1]
			_, match := wSet[str]
			if dp[j] == true && match {
				dp[i] = match
				break
			}
		}

	}
	return dp[len(w)-1]
}

// logic: determine the len of each word, sort the len array
// for each word search its en in the len array, for lens < it's len
// search substring in the array using dp approach,
// dp[i] = true if either  s[0..i] matches or
// 					dp[j] ==true and s[j+1..i] macthes for all j
// if dp[k] == true, whole concat string matched
// trick at each stage substring of len < minWlen should give false
// also for cases when entire word matched, keep a count of total matches

// complexity: T = O(nlogn + nk^2) for sorting and O(k^2) for each word
// you can remove the sorting by storing the entire wordlist in memory
