package practice

import (
	"fmt"
	"testing"
)

func TestConvertToRoman(t *testing.T) {
	num := 49

	fmt.Println("Num", num, "Roman", ConvertToRoman(num))

}
