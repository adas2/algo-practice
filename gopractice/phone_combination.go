package practice

import (
	"fmt"
	"strings"
)

func letterCombinations(digits string) []string {

	// make map of phone number encodings

	keyMap := make(map[string]string)
	keyMap["2"] = "abc"
	keyMap["3"] = "def"
	keyMap["4"] = "ghi"
	keyMap["5"] = "jkl"
	keyMap["6"] = "mno"
	keyMap["7"] = "pqrs"
	keyMap["8"] = "tuv"
	keyMap["9"] = "wxyz"

	result := make([]string, 0)

	used_map := make(map[string][]bool)
	nums := "23456789"
	for i := 0; i < len(nums); i++ {
		used_map[string(nums[i])] = make([]bool, len(keyMap[string(nums[i])]))
		for j := 0; j < len(keyMap[string(nums[i])]); j++ {
			used_map[string(nums[i])][j] = false
		}
	}
	var candidate string = ""

	// genCombination(digits, len(digits), candidate, keyMap, used_map, &result)
	genCombination2(digits, len(digits), candidate, keyMap, &result)

	return result
}

// No reuse of the letters for same key
func genCombination(digits string, digLen int, candidate string, keyMap map[string]string, used_map map[string][]bool, result *[]string) {
	// termination case:
	if digLen == 0 {
		// don't add empty strings
		if len(candidate) > 0 {
			*result = append(*result, candidate)
		}
		return
	}

	for i := 0; i < len(digits); i++ {

		for idx, val := range used_map[string(digits[i])] {
			if val == false {
				candidate = candidate + string(keyMap[string(digits[i])][idx])
				used_map[string(digits[i])][idx] = true
				genCombination(digits[i+1:], digLen-1, candidate, keyMap, used_map, result)
				used_map[string(digits[i])][idx] = false
				candidate = candidate[:len(candidate)-1]
			}
		}
	}
}

// same letter can be reused for same number key, e.g. aa, bb etc
func genCombination2(digits string, digLen int, candidate string, keyMap map[string]string, result *[]string) {

	// termination case:
	if digLen == 0 {
		// don't add empty strings
		if len(candidate) > 0 {
			*result = append(*result, candidate)
		}
		return
	}

	for i := 0; i < len(digits); i++ {
		if str, ok := keyMap[string(digits[i])]; ok {
			strList := strings.Split(str, "")
			for _, letter := range strList {
				candidate += letter
				genCombination2(digits[i+1:], digLen-1, candidate, keyMap, result)
				candidate = candidate[:len(candidate)-1]
			}
		} else {
			// error case
			err := fmt.Errorf("Digit %s is not found", string(digits[i]))
			fmt.Println(err.Error())
			return
		}
	}
}

// for give number, create all possible sunstrings using recursive backtracking
// e.g. "34" --> map[3] = "def" map[4] ="ghi"
// genCombination(digits, len(digits), used_map[int][]bool, *result)
// backtracking step:
// for i:=0; i<= len(digits); i++ {
//		for key, val:= range (used[int(digits[i])]); val == false {
// used[key] = true
// genCombination(digits[i:], len(digits)-1, used_map, res)
// used[key] = false
// }
// }
// Note: if reuse is accepted, ignore the used_map (2nd impl)
