package practice

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestConvertToString(t *testing.T) {
	var input = `the testing    is done,  
				 thanks !`

	fmt.Println(reflect.TypeOf(input))
	r := strings.NewReader(input)
	str := ConvertToString(r)
	fmt.Printf("Converted: %s\n", str)

	// tokenize
	fmt.Println(ParseIntoTokens(str, '!', ',', '\n', ' ', '\t'))

}
