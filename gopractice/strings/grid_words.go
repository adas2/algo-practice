package strs

import (
	"fmt"

	custom "adas2.io/practice/custom_impl"
)

// E.g. grid
// 	 T I P
//	 O Y G
//	 N S K

// complexity: with trie based dictionary
// traversal is like a dfs
// starting from every letter only adjacent letter can be used
// longest word len: m*m = m^2
// num starting positions = m^2
// hence worst case nums letters to search = O(m^4)
// using a trie each lookup is O(k) where k is the len of key
// total time for search := O(mk); worst case = O(m^3), average case ~ O(m^2), best O(m) i.e. no match
// trie build O(N.w) N= total words, w = average words len

// assumes dict is given in the form a trie
func findValidWords(grid [][]string, dict *custom.Trie) []string {
	m := len(grid)
	out := []string{}
	// init visited to false
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	// for all starting index
	for i := range grid {
		for j := range grid[0] {
			candidate := ""
			// optimization: if grid(i,j) is in trie then proceed
			if isCharInTrie(dict.Root, grid[i][j]) {
				dfsGridUtil(grid, i, j, visited, candidate, dict.Root)
			}
		}
	}

	return out
}

// dfs traversal of grid: from a starting point create a string using dfs
func dfsGridUtil(grid [][]string, i, j int, visited [][]bool, partial string, root *custom.TrieNode) {

	m := len(grid)
	// mark src as visited
	visited[i][j] = true
	// create candidate
	partial += grid[i][j]
	// fmt.Printf("%s\n", partial)

	r := []rune(grid[i][j])
	index := r[0] - 'a'

	// letter match
	if root.Childrens[index].IsWordEnd {
		// found a word match
		fmt.Println("word match:", partial)
	}

	// offsets for index
	row := []int{-1, 1, 0, 0}
	col := []int{0, 0, -1, 1}

	// for all safe neighbors recursive call
	for k := range row {
		if isSsafeNeighbor(m, m, i+row[k], j+col[k], visited) {
			// check is neighbor is in trie

			if isCharInTrie(root.Childrens[index], grid[i+row[k]][j+col[k]]) {
				// recurse
				dfsGridUtil(grid, i+row[k], j+col[k], visited, partial, root.Childrens[index])
			}
		}
	}
	// backtrack so that all other permutations are considered
	visited[i][j] = false
	partial = removeLastRune(partial)

}

// return the tuple of indices for valid neighbors
func isSsafeNeighbor(rSize, cSize, i, j int, visited [][]bool) bool {
	// if index is valid and not visited
	if i >= 0 && i < rSize && j >= 0 && j < cSize && visited[i][j] == false {
		return true
	}
	return false
}

// util
func removeLastRune(s string) string {
	if len(s) == 0 {
		fmt.Println("ERROR")
		return ""
	}
	r := []rune(s)
	return string(r[:len(r)-1])
}

// check if starting char is in trie
func isCharInTrie(root *custom.TrieNode, s string) bool {
	if root == nil {
		return false
	}
	r := []rune(s)
	index := r[0] - 'a'
	// fmt.Println("string: ", s, index)
	if root.Childrens[index] == nil {
		return false
	}
	return true
}
