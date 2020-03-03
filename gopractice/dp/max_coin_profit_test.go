package dp

import (
	"fmt"
	"testing"
)

func TestMaxCoinProfit(t *testing.T) {

	C := []int{10, 25, 5, 1, 10, 5}

	result := maxCoinProfit(C)

	fmt.Println("Max profit for Player 1:", result)
}
