package practice

// LC : 953
// find if the alien dict is sorted or not based in the letter ordering
func isAlienSorted(words []string, order string) bool {

	// create map of the ordering of the alphabets based on the indices they appeat in
	wMap := map[rune]int{}
	r := []rune(order)
	for i, v := range r {
		wMap[v] = i
	}

	// now traverse the words linearly and check if ith word is <= i+1th word
	for idx := 0; idx < len(words)-1; idx++ {
		if diff(words[idx], words[idx+1], wMap) > 0 {
			return false
		}
	}

	return true
}

// helper: find s1-s2
func diff(s1, s2 string, wMap map[rune]int) int {
	r1, r2 := []rune(s1), []rune(s2)
	i, j := 0, 0
	// compare till end of a str
	for i < len(r1) && j < len(r2) {
		if r1[i] == r2[j] {
			i++
			j++
		} else if wMap[r1[i]]-wMap[r2[j]] <= 0 {
			return 0
		} else {
			return 1
		}
	}

	// note: smaller words with shared prefix should come first
	if j == len(r2) && i != len(r1) {
		// second str is finished, but first str still remaining
		return 1
	}

	return 0

}
