package practice

import "fmt"

// Given two numbers in reverse order (as a linked list with single digit),
// add the two numbers; no leading zeros
// 2->4->7 + 5->6->4 = ( 7->0->2->1 ) i.e. 742+465 = 1207
// [Leetcode med # 2]

// struct to hold the digits
type nDig struct {
	digit int
	next  *nDig
}

// list for whole number
type nList *nDig

// numAdd add such two nums
func numAdd(head1, head2 nList) nList {
	// traverse list and add digits
	carry := 0
	var res, curr, prev *nDig

	// terminate when all are nil and no carry
	for head1 != nil || head2 != nil || carry > 0 {
		v1, v2 := 0, 0

		if head1 != nil {
			v1 = head1.digit
			head1 = head1.next
		}

		if head2 != nil {
			v2 = head2.digit
			head2 = head2.next
		}

		sum := (v1 + v2 + carry) % 10
		carry = (v1 + v2 + carry) / 10

		curr = &nDig{sum, nil}

		// head of the final list is the first node created
		if res == nil {
			res = curr
		}
		if prev != nil {
			prev.next = curr
		}
		prev = curr
	}

	return res
}

// helper: createNum creates a representation linked list reverse order
func createNum(n int) nList {
	var d int
	var head, last, curr *nDig = nil, nil, nil
	for num := n; num > 0; {
		d = num % 10
		curr = &nDig{d, nil}
		if head == nil {
			head = curr
		}
		if last != nil {
			last.next = curr
		}
		last = curr
		num /= 10
	}

	return head
}

// printNum iteratively travreses and prints digits in num
func printNum(num nList) {
	var head nList
	for head = num; head != nil; head = head.next {
		fmt.Printf("%d->", head.digit)
	}
	fmt.Printf("nil\n")
}
