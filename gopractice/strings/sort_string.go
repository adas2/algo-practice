package strs

import "sort"

// Sort characters in a string

// define rune type for string operations
type str []rune

// define interface
func (r str) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r str) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r str) Len() int {
	return len(r)
}

// SortString is a better way other than splitting the string into characters
func SortString(s string) string {
	// convert to rune
	r := str(s)

	sort.Sort(r)

	return string(r)
}
