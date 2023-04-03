package strs

// LC # 79 (grid problemm similar to `grid_words.go``)

func exist(board [][]byte, word string) bool {

	ans := false
	// var candidate string = ""
	for i := range board {
		for j := range board[0] {
			if board[i][j] == byte(word[0]) {
				board[i][j] = '#'
				existUtil(board, i, j, string(word[0]), word[1:], &ans)
				board[i][j] = word[0]
			}
		}
	}
	return ans
}

// TODO: fix the recursion to include the candidate="" case
// instead of candidate, change to index of the main words,
// terminate when index == len(word), this way less str copy every recursion
func existUtil(board [][]byte, i, j int, candidate, word string, matched *bool) {

	// fmt.Println(candidate, word)
	if len(word) == 0 {
		// full match
		*matched = true
		return
	}

	roff := []int{1, -1, 0, 0}
	coff := []int{0, 0, 1, -1}

	for k := 0; k < 4; k++ {
		if i+roff[k] >= 0 && i+roff[k] < len(board) && j+coff[k] >= 0 && j+coff[k] < len(board[0]) &&
			board[i+roff[k]][j+coff[k]] == word[0] {
			// mark cell as chosen
			board[i+roff[k]][j+coff[k]] = '#' // since # is not in alphabet
			existUtil(board, i+roff[k], j+coff[k], candidate+string(word[0]), word[1:], matched)
			board[i+roff[k]][j+coff[k]] = word[0]
		}
	}

}
