package strs

// LC # 678

// Tricky:
func checkValidString(s string) bool {

	r := []rune(s)

	stk1 := []int{}
	stk2 := []int{}

	for idx, c := range r {
		if c == '(' {
			stk1 = append(stk1, idx)
		}
		if c == '*' {
			stk2 = append(stk2, idx)
		}
		if c == ')' {
			// first check if stk1 has matching paran
			if len(stk1) > 0 {
				// pop
				stk1 = stk1[:len(stk1)-1]
			} else if len(stk2) > 0 {
				// pop
				stk2 = stk2[:len(stk2)-1]
			} else {
				// no matching paran
				return false
			}
		}
	}

	// end of str
	// case: when non-empty startParan but less * to match it
	if len(stk1) != 0 && len(stk1) > len(stk2) {
		return false
	}

	// now for every unmatched paran there has to be a * coming after it (with > index)
	// note len stk2 is >= len stk1, and indices are in descending order
	for len(stk1) > 0 {
		// pop stk1 and match is pop stk2
		posParan := stk1[len(stk1)-1]
		stk1 = stk1[:len(stk1)-1]

		posStar := stk2[len(stk2)-1]
		stk2 = stk2[:len(stk2)-1]

		// if ( comes after * then unbalanced
		if posParan > posStar {
			return false
		}
	}

	return true
}

// valid string with paranthesis with a * wildcard which can represent 1 paran
// define two stacks, one for '(' and one for '*',
// when ')' is encountered check '(' stk, if empty pop '*' stk, if both empty false
//  at end of str, check if len(stk1) == 0 or len(stk1) <= len(stk2), i.e. remaning *'s can match ('s
