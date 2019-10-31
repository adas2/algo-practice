package practice

import (
	"fmt"
	"sort"
	"strings"
)

// define rune type for string operations
type str []rune

// define interface
func (r str) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r str) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r str) Len() int {
	return len(r)
}

// Given a sentence and find the missing letters
func FindMissingLetters(s string) []string {
	// Create alphabet Map
	aMap := map[rune]bool{}
	for i := 1; i <= 26; i++ {
		aMap[toChar(i)] = false
	}
	// convert the given string to lower case
	// and convert to rune array
	str := []rune(strings.ToLower(s))
	fmt.Println(s)
	// Traverse the sentence string and mark the chars visited
	for j := range str {
		if _, present := aMap[str[j]]; present {
			aMap[str[j]] = true
		}
	}

	// fmt.Println(aMap)
	var output []string

	// retrun the unvisited chars with zero count
	for k, v := range aMap {
		if used := v; used == false {
			// fmt.Println(string(k))
			output = append(output, string(k))
		}
	}

	return output
}

// Given a string of words, find which belong to a dictionary
func FindDictWord(s string, dict []string) []string {
	// outp
	outp := []string{}

	// sort the dict
	sort.Strings(dict)
	fmt.Println(dict)

	// sort the individual words of the dict
	// for _, w := range dict {
	// 	r := []rune(w)
	// 	sort.Sort(str(r))
	// 	fmt.Println(string(r))
	// }

	// create map of the sorted dic strings
	// split sentence into words (space separate)
	wList := strings.Split(strings.ToLower(s), " ")
	fmt.Println(wList)
	// check each word in the sorted dict
	for i := range wList {
		index := sort.SearchStrings(dict, wList[i])
		if index < len(dict) && dict[index] == wList[i] {
			// fmt.Println(wList[i])
			outp = append(outp, wList[i])
		}
	}
	// return list of strings that are not in sorted dic
	return outp
}

// rune is the code point conversion of a char
func toChar(i int) rune {
	return rune('a' - 1 + i)
}
