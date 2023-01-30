package practice

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

	curr := head
	// if at least two nodes
	if curr != nil && curr.Next != nil {
		// save head of next recursion
		temp2 := curr.Next.Next

		// swap curr and curr->next
		temp := curr
		curr = curr.Next
		curr.Next = temp
		// recurse over the rest of the list
		curr.Next.Next = swapPairs(temp2)
	}

	return curr
}

// given a linked list, swap the adjacent pairs
// without changing the value at the nodes
// e.g. 1->2->3->4 ==> 2->1->4->3
// 7->8->9 == > 8->7->9
// logic: recursive solution
// if at least two nodes in list
// 		swap(curr, curr->next)
// 		followed by swapPair(curr->next->next)
// else if 1 node
// 		return curr
