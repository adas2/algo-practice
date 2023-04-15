package strs

// LC # 139
func wordBreak(s string, wordDict []string) bool {
	// map for O(1) lookup
	wordMap := make(map[string]struct{})
	for _, w := range wordDict {
		wordMap[w] = struct{}{}
	}

	// call util
	return wordSearch(wordMap, s, 0)
}

func wordSearch(wordMap map[string]struct{}, s string, start int) bool {
	// base case
	if start == len(s) {
		return true
	}

	// else check prefix ending in start+1 to len(s)-1
	for end := start + 1; end <= len(s); end++ {
		// check if prefix is a word
		if _, exists := wordMap[s[start:end]]; exists && wordSearch(wordMap, s, end) == true {
			// fmt.Printf("Matched word %s in Trie at index ending %d\n", string(r[start:i]), i)
			return true
		}
	}
	return false
}

// logic:
// create a trie with dictionary words
// traverse str from start to end, at every index which hits word end in dict, search(trie, index+1)
// base case: the entire string matches; return true, by default return false
// T = O(2^n*k), where k is the avg len of a dict word; since k ~= 20, O(2^n)
// Space = O(m) where m is the number of dict words

// Update: we have subproblems which can be opimized by dp/memoization
// make the above into dp solution:
// dp[i] --> true means len i from 0th index i.e. s[:i] matched
// dp[0] --> true (since nil string always match)
// for len j := 1 ..len(s)
// we split in every index
// dp[i+1] --> dp[i] && exists(wordDict, s[i:])

func wordBreakDp(s string, wordDict []string) bool {

	// dp array to contain result for substr form len 0 to n
	dp := make([]bool, len(s)+1)

	wordMap := make(map[string]struct{})
	for _, w := range wordDict {
		wordMap[w] = struct{}{}
	}

	// base case
	dp[0] = true

	// for all substr len i
	for i := 1; i <= len(s); i++ {
		// partition/break substr into [0..j] and [j+1...i-1]
		for j := 0; j < i; j++ {
			// we check whether any of this subproblem matched
			//  strings created s[0:j]..s[j:i]
			if dp[j] && exists(wordMap, s[j:i]) == true {
				dp[i] = true
			}
		}
	}

	return dp[len(s)]

}

func exists(wordMap map[string]struct{}, s string) bool {
	if _, exists := wordMap[s]; exists {
		return true
	}
	return false
}
