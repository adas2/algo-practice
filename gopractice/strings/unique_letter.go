package strs

// LC # 828 (Hard)
func uniqueLetterString(s string) int {

	num_chars := 128
	offset := byte('A')

	// define 2 arrays
	lastPosition := make([]int, num_chars)
	contribution := make([]int, num_chars)

	// init lastPosiiton for each char to -1 since it didn;t appear yet
	for i := range lastPosition {
		lastPosition[i] = -1
	}

	// base case: 0th index pre-compute
	var curr, prev int

	res := prev

	for i := 0; i < len(s); i++ {
		// num uniqe = prev num unique + new contribution - redundant contribution
		curr = prev + (i - lastPosition[s[i]-offset]) - contribution[s[i]-offset]
		contribution[s[i]-offset] = (i - lastPosition[s[i]-offset])
		lastPosition[s[i]-offset] = i
		res += curr
		// fmt.Printf("Prev [%d] : %d, Curr [%d] : %d, res : %d\n", i-1, prev, i, curr, res)
		prev = curr
	}

	return res
}

// tricky dp type logic:
// string ending in ith index has unique letters
// for each index track the following:
// 	contribution[s[i]] = i-showLastPosition[s[i]]
// 	showLastPosition[s[i]] = -1 default, i elsewhere
// we update these once we process i in the current iteration
// at each iteration i, unique letters
//  curr[i] = curr[i-1] + (i- showLastPosition[s[i]]) - contribution[s[i]]
// since we only need i, i-1 values we store in prev and curr

// LC # 2622
// naive way to impl the above (time limit exceeds for large inputs)
// func to find appeal(unique letter) of substrings of a given str
func appealSum(s string) int64 {
	var n int64 = 0
	// for each len find a str
	for i := 1; i <= len(s); i++ {
		for j := 0; j <= len(s)-i; j++ {
			n += countLetters(s[j : j+i])
		}
	}

	return n
}

// util to count unique letters in a str
func countLetters(s string) int64 {

	r := []rune(s)
	letterMap := make(map[rune]struct{})

	for i := 0; i < len(s); i++ {
		if _, exists := letterMap[r[i]]; !exists {
			letterMap[r[i]] = struct{}{}
		}
	}

	return int64(len(letterMap))
}
