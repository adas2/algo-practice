package strs

// LC # 316 (hard)

func removeDuplicateLetters(s string) string {

	// create stack of bytes/chars to keep result
	stk := []byte{}
	// alpha map: track of last indices of the char
	aMap := map[byte]int{}
	for i := range s {
		aMap[s[i]] = i
	}
	// set to keep track of elements already seen in result
	aSet := map[byte]bool{}

	// traverse the string
	// i := 0
	for i := range s {
		// if stk is empty add byte
		if len(stk) == 0 {
			stk = append(stk, s[i])
			// update set of chars already seen
			aSet[s[i]] = true
		} else if seen, exists := aSet[s[i]]; exists && seen { // if current char is already in result, ignore (tricky)
			continue
		} else if s[i] > stk[len(stk)-1] { // if curr char is > top of stack, add
			stk = append(stk, s[i])
			aSet[s[i]] = true
		} else {
			// case curr char is < top of stack and the top of stack appears later in the string
			// pop all chars till we have st.tp < curr_char
			for len(stk) > 0 && s[i] < stk[len(stk)-1] && aMap[stk[len(stk)-1]] > i {
				// remove fro seen set and pop
				aSet[stk[len(stk)-1]] = false
				stk = stk[:len(stk)-1]
			}
			// add the curr char
			stk = append(stk, s[i])
			aSet[s[i]] = true
		}
	}

	return string(stk)
}

// Idea: compress the string when you see a char which is smaller and remove ones which appear later
// Space : = O(N) Map and Set each N long
// Time: O(N) linear pass through string
