package strs

var alpha = [][]rune{
	[]rune("abcde"),
	[]rune("fghij"),
	[]rune("klmno"),
	[]rune("pqrst"),
	[]rune("uvwxy"),
	[]rune("z"),
}

func alphabetBoardPath(target string) string {

	t := []rune(target)

	// start
	i, j := 0, 0

	var final string

	// walk the board as needed
	for _, char := range t {
		// destination letter
		row := int((char - 'a') / 5)
		col := int((char - 'a') % 5)

		final += getBoardPath(i, j, row, col)
		// update curr position
		i, j = row, col
	}

	return final
}

// util to get path from (r1,c1) --> (r2,c2)
func getBoardPath(r1, c1, r2, c2 int) string {

	// fmt.Printf("row: %d, col: %d, new row: %d, new col: %d\n", r1, c1, r2, c2)
	res := ""

	for r2 < r1 {
		res += "U"
		r1--
	}
	for c2 > c1 {
		res += "R"
		c1++
	}
	for c2 < c1 {
		res += "L"
		c1--
	}
	// walk down last to handle case dest='z'
	for r2 > r1 {
		res += "D"
		r1++
	}

	// reaches here when r1==r2 & c1==c2
	return res + "!"
}

// Greedy scheme: find the relative position of the next letter needed
// [][]rune{rune(abcde), rune(fghif)....rune(z)}
// (char-'a')/5 = r, (char-'a')%5 = c
// next move for  r',c'
// while r'<r --> U, while c'<c --> L
// while r'>r --> D, while c'>c --> R
// r==r' && c==c' --> !
