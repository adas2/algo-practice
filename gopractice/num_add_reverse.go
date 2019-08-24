package practice

import "fmt"

// Given two numbers in reverse order (as a linked list with single digit),
// add the two numbers; no leading zeros
// 2->4->7 + 5->6->4 = ( 7->0->2->1 ) i.e. 742+465 = 1207
// [Leetcode med]

// struct to hold the digits
type nDig struct {
	digit int
	next  *nDig
}

// list for whole number
type nList *nDig

// CreateNum creates a representation linked list reverse order
func CreateNum(n int) nList {
	div := 10
	var d int
	var head, last, curr *nDig = nil, nil, nil
	for num := n; num > 0; {
		d = num % div
		curr = &nDig{d, nil}
		if head == nil {
			head = curr
		}
		if last != nil {
			last.next = curr
		}
		last = curr
		num /= div
	}

	return head
}

// PrintNum iteratively travreses and prints digits in num
func PrintNum(num nList) {
	var head nList
	for head = num; head != nil; head = head.next {
		fmt.Printf("%d->", head.digit)
	}
	fmt.Printf("nil\n")
}

// NumAdd add such two nums
func NumAdd(a, b nList) nList {
	// traverse list and add digits
	var carry int = 0
	var curr, prev *nDig = nil, nil
	var head1, head2 nList = a, b
	for head1, head2 = a, b; head1 != nil && head2 != nil; head1, head2 = head1.next, head2.next {
		newd := (head1.digit + head2.digit + carry) % 10
		carry = (head1.digit + head2.digit + carry) / 10
		curr = &nDig{newd, nil}
		curr.next = prev
		prev = curr
	}
	for head1 != nil {
		newd := (head1.digit + carry) % 10
		carry = (head1.digit + carry) / 10
		curr = &nDig{newd, nil}
		curr.next = prev
		head1 = head1.next

		prev = curr
	}

	for head2 != nil {
		newd := (head2.digit + carry) % 10
		carry = (head2.digit + carry) / 10
		curr = &nDig{newd, nil}
		curr.next = prev
		head2 = head2.next

		prev = curr
	}

	if carry > 0 {
		curr = &nDig{carry, nil}
		curr.next = prev
	}

	return curr
}
