package trees

import (
	"fmt"
)

// Problem (EPI): A binary tree is defined as k-balanced if at any node, the differnce between
//	the num nodes in left and right subtree is <= k;
//  Given a binary tree and a positive int k return the first node which which is NOT k-balanced,
//  all of its descendents are k-balanced.
//  E.g.
// 								           A
//                             /                  \
//
//					B                                 E
// 			    /      \                            /
//													F
//           C            D                       /    \
//                                              G         I
//                                                \
//                                                H
//
//
// Output: Node E (when k=2)
// returns the count at the node and whether the tree is k-balanced
func findKbalancedUtil(root *btreeNode, k int) (bool, int) {
	// traverse root recrusively and keep track of count

	// nil node is k-balanced
	if root == nil {
		return true, 0
	}

	// non nil node
	lcheck, lcount := findKbalancedUtil(root.left, k)
	rcheck, rcount := findKbalancedUtil(root.right, k)

	if lcheck == true && rcheck == true && abs(lcount-rcount) <= k {
		return true, lcount + rcount + 1
	}
	//any one of the above is failing
	// TODO: return the node at which k-imbalance occurs
	return false, -1 //error code is returned instead of count to avoid confusion
}

// alternative for math.Abs for ints
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// main func
func findKbalancedMain(root *btreeNode) {
	if check, _ := findKbalancedUtil(root, 2); check == false {
		fmt.Printf("The tree is not k-balanced\n")
	} else {
		fmt.Printf("The tree is k-balanced\n")
	}
}
