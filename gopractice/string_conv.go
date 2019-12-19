package practice

import (
	"bytes"
	"io"
	"strings"
)

func ConvertToString(stream io.Reader) string {
	var buf bytes.Buffer
	buf.ReadFrom(stream)
	return buf.String()
}

func ParseIntoTokens(input string, delims ...rune) []string {

	// define delim func
	f := func(c rune) bool {
		var flag bool
		for _, d := range delims {
			flag = flag || c == d
		}
		return flag
	}
	// can specify fixed number of delimeters
	// f2 := func(c rune) bool {
	// 	return c == ',' || c == ' ' || c == '!'
	// }
	// fmt.Printf("Intermediate: %q\n", strings.FieldsFunc(input, f))

	out := strings.FieldsFunc(input, f)
	return out
}
