package practice

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {

	fmt.Printf("Multi Transaction Max profit:  \n")

	p := []int{7, 1, 5, 3, 6, 4}
	p2 := []int{1, 2, 3, 4, 5}

	fmt.Println(maxProfit(p))
	fmt.Println(maxProfit(p2))

}

func TestOneMaxProfit(t *testing.T) {
	fmt.Printf("Single Transaction Max profit:  \n")

	p := []int{7, 1, 5, 3, 6, 4}
	p2 := []int{3, 2, 1}

	fmt.Println(oneMaxProfit(p))
	fmt.Println(oneMaxProfit(p2))
}
