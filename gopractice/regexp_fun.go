package practice

import (
    "fmt"
    "regexp"
)

func commandCount(commands []string) []int32 {
    // Write your code here
    out := []int32{}
    // Find all regexp matches for each cmd
    for _, cmd := range commands {
        fmt.Printf("Input: %v\n", cmd)
        // matches := exp.FindAllString(cmd, -1)
        // fmt.Println(matches)
        matches := cmdCountHelper(cmd)
        out = append(out, matches)
    }

    return out
}

// helper func to calcuate the number of
// overlapping regexp matches in a given string
func cmdCountHelper(str string) int32 {
    // find all sunstrings in str
    // min lenght of substring for a match = 4
    strList := findSubstringsLenK(str, 4)
    // match regexp for each substr
    var count int32 = 0
    // this is tricky
    exp, err := regexp.Compile("^[a-z][a-z0-9:]+/[a-z0-9]+\\\\[a-z]+")
    if err != nil {
        fmt.Printf("Error in regexp")
        return -1
    }
    // fmt.Printf("Reg Exp: %v\n", exp)
    for _, sub := range strList {
        matches := exp.FindAllString(sub, -1)
        // fmt.Printf("String: %s\t Matches: %v\n", sub, matches)
        count += int32(len(matches))
    }

    return count
}

// utlity to get all sunstrings of len >=4
// since regexp can match only substrings of len 4+
func findSubstringsLenK(str string, k int) []string {
    if k > len(str) {
        // error
        return nil
    }

    var result []string
    r := []rune(str)
    // from larger to smaller lens
    for ln := len(str); ln >= k; ln-- {
        // start pos from 0 to n-k
        for start := 0; start <= len(str)-ln; start++ {
            sub := string(r[start : start+ln])
            result = append(result, sub)
        }
    }
    return result
}
