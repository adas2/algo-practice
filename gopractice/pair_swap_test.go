package practice

import (
	"fmt"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	h := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}

	res := swapPairs(&h)

	for res != nil {
		fmt.Printf("%v --> ", res.Val)
		res = res.Next
	}
	fmt.Printf("Nil\n")
}
