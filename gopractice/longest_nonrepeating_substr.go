package practice

// leetcode#

func lengthOfLongestSubstring(s string) int {
    // set of alphabet runes
    aSet := make(map[rune]int)
    maxLen := 0 // max len seen
    startIndex := 0

    rStr := []rune(s)

    // walk through O(n)
    for i, r := range rStr {
        // check if entry present
        if lastIndex, present := aSet[r]; present {
            // update len from last seen index for same alphabet
            startIndex = max(startIndex, lastIndex+1)
            // fmt.Println("start:", startIndex, "last:", lastIndex)
        }
        // update/create index
        aSet[r] = i
        maxLen = max(maxLen, i-startIndex+1)
        // fmt.Println("max:", maxLen)
    }

    return maxLen

}

// logic define a set and check membership every time
// if alphabet not present add and maxLen++
// else clear map and start from the current element again
