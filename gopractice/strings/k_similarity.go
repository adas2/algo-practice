package strs

import (
	"fmt"
	"strings"
)

// Hard Problem: LC # 854 (K-similarity of strings)

func kSimilarity(s1 string, s2 string) int {

	// set of seen strings
	seen := make(map[string]struct{})

	// queue of str nodes in similarity tree
	dqueue := make([]qNode, 0)

	// bfs traversal starting with node s1
	dqueue = append(dqueue, qNode{s1, 0})
	seen[s1] = struct{}{}

	for len(dqueue) > 0 {
		// pop node
		node := dqueue[0]
		dqueue = dqueue[1:]

		// if matches
		if node.str == s2 {
			return node.d
		}

		// find adjacent nodes which are not seen
		adjs := adjStrs(node.str, s2)
		for _, s := range adjs {
			// if !seen.Exists(s) {
			if _, exists := seen[s]; !exists {
				// add node to seen set and queue
				dqueue = append(dqueue, qNode{s, node.d + 1})
				seen[s] = struct{}{}
			}
		}
	}

	return -1
}

type qNode struct {
	str string
	d   int //depth
}

// util to find adjacent nodes in similarity tree
// somewhat similar to recursiver permuations
func adjStrs(s1, s2 string) []string {

	// strings in adjacent
	res := []string{}

	r1, r2 := []rune(s1), []rune(s2)
	// traverse to
	i := 0
	for r1[i] == r2[i] {
		// skip the already matching part
		i++
	}

	// swicth the first letter which differs
	for j := i + 1; j < len(r1); j++ {
		// match with j-th letter in s2 and when swapped does't break j's order (tricky)
		if r1[j] == r2[i] && r2[j] != r1[j] {
			// swap i and j
			r1[i], r1[j] = r1[j], r1[i]
			// add to result
			res = append(res, string(r1))
			// swap back to backtrack
			r1[i], r1[j] = r1[j], r1[i]
		}
	}

	// fmt.Println(res)
	return res
}

// Logic: create a similarity graph for string s1 such that:
// directed edge between s1 ans s1' where s1' is one closer to s2 after a single swap
// similar to calculating permutations of s1 which are simialar to s2
// once the graph is created, do a regular BFS and return depth when node val == s2

/*	*	*	*	* 	*/

// Non-optimal solution: simple recursive impl (this just finds a k not min k)
// Note char are in {a,b,c,d,e,f} find option to simplify
func kSimilarity2(s1 string, s2 string) int {

	// error case
	if len(s1) != len(s2) {
		return -1
	}

	// base case
	if s1 == s2 {
		return 0
	}
	// case where 2 letter anagram
	if len(s1) == 2 {
		return 1
	}

	// if len > 2
	r1, r2 := []rune(s1), []rune(s2)
	// to swicth the first letter which differs
	for i := range r1 {
		if r1[i] != r2[i] {
			// find the index of the letter in the remaing unmatched str
			idx := strings.IndexRune(s1[i+1:], r2[i])
			fmt.Println("Index:", idx)
			// swap
			idx = idx + i + 1 //adjust
			r1[i], r1[idx] = r1[idx], r1[i]
			fmt.Printf("Transform: %s\n", string(r1))
			// recurse
			fmt.Printf("Recurse over %s, %s\n", string(r1[i+1:]), string(r2[i+1:]))
			return 1 + kSimilarity2(string(r1[i+1:]), string(r2[i+1:]))
		}
	}

	//
	return 0
}

// Logic:
// Recursive version: assumption all strs are anagrams
// for len == 2; if s1==s2, k=0, else k=1
// for len > 2; swap the first letter of s1 with first letter of s2 in s1
// recurse over k-similarity (s1[1:], s2[1:])
// abcd --> dcba 2
// abcd --> dabc 3
// T = O(k) where k = str_len, S = O(k) stack space
