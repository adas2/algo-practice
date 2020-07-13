package practice

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRand7(t *testing.T) {

	// pass rand5 as a func variable
	rand5 := func() int {
		// generate seed based on time
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(5)
	}

	fmt.Println("rand7: ", Rand7(rand5))

}
