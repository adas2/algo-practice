package practice

import "sort"

// FindAllSubstrings uses recursive backtracking to find all substring permutations of len == k
func FindAllSubstrings(str, candidate string, k int, used []bool, strList *[]string) {
	// terminating case
	if k == 0 {
		*strList = append(*strList, candidate)
		return
	}

	// create candidate from unused alphabets
	for i := 0; i < len(str); i++ {
		if used[i] == false {
			candidate = candidate + string(str[i])
			used[i] = true
			// recurse
			FindAllSubstrings(str, candidate, k-1, used, strList)
			// backtrack
			used[i] = false
			candidate = candidate[:len(candidate)-1]
		}

	}

}

// FindSubtringPermutations returns
// a list of strings in decaresing len len containing letters of given str
func FindSubtringPermutations(str string) []string {
	// bool slice inited to false
	used := make([]bool, len(str))

	var strList []string

	var candidate string
	// from len(str) to 1 len find all permutations
	for k := len(str); k >= 1; k-- {
		FindAllSubstrings(str, candidate, k, used, &strList)
	}

	return strList
}

// FindApproximateStringMatch :
// Given a dictionary and a word find the longest approximate match
// If more than one match, display all
func FindApproximateStringMatch(dict []string, str string) []string {
	// sort the dictionary
	sort.Strings(dict)

	// find all possible approximate strings
	subStrList := FindSubtringPermutations(str)

	// for each substring large to small find which ones matched
	// keep maxLen counter to monitor last matched size
	var outList []string
	maxLen := 0

	for _, s := range subStrList {
		index := sort.SearchStrings(dict, s)
		// check if result is valid
		if index < len(dict) && dict[index] == s {
			if len(s) >= maxLen {
				// larger string, update lists and maxLen
				outList = append(outList, s)
				maxLen = len(s)
			} else { // smaller string encountered
				break
			}
		}
	}
	return outList
}
