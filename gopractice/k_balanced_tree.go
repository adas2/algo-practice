package practice

// Problem (EPI): A binary tree is defined as k-balanced if at any node, the doffernce between
//	the num nodes in left and right subtree is <= k;
//  Given a binary tree and a positive int k return the first node which which is NOT k-balanced,
//  all of its descendents are k-balanced.
//  E.g.
// 								           A
//                             /                  \
// 					B                                 E
// 			    /      \                            /
//													F
//           C            D                       /    \
//                                              G         I
//                                                \
//                                                H
//
//
// Output: Node E
