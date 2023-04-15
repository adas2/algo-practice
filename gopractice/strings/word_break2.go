package strs

func wordBreak2(s string, wordDict []string) []string {
	// map for O(1) lookup
	wordMap := make(map[string]struct{})
	for _, w := range wordDict {
		wordMap[w] = struct{}{}
	}

	strList := []string{}

	// call util
	wordBreakUtil(wordMap, s, "", 0, &strList)

	return strList
}

func wordBreakUtil(wordMap map[string]struct{}, s, candidate string, start int, strList *[]string) {
	// base case
	if start == len(s) {
		// add full string (trim right space)
		*strList = append(*strList, candidate[:len(candidate)-1])
	}

	// else check prefix ending in start+1 to len(s)-1
	for end := start + 1; end <= len(s); end++ {
		// check if prefix is a word
		if _, exists := wordMap[s[start:end]]; exists {
			wordBreakUtil(wordMap, s, candidate+s[start:end]+" ", end, strList)

		}
	}
}
