package practice

import (
	"fmt"
	"sort"
	"strings"
)

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
func FindLongestMatch(s string, dict []string) []string {

	sort.Strings(dict)
	// fmt.Println(dict)

	// find all possible substrings of sorted string: longest to shortest
	subStrList := FindSubstrings(s)
	// fmt.Println(subStrList)

	// for all substrings search in sorted dictionary
	// var maxLen int = 0
	var out []string
	var matchLen int = 0
	for _, sub := range subStrList {
		index := sort.SearchStrings(dict, sub)
		if index < len(dict) && dict[index] == sub {
			// fmt.Println("Match:", sub)
			// matches are from longer to shorter strings
			if matchLen <= len(sub) {
				matchLen = len(sub)
				out = append(out, sub)
			} else { // shorter matches
				break
			}

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

	// pick a len long to short
	// for ln := 1; ln <= len(s); ln++ {
	for ln := len(s); ln > 0; ln-- {
		// pick indices
		for start := 0; start <= len(s)-ln; start++ {
			subS := string(str[start : start+ln])
			// fmt.Println(subS)
			out = append(out, subS)
		}
	}

	return out
}
