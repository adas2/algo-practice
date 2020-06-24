package practice

import "fmt"

// leetcode # 30
// non overlapping version
func findSubstring(s string, words []string) []int {

	if len(words) == 0 || len(s) == 0 {
		return nil //error case
	}

	wLen := len(words[0])
	wListLen := len(words)
	// var startIdx, currIndex int
	// var match bool // indicates when first match has started
	// rune format string
	r := []rune(s)
	result := []int{}

	wSet := initWordSet(words)
	startIdx, currIndex := 0, 0
	matchCount := 0

	// loop till end of string
	for startIdx < len(r) && currIndex < len(r)-wLen+1 {
		// currIndex = startIdx
		str := string(r[currIndex : currIndex+wLen])
		fmt.Println(str)
		fmt.Printf("word set: %v\n", *wSet)

		/*
			// check if match
			if match := checkMatch(*wSet, str); match == 0 {
				(*wSet)[str]++
				matchCount++      // mark presence
				currIndex += wLen //move by 1 word
			} else if match == 1 {
				fmt.Println("word matched before")
				wSet = initWordSet(words) //clears the word map
				(*wSet)[str]++            // maek present
				currIndex += wLen         // move by one word
				matchCount = 1            // add this as first word
			} else { // = 2
				// no match
				matchCount = 0
				fmt.Println("word mismatch")
				wSet = initWordSet(words)
				currIndex++
			}
		*/

		match := updateMatch(wSet, str)
		if match == 1 {
			matchCount++      // mark presence
			currIndex += wLen //move by 1 word
		} else if match == 2 {
			fmt.Println("word count exceeded")
			wSet = initWordSet(words) //clears the word map
			// currIndex += wLen         // move back for a fresh match
			// matchCount = 1            // add this as first word
			// currIndex -= (wLen * (wListLen - 1)) // shift left for a fresh match
			
		} else { //no match
			matchCount = 0
			fmt.Println("word mismatch")
			wSet = initWordSet(words)
			currIndex++
			continue
		}

		// check if elements matched in sequence
		// guaranteed with al elements in set
		// without being deleted and reinitialized
		if matchCount == wListLen && allMatched(*wSet) {
			fmt.Println("all words matched")
			startIdx = currIndex - (wLen * wListLen)
			result = append(result, startIdx)
			// reset for next iteration
			matchCount = 0
			currIndex = startIdx + wLen
		}
	}

	return result
}

// decrement when a word count when a word matches
// returns false if no match
func updateMatch(wSet *map[string]int, s string) int {
	if _, exists := (*wSet)[s]; exists {
		if (*wSet)[s] > 0 {
			(*wSet)[s]--
			return 1
		}
		// count is 0
		return 2
	}
	// not present
	return 0
}

// reinitialize the map with contents of words[]
func initWordSet(words []string) *map[string]int {
	// wordSet := map[string]bool{}
	wordSet := make(map[string]int)
	for _, str := range words {
		wordSet[str]++
	}
	return &wordSet
}

// check if all words are covered
func allMatched(wSet map[string]int) bool {
	for _, val := range wSet {
		if val != 0 {
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
		}
		return 0
	}
	return 2
}

// logic:
// create a map of the given words, will bool flag indicating seen/unseen
// create two indices: startIndex and currIndex to traverse the string
// if there is s match increment currIdx+Len(word)
// if mismatch adjust startIdx and currIdx
// if all words seen keep track of the index, adjust startIdx
