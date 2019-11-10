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

// Simple: Given a string of words, find which belong to the given dictionary
func FindDictWord(s string, dict []string) []string {
	// outp
	outp := []string{}

	// sort the dict
	sort.Strings(dict)
	fmt.Println(dict)

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

// Intermediate: Given a string and dictionary, find the longest substring match
// E.g. Dict: {"to", "toe", "note, "tone", "ones", "toner"},
// Input str: "stones" --> out: {"tone", "ones"}; all same len strings
func FindClosestWord(s string, dict []string) []string {

	var sortedDict []string

	// sort the individual words of the dict
	for _, w := range dict {
		r := str(w)
		sort.Sort(r)
		// add to sortedDict
		sortedDict = append(sortedDict, string(r))
	}

	// sort the sorted dictionary entries
	sort.Strings(sortedDict)
	fmt.Println(sortedDict)

	// sort the given string
	// r := str(s)
	// sort.Sort(r)
	// s = string(r)
	// fmt.Println("Sorted string:", s)

	// find all substrings of sorted string: longest to shortest
	subStrList := FindSubstrings(s)
	// subStr := FindPermutaionsMain(s)

	// sort each substring
	for i := range subStrList {
		r := str(subStrList[i])
		sort.Sort(r)
		subStrList[i] = string(r)
	}
	fmt.Println(subStrList)

	// for all substrings search in sorted dictionary
	// var maxLen int = 0
	var out []string
	for _, sub := range subStrList {
		index := sort.SearchStrings(sortedDict, sub)
		if index < len(dict) && sortedDict[index] == sub {
			// fmt.Println(wList[i])
			out = append(out, sub)
		}
	}

	return out
}

// Given a string find all possible substrings
func FindSubstrings(s string) []string {
	// null string
	if len(s) == 0 {
		return nil
	}

	str := []rune(s)
	var out []string

	// pick a low to high
	for ln := 1; ln <= len(s); ln++ {
		for start := 0; start <= len(s)-ln; start++ {
			subS := string(str[start : start+ln])
			// fmt.Println(subS)
			out = append(out, subS)
		}
	}

	return out
}
