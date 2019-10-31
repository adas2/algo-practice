package practice

// Find permutations of a string using recursion,
// return a list of strings, (pass output by reference)
func RecursivePermutations(s string, start, end int, out *[]string) {
	// nill string
	if len(s) == 0 {
		return
	}

	if start == end {
		// append to out
		*out = append(*out, s)
		return
	}

	// ieterate on all elements
	for i := start; i <= end; i++ {
		// swap s[0] with s[i]
		// s[start], s[i] = s[i], s[start]
		s = swapChars(s, i, start)

		// call permutation s[1..n]
		RecursivePermutations(s, start+1, end, out)
		// swap back to original string
		// s[i], s[start] = s[start], s[i]
		s = swapChars(s, i, start)
	}

}

func swapChars(s string, i, j int) string {
	// convert to rune
	str := []rune(s)
	str[i], str[j] = str[j], str[i]
	return string(str)
}

// main func using
func FindPermutaionsMain(s string) []string {

	perm_strings := []string{}

	RecursivePermutations(s, 0, len(s)-1, &perm_strings)

	return perm_strings
}

// Logic:
// abcd --> a + perm(bcd), b + perm(acd), c+perm(bad), d+perm(bca)
// bcd --> bcd, bdc, cbd, cdb, dcb, dbc
// cd --> cd, dc
