package practice

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRand7(t *testing.T) {

	// pass rand5 as a func variable
	rand5 := func() int {
		return rand.Intn(5)
	}

	fmt.Println("rand7: ", Rand7(rand5))

}
