package dp

import "fmt"

// LongestRepeatingSubstring return length; overlapping allowed
func LongestRepeatingSubstring(S string) int {

	// create/init LRS 2D slice
	lrs := make([][]int, len(S))
	for i := range S {
		lrs[i] = make([]int, len(S))
	}

	// fill out the 0th row
	for j := 1; j < len(S); j++ {
		if S[j] == S[0] {
			lrs[0][j] = 1
		}
	}

	var maxLen int
	var maxSubtr string
	// fill out the rest of the entries
	// i : 0 -> len-1
	for i := 1; i < len(S); i++ {
		for j := i + 1; j < len(S); j++ {
			// fmt.Println(i, j)
			if S[i] == S[j] {
				lrs[i][j] = lrs[i-1][j-1] + 1
				// maxLen = Max(maxLen, lrs[i][j])
				// update when substr is larger
				if maxLen < lrs[i][j] {
					maxLen = lrs[i][j]
					maxSubtr = S[j-maxLen+1 : j+1] //tricky
				}
			} else {
				lrs[i][j] = 0
			}
		}
	}

	fmt.Printf("Longest repeated substr: %s\n", maxSubtr)
	// return the max len
	return maxLen
}

// Dynamic Porgramming approach (overlap allowed):
// LRS[i][j] --> len of suffix in s[0..j] that matches the prefix in s[0..i]
// if s[i]==s[j] --> LRS[i][j] = LRS[i-1][j-1] + 1
// else LRS[i][j] = 0  ; i.e. of no suffix exists ending in j that matches prefix s[:j]
// NO overlap: check if the length of the match doesn't exceed distance beween i & j
// i.e. if (j-i) > LRS[i-1][j-1] then only update value, else 0
// Note: LRS[n][n] does give the final max value, traverse 2D to get the max value
