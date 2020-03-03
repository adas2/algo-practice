package practice

import "fmt"

// leetcode # 30
// non overlapping version
func findSubstring(s string, words []string) []int {

	if len(words) == 0 || len(s) == 0 {
		return nil //error case
	}

	// vars
	wLen := len(words[0])
	wListLen := len(words)
	var startIdx, currIndex int
	// var match bool // indicates when first match has started
	// rune format string
	r := []rune(s)
	result := []int{}
	startIdx = 0

	wSet := initWordSet(words)
	startIdx = 0
	currIndex = 0
	matchCount := 0

	// loop till end of string
	for startIdx < len(r) && currIndex < len(r)-wLen {

		// currIndex = startIdx
		str := string(r[currIndex : currIndex+wLen])
		fmt.Println(str)

		// check if match
		if match := checkMatch(*wSet, str); match == 0 {
			(*wSet)[str] = true
			matchCount++
			currIndex += wLen
		} else if match == 1 {
			fmt.Println("word matched before")
			wSet = initWordSet(words) //re-init
			(*wSet)[str] = true
			currIndex += wLen
			matchCount = 1 // add this as first word
		} else { // = 2
			// no match
			matchCount = 0
			fmt.Println("word mismatch")
			wSet = initWordSet(words)
			currIndex++
		}

		// check if elements matched in sequence
		// guaranteed with al elements in set
		// without being deleted and reinitialized
		if matchCount == wListLen && allMatched(*wSet) {
			fmt.Println("all words matched")
			startIdx = currIndex - (wLen * wListLen)
			result = append(result, startIdx)
			// matchCount = 0
		}
	}

	return result
}

// reinitialize the map with contents of words[]
func initWordSet(words []string) *map[string]bool {
	// wordSet := map[string]bool{}
	wordSet := make(map[string]bool)
	for _, str := range words {
		wordSet[str] = false
	}
	return &wordSet
}

// check if all words covered
func allMatched(wSet map[string]bool) bool {
	for _, val := range wSet {
		if val != true {
			return false
		}
	}
	return true
}

// check if match, delete entry
// two cases: 1. already matched in this pass
// 2. not matched
// return 0 matched first time
// return 1. reinit the map, and start counting from current word
// return 2. reninit the map, don't count
func checkMatch(wSet map[string]bool, s string) int8 {
	if _, exists := wSet[s]; exists {
		if wSet[s] == true {
			return 1
		} else {
			return 0
		}

	}
	return 2
}

// logic:
// create a  of the given words, will bool flag indicating seen/unseen
// create two indices: startIndex and currIndex to traverse the string
// if there is s match increment currIdx+Len(word)
// if mismatch adjust startIdx and currIdx
// if all words seen keep track of the index, adjust startIdx
